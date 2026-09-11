package main

import (
	"fmt"
	"golang-cqrs/event"
	"golang-cqrs/facade"
)

type BankAccount struct {
	ID      string
	Balance int
}

var accounts = &BankAccount{ID: "acc-1", Balance: 0}

func NewBookFacade() *facade.Facade {
	eventBus := event.NewEventBus()

	balanceLogger := event.NewEventHandler(*event.NewEvent("BalanceUpdated", func(data interface{}) {
		fmt.Printf("[EVENT] Notification abonné: solde actuel = %d centimes\n", accounts.Balance)
	}), eventBus)

	_ = balanceLogger

	// cmd side

	return facade.NewFacade(
	// ...
	)
}

func main() {
}
