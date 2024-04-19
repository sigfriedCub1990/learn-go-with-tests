package wallet

import "testing"

func TestWallet(t *testing.T) {

    t.Run("deposit", func(t *testing.T) {
        wallet := Wallet{}
        wallet.Deposit(Bitcon(10))
        assertBalance(t, wallet, Bitcon(10))
    })

    t.Run("withdraw with funds", func(t *testing.T) {
        wallet := Wallet{balance: Bitcon(20)}
        err := wallet.Withdraw(Bitcon(5))

        assertNoError(t, err)
        assertBalance(t, wallet, Bitcon(15))
    })

    t.Run("withdraw insufficient funds", func(t *testing.T) {
        startingBalance := Bitcon(20)
        wallet := Wallet{startingBalance}
        err := wallet.Withdraw(Bitcon(100))

        assertError(t, err, ErrInsufficientFunds)
        assertBalance(t, wallet, startingBalance)
    })
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcon) {
    t.Helper()
    got := wallet.Balance()

    if got != want {
        t.Errorf("got %s want %s", got, want)
    }
}

func assertNoError(t testing.TB, got error) {
    t.Helper()

    if got != nil {
        t.Fatal("got an error but din't want one")
    }
}

func assertError(t testing.TB, got error, want error) {
    t.Helper()
    if got == nil {
        t.Fatal("wanted an error but didn't get one")
    }

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}

