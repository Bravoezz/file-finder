package main

import (
	"flag"
	"fmt"
	"github.com/manifoldco/promptui"
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

	if *fileName == "nil" {
		flag.Usage()
		fmt.Println("**Exit program**")
		os.Exit(0)
	}

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

	foundFiles := []string{}
	for str := range pathFoundFile {
		name := filepath.Base(str)
		foundFiles = append(foundFiles, fmt.Sprintf("%s path -> %s", name, str))
	}
	timeDuration := time.Since(startDate)

	if len(foundFiles) == 0 {
		fmt.Println("Not found file")
		fmt.Println("**Exit program**")
		os.Exit(0)
	}

	////test nomas
	//Map(foundFiles, func(srt string) string {
	//	return strings.Split(srt, "::")[0]
	//})

	selectUi := promptui.Select{
		Label: "Select one file for open with file explorer",
		Items: foundFiles,
	}

	_, selected, err := selectUi.Run()
	if err != nil {
		fmt.Println("Invalid selection")
	}

	fileIndex := slices.Index(foundFiles, selected)

	fmt.Println("Opening file...")
	exec.Command("explorer", strings.Split(foundFiles[fileIndex], " path -> ")[1]).CombinedOutput()

	fmt.Printf("Exec time: %v\n", timeDuration)
	fmt.Println("**Exit program**")
}

func searchFile(path, findFile string, c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	files, _ := os.ReadDir(path)

	for _, file := range files {
		blinkFiles := []string{"node_modules", "dist", "build", ".git", ".idea"}
		if slices.Contains(blinkFiles, file.Name()) {
			continue
		}

		if strings.Contains(strings.ToLower(file.Name()), strings.ToLower(findFile)) {
			c <- filepath.Join(path, file.Name())
		}

		if file.IsDir() {
			wg.Add(1)
			go searchFile(filepath.Join(path, file.Name()), findFile, c, wg)
		}

	}

}
