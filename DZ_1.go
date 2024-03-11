package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
		if len(os.Args) < 2 {
		log.Fatal("Укажите полный путь до файла вторым аргументом")
	}
	filePth := os.Args[1]
	var fileName, fileExt string
	indE := strings.LastIndex(filePth, ".")
	fileExt = filePth[indE+1:]
	indN := strings.LastIndex(filePth, "/")
	fileName = filePth[indN+1:indE]
	// Напишите код, который выведет следующее
	// filename: <name>
	// extension: <extension>

	// Подсказка. Возможно вам понадобится функция strings.LastIndex
	// Для проверки своего решения используйте функции filepath.Base() filepath.Ext(
	// ) Они могут помочь для проверки решения
	fmt.Printf("filename: %s\n", fileName)
	fmt.Printf("extension: %s\n", fileExt)
}
