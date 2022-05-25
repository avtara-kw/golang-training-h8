package params

type Order struct {
	OrderedAt    string `json:"ordered_at"`
	CustomerName string `json:"customer_name"`
	Items        []struct {
		ItemCode    string `json:"item_code"`
		Description string `json:"description"`
		Quantity    int    `json:"quantity"`
	} `json:"items"`
}
