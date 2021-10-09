package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)
}

type UrlPath struct {
	Path     string   `json:"path"`
	Versions []string `json:"Versions"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Path[1:]

	//runCommand(q)
	fmt.Fprint(w, q)
}

func runCommand(path string) {
	versions := runTipOne(path)

	runTipTwoAndThree(path + "@" + versions[0])
}

func runTipOne(path string) []string {
	out, _ := exec.Command("go", "list", "-json", "-m", "-versions", path).Output()
	u := new(UrlPath)
	err := json.Unmarshal(out, u)
	if err != nil {
		fmt.Println(err)
	}

	return u.Versions
}

func runTipTwoAndThree(path string) {
	exec.Command("go", "get", path)
	filePath, _ := exec.Command("go", "list", "-f", "\"{{.Dir}}\"", "-m", path).Output()
	result, _ := exec.Command("go", "vet", "-json", string(filePath)).Output()
	fmt.Println(string(filePath))
	fmt.Println(result)
}
