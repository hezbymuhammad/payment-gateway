package sqlite_test

import (
        "context"
        "fmt"
	"testing"
        "regexp"

        "github.com/stretchr/testify/assert"
        sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/hezbymuhammad/payment-gateway/domain"
	transactionRepo "github.com/hezbymuhammad/payment-gateway/transaction/repository/sqlite"
)

func TestGetByIDSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
        if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

        data := domain.Transaction{
                ID: 1,
                MerchantID: 1,
                ParentMerchantID: 2,
                SettingID: 1,
                Status: true,
        }

        rows := sqlmock.NewRows([]string{"id", "merchant_id", "parent_merchant_id", "setting_id", "status"}).AddRow(data.ID, data.MerchantID, data.ParentMerchantID, data.SettingID, 1)
        query := regexp.QuoteMeta("SELECT id, merchant_id, parent_merchant_id, setting_id, status FROM transactions WHERE id=? LIMIT 1")

        mock.ExpectQuery(query).WillReturnRows(rows)
        tr := transactionRepo.NewTransactionRepository(db)

        res, err := tr.GetByID(context.TODO(), 1)
        assert.NoError(t, err)
        assert.Equal(t, res, data)
}

func TestGetByIDError(t *testing.T) {
	db, mock, err := sqlmock.New()
        if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

        query := regexp.QuoteMeta("SELECT id, merchant_id, parent_merchant_id, setting_id, status FROM transactions WHERE id=? LIMIT 1")

        mock.ExpectQuery(query).WillReturnError(fmt.Errorf("some error"))
        tr := transactionRepo.NewTransactionRepository(db)

        _, err = tr.GetByID(context.TODO(), 1)
        assert.Error(t, err)
}

func TestGetByIDNoRows(t *testing.T) {
	db, mock, err := sqlmock.New()
        if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

        rows := sqlmock.NewRows([]string{"id", "merchant_id", "parent_merchant_id", "setting_id", "status"})
        query := regexp.QuoteMeta("SELECT id, merchant_id, parent_merchant_id, setting_id, status FROM transactions WHERE id=? LIMIT 1")

        mock.ExpectQuery(query).WillReturnRows(rows)
        tr := transactionRepo.NewTransactionRepository(db)

        res, err := tr.GetByID(context.TODO(), 1)
        assert.Equal(t, res, domain.Transaction{})
}

func TestStoreSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
        if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

        data := &domain.Transaction{
                MerchantID: 1,
                ParentMerchantID: 2,
                SettingID: 1,
                Status: false,
        }
        query := regexp.QuoteMeta("INSERT INTO transactions (merchant_id, parent_merchant_id, setting_id, status) values (?, ?, ?, ?)")

        prep := mock.ExpectPrepare(query)
        prep.ExpectExec().WithArgs(data.MerchantID, data.ParentMerchantID, data.SettingID, 0).WillReturnResult(sqlmock.NewResult(12, 1))
        tr := transactionRepo.NewTransactionRepository(db)

        err = tr.Store(context.TODO(), data)
        assert.NoError(t, err)
        assert.Equal(t, data.ID, int64(12))
}

func TestStoreError(t *testing.T) {
	db, mock, err := sqlmock.New()
        if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

        data := &domain.Transaction{
                MerchantID: 1,
                ParentMerchantID: 2,
                SettingID: 1,
                Status: false,
        }
        query := regexp.QuoteMeta("INSERT INTO transactions (merchant_id, parent_merchant_id, setting_id, status) values (?, ?, ?, ?)")

        prep := mock.ExpectPrepare(query)
        prep.ExpectExec().WithArgs(data.MerchantID, data.ParentMerchantID, data.SettingID, 0).WillReturnError(fmt.Errorf("some error"))
        tr := transactionRepo.NewTransactionRepository(db)

        err = tr.Store(context.TODO(), data)
        assert.Error(t, err)
}

func TestUpdateSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
        if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

        data := &domain.Transaction{
                ID: 1,
                MerchantID: 1,
                ParentMerchantID: 2,
                SettingID: 1,
                Status: true,
        }
        query := regexp.QuoteMeta("UPDATE transactions SET merchant_id=?, parent_merchant_id=?, setting_id=?, status=? WHERE id=?")

        prep := mock.ExpectPrepare(query)
        prep.ExpectExec().WithArgs(data.MerchantID, data.ParentMerchantID, data.SettingID, data.Status, data.ID).WillReturnResult(sqlmock.NewResult(12, 1))
        tr := transactionRepo.NewTransactionRepository(db)

        err = tr.Update(context.TODO(), data)
        assert.NoError(t, err)
}

func TestUpdateError(t *testing.T) {
	db, mock, err := sqlmock.New()
        if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

        data := &domain.Transaction{
                ID: 1,
                MerchantID: 1,
                ParentMerchantID: 2,
                SettingID: 1,
                Status: true,
        }
        query := regexp.QuoteMeta("UPDATE transactions SET merchant_id=?, parent_merchant_id=?, setting_id=?, status=? WHERE id=?")

        prep := mock.ExpectPrepare(query)
        prep.ExpectExec().WithArgs(data.MerchantID, data.ParentMerchantID, data.SettingID, data.Status, data.ID).WillReturnError(fmt.Errorf("some error"))
        tr := transactionRepo.NewTransactionRepository(db)

        err = tr.Update(context.TODO(), data)
        assert.Error(t, err)
}
