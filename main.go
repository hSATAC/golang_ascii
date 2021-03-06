package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"strings"
	"io/ioutil"
	"github.com/wsxiaoys/terminal/color"
)

func readAsciiLines() []string {
	content, err := ioutil.ReadFile("golang.txt")
	if err != nil {
		fmt.Println("Error: %s\n", err)
	}

	lines := strings.Split(string(content), "\n")

	return lines
}

func welcome(w http.ResponseWriter, r *http.Request) {

	if !strings.Contains(r.UserAgent(), "curl") {
		html, _ := ioutil.ReadFile("golang.html")
		fmt.Fprintf(w, string(html))
		return
	}

	lines := readAsciiLines()

	colorBufferPrintln(w, "\x1b[2J\x1b[1;1H")
	colorBufferPrintln(w, "\x1b[1F@{wB}WELCOME TO GOLANG.TW")
	tick()

	// show loading
	loading_symbols := [...]string{"\\","-","|","/"}
	colorBufferPrintln(w, "@{r}now loading page......")
	tick()

	for i := 0; i < 10; i++ {
		index := i % len(loading_symbols)
		str := fmt.Sprintf("%s%s", "\x1b[1F@{r}now loading page......", loading_symbols[index])
		colorBufferPrintln(w, str);
		tick()
	}


	for _, line := range lines {
		tick()
		colorBufferPrintln(w, line)
    }
}

func colorBufferPrintln(w http.ResponseWriter, s string) {
	color.Fprintln(w, s)
	w.(http.Flusher).Flush()
}

func tick() {
	time.Sleep(200 * time.Millisecond)
}

func main() {
	http.HandleFunc("/", welcome)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
