package main

import (
	"fmt"
	"log"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
	"github.com/vishvananda/netlink"
)

func main() {
	objectFilePath := "ebpf.o"

	// Load the ebpf object file
	spec, err := ebpf.LoadCollectionSpec(objectFilePath)
	catchErr(err)

	fmt.Printf("Loaded eBPF spec: %+v\n", spec)

	// load the pogram to kernal
	coll, err := ebpf.NewCollection(spec)
	catchErr(err)

	program := coll.Programs["xdp_program"]
	if program == nil {
		log.Fatalf("Failed to find XDP program in collection")
	}

	// Attach the program to the network interface
	interfaceName := "wlp2s0"

	interfaces, err := netlink.LinkByName(interfaceName)
	if err != nil {
		log.Fatalf("Failed to get network interface %s: %v", interfaceName, err)
	}

	link, err := link.AttachXDP(link.XDPOptions{
		Program:   program,
		Interface: interfaces.Attrs().Index,
		Flags:     link.XDPGenericMode,
	})
	if err != nil {
		log.Fatalf("Failed to attach XDP program to interface %s: %v", interfaceName, err)
	}
	defer link.Close()

	fmt.Printf("Attached XDP program to interface %s (index %d)\n", interfaceName, interfaces.Attrs().Index)

	// Keep the program running
	fmt.Println("Program running, press Ctrl+C to exit")
	select {}
}

func catchErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
