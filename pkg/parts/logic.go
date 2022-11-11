package parts

import (
	"log"
	"strconv"
)

func GetAllParts(data [][]string) []Part {
	// convert csv lines to array of structs
	var partsList []Part
	for i, line := range data {
		if i > 0 { // omit header line
			var rec Part
			for j, field := range line {
				if j == 0 {
					rec.ID = field
				} else if j == 1 {
					rec.Description = field
				} else if j == 2 {
					rec.Brand = field
				} else if j == 3 {
					rec.Color = field
				} else if j == 4 {
					float, err := strconv.ParseFloat(field, 32)
					if err != nil {
						log.Panic(err)
					}
					rec.Price = float32(float)
				}
			}
			partsList = append(partsList, rec)
		}
	}
	return partsList
}
