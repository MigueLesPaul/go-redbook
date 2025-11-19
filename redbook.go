package main

import (
	fmt "fmt"

	obf "redbook/obsidianfrontmatter"
	fmtstats "redbook/pkg/services/frontmatterstats"
)

func main() {
	// test_path := "/data/data/com.termux/files/home/storage/shared/Obsidian/Journal"
	test_path := "/home/miguel/Documents/Obsidian/Journal"
	dailyNotes, error := obf.LoadFrontMattersFromDir(test_path)
	variable := "mood"
	dailyBalances, error := fmtstats.GetVariableNdayValues(dailyNotes, variable, 30)
	total, error := fmtstats.GetVariableNdayTotal(dailyBalances)

	if error != nil {
		panic("Panic at the Dico!")
	}
	fmt.Printf("Total de '%s'\n", variable)
	fmt.Println(dailyBalances)
	fmt.Println(total)

	variable = "finances"

	dailyBalances, error = fmtstats.GetVariableNdayValues(dailyNotes, variable, 30)
	total, error = fmtstats.GetVariableNdayTotal(dailyBalances)
}
