package main 

import (
	obf "redbook/obsidianfrontmatter"
	fmtstats "redbook/pkg/services/frontmatterstats"
	fmt "fmt"
)

func main(){
	test_path := "/data/data/com.termux/files/home/storage/shared/Obsidian/Journal"
	dailyNotes,error := obf.LoadFrontMattersFromDir(test_path)
	variable := "mood"
	dailyBalances,error := fmtstats.GetVariableNdayValues(dailyNotes,variable,30)
	if error != nil {
		panic("Panic at the Dico!")
	}
	fmt.Println(dailyBalances)
//fmt.Println(obf.LoadFrontMattersFromDir(test_path))


}
