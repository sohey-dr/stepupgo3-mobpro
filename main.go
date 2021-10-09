package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
)

func main() {
	//http.HandleFunc("/", handler)
	//
	//http.ListenAndServe(":8080", nil)

	q := "github.com/gostaticanalysis/skeleton/v2"
	runCommand(q)
}

type UrlPath struct {
	Path     string   `json:"path"`
	Versions []string `json:"Versions"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Path[1:]

	out, _ := exec.Command("go", "list", "-json", "-m", "-versions", q).Output()
	fmt.Printf("ls result: \n%s", string(out))

	fmt.Fprint(w, q)
}

func runCommand(path string) {
	out, _ := exec.Command("go", "list", "-json", "-m", "-versions", path).Output()
	//var u UrlPath`
	u := new(UrlPath)
	err := json.Unmarshal(out, u)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u)
}
