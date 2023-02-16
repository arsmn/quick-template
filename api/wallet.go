package api

import (
	"QUICK-Template/models"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const (
	WalletIDParam = "wallet_id"
)

func (h *Handler) walletsRouter(r chi.Router) {
	r.Use(h.AuthorizeMiddleware)
	r.Get("/{wallet_id}/balance", h.convert(h.getWalletBalance))
	r.Post("/{wallet_id}/credit", h.convert(h.creditWallet))
	r.Post("/{wallet_id}/debit", h.convert(h.debitWallet))
}

// GetWalletBalance godoc
// @Summary      Get Wallet Balance
// @Description  Retrieves the balance of a given wallet id
// @Tags         balance
// @Produce      json
// @Param        wallet_id  path uint true "Wallet ID"
// @Success      200  {object}  api.PayloadResponse[models.GetWalletResponse]
// @Failure      400  {object}  herodot.DefaultError
// @Failure      401  {object}  herodot.DefaultError
// @Failure      500  {object}  herodot.DefaultError
// @Router       /wallets/{wallet_id}/balance [GET]
func (h *Handler) getWalletBalance(rw http.ResponseWriter, r *http.Request) error {
	wid, err := walletIDParam(r)
	if err != nil {
		return err
	}

	req := models.GetWalletRequest{
		AuthorizedRequest: authorizedRequest(r.Context()),
		WalletID:          wid,
	}

	res, err := h.service.Balance(r.Context(), req)
	if err != nil {
		return err
	}

	return h.writer.Write(rw, r, OkPayloadResponse(res))
}

// CreditWallet godoc
// @Summary      Credit wallet balance
// @Description  Credits money on a given wallet id
// @Tags         credit
// @Accept       json
// @Produce      json
// @Param        wallet_id  path uint true "Wallet ID"
// @Param        request body models.CreditWalletRequest true "payload"
// @Success      200  {object}  api.Response
// @Failure      400  {object}  herodot.DefaultError
// @Failure      401  {object}  herodot.DefaultError
// @Failure      500  {object}  herodot.DefaultError
// @Router       /wallets/{wallet_id}/credit [POST]
func (h *Handler) creditWallet(rw http.ResponseWriter, r *http.Request) error {
	req, err := ParseJSON[models.CreditWalletRequest](r)
	if err != nil {
		return err
	}

	wid, err := walletIDParam(r)
	if err != nil {
		return err
	}

	req.WalletID = wid
	req.AuthorizedRequest = authorizedRequest(r.Context())

	if err := h.service.Credit(r.Context(), req); err != nil {
		return err
	}

	return h.writer.Write(rw, r, OkResponse)
}

// DebitWallet godoc
// @Summary      Debit wallet balance
// @Description  Debits money from a given wallet id
// @Tags         debit
// @Accept       json
// @Produce      json
// @Param        wallet_id  path uint true "Wallet ID"
// @Param        request body models.DebitWalletRequest true "payload"
// @Success      200  {object}  api.Response
// @Failure      400  {object}  herodot.DefaultError
// @Failure      401  {object}  herodot.DefaultError
// @Failure      500  {object}  herodot.DefaultError
// @Router       /wallets/{wallet_id}/debit [POST]
func (h *Handler) debitWallet(rw http.ResponseWriter, r *http.Request) error {
	req, err := ParseJSON[models.DebitWalletRequest](r)
	if err != nil {
		return err
	}

	wid, err := walletIDParam(r)
	if err != nil {
		return err
	}

	req.WalletID = wid
	req.AuthorizedRequest = authorizedRequest(r.Context())

	if err := h.service.Debit(r.Context(), req); err != nil {
		return err
	}

	return h.writer.Write(rw, r, OkResponse)
}
