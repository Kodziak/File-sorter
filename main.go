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
	dirname := "./Downloads"

	images := []string{"jpg", "jpeg", "png", "bmp"}
	documents := []string{"pdf", "doc", "docx", "txt", "ppt", "rtf", "xlsx", "xls"}
	movies := []string{"avi", "mov"}
	archive := []string{"rar", "zip", "7z"}

	log.Info("Create new cron")
	c := cron.New()
	c.AddFunc("@every 1m", func() {
		log.Info("[Job 1]Every minute job\n")

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
			extension := strings.ToLower(extensionArray[len(extensionArray)-1])

			if contains(images, extension) {
				printMoveFile(name)
				moveFile(dirname, name, "Images")
			} else if contains(documents, extension) {
				printMoveFile(name)
				moveFile(dirname, name, "Documents")
			} else if contains(movies, extension) {
				printMoveFile(name)
				moveFile(dirname, name, "Movies")
			} else if contains(archive, extension) {
				printMoveFile(name)
				moveFile(dirname, name, "Archives")
			}
		}
	})

	// Start cron with one scheduled job
	log.Info("Start cron")
	c.Start()
	printCronEntries(c.Entries())
	time.Sleep(60 * 24 * time.Minute)
}

func printCronEntries(cronEntries []cron.Entry) {
	log.Infof("Cron Info: %+v\n", cronEntries)
}

func printMoveFile(file string) {
	log.Infof("Moving file: " + file)
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func moveFile(baseDirectory string, fileName string, targetDirectory string) {
	if _, err := os.Stat(baseDirectory + targetDirectory); os.IsNotExist(err) {
		os.Mkdir(baseDirectory+"/"+targetDirectory+"/", 0777)
	}

	error := os.Rename(baseDirectory+"/"+fileName, baseDirectory+"/"+targetDirectory+"/"+fileName)
	if error != nil {
		fmt.Println("Error rename", error)
	}
}
