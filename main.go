package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/fatih/color"
)

// Header determines the styles for response header keys.
var Header = color.New(color.Bold, color.BgYellow, color.FgBlack).SprintFunc()

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
		return
	}

	s := fmt.Sprintf("%s:\n  %s\n\n", Header(key), val)
	fmt.Fprint(color.Output, s)
}

func printJSON(hs headers) {
	data, err := json.Marshal(hs)
	if err != nil {
		log.Fatalf("whoops, unable to convert data into JSON: %s\n", err)
	}
	fmt.Println(string(data))
}

func main() {
	fmt.Printf("\n")

	if len(os.Args) < 2 {
		fmt.Println("Please provide a URL (e.g.  https://www.fastly.com/)")
		os.Exit(1)
	}

	help := flag.Bool("help", false, "show available flags")
	filter := flag.String("filter", "", "comma-separated list of headers to be displayed\n\te.g. cache,vary")
	plain := flag.Bool("plain", false, "output is formatted without any extraneous spacing or ANSI colour codes")
	jsonv := flag.Bool("json", false, "output is formatted into JSON for easy parsing")
	flag.Parse()

	if *help == true {
		flag.PrintDefaults()
		os.Exit(1)
	}

	url := os.Args[len(os.Args)-1]

	client := &http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	req.Header.Set("Fastly-Debug", "true")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	hs := headers{}

	for k, v := range resp.Header {
		hs = append(hs, header{k, v})
	}

	sort.Sort(hs)

	if *filter != "" {
		filters := strings.Split(*filter, ",")

		jh := headers{}

		for _, h := range hs {
			for _, v := range filters {
				matched, _ := regexp.MatchString("(?i)"+v, h.Key)
				if matched {
					if *jsonv {
						jh = append(jh, h)
					} else {
						print(h.Key, h.Val, *plain)
					}
				}
			}
		}

		if *jsonv {
			printJSON(jh)
			return
		}

		return
	}

	if *jsonv {
		printJSON(hs)
		return
	}
	for _, header := range hs {
		print(header.Key, header.Val, *plain)
	}

	fmt.Printf("Status Code: %s\n\n", resp.Status)
}
