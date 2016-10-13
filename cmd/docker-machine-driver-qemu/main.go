package main

import (
	"github.com/moretea/docker-machine-driver-qemu"
	"github.com/docker/machine/libmachine/drivers/plugin"
)

func main() {
	plugin.RegisterDriver(qemu.NewDriver("default", "path"))
}
