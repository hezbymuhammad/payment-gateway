package sqlite

import (
	"context"
        "database/sql"
        "log"

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
                log.Println(query)
                log.Println(err)
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
        query := "INSERT INTO merchants(name) VALUES(?)"

        stmt, err := mr.DB.PrepareContext(ctx, query)
        if err != nil {
                log.Println(query)
                log.Println(err)
                return err
        }

        res, err := stmt.ExecContext(ctx, m.Name)
        if err != nil {
                log.Println(query)
                log.Println(err)
                return err
        }

        lastID, err := res.LastInsertId()
        if err != nil {
                log.Println(query)
                log.Println(err)
                return err
        }

        m.ID = lastID
        return nil
}

func (mr *sqliteMerchantRepo) InitSetting(ctx context.Context, m *domain.Merchant) error {
        query := "INSERT INTO settings(merchant_id, color, payment_type, payment_name) VALUES(?, ?, ?, ?)"

        stmt, err := mr.DB.PrepareContext(ctx, query)
        if err != nil {
                log.Println(query)
                log.Println(err)
                return err
        }

        _, err = stmt.ExecContext(ctx, m.ID, "RED", "CARD", "VISA")
        if err != nil {
                log.Println(query)
                log.Println(err)
                return err
        }
        return nil
}

func (mr *sqliteMerchantRepo) SetChild(ctx context.Context, mg *domain.MerchantGroup) error {
        query := "INSERT INTO merchant_groups(parent_merchant_id, child_merchant_id) VALUES(?, ?)"

        stmt, err := mr.DB.PrepareContext(ctx, query)
        if err != nil {
                log.Println(query)
                log.Println(err)
                return err
        }

        _, err = stmt.ExecContext(ctx, mg.ParentMerchantID, mg.ChildMerchantID)
        if err != nil {
                log.Println(query)
                log.Println(err)
                return err
        }

        return nil
}
