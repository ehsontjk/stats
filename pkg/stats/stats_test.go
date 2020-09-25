package stats


import (
	"github.com/ehsontjk/bank/v2/pkg/types"
	"testing"
	"reflect"
)
func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{ID:1, Category: "auto", Amount: 1_000_00},
		{ID:2, Category: "food", Amount: 2_000_00},
		{ID:8, Category: "food", Amount: 2_000_00},
		{ID:7, Category: "food", Amount: 2_000_00},
		{ID:10, Category: "auto", Amount: 1_000_00},
		{ID:9, Category: "auto", Amount: 1_000_00},
		{ID:3, Category: "auto", Amount: 2_000_00},
		{ID:70, Category: "food", Amount: 2_000_00},
		{ID:12, Category: "food", Amount: 2_000_00},
		{ID:77, Category: "food", Amount: 14_000_00},
		
	}
	expected := map[types.Category]types.Money{
		"auto": 3_000_00 ,
		"food": 24_000_00,
		
	}
	result := CategoriesAvg(payments)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}
func TestFilterByCategory_foundOne(t *testing.T) {
	payments := []types.Payment{
		{ID:1, Category: "auto", Amount: 1_000_00},
		{ID:2, Category: "food"},
		
		
		
	}
	expected := []types.Payment{
	
		{ID: 2, Category: "food"},
		
	}
	result := FilterByCategory(payments, "food")	
	if !reflect.DeepEqual(expected, result) {
		t.Error("invalid result")
	}
}
