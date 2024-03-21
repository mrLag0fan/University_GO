package checker

import "C3/internal/hash_table"

func EqualHashTables(table1, table2 [hash_table.HashTableSize][]string) bool {
	for i, row1 := range table1 {
		if i >= len(table2) {
			return false
		}
		if !equalSlices(row1, table2[i]) {
			return false
		}
	}

	return true
}

func equalSlices(slice1, slice2 []string) bool {
	if len(slice1) == 0 && len(slice1) == 0 {
		return true
	}
	if slice1[0] != slice2[0] {
		return false
	}
	return true
}
