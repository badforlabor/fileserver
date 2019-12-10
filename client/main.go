/**
 * Auth :   liubo
 * Date :   2019/12/9 17:06
 * Comment:
 */

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fileUrl := "http://127.0.0.1:7088/download/avatar.jpg"

	if err := DownloadFile("avatar.jpg", fileUrl); err != nil {
		panic(err)
	}
	fmt.Println("succ")
}
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
