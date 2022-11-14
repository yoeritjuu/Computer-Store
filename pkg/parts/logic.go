package parts

import (
	"log"
	"math"
	"os"
	"strconv"
)

type partService struct {
	partRepository PartsRepository
}

func NewPartService(u PartsRepository) PartsService {
	return &partService{
		partRepository: u,
	}
}

func (s *partService) GetAllParts() ([]Part, error) {
	list, err := s.partRepository.GetParts()
	if err != nil {
		log.Panic(err)
	}

	// // convert csv lines to array of structs
	// var partsList []Part
	// for i, line := range data {
	// 	if i > 0 { // omit header line
	// 		var rec Part
	// 		for j, field := range line {
	// 			if j == 0 {
	// 				rec.ID = field
	// 			} else if j == 1 {
	// 				rec.Description = field
	// 			} else if j == 2 {
	// 				rec.Brand = field
	// 			} else if j == 3 {
	// 				rec.Color = field
	// 			} else if j == 4 {
	// 				float, err := strconv.ParseFloat(field, 64)
	// 				if err != nil {
	// 					log.Panic(err)
	// 				}
	// 				rec.NoTaxPrice = float
	// 				rec.Price = AddTax(float)
	// 			}
	// 		}
	// 		partsList = append(partsList, rec)
	// 	}
	// }
	// return partsList

	return list, nil
}

func AddTax(rawPrice float64) float64 {
	tax, err := strconv.ParseFloat(os.Getenv("TAX_PERCENTAGE"), 64)
	if err != nil {
		log.Panic(err)
	}
	tax = (tax / 100) + 1
	newprice := rawPrice * float64(tax)
	rounded := math.Round(newprice*100) / 100
	return rounded
}
