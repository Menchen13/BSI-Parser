// main parser package
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	pPath := flag.String("input", "", "Path of file to be processed")
	flag.Parse()
	path := *pPath
	if path == "" {
		fmt.Fprintf(os.Stderr, "Error: empty path!\n")
		return
	}

	var title string = processHTML(path)

	fmt.Println(title)
	//assuming every title has "[CERT-Bund] Schwachstellenmanagement:"
	//this can be cut out! Index for cut is 38

	/*for index, value := range title {
		if value == ':' {
			fmt.Println(index)
		}
	}*/

	var trimmed_title string = title[38:]
	fmt.Println(trimmed_title)
	rename(path, "Test.html")
	time.Sleep(2 * time.Second)
}

func processHTML(path string) string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		panic(err)
	}

	title := doc.Find("title").Text()
	return title
}

func rename(path string, name string) {
	var a int
	for i := range path {
		if path[i] == '/' || path[i] == '\\' {
			a = i
		}
	}
	a++

	var result strings.Builder

	result.WriteString(path[:a])
	result.WriteString(name)

	err := os.Rename(path, result.String())
	if err != nil {
		panic(err)
	}

}
