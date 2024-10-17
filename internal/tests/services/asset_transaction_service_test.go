package services_test

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Manuel1Aguilar/portf_mger/internal/models"
	"github.com/Manuel1Aguilar/portf_mger/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	portfolioHoldingService := services.NewPortfolioHoldingService(db)
	assetTransactionService := services.NewAssetTransactionService(db, portfolioHoldingService)

	assetTransaction := &models.AssetTransaction{
		AssetID:         1,
		TransactionType: "BUY",
		ValueUSD:        1.0,
		Units:           1,
		UnitPrice:       1,
		DateTransacted:  time.Now(),
	}
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO asset_transaction(asset_id, transaction_type, valueUSD, units, unit_price, date_transacted) VALUES (?, ?, ?, ?, ?, ?)")).
		WithArgs(assetTransaction.AssetID,
			assetTransaction.TransactionType,
			assetTransaction.ValueUSD,
			assetTransaction.Units,
			assetTransaction.UnitPrice,
			assetTransaction.DateTransacted).WillReturnResult(sqlmock.NewResult(1, 1))

	err = assetTransactionService.Add(assetTransaction)
	assert.NoError(t, err)

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("There were unfulfilled expectations: %s", err)
	}
}
