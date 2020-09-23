
package stats

import (
	"github.com/ehsontjk/bank/v2/pkg/types"
)
func Avg(payments []types.Payment) types.Money  {
	sum := types.Money(0)
	max := types.Money(0)
	for _, payment := range payments{
	
		if payment.Status == "FAILED"{
		}else {
		sum += payment.Amount
		max = sum / types.Money(len(payments) - 1)
	}
}

return max

}
func TotalInCategory(payments []types.Payment, status types.Status) types.Money {

	sum := types.Money(0)
	
	for _, payment := range payments{
		
		if payment.Status == "StatusOk" {
		
		
		
			sum += payment.Amount
		}else if payment.Status == "StatusInProgress" {
			sum += payment.Amount
		}
		
	}

return sum
}
