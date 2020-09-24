
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
			Category: "cafe",
			Status: "StatusOk",
		},
		{
			ID: 13,
			Amount: 20_000_00,
			Category: "auto",
			Status: "StatusOk",
			
		},
		{
			ID: 34,
			Amount: 48_000_00,
			Category: "auto",
			Status: "StatusFail",
				
		},
		{
ID: 5,
Amount: 48_000_00,
Category: "auto",
Status: "StatusInProgress",
	},
	{
		ID: 45,
		Amount: 41_000_00,
		Category: "auto",
		Status: "StatusFail",
		
			},

			
		}
	

	
	totalInCategory := TotalInCategory(payments, "auto")
	fmt.Println(totalInCategory)

	// Output: 6800000
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
result := Avg(payments)
fmt.Println(result)
//Output: 20

}
