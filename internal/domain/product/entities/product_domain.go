package entities

type Product struct {
	ID    int
	Name  string
	Price float64
	Count int
}

func (p *Product) Validator() bool {
	if p.Name == "" || p.Count == 0 || p.Price == 0 {
		return false
	}
	return true
}
