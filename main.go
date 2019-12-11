package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
}

func main() {
	images := []string{"jpg", "jpeg", "png", "bmp"}
	documents := []string{"pdf", "doc", "docx"}
	movies := []string{"avi", "mov"}

	log.Info("Create new cron")
	c := cron.New()
	c.AddFunc("@every 1m", func() {
		log.Info("[Job 1]Every minute job\n")

		// error := os.Rename("hello.txt", "testdir/hello.txt")
		// if error != nil {
		// 	fmt.Println("Error rename", error)
		// }

		dirname := "../../../../../Users/kodziak/Downloads"

		f, err := os.Open(dirname)
		if err != nil {
			log.Fatal(err)
		}
		files, err := f.Readdir(-1)
		f.Close()
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			name := file.Name()
			extensionArray := strings.Split(name, ".")
			extension := extensionArray[len(extensionArray)-1]

			if contains(images, extension) {
				fmt.Println("This is image file: " + name)
				moveFile(name, "../../../../../Users/kodziak/Downloads/Images/")
			} else if contains(documents, extension) {
				fmt.Println("This is document file: " + name)
				moveFile(name, "../../../../../Users/kodziak/Downloads/Documents/")

			} else if contains(movies, extension) {
				fmt.Println("This is movie file: " + name)
				moveFile(name, "../../../../../Users/kodziak/Downloads/Movies/")
			}
		}
	})

	// Start cron with one scheduled job
	log.Info("Start cron")
	c.Start()
	printCronEntries(c.Entries())
	time.Sleep(20 * time.Minute)
}

func printCronEntries(cronEntries []cron.Entry) {
	log.Infof("Cron Info: %+v\n", cronEntries)
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func moveFile(fileName string, directory string) {
	error := os.Rename(fileName, directory+fileName)
	if error != nil {
		fmt.Println("Error rename", error)
	}
}
