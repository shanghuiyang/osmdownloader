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

	argCount := len(os.Args)
	if argCount != 5 && argCount != 7 {
		fmt.Println("error: invalid args")
		fmt.Println(`usage: osmdownloader -f "test.osm" -n 111,222 -w 333,444`)
		return
	}
	var file string
	var nodeIDs string
	var wayIDs string
	for i, arg := range os.Args {
		if arg == "-f" || arg == "-F" {
			file = os.Args[i+1]
		}
		if arg == "-n" || arg == "-N" {
			nodeIDs = os.Args[i+1]
		}
		if arg == "-w" || arg == "-W" {
			wayIDs = os.Args[i+1]
		}
	}

	if file == "" {
		fmt.Printf("need to specify a file for saving map data\n")
		return
	}
	if nodeIDs == "" && wayIDs == "" {
		fmt.Printf("need to specify the ids for node or way\n")
		return
	}

	fmt.Printf("file: %v\n", file)
	arrayNodeIDs := []string{}
	arrayWayIDs := []string{}
	if nodeIDs != "" {
		fmt.Printf("nodes: %v\n", nodeIDs)
		arrayNodeIDs = strings.Split(nodeIDs, ",")
	}
	if wayIDs != "" {
		fmt.Printf("ways: %v\n", wayIDs)
		arrayWayIDs = strings.Split(wayIDs, ",")
	}
	fmt.Printf("--------\n")

	files := map[string]string{}
	var error int
	for _, id := range arrayNodeIDs {
		url := fmt.Sprintf("%v/node/%v", api, id)
		fmt.Printf("node %v\t", id)
		f := fmt.Sprintf("n%v.osm", id)
		cmd := exec.Command("wget", url, "-O", f)
		if err := cmd.Run(); err != nil {
			fmt.Printf("not found\n")
			error++
			continue
		}
		if isEmptyFile(f) {
			fmt.Printf("not found\n")
			error++
			continue
		}
		fmt.Printf("ok\n")
		files[f] = id
	}

	for _, id := range arrayWayIDs {
		url := fmt.Sprintf("%v/way/%v/full", api, id)
		fmt.Printf("way %v\t", id)
		f := fmt.Sprintf("w%v.osm", id)
		cmd := exec.Command("wget", url, "-O", f)
		if err := cmd.Run(); err != nil {
			fmt.Printf("not found\n")
			error++
			continue
		}
		if isEmptyFile(f) {
			fmt.Printf("not found\n")
			error++
			continue
		}
		fmt.Printf("ok\n")
		files[f] = id
	}

	var args []string
	for f := range files {
		args = append(args, "--read-xml")
		args = append(args, f)
	}
	for i := 0; i < len(files)-1; i++ {
		args = append(args, "--merge")
	}

	if len(files) > 0 {
		args = append(args, "--write-xml")
		args = append(args, file)

		fmt.Printf("generate %v\n", file)
		cmd := exec.Command("osmosis", args...)
		if err := cmd.Run(); err != nil {
			fmt.Printf("failed to merge *.osm\n")
			error++
		}
	}

	error += clear("n", arrayNodeIDs)
	error += clear("w", arrayWayIDs)

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

func clear(elementType string, ids []string) int {
	error := 0
	for _, id := range ids {
		file := fmt.Sprintf("%v%v.osm", elementType, id)
		if _, err := os.Stat(file); err == nil {
			if e := os.Remove(file); e != nil {
				fmt.Printf("failed to delete %v, err: %v\n", file, e)
				error++
			}
		}
	}
	return error
}
