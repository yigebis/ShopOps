package Repository

import (
	"ShopOps/Domain"
	"ShopOps/UseCase"

	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type TokenRepository struct{
	DbCtx context.Context
	Collection *mongo.Collection
}

func NewTokenRepository(ctx context.Context, collection *mongo.Collection) UseCase.ITokenRepository {
	return &TokenRepository{
		DbCtx: ctx,
		Collection: collection,
	}
}

// InsertRefresher implements UseCase.ITokenRepository.
func (t *TokenRepository) InsertRefresher(email string, refresher string) error {
	refresher_data := Domain.RefresherTable{
		Email : email,
		Refresher: refresher,
	}

	_, err := t.Collection.InsertOne(t.DbCtx, refresher_data)
	return err
}

// CheckRefresher implements UseCase.ITokenRepository.
func (t *TokenRepository) CheckRefresher(email string, refresher string) error {
	panic("unimplemented")
}

// DeleteAllRefreshers implements UseCase.ITokenRepository.
func (t *TokenRepository) DeleteAllRefreshers(email string) error {
	panic("unimplemented")
}

// DeleteRefresher implements UseCase.ITokenRepository.
func (t *TokenRepository) DeleteRefresher(email string, refresher string) error {
	panic("unimplemented")
}

// InvalidateResetToken implements UseCase.ITokenRepository.
func (t *TokenRepository) InvalidateResetToken(email string) error {
	panic("unimplemented")
}

// LogoutToken implements UseCase.ITokenRepository.
func (t *TokenRepository) LogoutToken(token string) error {
	panic("unimplemented")
}

// StoreResetToken implements UseCase.ITokenRepository.
func (t *TokenRepository) StoreResetToken(email string, resetToken string) error {
	panic("unimplemented")
}

// UpdateRefresher implements UseCase.ITokenRepository.
func (t *TokenRepository) UpdateRefresher(email string, refresher string) error {
	panic("unimplemented")
}
