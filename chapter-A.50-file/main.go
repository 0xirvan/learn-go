package main

import (
	"fmt"
	"os"
	"os/exec"
)

var path string

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func createFile() {
	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
		fmt.Println("File berhasil dibuat", path)
	}

}

func writeFile() {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	_, err = file.WriteString("Hello\n")
	if isError(err) {
		return
	}
	_, err = file.WriteString("World\n")
	if isError(err) {
		return
	}

	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("File berhasil ditulis")
}

func readFile() {
	data, err := os.ReadFile(path)
	if isError(err) {
		return
	}
	fmt.Println("=> File berhasil dibaca")
	fmt.Println(string(data))
}

func deleteFile() {
	var err = os.Remove(path)
	if isError(err) {
		return
	}
	fmt.Println("=> File berhasil dihapus")
}

func main() {
	output, _ := exec.Command("pwd").Output()

	path = string(output)
	path = path[:len(path)-1]
	path = path + "/file.txt"

	createFile()
	writeFile()
	readFile()
	deleteFile()
}
