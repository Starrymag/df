package main

import (
	"fmt"
	"golang.org/x/exp/slices"
)

func getDefaultData(m []Mount) ([]Mount, error) {
	var resultMounts []Mount
	var uniqId []string
	for i := 0; i < len(m); i++ {
		if ((m[i].DeviceType == localDevice || (m[i].Type == "tmpfs" && m[i].Fstype != "devtmpfs")) && m[i].Total != 0) {
			if !slices.Contains(uniqId, m[i].Opts) {
				uniqId = append(uniqId, m[i].Opts)
				resultMounts = append(resultMounts, m[i])
			}
			// fmt.Println(m[i].DeviceType, m[i].Device, m[i].Fstype, m[i].Type)
		}
	}
	return resultMounts, nil
}

func getHumanData(m []Mount) ([]Mount, error) {
	// devider := float64(1024 * 1024)
	resultMounts := m
	for i := 0; i < len(m); i++ {
		// tot, _ := strconv.Atoi(ByteCountBinary(int64(m[i].Total)))
		// use, _ :=  strconv.Atoi(ByteCountBinary(int64(m[i].Used)))
		// fre, _ :=  strconv.Atoi(ByteCountBinary(int64(m[i].Free)))
		// blo, _ :=strconv.Atoi(ByteCountBinary(int64(m[i].Blocks)))
		
		// resultMounts[i].Total = uint64(tot)
		// resultMounts[i].Used =  uint64(use)
		// resultMounts[i].Free =  uint64(fre)
		// resultMounts[i].Blocks =uint64(blo) 
		
	}
	return resultMounts, nil
}

func ByteCountBinary(b int64) string {
        const unit = 1024
        if b < unit {
                return fmt.Sprintf("%d B", b)
        }
        div, exp := int64(unit), 0
        for n := b / unit; n >= unit; n /= unit {
                div *= unit
                exp++
        }
        return fmt.Sprintf("%.1f %ciB", float64(b)/float64(div), "KMGTPE"[exp])
}