package computer

import (
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type SysInfo struct {
	Hostname     string `json:"hostname"`
	Platform     string `json:"platform"`
	Uptime       uint64 `json:"uptime""`
	RAM          uint64 `json:"ram""`
	RAMAvailable uint64 `json:"ram available""`
	RAMFree      uint64 `json:"ram free""`
	Disk         uint64 `json:"disk"`
}

func Info() *SysInfo {
	hostStat, _ := host.Info()
	vmStat, _ := mem.VirtualMemory()
	diskStat, _ := disk.Usage("\\") // If you're in Unix change this "\\" for "/"

	info := new(SysInfo)

	info.Hostname = hostStat.Hostname
	info.Platform = hostStat.Platform
	info.Uptime = hostStat.Uptime
	info.RAM = vmStat.Total / 1024 / 1024
	info.RAMAvailable = vmStat.Available / 1024 / 1024
	info.RAMFree = vmStat.Free / 1024 / 1024
	info.Disk = diskStat.Total / 1024 / 1024

	return info
}
