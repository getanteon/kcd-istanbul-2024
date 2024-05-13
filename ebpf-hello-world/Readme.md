## Compile eBPF programs

```bash
clang-14 -O2 -g  -Wall -Werror -target bpf -c bpf.c -I ../
```

This command should generate bpf object files bpf.o

```bash
ls
bpf.c  bpf.o 
```

# Load bpf program with cilium's ebpf lib
```bash
go run --exec=sudo main.go
```

## Load programs into kernel with bpftool 

Using ```bpftool```, we'll try to load into kernel.
At this point verifier will run its checks on them.

// You'll need root permission

Load the prog into kernel
```bash
sudo bpftool prog load bpf.o /sys/fs/bpf/hello autoattach
```

Unload prog
```bash
sudo rm /sys/fs/bpf/hello
```