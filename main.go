package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/yoeritjuu/Computer-Store/pkg/parts"
)

func main() {
	f, err := os.Open("computer_parts.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	partsList := parts.GetAllParts(data)

	parts.ReturnHTML(partsList)

	// jsonData, err := json.MarshalIndent(partsList, "", "  ")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(jsonData))
}
