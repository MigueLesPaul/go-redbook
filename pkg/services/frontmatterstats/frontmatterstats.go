package frontmatterstats

import (
	fmt "fmt"
)

func GetVariableNdayValues(frontmatterList []map[string]interface{}, variable string, ndays int) ([]float32, error) {
	// Provides a list of n values that represents the daily balance for a specific variable from the frontmatter.
	// The variables have to be of object(map) type in the yaml
	// Asumes the list of Daily maps are sorted cronologically
	var balances []float32
	fmt.Println(ndays)
	for _, dailyf := range frontmatterList {
		if dailyf[variable] != nil {
			fmt.Printf(": %d\n",     dailyf[variable])
			switch v := dailyf[variable].(type) {
			case int:
				balances = append(balances, float32(v))
			case map[string]interface{}:
				var balanceHoy float32 = 0.0
				for _,sp := range v{
          //balanceHoy+=sp["value"].(float32)
					m, ok := sp.(map[string]interface{})
						if ok {
    					if val, ok := m["value"].(float32); ok {
        				balanceHoy += val
    					}
							if val, ok := m["amount"].(float32); ok {
								balanceHoy+= val
							}

						}
				}
				balances = append(balances, balanceHoy)
			}
		}
	}
	return balances, nil
}

func GetVariableNdayTotal(list []float32) (float32, error) {
	// Get the total from a list. basically df.sum() :)
	var sum float32

	for _, v := range list {
		sum += v
	}

	return sum, nil
}
