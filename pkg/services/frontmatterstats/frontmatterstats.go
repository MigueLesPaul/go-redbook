package frontmatterstats

import (
	fmt "fmt"
)

func GetVariableNdayValues(frontmatterList []map[string]interface{}, variable string, ndays int) ([]float64, error) {
	// Provides a list of n values that represents the daily balance for a specific variable from the frontmatter.
	// The variables have to be of object(map) type in the yaml
	// Asumes the list of Daily maps are sorted cronologically
	var balances []float64
	fmt.Println(ndays)
	for _, dailyf := range frontmatterList {
		if dailyf[variable] != nil {
			switch v := dailyf[variable].(type) {
			case int:
				balances = append(balances, float64(v))
			case []interface{}:
				var balanceHoy float64 = 0.0
				for _, sp := range v {
					// balanceHoy+=sp["value"].(float32)
					// fmt.Printf(": %T\n", sp)
					m, ok := sp.(map[string]interface{})
					if ok {
						// fmt.Println(m["amount"])
						if m["value"] != nil {
							value, ok := m["value"].(float64)
							if !ok {
								fmt.Println(ok)
							}
							balanceHoy += value
						}
						if m["amount"] != nil {
							value, ok := m["amount"].(float64)
							if !ok {
								if i, ok := m["amount"].(int); ok {
									value = float64(i)
								}
							}
							// fmt.Println(value, ok)
							balanceHoy += value
						}
					}
				}
				balances = append(balances, balanceHoy)
			}
		}
	}
	return balances, nil
}

func GetVariableNdayTotal(list []float64) (float64, error) {
	// Get the total from a list. basically df.sum() :)
	var sum float64

	for _, v := range list {
		sum += v
	}

	return sum, nil
}
