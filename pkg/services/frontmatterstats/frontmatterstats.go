package frontmatterstats

import (
	"time"
)

// Fieldfilters for second level maps filters
type Fieldfilters struct {
	Field string
	Value string
}

func GetVariableNdayValues(frontmatterList []map[string]interface{}, variable string, ndays int, filter Fieldfilters) ([]float64, error) {
	// Provides a list of n values that represents the daily balance for a specific variable from the frontmatter.
	// The variables have to be of object(map) type in the yaml
	// Asumes the list of Daily maps are sorted cronologically
	var balances []float64
	// fmt.Println(ndays)
	nDate := time.Now().AddDate(0, 0, -ndays)

	for _, dailyf := range frontmatterList {

		if dailyf["created"] != nil {
			dailyfDate := dailyf["created"].(time.Time)
			if nDate.After(dailyfDate) {
				continue
			}
		} else {
			continue
		}

		if dailyf[variable] != nil {
			switch v := dailyf[variable].(type) {
			case int:
				balances = append(balances, float64(v))
			case []interface{}:
				var balanceHoy float64 = 0.0
				var value float64 = 0.0
				for _, sp := range v {
					m, ok := sp.(map[string]interface{})
					if ok {
						if m["value"] != nil {
							value, ok = m["value"].(float64)
							if !ok {
								if i, ok := m["value"].(int); ok {
									value = float64(i)
								}
							}
						}
						if m["amount"] != nil {
							value, ok = m["amount"].(float64)
							if !ok {
								if i, ok := m["amount"].(int); ok {
									value = float64(i)
								}
							}
						}

						if filter.Field == "" || m[filter.Field] == filter.Value {
							balanceHoy += value
						}

					}
				}
				// fmt.Println(dailyf["created"], balanceHoy)
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

func RemoveValue(slice []float64, value float64) []float64 {
	for i, v := range slice {
		if v == value {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	// Return original slice if value not found
	return slice
}
