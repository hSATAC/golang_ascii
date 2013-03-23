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

func readAscii() []string {
	content, err := ioutil.ReadFile("golang.txt")
	if err != nil {
		fmt.Println("Error: %s\n", err)
	}

	lines := strings.Split(string(content), "\n")

	return lines
}

func welcome(w http.ResponseWriter, r *http.Request) {

	lines := readAscii()

	if !strings.Contains(r.UserAgent(), "curl") {
		fmt.Fprintln(w, "Try `curl ascii.golang.tw`")
		for _, line := range lines {
			fmt.Fprintln(w, line)
		}
		return
	}
	printColorBuffer(w, "@{wB}WELCOME TO GOLANG.TW")
	tick()
	printColorBuffer(w, "@{r}now loading page......")
	tock()


	for _, line := range lines {
		tick()
		printColorBuffer(w, line)
    }
}

func printColorBuffer(w http.ResponseWriter, s string) {
	color.Fprintln(w, s)
	w.(http.Flusher).Flush()
}

func tick() {
	time.Sleep(200 * time.Millisecond)
}

func tock() {
	time.Sleep(2 * time.Second)
}

func main() {
	http.HandleFunc("/", welcome)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
