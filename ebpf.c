#include <linux/bpf.h>
#include <linux/if_ether.h>
#include <linux/ip.h>
#include <linux/tcp.h>
#include <bpf/bpf_helpers.h>

// Set Default 
#define Default_Port 4040
#define IPPROTO_TCP 6
#define bpf_htons(x) ((__be16)___constant_swab16((x)))

SEC("xdp")
int xdp_program(struct xdp_md *ctx) {
    void *data_end = (void *)(long)ctx->data_end;
    void *data = (void *)(long)ctx->data;

    // parsing ethernet header and validate 
    struct ethhdr *ethernet = data;
    if ((void *)(ethernet + 1) > data_end) {
        return XDP_PASS;
    }

    // check packet ip is ipv4 type by utilize the ethernet header 
    // conver constend value to 16 for checking
    if (ethernet->h_proto != bpf_htons(ETH_P_IP)){
        return XDP_PASS;
    }

    struct iphdr *ip = data + sizeof(struct ethhdr);
    if ((void *)(ip + 1) > data_end) {
        return XDP_PASS;
    }

    // checking if tcp protocol or not
    if (ip->protocol == IPPROTO_TCP) {
        // if it is tcp get the tcp header and 
        // validate with packet have enough data
        struct tcphdr *tcp = data + sizeof(struct ethhdr) + sizeof(struct iphdr);
        if ((void *)(tcp + 1) > data_end) {
            return XDP_PASS;
        }

        // if port are equal drop the packet from xdp level without reaching to kernel
        if (tcp->dest == bpf_htons(Default_Port)){
            return XDP_DROP;
        }
    }

    return XDP_PASS;
}

char _license[] SEC("license") = "GPL";
