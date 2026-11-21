package xml

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"github.com/BinaryKenny/task-3/json"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

type ValCurs struct {
	Date    string      `xml:"Date,attr"`
	Name    string      `xml:"name,attr"`
	Valutes []XMLValute `xml:"Valute"`
}

func ParseXMLFile(path string) ValCurs {
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintln("xml file is uncorrect " + err.Error()))
	}

	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	decoder := xml.NewDecoder(file)

	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		return transform.NewReader(input, charmap.Windows1251.NewDecoder()), nil
	}

	var valCurs ValCurs

	err = decoder.Decode(&valCurs)
	if err != nil {
		panic(fmt.Sprintln("The structure of xml file is wrong" + err.Error()))
	}

	return valCurs
}

func ValuteSliceParseToJSON(curs ValCurs) ([]json.JSONValute, error) {
	valutes := make([]json.JSONValute, 0, len(curs.Valutes))

	for _, val := range curs.Valutes {
		valute, err := ValuteParseToJSON(val)
		if err != nil {
			return nil, err
		}

		valutes = append(valutes, valute)
	}

	return valutes, nil
}
