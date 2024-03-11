package main

import (
	"fmt"
	"log"
	//"os"
	"strings"
	"path/filepath"
)

func main() {
	// if len(os.Args) < 2 {
	// log.Fatal("Укажите полный путь до файла вторым аргументом")
	// }
	filePth := "./test.txt"
	abs_fname, err := filepath.Abs(filePth)
	if err != nil {
	log.Fatal(err)
	}
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
	fmt.Println("Полный путь: ", abs_fname)
	fmt.Printf("Base: %s\n", filepath.Base(filePth))
	fmt.Printf("Ext %s\n", filepath.Ext(filePth))
}
