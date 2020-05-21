package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type element string

const (
	api = "https://www.openstreetmap.org/api/0.6"

	node     element = "node"
	way      element = "way"
	relation element = "relation"
)

func main() {

	argCount := len(os.Args)
	if argCount != 5 && argCount != 7 && argCount != 9 {
		fmt.Println("error: invalid args")
		fmt.Println(`usage: osmdownloader -f "test.osm" -n 111,222 -w 333,444 -r 555,666`)
		os.Exit(-1)
	}
	var (
		file        string
		nodeids     string
		wayids      string
		relationids string
	)
	for i, arg := range os.Args {
		if arg == "-f" || arg == "-F" {
			file = os.Args[i+1]
		}
		if arg == "-n" || arg == "-N" {
			nodeids = os.Args[i+1]
		}
		if arg == "-w" || arg == "-W" {
			wayids = os.Args[i+1]
		}
		if arg == "-r" || arg == "-R" {
			relationids = os.Args[i+1]
		}
	}

	if file == "" {
		fmt.Printf("need to specify a file for saving map data\n")
		os.Exit(-1)
	}
	if nodeids == "" && wayids == "" && relationids == "" {
		fmt.Printf("need to specify the element ids\n")
		os.Exit(-1)
	}

	var files []string
	if nodeids != "" {
		f := download(node, nodeids)
		files = append(files, f...)
	}

	if wayids != "" {
		f := download(way, wayids)
		files = append(files, f...)
	}
	if relationids != "" {
		f := download(relation, relationids)
		files = append(files, f...)
	}

	fmt.Printf("---------------------------------------------\n")
	if len(files) == 0 {
		fmt.Printf("nothing was downloaded\n")
		os.Exit(-1)
	}

	var args []string
	for _, f := range files {
		args = append(args, "--read-xml")
		args = append(args, f)
	}
	for i := 0; i < len(files)-1; i++ {
		args = append(args, "--merge")
	}

	args = append(args, "--write-xml")
	args = append(args, file)

	fmt.Printf("generate %v\n", file)
	cmd := exec.Command("osmosis", args...)
	if err := cmd.Run(); err != nil {
		fmt.Printf("failed to merge *.osm, err: %v\n", err)
		clear(files)
		os.Exit(-1)
	}

	clear(files)
	fmt.Printf("done in success\n")
	os.Exit(0)
}

func download(e element, idlist string) []string {

	if idlist == "" {
		return nil
	}
	var (
		succes []string
		failed []string
	)
	defer func() {
		clear(failed)
	}()

	ids := strings.Split(idlist, ",")
	for _, id := range ids {
		url := fmt.Sprintf("%v/%v/%v", api, e, id)
		if e == way || e == relation {
			url += "/full"
		}

		fmt.Printf("%-15s%-15s", e, id)
		f := fmt.Sprintf("%v%v.osm", e, id)
		cmd := exec.Command("wget", url, "-O", f)
		if err := cmd.Run(); err != nil {
			fmt.Printf("%-15s\n", "[not found]")
			failed = append(failed, f)
			continue
		}

		if isEmptyFile(f) {
			fmt.Printf("%-15s\n", "[not found]")
			failed = append(failed, f)
			continue
		}
		fmt.Printf("%-15s\n", "[ok]")
		succes = append(succes, f)
	}
	return succes
}

func isEmptyFile(file string) bool {
	fi, err := os.Stat(file)
	if err != nil {
		return false
	}
	return fi.Size() == 0
}

func clear(files []string) {
	for _, file := range files {
		if _, err := os.Stat(file); err != nil {
			fmt.Printf("failed to remove %v, err: %v\n", files, err)
			continue
		}
		if err := os.Remove(file); err != nil {
			fmt.Printf("failed to remove %v, err: %v\n", file, err)
			continue
		}
	}
}
