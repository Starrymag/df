package main

import (
	// "fmt"
	"flag"
)

func main() {
	// setup flags and parse them all
	hrFlag := flag.Bool("h", false, "prints in human readable format")
	allFlag := flag.Bool("a", false, "prints all mounted fs")
	flag.Parse()

	// get list of all mounted fs
	allData, _, err := mounts()
	if err != nil {
		panic(err)
	}
	// if len(warnings) != 0 { 
	// 	fmt.Println(warnings)
	// }

	data, _ := getDefaultData(allData)

	// process flags
	if *allFlag {
		data = allData
	} 
	if *hrFlag {
		data, _ = getHumanData(data)
	}

	printRes(data)
}
