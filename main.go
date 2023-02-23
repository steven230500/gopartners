package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Ticket struct {
    ID          int       `json:"id"`
    User        string    `json:"user"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
    Status      string    `json:"status"`
}

var tickets []Ticket

func createTicket(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var ticket Ticket
    _ = json.NewDecoder(r.Body).Decode(&ticket)
    ticket.CreatedAt = time.Now()
    ticket.UpdatedAt = time.Now()
    tickets = append(tickets, ticket)

    json.NewEncoder(w).Encode(ticket)
}

func getTickets(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tickets)
}

func getTicketByID(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for _, ticket := range tickets {
        if strconv.Itoa(ticket.ID) == params["id"] {
            json.NewEncoder(w).Encode(ticket)
            return
        }
    }
    json.NewEncoder(w).Encode(&Ticket{})
}

func updateTicketByID(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for i, ticket := range tickets {
        if strconv.Itoa(ticket.ID) == params["id"] {
            tickets = append(tickets[:i], tickets[i+1:]...)

            var updatedTicket Ticket
            _ = json.NewDecoder(r.Body).Decode(&updatedTicket)
            updatedTicket.ID = ticket.ID
            updatedTicket.CreatedAt = ticket.CreatedAt
            updatedTicket.UpdatedAt = time.Now()
            tickets = append(tickets, updatedTicket)

            json.NewEncoder(w).Encode(updatedTicket)
            return
        }
    }
    json.NewEncoder(w).Encode(&Ticket{})
}

func deleteTicketByID(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for i, ticket := range tickets {
        if strconv.Itoa(ticket.ID) == params["id"] {
            tickets = append(tickets[:i], tickets[i+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(tickets)
}

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/tickets", createTicket).Methods("POST")
    router.HandleFunc("/tickets", getTickets).Methods("GET")
    router.HandleFunc("/tickets/{id}", getTicketByID).Methods("GET")
    router.HandleFunc("/tickets/{id}", updateTicketByID).Methods("PUT")
    router.HandleFunc("/tickets/{id}", deleteTicketByID).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8000", router))
}