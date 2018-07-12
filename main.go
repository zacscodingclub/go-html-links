package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/zacscodingclub/go-html-links/link"
)

func main() {
	parseExamples()
}

func parseExamples() {
	files := []string{"one", "two", "three", "four"}

	for i := 0; i < len(files); i++ {
		fn := fmt.Sprintf("./pages/%s.html", files[i])
		rf, err := ioutil.ReadFile(fn)
		if err != nil {
			panic(err)
		}

		r := strings.NewReader(string(rf))
		links, _ := link.Parse(r)
		for _, link := range links {
			fmt.Println(link.Text)
		}
	}
}
