package wallet

import (
	"context"
	"fmt"
	"github.com/julo/walletsvc/internal/domain/entity"
	"github.com/julo/walletsvc/internal/domain/enums"
	"strconv"
)

func (m *module) Store(ctx context.Context, entity entity.Wallet) error {
	_, err := m.BaseRepository.DB(ctx).NamedExecContext(ctx, sqlStore, entity)
	return err
}

func (m *module) UpdateStatusByUserID(ctx context.Context, status enums.StatusWallet, epocMilli int64, userID string) error {
	column1 := "enabled_at"
	column2 := "disabled_at"

	if status == enums.Disabled {
		column1 = "disabled_at"
		column2 = "enabled_at"
	}

	sqlStatement := fmt.Sprintf(sqlUpdateStatusByUserID, column1, column2)

	_, err := m.BaseRepository.DB(ctx).ExecContext(ctx, sqlStatement, status, strconv.FormatInt(epocMilli, 10), userID, status.Against())
	if err != nil {
		return err
	}

	return err
}

func (m *module) UpdateBalance(ctx context.Context, entity entity.Wallet) error {
	//TODO implement me
	panic("implement me")
}
