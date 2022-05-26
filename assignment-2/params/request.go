package params

type CreateOrder struct {
	OrderedAt    string `json:"ordered_at"`
	CustomerName string `json:"customer_name"`
	Items        []struct {
		ItemCode    string `json:"item_code"`
		Description string `json:"description"`
		Quantity    int    `json:"quantity"`
	} `json:"items"`
}

type UpdateOrder struct {
	OrderedAt    string `json:"ordered_at"`
	CustomerName string `json:"customer_name"`
	Items        []struct {
		LineItemID  int    `json:"line_item_id"`
		ItemCode    string `json:"item_code"`
		Description string `json:"description"`
		Quantity    int    `json:"quantity"`
	} `json:"items"`
}
