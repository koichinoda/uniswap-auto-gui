package services

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"

	gosxnotifier "github.com/deckarep/gosx-notifier"
	"github.com/uniswap-auto-gui/data"
)

func Notify(title string, message string, link string) {
	logo := canvas.NewImageFromResource(data.FyneScene)
	note := gosxnotifier.NewNotification(message)
	note.Title = title
	note.Sound = gosxnotifier.Default
	note.Link = link
	note.AppIcon = logo.Resource.Name()
	note.Push()
}

func Alert(title string, message string) {
	fyne.CurrentApp().SendNotification(&fyne.Notification{
		Title:   title,
		Content: message,
	})
}

func ReadPairs() []string {
	path := absolutePath() + "/pairs.txt"
	lines, err := readLines(path)
	if err != nil {
		lines = []string{"0x9d9681d71142049594020bd863d34d9f48d9df58", "0x7a99822968410431edd1ee75dab78866e31caf39"}
	}
	return lines
}

func IsExist(pair string) bool {
	pairs := ReadPairs()
	for _, v := range pairs {
		if v == pair {
			return true
		}
	}
	return false
}

func WritePairs(lines []string) error {
	path := absolutePath() + "/pairs.txt"
	err := writeLines(lines, path)
	return err
}

func absolutePath() string {
	ex, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func writeOnePair(pair string) error {
	path := absolutePath() + "/pairs.txt"
	pairs, _ := readLines(path)
	pairs = append(pairs, pair)
	err := writeLines(pairs, path)
	return err
}
