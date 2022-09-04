package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

var (
	fileName          string
	fullURLFile       string
	fileNameForDelete string
)

func main() {
	fileNameForDelete := DownloadFile()
// 	fmt.Println(fileNameForDelete)

	defer deleteFile(fileNameForDelete)

}

func labSetup() {
	fmt.Println("inside the lab setup block")
	cmdStr := "docker-compose up "
	out, _ := exec.Command("/bin/sh", "-c", cmdStr).Output()
	fmt.Printf("%s", out)

}

func deleteFile(fileNameForDelete string) {
	removeFile := os.Remove(fileNameForDelete)
	if removeFile != nil {
		log.Fatal(removeFile)
	}
	fmt.Printf("%s", "file removed")

}

func DownloadFile() (fileName string) {
// 	place where the url can be changed for packging the binary
	fullURLFile = "https://raw.githubusercontent.com/Srivignesh31/docker-compose/main/docker-compose.yaml"
	//fullURLFile = "https://raw.githubusercontent.com/Srivignesh31/docker-compose/main/docker.yaml"
	// Build fileName from fullPath
	fileURL, err := url.Parse(fullURLFile)
	if err != nil {
		log.Fatal(err)
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName = segments[len(segments)-1]

	// Create blank file
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	// Put content on file
	resp, err := client.Get(fullURLFile)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	defer file.Close()

	fmt.Printf("Downloaded a file %s with size %d", fileName, size)

	labSetup()

	return fileName

}
