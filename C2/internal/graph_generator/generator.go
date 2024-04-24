package graph_generator

import (
	"C2/internal/data"
	"sort"
)

func contains(graph [][2]string, edge [2]string) bool {
	for _, e := range graph {
		if e[0] == edge[0] && e[1] == edge[1] {
			return true
		}
	}
	return false
}

func sortM() *[]string {
	keys := make([]string, 0, len(data.M))
	for key := range data.M {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return &keys
}

func GenerateGraph() {
	keys := *sortM()
	for i, key1 := range keys {
		for _, key2 := range keys[i+1:] {
			for _, v1 := range data.M[key1] {
				for _, v2 := range data.M[key2] {
					if contains(data.V, [2]string{v1, v2}) {
						data.R[[2]string{key1, key2}]++
					}
				}
			}
		}
	}
}
