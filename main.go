package main

import (
	"fmt"
	"net/http"
)

type TicketSale struct {
	ID          int     `json:"id"`
	EventName   string  `json:"event_name"`
	TicketPrice float64 `json:"ticket_price"`
	Quantity    int     `json:"quantity"`
}

var ticketSales []TicketSale

func createTicketSale(w http.ResponseWriter, r *http.Request) {
	newTicketSale := TicketSale{
		ID:          len(ticketSales) + 1,
		EventName:   "Rep Festival",
		TicketPrice: 149.90,
		Quantity:    1000,
	}
	ticketSales = append(ticketSales, newTicketSale)
	fmt.Println(w, "Venda de ingresso criada com sucesso!")
}

func getAllTicketSales(w http.ResponseWriter, r *http.Request) {
	for _, sale := range ticketSales {
		fmt.Fprintf(w, "ID: %d, Evento: %s, Preço: %.2f, Quantidade: %d\n", sale.ID, sale.EventName, sale.TicketPrice, sale.Quantity)
	}
}

func main() {
	http.HandleFunc("/ticket-sales", createTicketSale)
	http.HandleFunc("/ticket-sales/all", getAllTicketSales)

	fmt.Println("Servidor rodando na porta 8080..")
	http.ListenAndServe(":8080", nil)
}
