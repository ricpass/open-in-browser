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
			errMsg := fmt.Sprintf("Error: %s", errDecode)
			fmt.Println(errMsg)
			fmt.Fprintln(w, errMsg)
			return
		}

		url := string(b)
		fmt.Printf("url=%s\n", url)

		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			errMsg := fmt.Sprintf("INVALID_URL: %s", url)
			fmt.Println(errMsg)
			fmt.Fprintln(w, errMsg)
			return
		}

		errOpen := exec.Command("open", url).Run()

		if errOpen != nil {
			errMsg := fmt.Sprintf("Error: %s", errOpen)
			fmt.Println(errMsg)
			fmt.Fprintln(w, errMsg)
			return
		}

		fmt.Fprintf(w, "OK")
	})

	http.ListenAndServe(":8080", nil)

}
