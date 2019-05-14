package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	api = "https://www.openstreetmap.org/api/0.6"
)

func main() {

	if len(os.Args) != 5 {
		fmt.Println("error: invalid args")
		fmt.Println("usage: geojson2poly -i infile.json -o outfile.poly")
		return
	}
	var elementType string
	var ids string
	for i, arg := range os.Args {
		if arg == "-t" || arg == "-T" {
			elementType = os.Args[i+1]
		}
		if arg == "-i" || arg == "-I" {
			ids = os.Args[i+1]
		}
	}

	fmt.Printf("type: %v\n", elementType)
	fmt.Printf("ids: %v\n", ids)

	files := map[string]string{}
	var error int
	arrayIDs := strings.Split(ids, ",")
	for _, id := range arrayIDs {
		var url string
		if elementType == "node" {
			url = fmt.Sprintf("%v/node/%v", api, id)
		} else if elementType == "way" {
			url = fmt.Sprintf("%v/way/%v/full", api, id)
		} else {
			fmt.Printf("invalid element type\n")
			return
		}

		fmt.Printf("%v\t", id)
		file := fmt.Sprintf("%v.osm", id)
		cmd := exec.Command("wget", url, "-O", file)
		if err := cmd.Run(); err != nil {
			fmt.Printf("not found\n")
			error++
			continue
		}
		if isEmptyFile(file) {
			fmt.Printf("not found\n")
			error++
			continue
		}
		fmt.Printf("ok\n")
		files[file] = id
	}

	var args []string
	for file := range files {
		args = append(args, "--read-xml")
		args = append(args, file)
	}
	for i := 0; i < len(files)-1; i++ {
		args = append(args, "--merge")
	}

	if len(files) > 0 {
		args = append(args, "--write-xml")
		args = append(args, "merge.osm")

		fmt.Printf("merge...\n")
		cmd := exec.Command("osmosis", args...)
		if err := cmd.Run(); err != nil {
			fmt.Printf("failed to merge *.osm\n")
			error++
		}
	}

	error += clear(arrayIDs)

	if error == 0 {
		fmt.Printf("done in success\n")
	} else {
		fmt.Printf("done with %v errors\n", error)
	}
	return
}

func isEmptyFile(file string) bool {
	fi, err := os.Stat(file)
	if err != nil {
		return false
	}
	return fi.Size() == 0
}

func clear(ids []string) int {
	error := 0
	for _, id := range ids {
		file := fmt.Sprintf("%v.osm", id)
		if _, err := os.Stat(file); err == nil {
			if e := os.Remove(file); e != nil {
				fmt.Printf("failed to delete %v, err: %v\n", file, e)
				error++
			}
		}
	}
	return error
}
