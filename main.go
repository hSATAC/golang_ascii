package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"strings"
	"io/ioutil"
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
	fmt.Fprintln(w, "\x1b[4m\x1b[5m\x1b[44mWELCOME TO GOLANG.TW\x1b[0m")
	fmt.Fprintln(w, "\x1b[5m\x1b[32mnow loading page......\x1b[0m\n")
	w.(http.Flusher).Flush()
	time.Sleep(2*time.Second)


	for _, line := range lines {
		time.Sleep(200 * time.Millisecond)
		fmt.Fprintln(w, line)
		w.(http.Flusher).Flush()
    }
}
func main() {
	http.HandleFunc("/", welcome)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
