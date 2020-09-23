
package stats


import (
	"github.com/ehsontjk/bank/pkg/bank/types"
	"fmt"
)
func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID: 12,
			Amount: 40_000_00,
			Category: "bank",
		},
		{
			ID: 13,
			Amount: 20_000_00,
			Category: "auto",
		},
		{
			ID: 34,
			Amount: 48_000_00,
			Category: "auto",	
		},
		{
ID: 5,
Amount: 48_000_00,
	},
	{
		ID: 45,
		Amount: 41_000_00,
		Category: "bank",
			},

			
		}
	

	
	totalInCategory := TotalInCategory(payments, "auto")
	fmt.Println(totalInCategory)
	// Output: 6800000
}

