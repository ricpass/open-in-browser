package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[1:]

		fmt.Printf("path=%s\n", path)
		b, errDecode := base64.URLEncoding.DecodeString(path)

		if errDecode != nil {
			fmt.Fprintf(w, "Error: %s", errDecode)
			return
		}

		url := string(b)
		fmt.Printf("url=%s\n", url)

		if !strings.HasPrefix(url, "http://") && strings.HasPrefix(url, "https://") {
			fmt.Fprintf(w, "INVALID_URL: %s", url)
			return
		}

		errOpen := exec.Command("xdg-open", url).Run()

		if errOpen != nil {
			fmt.Fprintf(w, "Error: %s", errOpen)
			return
		}

		fmt.Fprintf(w, "OK")
	})

	http.ListenAndServe(":8080", nil)

}
