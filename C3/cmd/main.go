package main

import (
	"C3/internal/checker"
	"C3/internal/hash_table"
	"C3/internal/scanner"
	"fmt"
)

func main() {
	scanner.Scan(&hash_table.TableFile1, "./data/first.txt")
	scanner.Scan(&hash_table.TableFile2, "./data/second.txt")
	fmt.Println(hash_table.TableFile1)
	fmt.Println(hash_table.TableFile2)
	fmt.Println(checker.EqualHashTables(hash_table.TableFile1, hash_table.TableFile2))
}
