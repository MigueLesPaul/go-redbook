package main

import (
	fmt "fmt"
	"image/color"
	"log"
	"math"
	"slices"

	"redbook/internal/config"
	"redbook/internal/constants"
	obf "redbook/obsidianfrontmatter"
	fmtstats "redbook/pkg/services/frontmatterstats"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
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
	maxValue := slices.Max(dailyBalances)

	if maxValue > constants.WeeklyEuroTarget {
		dailyBalances = fmtstats.RemoveValue(dailyBalances, maxValue)
	}

	total, error = fmtstats.GetVariableNdayTotal(dailyBalances)

	log.Printf("Total de '%s'\n", variable)
	fmt.Println(dailyBalances)
	fmt.Println(total)
	totalBalance := math.Abs(total)
	variable = "bike"
	filter = fmtstats.Fieldfilters{}
	dailyBalances, error = fmtstats.GetVariableNdayValues(dailyNotes, variable, 30, filter)
	total, error = fmtstats.GetVariableNdayTotal(dailyBalances)

	fmt.Printf("Total de '%s'\n", variable)
	// fmt.Println(dailyBalances)
	fmt.Println(total)

	a := app.New()
	w := a.NewWindow("Expenses")

	// Main Title label
	title := widget.NewLabelWithStyle("Expenses", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	// Three labels below the title
	// formatted1 := fmt.Sprintf("Total Spent: %.2f/%d", totalBalance, constants.MonthlyEuroTarget)
	formattedfull := fmt.Sprintf("Total Spent: %.2f/%d \n Estimated Saving: %.2f", totalBalance, constants.MonthlyEuroTarget, constants.MonthlyEuroTarget-totalBalance)
	label1 := widget.NewLabel("")
	label1.SetText(formattedfull)
	// label2 := widget.NewLabel("")
	// label2Text := fmt.Sprintf("Estimated Saving: %.2f", constants.MonthlyEuroTarget-totalBalance)
	// label2.SetText(label2Text)
	// label3 := widget.NewLabel("Label 3: Additional Info")

	// labelsContainer := container.NewVBox(label1, label2, label3)
	labelsContainer := container.NewVBox(label1)
	// Placeholder for graph area - using a simple colored rectangle for demo
	graphArea := canvas.NewRectangle(color.RGBA{R: 100, G: 149, B: 237, A: 255}) // Cornflower Blue
	graphArea.SetMinSize(fyne.NewSize(4, 2))

	// Compose the final layout:
	content := container.NewVBox(
		title,
		labelsContainer,
		layout.NewSpacer(), // Add some space between labels and graph
		graphArea,
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(450, 400))
	w.ShowAndRun()
}
