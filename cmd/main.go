package main

import (
	"fmt"
	"golang-cqrs/command"
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
	deposit := func(args []interface{}) {
		amount := args[0].(int)
		accounts.Balance += amount
		fmt.Printf("[COMMAND] Dépôt de %d centimes, nouveau solde: %d\n", amount, accounts.Balance)
		eventBus.Publish(*event.NewEvent("BalanceUpdated", nil))
	}

	withdraw := func(args []interface{}) {
		amount := args[0].(int)
		if amount > accounts.Balance {
			fmt.Printf("[COMMAND] Retrait refusé: solde insuffisant")
			return
		}
		accounts.Balance -= amount
		fmt.Printf("[COMMAND] Retrait de %d centimes, nouveau solde: %d\n", amount, accounts.Balance)
		eventBus.Publish(*event.NewEvent("BalanceUpdated", nil))
	}

	depositHandler := command.NewHandler(deposit)
	withdrawHandler := command.NewHandler(withdraw)

}

func main() {
}
