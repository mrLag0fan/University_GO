package main

import (
	memory_manager "C1/internal"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please set memory size and max output width (separated by space):")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	parts := strings.Split(input, " ")

	if len(parts) != 2 {
		fmt.Println("Invalid input. Please provide both memory size and max output width.")
		return
	}

	memory_manager.TotalSize, _ = strconv.Atoi(parts[0])

	maxOutputWidth, _ := strconv.Atoi(parts[1])

	fmt.Println("Type 'help' for additional info.")

	for {
		fmt.Print("> ")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)
		parts := strings.Split(command, " ")

		switch parts[0] {
		case "help":
			memory_manager.Help()
		case "exit":
			fmt.Println("Exiting program...")
			return
		case "print":
			memory_manager.PrintMemoryBlocks(memory_manager.MemoryBlocks, memory_manager.TotalSize, maxOutputWidth)
		case "allocate":
			if len(parts) < 2 {
				fmt.Println("Please provide the number of cells to allocate.")
				continue
			}
			size, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Invalid size provided.")
				continue
			}
			blockID := memory_manager.Allocate(size)
			fmt.Println(blockID)
		case "free":
			if len(parts) < 2 {
				fmt.Println("Please provide the block ID to free.")
				continue
			}
			blockID, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Invalid block ID provided.")
				continue
			}
			memory_manager.Free(blockID)
		default:
			fmt.Println("Invalid command. Type 'help' for available commands.")
		}
	}
}
