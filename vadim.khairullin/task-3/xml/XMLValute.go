package xml

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/BinaryKenny/task-3/json"
)

type XMLValute struct {
	ID        string `xml:"id,attr"`
	NumCode   int    `xml:"NumCode"`
	CharCode  string `xml:"CharCode"`
	Nominal   int    `xml:"Nominal"`
	Name      string `xml:"Name"`
	Value     string `xml:"Value"`
	VunitRate string `xml:"VunitRate"`
}

func ValuteParseToJSON(val XMLValute) (json.JSONValute, error) {
	num := val.Value
	num = strings.ReplaceAll(num, ",", ".")

	result, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return json.JSONValute{}, fmt.Errorf("failed to parse to float: %w", err)
	}

	return json.JSONValute{
		NumCode:  val.NumCode,
		CharCode: val.CharCode,
		Value:    result,
	}, nil
}
