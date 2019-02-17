package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/policypalnet/go-test/service/namesvc"
)

//go:generate go run api.go
func main() {
	fmt.Println("hello world")
	err := filepath.Walk("./service", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		// if info.IsDir() && info.Name() == subDirToSkip {
		//         fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
		//         return filepath.SkipDir
		// }
		if info.IsDir() {
			fmt.Printf("visited dir: %q\n", path)
		} else {
			fmt.Println(info.Name())
			fmt.Printf("visited file: %q\n", path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(namesvc.Name())
}
