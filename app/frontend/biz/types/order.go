package types

type OrderItems struct {
	ProductName string
	Picture     string
	Qty         uint32
	Cost        float32
}

type Order struct {
	OrderId     string
	CreatedDate string
	Qty         uint32
	Cost        float32
	Items       []OrderItems
}
