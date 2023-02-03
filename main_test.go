package main

import "testing"

func TestCreateAcct(t *testing.T) {
	fidelityAcct := NewFidelity()
	zenithAcct := NewZenith()

	expectedFidelityAcct := &Fidelity{balance: 0}
	expectedZenithAcct := &Zenith{balance: 0}

	if *fidelityAcct != *expectedFidelityAcct {
		t.Errorf("Expected an Empty Fidelity acct")
	}

	if *zenithAcct != *expectedZenithAcct {
		t.Errorf("Expected an Empty Zenith acct")
	}

}

func TestDeposit(t *testing.T) {
	amount := 500
	fidelityAcct := &Fidelity{balance: 0}
	zenithAcct := &Fidelity{balance: 0}

	fidelityAcct.Deposit(amount)
	fidelityAcct.Deposit(amount)
	zenithAcct.Deposit(amount)
	zenithAcct.Deposit(amount)

	if fidelityAcct.balance != amount*2 {
		t.Errorf("Expected Fidelity balance (%v) to be equal to amount (%v)", fidelityAcct.balance, amount*2)
	}

	if zenithAcct.balance != amount*2 {
		t.Errorf("Expected Zenith balance (%v) to be equal to amount (%v)", zenithAcct.balance, amount*2)
	}
}

func TestWithdrawal(t *testing.T) {
	amount := 500
	fidelityAcct := &Fidelity{balance: amount}
	zenithAcct := &Fidelity{balance: amount}

	fidelityAcct.Withdraw(amount)
	zenithAcct.Withdraw(amount)

	if fidelityAcct.balance != 0 {
		t.Errorf("Expected Fidelity balance (%v) to be equal to zero", fidelityAcct.balance)
	}

	if zenithAcct.balance != 0 {
		t.Errorf("Expected Zenith balance (%v) to be equal to zero", zenithAcct.balance)
	}
}
