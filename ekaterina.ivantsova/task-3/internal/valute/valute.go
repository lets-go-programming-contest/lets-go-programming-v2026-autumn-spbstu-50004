package valute

import "sort"

type Item struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func SortByValueDesc(items []Item) {
	sort.SliceStable(items, func(left, right int) bool {
		return items[left].Value > items[right].Value
	})
}
