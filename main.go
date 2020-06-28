package main

import (
	"flag"
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
	file := flag.String("f", "", "output file.")
	nids := flag.String("n", "", "node id list.")
	wids := flag.String("w", "", "way id list.")
	rids := flag.String("r", "", "relation id list.")
	flag.Parse()

	if *file == "" {
		fmt.Printf("parameter error, missing -f\n")
		usage()
		os.Exit(-1)
	}
	if *nids == "" && *wids == "" && *rids == "" {
		fmt.Printf("parameter error, missing -n or -w or -r\n")
		usage()
		os.Exit(-1)
	}

	var files []string
	if *nids != "" {
		f := download(node, nids)
		files = append(files, f...)
	}

	if *wids != "" {
		f := download(way, wids)
		files = append(files, f...)
	}
	if *rids != "" {
		f := download(relation, rids)
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
	args = append(args, *file)

	fmt.Printf("generate %v\n", *file)
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

func download(e element, ids *string) []string {

	if *ids == "" {
		return nil
	}
	var (
		succes []string
		failed []string
	)
	defer func() {
		clear(failed)
	}()

	idArray := strings.Split(*ids, ",")
	for _, id := range idArray {
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

func usage() {
	fmt.Printf("-----------------------------\n")
	fmt.Printf("Usage\n")
	fmt.Printf("example: osmdownloader -f test.osm -n 111,222 -w 333,444 -r 555,666\n")
	fmt.Printf("-f\tthe output file, required\n")
	fmt.Printf("-n\tnode id list\n")
	fmt.Printf("-w\tway id list\n")
	fmt.Printf("-r\trelation id list\n")
}
