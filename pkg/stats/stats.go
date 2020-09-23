
package stats

import (
	"github.com/ehsontjk/bank/pkg/types"
)
func Avg(payments []types.Payment) types.Money  {
	sum := types.Money(0)
	max := types.Money(0)
	for _, payment := range payments{
	 
		sum += payment.Amount
		max = sum / types.Money(len(payments))
	}

return max

}
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {

	sum := types.Money(0)
	
	for _, payment := range payments{
		if payment.Category == category {
		
		sum += payment.Amount
		}
		
	}

return sum
}
