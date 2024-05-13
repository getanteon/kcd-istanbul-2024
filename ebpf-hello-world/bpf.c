#include "vmlinux.h"

#include <bpf/bpf_helpers.h>

struct {
    __uint(type, BPF_MAP_TYPE_PERF_EVENT_ARRAY);
    __uint(key_size, sizeof(int));
    __uint(value_size, sizeof(int));
} write_buffers SEC(".maps");

struct buffer_event{
    unsigned char payload[256];
};


struct trace_event_raw_sys_enter_write {
	struct trace_entry ent;
    __s32 __syscall_nr;
    __u64 fd;
    char * buf;
    __u64 count;
};

SEC("tracepoint/syscalls/sys_enter_write")
int sys_enter_write(struct trace_event_raw_sys_enter_write* ctx) {
    __u64 id = bpf_get_current_pid_tgid();
    __u32 pid = id >> 32;

    // change pid to see write content of client or the server
    if (pid == 533973){    
    bpf_printk("write: fd:%d , count:%d\n", ctx->fd, ctx->count);
    // copy buffer to user space

    struct buffer_event e = {};
    bpf_probe_read(&e.payload, sizeof(e.payload), (const void *)ctx->buf);

    long r = bpf_perf_event_output(ctx, &write_buffers, BPF_F_CURRENT_CPU, &e, sizeof(e));
    if (r < 0) {
        bpf_printk("failed to pass buffer to user space: %d\n", r);
    }
   }

   return 0;
}

char __license[] SEC("license") = "Dual MIT/GPL";
