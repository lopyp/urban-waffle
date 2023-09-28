package main

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: wget <url>")
		return
	}

	url := os.Args[1]
	err := download(url)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func download(url string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP Error: %s", response.Status)
	}

	// Извлекаем имя файла из URL
	fileName := getFileNameFromURL(url)

	// Создаем файл для сохранения данных
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Копируем данные из HTTP-ответа в файл
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Downloaded %s\n", fileName)
	return nil
}

func getFileNameFromURL(url string) string {
	parts := strings.Split(url, "/")
	fileName := parts[len(parts)-1]
	if fileName == "" {
		fileName = "index.html" // Если URL оканчивается на "/", используем "index.html"
	}
	return fileName
}
