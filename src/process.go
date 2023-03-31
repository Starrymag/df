package main

import (
	// "fmt"
	"strings"
)

func getDefaultData(m []Mount) ([]Mount, error) {
	var resultMounts []Mount
	for i := 0; i < len(m); i++ {
		if (m[i].DeviceType == localDevice || (m[i].Type == "tmpfs" && m[i].Fstype != "devtmpfs" && !strings.Contains(m[i].Mountpoint, "snap"))) {
			resultMounts = append(resultMounts, m[i])
			// fmt.Println(m[i].DeviceType, m[i].Device, m[i].Fstype, m[i].Type)
		}
	}
	return resultMounts, nil
}

func getHumanData(m []Mount) ([]Mount, error) {
	devider := float64(1024 * 1024)
	resultMounts := m
	for i := 0; i < len(m); i++ {
		resultMounts[i].Total = uint64(float64(m[i].Total) / devider)
		resultMounts[i].Used = uint64(float64(m[i].Used) / devider)
		resultMounts[i].Free = uint64(float64(m[i].Free) / devider)
		resultMounts[i].Blocks = uint64(float64(m[i].Blocks) / devider)
	}
	return resultMounts, nil
}
