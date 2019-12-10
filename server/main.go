/**
 * Auth :   liubo
 * Date :   2019/12/9 17:06
 * Comment:
 */

package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main1() {
	http.Handle("/download/", http.FileServer(Dir(".")))
	err := http.ListenAndServe(":7088", nil)
	fmt.Println(err.Error())
}
func main() {
	//http.Handle("/download/", http.FileServer(Dir(".")))
	//err := http.ListenAndServe(":7088", nil)

	handler := http.NewServeMux()
	handler.Handle("/download/", http.FileServer(Dir(".")))

	server := &http.Server{Addr: ":7088", Handler: handler}
	err := server.ListenAndServe()

	fmt.Println(err.Error())
}

type Dir string
func (d Dir) Open(name string) (http.File, error) {
	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) {
		return nil, errors.New("http: invalid character in file path")
	}
	dir := string(d)
	if dir == "" {
		dir = "."
	}
	fullName := "./avatar.jpg"// filepath.Join(dir, filepath.FromSlash(path.Clean("/"+name)))
	f, err := os.Open(fullName)
	if err != nil {
		return nil, err
	}
	return f, nil
}

