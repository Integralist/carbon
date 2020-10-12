package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
)

type header struct {
	Key string
	Val []string
}

type headers []header

// Satisfy the Sort interface
func (v headers) Len() int      { return len(v) }
func (v headers) Swap(i, j int) { v[i], v[j] = v[j], v[i] }
func (v headers) Less(i, j int) bool {
	return v[i].Key < v[j].Key
}

func main() {
	fmt.Printf("\n")

	if len(os.Args) < 2 {
		fmt.Println("Please provide a URL (e.g.  https://www.buzzfeed.com/?country=us)")
		os.Exit(1)
	}

	help := flag.Bool("help", false, "show available flags")
	filter := flag.String("filter", "", "comma-separated list of headers to be displayed\n\te.g. X-Siterouter-Upstream,X-Cache")
	flag.Parse()

	if *help == true {
		flag.PrintDefaults()
		os.Exit(1)
	}

	var url string
	if *filter != "" {
		url = os.Args[3]
	} else {
		url = os.Args[1]
	}

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	hs := headers{}

	for k, v := range response.Header {
		hs = append(hs, header{k, v})
	}

	sort.Sort(hs)

	if *filter != "" {
		filters := strings.Split(*filter, ",")
		for _, header := range hs {
			for _, v := range filters {
				matched, _ := regexp.MatchString("(?i)"+v, header.Key)
				if matched {
					fmt.Printf("%s:\n  %s\n\n", header.Key, header.Val)
				}
			}
		}
	} else {
		for _, header := range hs {
			fmt.Printf("%s:\n  %s\n\n", header.Key, header.Val)
		}
	}

	fmt.Printf("Status Code: %s\n\n", response.Status)
}
