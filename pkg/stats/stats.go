
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
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {

	sum := types.Money(0)
	
	for _, payment := range payments{
		if payment.Status == "FAIL" {

		}else if payment.Category == category {
		
			sum += payment.Amount
		
			
		}
	}
		return sum
	}
	
	func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {
		
		var filtered []types.Payment
		for _, payment := range payments {
			if payment.Category == category	{
				filtered = append(filtered, payment)
			}
			
		} 
		return filtered
	}



func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	
		categories := map[types.Category]types.Money{}
		for _, payment := range payments {
			
			
			result := FilterByCategory(payments,payment.Category)
			categories[payment.Category] += payment.Amount / types.Money(len(result))
		} 
		return categories
	}

	func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money {
		first = map[types.Category]types.Money{
			"auto": 2_999_99,
			"food": 3_000_00,
			
		}
		second = map[types.Category]types.Money{
			"auto": 5_999_99,
			"food": 7_000_00,
			
		}
		third := map[types.Category]types.Money{
			"auto": second["auto"] - first["auto"],
			"food": second["food"] - first["food"],
	   }
	   return third
	}