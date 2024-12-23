package store

const defaultTaxRate float64 = 0.2
const minThreshold = 10

var categoryMaxPrices = map[string]float64 {
	"Water Sports": 250,
	"Soccer": 150,
	"Chess": 50,
}

func init() {
	for category, price := range categoryMaxPrices {
		categoryMaxPrices[category] = price + (price * defaultTaxRate)
	}
}

/*
var categoryMaxPrices = map[string]float64 {
	"Water Sports": 250 + (250 * defaultTaxRate),
	"Soccer": 150 + (150 * defaultTaxRate),
	"Chess": 50 + (50 * defaultTaxRate),
}
*/

type taxRate struct {
	rate, threshold float64
}

func newTaxRate(rate, threshold float64) *taxRate {
	if (rate == 0) {
		rate = defaultTaxRate
	}
	if (threshold < minThreshold) {
		threshold = minThreshold
	}
	return &taxRate {rate, threshold}
}

func (taxRate *taxRate) calcTax(product *Product) (price float64) {
	if (product.price > taxRate.threshold) {
		price = product.price + (product.price * taxRate.rate)
	} else {
		price = product.price
	}
	if max, ok := categoryMaxPrices[product.Category]; ok && price > max {
		price = max
	}
	return
}

/*
func (taxRate *taxRate) calcTax(product *Product) float64 {
	if (product.price > taxRate.threshold) {
		return product.price + (product.price * taxRate.rate)
	}
	return product.price
}
*/