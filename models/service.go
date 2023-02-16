package models

import "context"

type Service interface {
	Signin(ctx context.Context, req SigninRequest) (SigninResponse, error)
	Signup(ctx context.Context, req SignupRequest) error
	Session(ctx context.Context, req GetSessionRequest) (GetSessionResponse, error)
	Balance(ctx context.Context, req GetWalletRequest) (GetWalletResponse, error)
	Credit(ctx context.Context, req CreditWalletRequest) error
	Debit(ctx context.Context, req DebitWalletRequest) error
}
