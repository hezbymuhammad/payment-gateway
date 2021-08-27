package sqlite_test

import (
        "context"
	"testing"
        "regexp"

        "github.com/stretchr/testify/assert"
        sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/hezbymuhammad/payment-gateway/domain"
	merchantRepo "github.com/hezbymuhammad/payment-gateway/merchant/repository/sqlite"
)

func TestIsAuthorizedParentSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
        if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

        rows := sqlmock.NewRows([]string{"authorized"}).AddRow(1)
        query := regexp.QuoteMeta("SELECT EXISTS (SELECT 1 FROM merchant_groups WHERE parent_merchant_id=? AND child_merchant_id=? LIMIT 1) as authorized")

        mock.ExpectQuery(query).WillReturnRows(rows)
        m := merchantRepo.NewMerchantRepository(db)
        data := &domain.MerchantGroup{
                ParentMerchantID: 1,
                ChildMerchantID: 2,
        }

        res, err := m.IsAuthorizedParent(context.TODO(), data)
        assert.NoError(t, err)
        assert.Equal(t, res, true)
}

func TestIsAuthorizedParentNotAuthorized(t *testing.T) {
	db, mock, err := sqlmock.New()
        if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

        rows := sqlmock.NewRows([]string{"authorized"}).AddRow(0)
        query := regexp.QuoteMeta("SELECT EXISTS (SELECT 1 FROM merchant_groups WHERE parent_merchant_id=? AND child_merchant_id=? LIMIT 1) as authorized")

        mock.ExpectQuery(query).WillReturnRows(rows)
        m := merchantRepo.NewMerchantRepository(db)
        data := &domain.MerchantGroup{
                ParentMerchantID: 1,
                ChildMerchantID: 2,
        }

        res, err := m.IsAuthorizedParent(context.TODO(), data)
        assert.NoError(t, err)
        assert.Equal(t, res, false)
}

func TestStore(t *testing.T) {
	db, mock, err := sqlmock.New()
        if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

        m := &domain.Merchant{Name: "lorem"}

        query := regexp.QuoteMeta("INSERT INTO merchants(name) VALUES(?)")

        prep := mock.ExpectPrepare(query)
        prep.ExpectExec().WithArgs(m.Name).WillReturnResult(sqlmock.NewResult(12, 1))
        mr := merchantRepo.NewMerchantRepository(db)

        err = mr.Store(context.TODO(), m)
        assert.NoError(t, err)
        assert.Equal(t, m.ID, int64(12))
}

func TestSetChild(t *testing.T) {
	db, mock, err := sqlmock.New()
        if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

        query := regexp.QuoteMeta("INSERT INTO merchant_groups(parent_merchant_id, child_merchant_id) VALUES(?, ?)")

        prep := mock.ExpectPrepare(query)
        prep.ExpectExec().WithArgs(1, 2).WillReturnResult(sqlmock.NewResult(12, 1))
        mr := merchantRepo.NewMerchantRepository(db)
        data := &domain.MerchantGroup{
                ParentMerchantID: 1,
                ChildMerchantID: 2,
        }

        err = mr.SetChild(context.TODO(), data)
        assert.NoError(t, err)
}
