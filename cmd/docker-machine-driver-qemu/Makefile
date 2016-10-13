default: build

build:
	GOOS=linux  GOGC=off go build -ldflags "-extldflags -static" -i -o docker-machine-driver-qemu
