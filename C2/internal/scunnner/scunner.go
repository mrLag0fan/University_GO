package scunnner

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadV() [][2]string {
	var V [][2]string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter a pair of vertices (or press Enter to finish): ")
		scanner.Scan()
		input := scanner.Text()
		if input == "" {
			break
		}
		vertices := strings.Fields(input)
		if len(vertices) != 2 {
			fmt.Println("Invalid input. Please enter two vertices separated by a space.")
			continue
		}
		V = append(V, [2]string{vertices[0], vertices[1]})
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}

	return V
}

func ReadM() map[string][]string {
	M := make(map[string][]string)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter the name of the set of vertices (or press Enter to finish): ")
		scanner.Scan()
		setName := scanner.Text()
		if setName == "" {
			break
		}

		fmt.Print("Enter the vertices for set ", setName, " (separated by spaces): ")
		scanner.Scan()
		vertices := strings.Fields(scanner.Text())
		M[setName] = vertices
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
	return M
}
