package entity

type Freight struct {
	PricePerKM float64
}

func NewFreight(pricePerKM float64) *Freight{
	return &Freight{
		PricePerKM: pricePerKM,
	}
}

func(f *Freight) Calculate(route *Route) {
	route.FreightPrice = route.Distance * f.PricePerKM
}