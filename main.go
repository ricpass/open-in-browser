package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[1:]

		if !strings.HasPrefix(path, "http://") && strings.HasPrefix(path, "https://") {
			fmt.Fprintf(w, "INVALID_URL: %s", path)
			return
		}

		err := exec.Command("xdg-open", path).Run()

		if err != nil {
			fmt.Fprintf(w, "Error: %s", err)
		} else {
			fmt.Fprintf(w, "OK")
		}
	})

	http.ListenAndServe(":8080", nil)

}
