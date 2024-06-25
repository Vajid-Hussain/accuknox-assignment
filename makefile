cbuild:
	clang -O2 -target bpf -c ebpf.c -o ebpf.o

defanition:
	readelf -a ebpf.o

gobuild:
	go build -o executable main.go

run:
	sudo ./executable