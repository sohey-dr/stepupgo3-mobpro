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

	if !strings.Contains(q, ".com") {
		fmt.Fprintln(w, "正しいモジュールのパスではありません")
		return
	}

	results, err := run(q)
	if err != nil {
		log.Println(err)
	}

	for _, result := range results {
		fmt.Fprintln(w, result.path)
		fmt.Fprintln(w, result.result)
	}
}

func run(path string) ([]Result, error) {
	versions, err := tipOne(path)
	if err != nil {
		log.Println(err)
	}

	var results []Result
	wg := sync.WaitGroup{}
	for _, version := range versions {
		wg.Add(1)

		v := version
		path, err := tipTwo(path + "@" + v)
		if err != nil {
			log.Println(err)
			continue
		}

		go func() {
			defer wg.Done()
			output, err := tipThree(path)
			if err != nil {
				log.Println(err)
			}
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

// 指定したモジュールのバージョンを取得して、スライスにして返す
func tipOne(path string) ([]string, error) {
	out, _ := exec.Command("go", "list", "-json", "-m", "-versions", path).Output()
	u := new(UrlPath)
	err := json.Unmarshal(out, u)
	if err != nil {
		return nil, err
	}

	return u.Versions, nil
}

// モジュールをgo getして保存先を取得してそのパスを返す
func tipTwo(path string) (string, error) {
	exec.Command("go", "get", path).Output()
	filePath, err := exec.Command("go", "list", "-f", "\"{{.Dir}}\"", "-m", path).Output()
	if err != nil {
		return "", err
	}

	return strings.Replace(strings.TrimSpace(string(filePath)), `"`, ``, -1), nil
}

//　指定したモジュールをge vetする
func tipThree(path string) (string, error) {
	cmd := exec.Command("go", "vet", "-json", path)
	out, err := cmd.CombinedOutput()
	if err != nil {
		// 失敗していなくても出ることがあるためreturnしない
		log.Println(err)
	}

	return string(out), err
}
