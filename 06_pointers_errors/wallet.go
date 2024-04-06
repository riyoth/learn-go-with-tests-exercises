package pointer

import (
	"fmt"
	"errors"
)

type Bitcoin int

type Wallet struct{
	balance Bitcoin
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Deposit(amount Bitcoin){
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error{
	if amount > w.balance{
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin{
	//fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)

}
