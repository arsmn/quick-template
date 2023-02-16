package app

import (
	"QUICK-Template/models"
	v "QUICK-Template/module/validation"
	"context"

	"github.com/ory/herodot"
)

func (a App) Balance(ctx context.Context, req models.GetWalletRequest) (models.GetWalletResponse, error) {
	if err := v.Validate(req); err != nil {
		return models.GetWalletResponse{}, err
	}

	fallback := func() (models.Wallet, error) {
		res, err := a.storage.Wallet().FindByUser(ctx, req.WalletID, req.UserID)
		if err != nil {
			return res, err
		}

		if err := a.walletCache.Set(ctx, req.WalletID, res); err != nil {
			a.logger.Error(err)
		}

		return res, nil
	}

	wallet, err := a.walletCache.GetFallback(ctx, req.WalletID, fallback)
	if err != nil {
		return models.GetWalletResponse{}, err
	}

	if !wallet.IsOwner(req.UserID) {
		return models.GetWalletResponse{}, herodot.ErrNotFound
	}

	return models.GetWalletResponse{
		ID:      wallet.ID,
		Balance: wallet.Balance,
		UserID:  wallet.UserID,
	}, nil
}

func (a App) Credit(ctx context.Context, req models.CreditWalletRequest) error {
	if err := v.Validate(req); err != nil {
		return err
	}

	wallet, err := a.storage.Wallet().FindByUser(ctx, req.WalletID, req.UserID)
	if err != nil {
		return err
	}

	if err := wallet.Credit(req.Amount); err != nil {
		return err
	}

	tx := models.Transaction{
		Amount:      req.Amount,
		Type:        models.Credit,
		Description: req.Description,
		WalletID:    wallet.ID,
		Wallet:      wallet,
	}

	err = a.storage.Transaction().Credit(ctx, tx)
	if err != nil {
		return err
	}

	err = a.walletCache.Del(ctx, req.WalletID)
	if err != nil {
		a.logger.Error(err)
	}

	return nil
}

func (a App) Debit(ctx context.Context, req models.DebitWalletRequest) error {
	if err := v.Validate(req); err != nil {
		return err
	}

	wallet, err := a.storage.Wallet().FindByUser(ctx, req.WalletID, req.UserID)
	if err != nil {
		return err
	}

	if err := wallet.Debit(req.Amount); err != nil {
		return err
	}

	tx := models.Transaction{
		Amount:      req.Amount,
		Type:        models.Debit,
		Description: req.Description,
		WalletID:    wallet.ID,
		Wallet:      wallet,
	}

	err = a.storage.Transaction().Debit(ctx, tx)
	if err != nil {
		return err
	}

	err = a.walletCache.Del(ctx, req.WalletID)
	if err != nil {
		a.logger.Error(err)
	}

	return nil
}
