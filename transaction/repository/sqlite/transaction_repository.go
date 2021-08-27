package sqlite

import (
	"context"
        "database/sql"
        "log"

	"github.com/hezbymuhammad/payment-gateway/domain"
)

type sqliteTransactionRepo struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) domain.TransactionRepository {
        return &sqliteTransactionRepo{
                DB: db,
        }
}

func (tr *sqliteTransactionRepo) GetByID(ctx context.Context, id int64) (domain.Transaction, error) {
        query := "SELECT id, merchant_id, parent_merchant_id, setting_id, status FROM transactions WHERE id=? LIMIT 1"

        rows, err := tr.DB.Query(query, id)
        if err != nil {
                log.Println(query)
                log.Println(err)
                return domain.Transaction{}, err
        }
        defer rows.Close()

        var rawStatus int
        data := domain.Transaction{}

        rows.Next()
        err = rows.Scan(
                &data.ID,
                &data.MerchantID,
                &data.ParentMerchantID,
                &data.SettingID,
                &rawStatus,
        )
        if err != nil {
                log.Println(query)
                log.Println(err)
                return domain.Transaction{}, err
        }

        data.Status = rawStatus != 0

        return data, nil
}
func (tr *sqliteTransactionRepo) Store(ctx context.Context, t *domain.Transaction) error {
        query := "INSERT INTO transactions (merchant_id, parent_merchant_id, setting_id, status) values (?, ?, ?, ?)"

        stmt, err := tr.DB.PrepareContext(ctx, query)
        if err != nil {
                log.Println(query)
                log.Println(err)
                return err
        }

        status := btoi(t.Status)
        res, err := stmt.ExecContext(ctx, t.MerchantID, t.ParentMerchantID, t.SettingID, status)
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

        t.ID = lastID

        return nil

}
func (tr *sqliteTransactionRepo) Update(ctx context.Context, t *domain.Transaction) error {
        query := "UPDATE transactions SET merchant_id=?, parent_merchant_id=?, setting_id=?, status=? WHERE id=?"

        stmt, err := tr.DB.PrepareContext(ctx, query)
        if err != nil {
                log.Println(query)
                log.Println(err)
                return err
        }

        _, err = stmt.ExecContext(ctx, t.MerchantID, t.ParentMerchantID, t.SettingID, t.Status, t.ID)
        if err != nil {
                log.Println(query)
                log.Println(err)
                return err
        }

        return nil
}

func btoi(b bool) int {
    if b {
        return 1
    }
    return 0
 }
