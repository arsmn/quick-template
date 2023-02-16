package models

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/ory/herodot"
	"github.com/shopspring/decimal"
)

func Test_Credit(t *testing.T) {
	testCases := []struct {
		desc   string
		wallet *Wallet
		amount decimal.Decimal
		err    error
		result *Wallet
	}{
		{
			desc:   "#1 zero amount",
			wallet: &Wallet{Balance: decimal.NewFromInt(1)},
			amount: decimal.Zero,
			err:    nil,
			result: &Wallet{Balance: decimal.NewFromInt(1)},
		},
		{
			desc:   "#2 valid amount",
			wallet: &Wallet{Balance: decimal.NewFromInt(1)},
			amount: decimal.NewFromInt(2),
			err:    nil,
			result: &Wallet{Balance: decimal.NewFromInt(3)},
		},
		{
			desc:   "#3 negative amount",
			wallet: &Wallet{Balance: decimal.NewFromInt(1)},
			amount: decimal.NewFromInt(-1),
			err:    herodot.ErrBadRequest,
			result: &Wallet{Balance: decimal.NewFromInt(1)},
		},
		{
			desc:   "#4 float amount",
			wallet: &Wallet{Balance: decimal.NewFromInt(1)},
			amount: decimal.NewFromFloat(0.5),
			err:    nil,
			result: &Wallet{Balance: decimal.NewFromFloat(1.5)},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := tC.wallet.Credit(tC.amount)

			if !cmp.Equal(err, tC.err, cmpopts.EquateErrors()) {
				t.Errorf("expected error [%v] but got error [%v]", tC.err, err)
			}

			if !cmp.Equal(tC.wallet, tC.result) {
				t.Errorf("expected [%v] but got error [%v]", tC.result, tC.wallet)
			}
		})
	}
}

func Test_Debit(t *testing.T) {
	testCases := []struct {
		desc   string
		wallet *Wallet
		amount decimal.Decimal
		err    error
		result *Wallet
	}{
		{
			desc:   "#1 zero amount",
			wallet: &Wallet{Balance: decimal.NewFromInt(1)},
			amount: decimal.Zero,
			err:    nil,
			result: &Wallet{Balance: decimal.NewFromInt(1)},
		},
		{
			desc:   "#2 valid amount",
			wallet: &Wallet{Balance: decimal.NewFromInt(3)},
			amount: decimal.NewFromInt(2),
			err:    nil,
			result: &Wallet{Balance: decimal.NewFromInt(1)},
		},
		{
			desc:   "#3 negative amount",
			wallet: &Wallet{Balance: decimal.NewFromInt(1)},
			amount: decimal.NewFromInt(-1),
			err:    herodot.ErrBadRequest,
			result: &Wallet{Balance: decimal.NewFromInt(1)},
		},
		{
			desc:   "#4 float amount",
			wallet: &Wallet{Balance: decimal.NewFromInt(1)},
			amount: decimal.NewFromFloat(0.5),
			err:    nil,
			result: &Wallet{Balance: decimal.NewFromFloat(0.5)},
		},
		{
			desc:   "#5 negative balance",
			wallet: &Wallet{Balance: decimal.NewFromInt(1)},
			amount: decimal.NewFromFloat(1.1),
			err:    herodot.ErrBadRequest,
			result: &Wallet{Balance: decimal.NewFromInt(1)},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := tC.wallet.Debit(tC.amount)

			if !cmp.Equal(err, tC.err, cmpopts.EquateErrors()) {
				t.Errorf("expected error [%v] but got error [%v]", tC.err, err)
			}

			if !cmp.Equal(tC.wallet, tC.result) {
				t.Errorf("expected [%v] but got error [%v]", tC.result, tC.wallet)
			}
		})
	}
}
