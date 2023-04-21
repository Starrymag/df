package main

import (
	"fmt"
	"math"
	"golang.org/x/exp/slices"
)

func getDefaultData() ([]Mount, error) {
	m, _, _ := mounts(readFromFile)
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

// func getHumanDataBin(m Mount) (Mount, error) {
// 	resultMount
// 	for i := 0; i < len(m); i++ {
// 		resultMount.Total = ByteCountBin(int64(m[i].Total))
// 		resultMount.Used = ByteCountBin(int64(m[i].Used))
// 		resultMounts.Free = ByteCountBin(int64(m[i].Free))
// 	}
// 	return resultMounts, nil
// }

// func getHumanDataDec(m []Mount) ([]MountString, error) {
// 	resultMounts := make([]MountString, 0, len(m))
// 	for i := 0; i < len(m); i++ {
// 		resultMounts[i].Total = ByteCountDec(int64(m[i].Total))
// 		resultMounts[i].Used = ByteCountDec(int64(m[i].Used))
// 		resultMounts[i].Free = ByteCountDec(int64(m[i].Free))
// 	}
// 	return resultMounts, nil
// }

func ByteCountDec(b uint64) string {
	if b == 0 {
		return "0"
	}
	bf := float64(b)
	for _, unit := range []string{"K", "M", "G", "T", "P", "E", "Z"} {
		if math.Abs(bf) < 1000.0 {
			return fmt.Sprintf("%3.1f%s", bf, unit)
		}
		bf /= 1000.0
	}
	return fmt.Sprintf("%.1f", bf)
}

func ByteCountBin(b uint64) string {
	if b == 0 {
		return "0"
	}
	bf := float64(b)
	for _, unit := range []string{"K", "M", "G", "T", "P", "E", "Z"} {
		if math.Abs(bf) < 1024.0 {
			return fmt.Sprintf("%3.1f%s", bf, unit)
		}
		bf /= 1024.0
	}
	return fmt.Sprintf("%.1f", bf)
}

func InodeCountDec(b uint64) string {
	if b == 0 {
		return "0"
	}
	bf := int(b)
	for _, unit := range []string{"", "K", "M", "B"} {
		if bf < 1000 {
			return fmt.Sprintf("%3d%s", bf, unit)
		}
		bf /= 1000
	}
	return fmt.Sprintf("%d", int(bf))
}

func InodeCountBin(b uint64) string {
	if b == 0 {
		return "0"
	}
	bf := int(b)
	for _, unit := range []string{"", "K", "M", "B"} {
		if bf < 1024 {
			return fmt.Sprintf("%d%s", bf, unit)
		}
		bf /= 1024
	}
	return fmt.Sprintf("%d", int(bf))
}