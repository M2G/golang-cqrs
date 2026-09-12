package main

import (
	"errors"
	"fmt"
	"golang-cqrs/command"
	"golang-cqrs/event"
	"golang-cqrs/facade"
	"golang-cqrs/query"
	"time"
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

	getBalance := query.NewHandler[int]("GetBalance", func(args []string) (int, error) {
		if args[0] != accounts.ID {
			return 0, errors.New("Account not found")
		}
		return accounts.Balance, nil
	})

	return facade.NewFacade(
		map[string]*command.CommandHandler{
			"Deposit":  depositHandler,
			"Withdraw": withdrawHandler,
		},
		map[string]query.Handler{
			"GetBalance": getBalance,
		},
		eventBus,
	)
}

func main() {
	f := NewBookFacade()

	if err := f.Dispatch(command.NewCommand("Deposit", 5000)); err != nil {
		fmt.Println("erreur :", err)
	}

	if err := f.Dispatch(command.NewCommand("Withdraw", 5000)); err != nil {
		fmt.Println("erreur :", err)
	}

	result, err := f.Ask(query.NewQuery("GetBalance", accounts.ID))
	if err != nil {
		fmt.Println("erreur :", err)
		return
	}

	balance := result.(int)
	fmt.Printf("[QUERY] Solde final: %d centimes\n", balance)
	time.Sleep(50 * time.Millisecond)

}
