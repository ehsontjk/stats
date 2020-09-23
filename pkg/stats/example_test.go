
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
			Status: "OK",
		},
		{
			ID: 34,
			Amount: 48_000_00,
			Category: "auto",
			Status: "FAILED",	
		},
		{
ID: 5,
Amount: 48_000_00,
Status: "OK",
	},
	{
		ID: 45,
		Amount: 41_000_00,
		Category: "bank",
		Status: "INPROGRESS",
			},

			
		}
	

	
	totalInCategory := TotalInCategory(payments, "OK")
	fmt.Println(totalInCategory)

	// Output: 8100000
}
func ExampleAvg() {
	payments := []types.Payment{
		{
			ID: 45,
		Amount: 41_000_00,
		Category: "bank",
		Status: "FAILED",
			},
			{
				ID: 45,
			Amount: 10,
			Category: "bank",
			Status: "OK",
				},
				{
ID: 45,
		Amount: 20,
		Category: "bank",
		Status: "INPROGRESS",
			},
			{
			ID: 45,
		Amount: 30,
		Category: "bank",
		Status: "OK",
			},
	}
result := Avg(payments, "OK")
fmt.Println(result)
//Output: 15

}
