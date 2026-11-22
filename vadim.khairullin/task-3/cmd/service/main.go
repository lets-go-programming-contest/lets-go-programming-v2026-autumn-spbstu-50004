package main

import (
	"github.com/BinaryKenny/task-3/config"
	"github.com/BinaryKenny/task-3/json"
	"github.com/BinaryKenny/task-3/xml"
)

func main() {
	conf, err := config.ParseConfig()
	if err != nil {
		panic("Failed to parse config: " + err.Error())
	}

	ValCurs := xml.ParseXMLFile(conf.InputFile)
	json.EncodeToJSON(ValCurs, conf.OutputFile)
}
