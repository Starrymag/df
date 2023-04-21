package main

import (
	"fmt"
	"math"
	"strconv"
)

func printTable(c Config) {
	var m []Mount
	if c.whichFiles == oneFile {
		m, _, _ = mounts(readFromArgs, c.singleFilePath)
	} else if c.whichFiles == allFiles {
		m, _, _ = mounts(readFromFile)
	} else if c.whichFiles == defaultFiles {
		m, _ = getDefaultData()
	}

	if c.displayFormat == sizeFormat {
		if c.notationType == humanBinNotation || c.notationType == humanDecNotation {
			fmt.Printf("Файл. система           Размер Использовано    Дост Использовано%% Смонтировано в\n")
		} else {
			fmt.Printf("Файл. система           1К-блоков Использовано Доступно Использовано%% Смонтировано в\n")
		}
		for _, mount := range m {
			device := mount.Device
			mountPoint := mount.Mountpoint
			total := mount.Total
			free := mount.Free
			used := mount.Used
			usagePercent := "-"
			// do block multiple 1K size
			blocks := mount.Blocks * (mount.BlockSize / 1024)
			if total != 0 {
				usagePercent = strconv.FormatFloat(math.Ceil((float64(used)/float64(total))*100), 'f', 0, 64) + "%"
			}
			if c.notationType == humanBinNotation || c.notationType == humanDecNotation {
				var freeStr, usedStr, blocksStr string
				if c.notationType == humanBinNotation {
					freeStr = ByteCountBin(free)
					usedStr = ByteCountBin(used)
					blocksStr = ByteCountBin(mount.Blocks * (mount.BlockSize / 1024))
				} else if c.notationType == humanDecNotation {
					freeStr = ByteCountDec(mount.Free)
					usedStr = ByteCountDec(mount.Used)
					blocksStr = ByteCountDec(mount.Blocks * (mount.BlockSize / 1024))
				}
				fmt.Printf("%-23s %6s %12s %7s %13s %-25s\n", device, blocksStr, usedStr, freeStr, usagePercent, mountPoint)
			} else {
				fmt.Printf("%-23s %9d %12d %8d %13s %-25s\n", device, blocks, used, free, usagePercent, mountPoint)
			}
		}
	} else if c.displayFormat == inodeFormat {
		if c.notationType == humanBinNotation || c.notationType == humanDecNotation {
			fmt.Printf("Файл. система           Iнодов IИспользовано Iсвободно IИспользовано%% Смонтировано в\n")
		} else {
			fmt.Printf("Файл. система             Iнодов IИспользовано Iсвободно IИспользовано%% Смонтировано в\n")
		}
		for _, mount := range m {
			device := mount.Device
			mountPoint := mount.Mountpoint
			total := mount.Inodes
			free := mount.InodesFree
			used := mount.InodesUsed
			usagePercent := "-"
			if total != 0 {
				usagePercent = strconv.FormatFloat(math.Ceil((float64(used)/float64(total))*100), 'f', 0, 64) + "%"
			}
			if c.notationType == humanBinNotation || c.notationType == humanDecNotation {
				var freeStr, usedStr, totalStr string
				if c.notationType == humanBinNotation {
					freeStr = InodeCountBin(free)
					usedStr = InodeCountBin(used)
					totalStr = InodeCountBin(total)
				} else if c.notationType == humanDecNotation {
					freeStr = InodeCountDec(free)
					usedStr = InodeCountDec(used)
					totalStr = InodeCountDec(total)
				}
				fmt.Printf("%-23s %6s %13s %9s %15s %-25s\n", device, totalStr, usedStr, freeStr, usagePercent, mountPoint)
			} else {
				fmt.Printf("%-23s %8d %13d %9d %15s %-25s\n", device, total, used, free, usagePercent, mountPoint)
			}
		}
	}
}
