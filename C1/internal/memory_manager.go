package memory_manager

import "fmt"

var TotalSize = 30
var MemoryBlocks = [][2]int{}

func Allocate(size int) int {
	if len(MemoryBlocks) == 0 {
		if size <= TotalSize {
			MemoryBlocks = append(MemoryBlocks, [2]int{0, size - 1})
			return 0
		}
		return -1
	}

	if MemoryBlocks[0][0] >= size {
		MemoryBlocks = append([][2]int{{0, size - 1}}, MemoryBlocks...)
		return 0
	}
	for i := 0; i < len(MemoryBlocks)-1; i++ {
		if MemoryBlocks[i+1][0]-MemoryBlocks[i][1]-1 >= size {
			MemoryBlocks = append(MemoryBlocks[:i+1], append([][2]int{{MemoryBlocks[i][1] + 1, MemoryBlocks[i][1] + size}}, MemoryBlocks[i+1:]...)...)
			return MemoryBlocks[i][1] + 1
		}
	}
	lastIndex := MemoryBlocks[len(MemoryBlocks)-1][1]
	if lastIndex+size < TotalSize {
		MemoryBlocks = append(MemoryBlocks, [2]int{
			lastIndex + 1,
			lastIndex + size,
		})
		return lastIndex + 1
	}
	return -1
}

func Free(index int) {
	for i := 0; i < len(MemoryBlocks); i++ {
		if MemoryBlocks[i][0] == index {
			MemoryBlocks = append(MemoryBlocks[:i], MemoryBlocks[i+1:]...)
		}
	}
}

func PrintMemoryBlocks(memoryBlocks [][2]int, totalSize int, maxOutputWidth int) {
	fmt.Println("Memory Blocks:")
	for i := 0; i < totalSize/maxOutputWidth; i++ {
		fmt.Print("|")
		for j := 0; j < maxOutputWidth; j++ {
			found := false
			for _, block := range memoryBlocks {
				if block[0] == j+i*maxOutputWidth {
					found = true
					if i*maxOutputWidth != block[0] {
						fmt.Print("|")
					}
					fmt.Print(block[0])
					break
				} else if block[0] < j+i*maxOutputWidth && block[1] >= j+i*maxOutputWidth {
					found = true
					fmt.Print("x")
					break
				}
			}
			if !found {
				fmt.Print(" ")
			}
		}
		fmt.Println("|")
	}
}

func Help() {
	fmt.Println("Available commands:")
	fmt.Println("help - show this help")
	fmt.Println("exit - exit this program")
	fmt.Println("print - print memory blocks map")
	fmt.Println("allocate <num> - allocate <num> cells. Returns block first cell number")
	fmt.Println("free <num> - free block with first cell number <num>")
}
