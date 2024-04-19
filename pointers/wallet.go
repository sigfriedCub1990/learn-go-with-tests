package wallet

import (
	"errors"
	"fmt"
)

type Bitcon int

type Wallet struct {
    balance Bitcon
}

func (w *Wallet) Deposit(amount Bitcon) {
    w.balance += amount
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcon) error {

    if amount > w.balance {
        return ErrInsufficientFunds
    }

    w.balance -= amount
    return nil
}

func (w *Wallet) Balance() Bitcon {
    return w.balance
}

type Stringer interface {
    String() string
}

func (b Bitcon) String() string {
    return fmt.Sprintf("%d BTC", b)
}
