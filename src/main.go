package main

import (
	"github.com/ogier/pflag"
)

// constants to describe different options
const (
	// which files
	oneFile = iota
	allFiles = iota
	defaultFiles = iota

	// display format
	sizeFormat = iota
	inodeFormat = iota

	// type of notation
	standartNotation = iota
	humanBinNotation = iota
	humanDecNotation = iota

	// reading source
	readFromFile = iota
	readFromArgs = iota
)

// config to handle all possible programm modes
type Config struct {
	whichFiles int
	displayFormat int
	notationType int
	singleFilePath []string
}

// parse flags and form config
func parseFlag() Config {
	// setup flags and parse them all
	hrFlag := pflag.BoolP("human-readable", "h", false, "prints in human readable format in power of 1024")
	HrFlag := pflag.BoolP("si", "H", false, "prints in human readable format in power of 1000")
	iFlag := pflag.BoolP("inodes", "i", false, "prints Inodes")
	allFlag := pflag.BoolP("all", "a", false, "prints all mounted fs")
	pflag.Parse()

	c := Config{}

	// fill the configuration
	if len(pflag.Args()) != 0 {
		c.whichFiles = oneFile
		c.singleFilePath = pflag.Args()	
	} else if *allFlag {
		c.whichFiles = allFiles
	} else {
		c.whichFiles = defaultFiles
	}

	if *hrFlag {
		c.notationType = humanBinNotation
	} else if *HrFlag {
		c.notationType = humanDecNotation 
	} else {
		c.notationType = standartNotation
	}

	if *iFlag {
		c.displayFormat = inodeFormat
	} else {
		c.displayFormat = sizeFormat
	}

	return c
}

func main() {
	config := parseFlag()

	printTable(config)
}
