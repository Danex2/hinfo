package main

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	host, _ := host.Info()
	cpu, _ := cpu.Info()
	mem, _ := mem.VirtualMemory()
	fmt.Printf(`				
General						
-------
Hostname: %v           
Platform: %v
Build/Version: %v   
Architecture: %v

CPU
-------
Model: %v 
Cores: %v
Frequency: %v Mhz

Memory
-------
Total: %v MB
Available: %v MB
Used: %v MB
Used Percent: %v%%`,
		host.Hostname,
		host.Platform,
		host.PlatformVersion,
		host.KernelArch,
		cpu[0].ModelName,
		cpu[0].Cores,
		cpu[0].Mhz,
		mem.Total/1024/1024,
		mem.Available/1024/1024,
		mem.Used/1024/1024,
		mem.UsedPercent,
	)
}
