package sysinfo

import (
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"runtime"
)

type InfoUtil interface {
	Info() *SysInfo
}

type InfoUtilImpl struct {
}

type SysInfo struct {
	Hostname     string `json:"hostname"`
	Platform     string `json:"platform"`
	Uptime       uint64 `json:"uptime"`
	RAM          uint64 `json:"ram"`
	RAMAvailable uint64 `json:"ram available"`
	RAMFree      uint64 `json:"ram free"`
	Disk         uint64 `json:"disk"`
}

func (util InfoUtilImpl) Info() *SysInfo {
	hostStat, _ := host.Info()
	vmStat, _ := mem.VirtualMemory()

	var path string
	os := runtime.GOOS

	if os == "windows" {
		path = "\\"
	} else {
		path = "/"
	}
	diskStat, _ := disk.Usage(path)
	info := new(SysInfo)

	if hostStat != nil {
		info.Hostname = hostStat.Hostname
		info.Platform = hostStat.Platform
		info.Uptime = hostStat.Uptime
	}

	if vmStat != nil {
		info.RAM = vmStat.Total / 1024 / 1024
		info.RAMAvailable = vmStat.Available / 1024 / 1024
		info.RAMFree = vmStat.Free / 1024 / 1024
	}

	if diskStat != nil {
		info.Disk = diskStat.Total / 1024 / 1024
	}

	return info
}
