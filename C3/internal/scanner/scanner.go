package scanner

import (
	"C3/internal/hash_table"
	"bufio"
	"fmt"
	"log"
	"os"
)

func Scan(hashTable *[hash_table.HashTableSize][]string, path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		hash_table.Add(hashTable, fileScanner.Text())
	}

	if err := fileScanner.Err(); err != nil {
		fmt.Println("Помилка читання файлу:", err)
	}
}
