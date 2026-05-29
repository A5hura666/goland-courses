package main

import "fmt"

const (
	Pending = iota
	Shipped
	Delivered
	Cancelled
)

type Order struct {
	Status int
	Price  float64
}

func main() {
	// Utile, car on ne sait pas à l'avance le nombre de commandes
	//orders := []Order{
	//	{Pending, 120.50},
	//	{Shipped, 75.99},
	//	{Delivered, 300.00},
	//	{Pending, 50.00},
	//	{Cancelled, 20.00},
	//	{Delivered, 180.25},
	//	{Pending, 90.10},
	//}

	// Utilisation de make si on veut optimiser l'ajout de commande dans le slice
	orders := make([]Order, 0, 10)
	orders = append(orders, Order{Pending, 120.50})
	orders = append(orders, Order{Shipped, 75.99})
	orders = append(orders, Order{Delivered, 300.00})
	orders = append(orders, Order{Pending, 50.00})
	orders = append(orders, Order{Cancelled, 20.00})
	orders = append(orders, Order{Delivered, 180.25})
	orders = append(orders, Order{Pending, 90.10})

	var totalPending float64
	var totalShipped float64
	var totalDelivered float64
	var totalCancelled float64

	countPending := 0
	countShipped := 0
	countDelivered := 0
	countCancelled := 0

	maxPending := 0.0
	maxShipped := 0.0
	maxDelivered := 0.0
	maxCancelled := 0.0

	for _, order := range orders {
		switch order.Status {

		case Pending:
			totalPending += order.Price
			countPending++
			if order.Price > maxPending {
				maxPending = order.Price
			}

		case Shipped:
			totalShipped += order.Price
			countShipped++
			if order.Price > maxShipped {
				maxShipped = order.Price
			}

		case Delivered:
			totalDelivered += order.Price
			countDelivered++
			if order.Price > maxDelivered {
				maxDelivered = order.Price
			}

		case Cancelled:
			totalCancelled += order.Price
			countCancelled++
			if order.Price > maxCancelled {
				maxCancelled = order.Price
			}
		}
	}

	fmt.Println("--- DESCRIPTION DES COMMANDES ---")

	if countPending > 0 {
		fmt.Println("Pending => total:", totalPending,
			"avg:", totalPending/float64(countPending),
			"max:", maxPending)
	}

	if countShipped > 0 {
		fmt.Println("Shipped => total:", totalShipped,
			"avg:", totalShipped/float64(countShipped),
			"max:", maxShipped)
	}

	if countDelivered > 0 {
		fmt.Println("Delivered => total:", totalDelivered,
			"avg:", totalDelivered/float64(countDelivered),
			"max:", maxDelivered)
	}

	if countCancelled > 0 {
		fmt.Println("Cancelled => total:", totalCancelled,
			"avg:", totalCancelled/float64(countCancelled),
			"max:", maxCancelled)
	}
}
