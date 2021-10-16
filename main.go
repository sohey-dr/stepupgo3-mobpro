package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"sync"
)

func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)
}

type UrlPath struct {
	Path     string   `json:"path"`
	Versions []string `json:"Versions"`
}

type Result struct {
	path   string
	result string
}

func handler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Path[1:]

	results, err := runCommand(q, w)
	if err != nil {
		log.Println(err)
	}

	for _, result := range results {
		fmt.Fprintln(w, result.path)
		fmt.Fprintln(w, result.result)
	}
}

func runCommand(path string, w http.ResponseWriter) ([]Result, error) {
	versions, err := runTipOne(path)
	if err != nil {
		log.Println(err)
	}

	var results []Result
	wg := sync.WaitGroup{}
	for _, version := range versions {
		wg.Add(1)

		v := version
		path := runTipTwo(path + "@" + v)
		go func() {
			defer wg.Done()
			output := runTipThree(path)
			result := Result{
				path:   path,
				result: output,
			}

			results = append(results, result)
		}()
	}

	wg.Wait()

	return results, nil
}

func runTipOne(path string) ([]string, error) {
	out, _ := exec.Command("go", "list", "-json", "-m", "-versions", path).Output()
	u := new(UrlPath)
	err := json.Unmarshal(out, u)
	if err != nil {
		return nil, err
	}

	return u.Versions, nil
}

func runTipTwo(path string) string {
	exec.Command("go", "get", path).Output()
	filePath, _ := exec.Command("go", "list", "-f", "\"{{.Dir}}\"", "-m", path).Output()
	fp := strings.Replace(strings.TrimSpace(string(filePath)), `"`, ``, -1)
	return string(fp)
}

func runTipThree(path string) string {
	cmd := exec.Command("go", "vet", "-json", path)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
	}
	return string(out)
}
