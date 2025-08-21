package main

import "fmt"

type Product struct {
	Name  string
	Price float32
	Qty   float32
}

type Order struct {
	ID           string
	Items        []Product
	Discount     float32
	DiscountType string
	Subtotal     float32
	Total        float32
}

func (o *Order) subtotal() {
	var subtotalItem float32 = 0
	for _, item := range o.Items {
		subtotalItem += item.Price * item.Qty
	}
	o.Subtotal = subtotalItem
}
func (o *Order) total() {
	var afterDiscount float32 = 0
	if o.DiscountType == "percentage" {
		afterDiscount = o.Subtotal - (o.Subtotal * (o.Discount / 100))
	} else {
		afterDiscount = o.Subtotal - o.Discount
	}
	o.Total = afterDiscount
}

func (o *Order) display() {
	str := "----items----\n"
	for i, item := range o.Items {
		str += fmt.Sprintf("%d.%s %2.f x %2.f=%2.f\n", i+1, item.Name, item.Price, item.Qty, item.Price*item.Qty)
	}
	str += fmt.Sprintf("Subtotal: %2.f\n", o.Subtotal)
	str += "-------------\n"
	str += fmt.Sprintf("Discount: %2.f\n", o.Discount)
	str += fmt.Sprintf("DiscountType: %s\n", o.DiscountType)
	str += "-------------\n "
	str += fmt.Sprintf("Total: %2.f", o.Total)
	fmt.Println(str)

}
func main() {
	products := []Product{
		{
			Name:  "Mi Cheat",
			Price: 4000,
			Qty:   10,
		},
		{
			Name:  "ABC",
			Price: 4500,
			Qty:   1,
		},
	}
	order := &Order{
		ID:           "1",
		Items:        products,
		Discount:     10,
		DiscountType: "percentage",
		Subtotal:     0,
		Total:        0,
	}
	order.subtotal()
	order.total()
	order.display()

}
