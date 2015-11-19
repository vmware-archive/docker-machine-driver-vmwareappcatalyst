package main

import (
	"github.com/docker/machine/libmachine/drivers/plugin"
	"github.com/vmware/docker-machine-driver-vmwareappcatalyst"
)

func main() {
	plugin.RegisterDriver(new(vmwareappcatalyst.Driver))
}
