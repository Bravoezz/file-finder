package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	"sync"
	"time"
)

func main() {
	startDate := time.Now()

	fileName := flag.String("name", "nil", "name of searching file")
	flag.Parse()
	fmt.Println("File name to search", *fileName)
	fmt.Println("Searching...")

	pathFoundFile := make(chan string)
	wg := new(sync.WaitGroup)
	rootPath, _ := os.Getwd()

	wg.Add(1)
	go searchFile(rootPath, *fileName, pathFoundFile, wg)

	go func() {
		wg.Wait()
		close(pathFoundFile)
	}()

	foundFile, open := <-pathFoundFile
	timeDuration := time.Since(startDate)
	if !open {
		fmt.Println("Not found file")
	} else {
		fmt.Println("Found file:", foundFile)

		var openToEx string
		fmt.Print("Open to file explorer (si/no) ->")
		fmt.Scanf("%s", &openToEx)

		if strings.ToLower(openToEx) == "si" {
			fmt.Println("Opening file...")
			cmd := exec.Command("explorer", foundFile)
			cmd.CombinedOutput()
		}
	}

	fmt.Println("**Exit program**")
	fmt.Printf("Exec time: %v\n", timeDuration)
}

func searchFile(path, findFile string, c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	files, _ := os.ReadDir(path)

	for _, file := range files {
		blinkFiles := []string{"node_modules", "dist", "build", ".git", ".idea"}
		if slices.Contains(blinkFiles, file.Name()) {
			continue
		}

		if strings.Contains(file.Name(), findFile) {
			c <- filepath.Join(path, file.Name())
		}

		if file.IsDir() {
			wg.Add(1)
			go searchFile(filepath.Join(path, file.Name()), findFile, c, wg)
		}

	}

}
