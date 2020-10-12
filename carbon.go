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

func print(key string, val []string, plain bool) {
	if plain == true {
		fmt.Printf("%s: %s\n", key, strings.Join(val, ", "))
	} else {
		fmt.Printf("%s:\n  %s\n\n", key, val)
	}
}

func main() {
	fmt.Printf("\n")

	if len(os.Args) < 2 {
		fmt.Println("Please provide a URL (e.g.  https://www.buzzfeed.com/?country=us)")
		os.Exit(1)
	}

	help := flag.Bool("help", false, "show available flags")
	filter := flag.String("filter", "", "comma-separated list of headers to be displayed\n\te.g. X-Siterouter-Upstream,X-Cache")
	plain := flag.Bool("plain", false, "output is formatted for easy piping")
	flag.Parse()

	if *help == true {
		flag.PrintDefaults()
		os.Exit(1)
	}

	var url string
	switch len(os.Args) {
	case 5:
		url = os.Args[4]
	case 4:
		url = os.Args[3]
	default:
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
					print(header.Key, header.Val, *plain)
				}
			}
		}
	} else {
		for _, header := range hs {
			print(header.Key, header.Val, *plain)
		}
	}

	fmt.Printf("Status Code: %s\n\n", response.Status)
}
