package main 

import (
	obf "redbook/obsidianfrontmatter"
	fmt "fmt"
)

func main(){
	test_path := "/data/data/com.termux/files/home/storage/shared/Obsidian/Journal/2025-11-19.md"

fmt.Println(obf.ReadFrontMatter(test_path))


}
