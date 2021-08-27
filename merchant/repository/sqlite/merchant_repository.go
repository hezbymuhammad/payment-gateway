package sqlite

import (
	"context"
        "database/sql"

	"github.com/hezbymuhammad/payment-gateway/domain"
)

type sqliteMerchantRepo struct {
	DB *sql.DB
}

func NewMerchantRepository(db *sql.DB) domain.MerchantRepository {
        return &sqliteMerchantRepo{
		DB: db,
	}
}

func (mr *sqliteMerchantRepo) IsAuthorizedParent(ctx context.Context, mg *domain.MerchantGroup) (bool, error) {
        query := "SELECT EXISTS (SELECT 1 FROM merchant_groups WHERE parent_merchant_id=? AND child_merchant_id=? LIMIT 1) as authorized"

        rows, err := mr.DB.Query(query, mg.ParentMerchantID, mg.ChildMerchantID)
        if err != nil {
                return false, err
        }
        defer rows.Close()

        var isExists string
        rows.Next()
        rows.Scan(&isExists)

        if isExists != "0" {
                return true, nil
        }
        return false, nil
}

func (mr *sqliteMerchantRepo) Store(ctx context.Context, m *domain.Merchant) error {
        query := "INSERT merchants SET name=?"

        stmt, err := mr.DB.PrepareContext(ctx, query)
        if err != nil {
                return err
        }

        res, err := stmt.ExecContext(ctx, m.Name)
        if err != nil {
                return err
        }

        lastID, err := res.LastInsertId()
        if err != nil {
                return err
        }

        m.ID = lastID
        return nil
}

func (mr *sqliteMerchantRepo) SetChild(ctx context.Context, mg *domain.MerchantGroup) error {
        query := "INSERT INTO merchant_groups SET parent_merchant_id=?, child_merchant_id=?"

        stmt, err := mr.DB.PrepareContext(ctx, query)
        if err != nil {
                return err
        }

        _, err = stmt.ExecContext(ctx, mg.ParentMerchantID, mg.ChildMerchantID)
        if err != nil {
                return err
        }

        return nil
}
