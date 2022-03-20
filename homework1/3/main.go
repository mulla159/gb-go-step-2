// Для закрепления практических навыков программирования,
// напишите программу, которая создаёт один миллион пустых файлов в известной,
// пустой директории файловой системы используя вызов os.Create.
// Ввиду наличия определенных ограничений операционной системы на число открытых файлов,
// такая программа должна выполнять аварийную остановку.
// Запустите программу и дождитесь полученной ошибки.
// Используя отложенный вызов функции закрытия файла, стабилизируйте работу приложения.
// Критерием успешного выполнения программы является успешное создание миллиона пустых файлов в директории

package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func createFile(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}

	defer file.Close()
}

func createAllFiles(dirPath string) {
	for i := 1; i <= 1000000; i++ {
		createFile(dirPath + "/files" + strconv.Itoa(i) + ".txt")
	}

	fmt.Println("Files created")
}

func main() {
	var fileName = flag.String("path", "files", "Path to dir")
	flag.Parse()

	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()

	fmt.Println("start")

	createAllFiles(*fileName)

	fmt.Println("end")
}
