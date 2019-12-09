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
			extension := strings.Split(name, ".")
			exst := extension[len(extension)-1]
			fmt.Println(extension[len(extension)-1])

			images := []string{"jpg", "jpeg", "png", "bmp"}
			documents := []string{"pdf", "doc", "docx"}
			movies := []string{"avi", "mov"}

			if contains(images, exst) {
				fmt.Println("This is image file: " + name)
			} else if contains(documents, exst) {
				fmt.Println("This is document file: " + name)
			} else if contains(movies, exst) {
				fmt.Println("This is movie file: " + name)
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
