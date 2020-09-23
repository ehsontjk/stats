
package stats


import (
	"github.com/ehsontjk/bank/v2/pkg/types"
	"fmt"
)
func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID: 12,
			Amount: 40_000_00,
			Category: "bank",
			Status: "OK",
		},
		{
			ID: 13,
			Amount: 20_000_00,
			Category: "auto",
			Status: "FAILED",
		},
		{
			ID: 34,
			Amount: 48_000_00,
			Category: "auto",
			Status: "OK",	
		},
		{
ID: 5,
Amount: 48_000_00,
	},
	{
		ID: 45,
		Amount: 41_000_00,
		Category: "bank",
		Status: "INPROGRESS",
			},

			
		}
	

	
	totalInCategory := TotalInCategory(payments, "auto")
	fmt.Println(totalInCategory)
	// Output: 8800000
}

