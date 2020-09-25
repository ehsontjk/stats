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
		{ID:3, Category: "auto", Amount: 3_000_00},
		{ID:4, Category: "auto", Amount: 4_000_00},
		{ID:5, Category: "fun", Amount: 5_000_00},
	}
	expected := map[types.Category]types.Money{
		"auto": 1_600_00,
		"food": 400_00,
		"fun": 1000_00,
	}
	result := CategoriesAvg(payments)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}