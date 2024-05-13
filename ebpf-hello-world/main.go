package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/perf"
)

type bpfObjects struct {
	bpfPrograms
	bpfMaps
}

type bpfPrograms struct {
	SysEnterWrite *ebpf.Program `ebpf:"sys_enter_write"`
}

type bpfMaps struct {
	WriteBuffers *ebpf.Map `ebpf:"write_buffers"`
}

func main() {
	// Load pre-compiled programs and maps into the kernel.
	reader := bytes.NewReader(_BpfBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		log.Fatalf("can't load bpf: %w", err)
	}

	obj := &bpfObjects{}

	err = spec.LoadAndAssign(obj, nil)
	if err != nil {
		log.Fatalf("can't load bpf: %w", err)
	}

	var writeBuffers *perf.Reader
	writeBuffers, err = perf.NewReader(obj.WriteBuffers, 4*os.Getpagesize())
	if err != nil {
		log.Fatalf("can't create perf reader: %v", err)
	}

	// Attach the program to the tracepoint.
	l1, err := link.Tracepoint("syscalls", "sys_enter_write", obj.SysEnterWrite, nil)
	if err != nil {
		log.Fatalf("error linking sys_enter_write tracepoint: %v", err)
	}

	log.Default().Print("Attached to sys_enter_write tracepoint.")
	defer l1.Close()
	for {
		record, err := writeBuffers.Read()
		if err != nil {
			log.Fatalf("error reading from perf buffer: %v", err)
		}

		// buf := [256]byte
		printBoxedLog(string(record.RawSample))
	}
}

//go:embed bpf.o
var _BpfBytes []byte

func printBoxedLog(message string) {
	// Define the horizontal border length based on the message length
	border := "+" + strings.Repeat("-", len(message)+2) + "+"

	fmt.Println(border)
	fmt.Printf("| %s |\n", message)
	fmt.Println(border)
}
