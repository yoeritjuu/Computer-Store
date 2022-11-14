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
	var partsWithTax []Part
	list, err := s.partRepository.GetParts()
	if err != nil {
		log.Panic(err)
	}

	for _, part := range list {
		part.Price = AddTax(part.NoTaxPrice)
		partsWithTax = append(partsWithTax, part)
	}

	return partsWithTax, nil
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
