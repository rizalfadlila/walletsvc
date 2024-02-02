package wallet

import (
	"context"
	"github.com/julo/walletsvc/internal/domain/entity"
)

func (m *module) FindByUserID(ctx context.Context, userID string) (*entity.Wallet, error) {
	var wallet entity.Wallet

	err := m.BaseRepository.DB(ctx).GetContext(ctx, &wallet, sqlFindByUserID, userID)
	if err != nil {
		return nil, err
	}

	return &wallet, nil
}
