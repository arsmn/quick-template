package models

import "context"

type Storage interface {
	User() UserPersister
	Wallet() WalletPersister
	Transaction() TransactionPersister
	Session() SessionPersister

	Close(ctx context.Context) error
}

type GenericPersister[T any] interface {
	Find(ctx context.Context, id uint) (T, error)
	Create(ctx context.Context, model T) (T, error)
	Update(ctx context.Context, id uint, model T) (T, error)
	Delete(ctx context.Context, id uint, model T) error
}

type UserPersister interface {
	GenericPersister[User]
	FindByEmail(ctx context.Context, email string) (User, error)
}

type WalletPersister interface {
	GenericPersister[Wallet]
	FindByUser(ctx context.Context, wid, uid uint) (Wallet, error)
}

type TransactionPersister interface {
	GenericPersister[Transaction]
	Credit(ctx context.Context, tx Transaction) error
	Debit(ctx context.Context, tx Transaction) error
}

type SessionPersister interface {
	GenericPersister[Session]
	FindByToken(ctx context.Context, token string) (Session, error)
}
