package main

import (
	fmt "fmt"

	"redbook/internal/config"
	obf "redbook/obsidianfrontmatter"
	fmtstats "redbook/pkg/services/frontmatterstats"
)

func main() {
	// test_path := "/data/data/com.termux/files/home/storage/shared/Obsidian/Journal"
	// test_path := "/home/miguel/Documents/Obsidian/Journal"

	configuracion := config.LoadConfig()

	NotesDir := configuracion["dirpath"].(string)
	dailyNotes, error := obf.LoadFrontMattersFromDir(NotesDir)

	variable := "mood"
	filter := fmtstats.Fieldfilters{}
	dailyBalances, error := fmtstats.GetVariableNdayValues(dailyNotes, variable, 30, filter)
	total, error := fmtstats.GetVariableNdayTotal(dailyBalances)

	if error != nil {
		panic("Panic at the Dico!")
	}
	fmt.Printf("Total de '%s'\n", variable)
	// fmt.Println(dailyBalances)
	fmt.Println(total)

	variable = "finances"
	filter = fmtstats.Fieldfilters{Field: "currency", Value: "EUR"}
	dailyBalances, error = fmtstats.GetVariableNdayValues(dailyNotes, variable, 30, filter)
	total, error = fmtstats.GetVariableNdayTotal(dailyBalances)

	fmt.Printf("Total de '%s'\n", variable)
	// fmt.Println(dailyBalances)
	fmt.Println(total)

	variable = "bike"
	filter = fmtstats.Fieldfilters{}
	dailyBalances, error = fmtstats.GetVariableNdayValues(dailyNotes, variable, 30, filter)
	total, error = fmtstats.GetVariableNdayTotal(dailyBalances)

	fmt.Printf("Total de '%s'\n", variable)
	// fmt.Println(dailyBalances)
	fmt.Println(total)
}
