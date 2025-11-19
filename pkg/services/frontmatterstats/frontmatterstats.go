package frontmatterstats

import (
	fmt "fmt"
)

func GetVariableNdayValues(frontmatterList []map[string]interface{} ,variable string,ndays int ) ([]float32, error)  {
	// Provides a list of n values that represents the daily balance for a specific variable from the frontmatter. 
	// The variables have to be of object(map) type in the yaml
	// Asumes the list of Daily maps are sorted cronologically
  var balances []float32  
	fmt.Println(ndays)
	for i,dailyf := range frontmatterList {
		if dailyf[variable] != nil {

			fmt.Printf("d %d: %d\n",i,     dailyf[variable])
		}
	}
   return balances,nil
}

/*
func GetVariableNdayTotal() (float32, error) {
	// Get the total from a list. basically df.sum() :)


}

*/
