package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getCategory(ext string) string {
	switch strings.ToLower(ext) {
	case ".jpg", ".png", ".jpeg":
		return "Images"
	case ".pdf", ".txt", ".docx":
		return "Documents"
	case ".mp3", ".wav":
		return "Audio"
	case ".mp4", ".mkv":
		return "Videos"
	default:
		return "Others"
	}
}

func organizeFiles(dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		name := file.Name()
		ext := filepath.Ext(name)
		category := getCategory(ext)

		folder := filepath.Join(dir, category)
		os.MkdirAll(folder, 0755)

		source := filepath.Join(dir, name)
		destination := filepath.Join(folder, name)

		if _, err := os.Stat(destination); err == nil {
			fmt.Println("Skipped:", name, "(already exists)")
			continue
		}

		if err := os.Rename(source, destination); err != nil {
			fmt.Println("Could not move:", name, err)
			continue
		}

		fmt.Println("Moved", name, "to", category)
	}

	fmt.Println("File organization completed!")
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- FILE ORGANIZER ---")
		fmt.Println("1. Organize files")
		fmt.Println("2. Exit")
		fmt.Print("Enter choice: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Print("Enter directory path: ")
			dir, _ := reader.ReadString('\n')
			dir = strings.TrimSpace(dir)

			info, err := os.Stat(dir)
			if err != nil || !info.IsDir() {
				fmt.Println("Invalid directory path.")
				continue
			}

			organizeFiles(dir)

		case "2":
			fmt.Println("Exiting File Organizer...")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}
