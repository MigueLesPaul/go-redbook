package main 

import (
	obf "redbook/obsidianfrontmatter"
	fmt "fmt"
)

func main(){
	test_path := "/data/data/com.termux/files/home/storage/shared/Obsidian/Journal"

fmt.Println(obf.LoadFrontMattersFromDir(test_path))


}
