package hash_table

import "crypto/md5"

const HashTableSize = 19

var TableFile1 [HashTableSize][]string
var TableFile2 [HashTableSize][]string

func Add(table *[HashTableSize][]string, value string) {
	index := hash(value) % HashTableSize
	table[index] = append(table[index], value)
}

func hash(value string) uint32 {
	h := md5.New()
	h.Write([]byte(value))
	return uint32(h.Sum(nil)[0])
}
