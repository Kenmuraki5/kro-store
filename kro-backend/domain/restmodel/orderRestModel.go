package restmodel

type Order struct {
	ProductId       string  `json:"productId"`
	Quantity        int     `json:"quantity"`
	Email           string  `json:"email"`
	OrderDate       string  `json:"orderDate"`
	Subtotal        float64 `json:"subtotal"`
	ShippingAddress string  `json:"shippingAddress"`
	ShippingMethod  string  `json:"shippingMethod"`
	Type            string  `json:"type"`
}
