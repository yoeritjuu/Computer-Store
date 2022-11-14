package parts

type Data struct {
	Items []Part
}

type Part struct {
	ID          string
	Description string
	Brand       string
	Color       string
	Price       float64
	NoTaxPrice  float64
}
