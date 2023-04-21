package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"strings"
	"path/filepath"

	"golang.org/x/sys/unix"
)

// position of needed data
const (
	// 36  35  98:0 /mnt1 /mnt2 rw,noatime master:1 - ext3 /dev/root rw,errors=continue
	// (0) (1) (2)   (3)   (4)      (5)      (6)   (7) (8)    (9)           (10)

	// value of st_dev
	mountinfoMajorMinor = 2
	// last column in final table
	mountinfoMountPoint = 4
	// optional fields
	mountinfoOptionalFields = 6
	// fs type
	mountinfoFsType = 8	
	// first column
	mountinfoMountSource = 9
)

// ds to handle ALL info about mount device
type Mount struct {
	Device     string
	DeviceType string
	Mountpoint string
	Fstype     string
	Type       string
	Opts       string
	Total      uint64
	Free       uint64
	Used       uint64
	Inodes     uint64
	InodesFree uint64
	InodesUsed uint64
	Blocks     uint64
	BlockSize  uint64
	Metadata   interface{}
}

func (m *Mount) Stat() unix.Statfs_t {
	return m.Metadata.(unix.Statfs_t) // cast to struct
}

// func parseMountInfo() []string {
// 	filename := "/proc/self/mountinfo"

// 	lines, err := readLines(filename)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return lines
// }

func mounts(readSource int, singleFilePath ...[]string) ([]Mount, []string, error) {
	var warnings []string

	filename := "/proc/self/mountinfo"

	lines, err := readLines(filename)
	if err != nil {
		return nil, nil, err
	}

	ret := make([]Mount, 0, len(lines))

	for _, line := range lines {
		i, fields := parseMountInfoLine(line)
		if i == 0 {
			continue
		}

		// check if number of fields matches with disered value
		if i != 11 {
			warnings = append(warnings, fmt.Sprintf("found invalid mountinfo line: %s", line))
			continue
		}

		// get desired fields from mountinfo file
		device := fields[mountinfoMountSource]
		mountPoint := fields[mountinfoMountPoint]
		fsType := fields[mountinfoFsType]
		stDev := fields[mountinfoMajorMinor]

		var stat unix.Statfs_t
		err := unix.Statfs(mountPoint, &stat)
		if err != nil {
			if err != os.ErrPermission {
				warnings = append(warnings, fmt.Sprintf("%s: %v", mountPoint, err))
				continue
			}
		}

		// create Mount entity for current mountpoint
		d := Mount{
			Device:     device,
			DeviceType: "",
			Mountpoint: mountPoint,
			Fstype:     fsType,
			Type:       fsTypeMap[int64(stat.Type)],
			Opts:       stDev,
			Total:      (uint64(stat.Blocks) * uint64(stat.Bsize)) / 1024,
			Free:       (uint64(stat.Bavail) * uint64(stat.Bsize)) / 1024,
			Used:       ((uint64(stat.Blocks) - uint64(stat.Bfree)) * uint64(stat.Bsize)) / 1024,
			Inodes:     stat.Files,
			InodesFree: stat.Ffree,
			InodesUsed: stat.Files - stat.Ffree,
			Blocks:     uint64(stat.Blocks),
			BlockSize:  uint64(stat.Bsize),
			Metadata:   stat,
		}
		d.DeviceType = deviceType(d)
		ret = append(ret, d)
	}

	if readSource == readFromArgs {
		var m []Mount 
		for _, v := range(singleFilePath[0]) {
			var fm []Mount
			fm, err = findMounts(ret, v)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			m = append(m, fm...)
		}
		ret = m
	}

	return ret, warnings, nil
}

// return number of parsed fields and their values in /proc/self/mountinfo
func parseMountInfoLine(line string) (int, [11]string) {
	var fields [11]string

	if len(line) == 0 || len(line) == 1 {
		return 0, fields
	}

	var i int
	for _, f := range strings.Fields(line) {
		if i == mountinfoOptionalFields {
			// loop until find separator
			if f != "-" {
				if fields[i] == "" {
					fields[i] += f
				} else {
					fields[i] += " " + f
				}

				continue
			}

			i++
		}

		fields[i] = f
		i++
	}

	return i, fields
}

// get all lines from file
func readLines(filename string) ([]string, error) {
	lines, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer lines.Close()

	var res []string
	scanner := bufio.NewScanner(lines)

	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	return res, scanner.Err()
}

func findMounts(mounts []Mount, path string) ([]Mount, error) {
	var err error
	path, err = filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	path, err = filepath.EvalSymlinks(path)
	if err != nil {
		return nil, err
	}

	_, err = os.Stat(path)
	if err != nil {
		return nil, err
	}

	var m []Mount
	for _, v := range mounts {
		if path == v.Device {
			return []Mount{v}, nil
		}

		if strings.HasPrefix(path, v.Mountpoint) {
			var nm []Mount

			// keep all entries that are as close or closer to the target
			for _, mv := range m {
				if len(mv.Mountpoint) >= len(v.Mountpoint) {
					nm = append(nm, mv)
				}
			}
			m = nm

			// add entry only if we didn't already find something closer
			if len(nm) == 0 || len(v.Mountpoint) >= len(nm[0].Mountpoint) {
				m = append(m, v)
			}
		}
	}

	return m, nil
}