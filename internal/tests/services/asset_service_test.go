package services_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Manuel1Aguilar/portf_mger/internal/models"
	"github.com/Manuel1Aguilar/portf_mger/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestCreateAsset(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	assetService := services.NewAssetService(db)

	asset := &models.Asset{
		Symbol:      "AAPL",
		Description: "Apple Inc.",
		AssetType:   "Stock",
	}

	mock.ExpectQuery(regexp.QuoteMeta("SELECT EXISTS(SELECT 1 FROM asset WHERE symbol = ?)")).
		WithArgs(asset.Symbol).
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO asset (symbol, description, asset_type) VALUES (?, ?, ?)")).
		WithArgs(asset.Symbol, asset.Description, asset.AssetType).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = assetService.CreateAsset(asset)
	assert.NoError(t, err)

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestGetAssetBySymbol(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	assetService := services.NewAssetService(db)

	symbol := "AAPL"
	mockAsset := &models.Asset{
		ID:          1,
		Symbol:      symbol,
		Description: "Apple Inc.",
		AssetType:   "Stock",
	}

	mock.ExpectQuery(regexp.QuoteMeta("SELECT a.id, a.symbol, a.description, a.asset_type FROM asset a WHERE a.symbol = ?")).
		WithArgs(symbol).
		WillReturnRows(sqlmock.NewRows([]string{"id", "symbol", "description", "asset_type"}).
			AddRow(mockAsset.ID, mockAsset.Symbol, mockAsset.Description, mockAsset.AssetType))

	asset, err := assetService.GetAssetBySymbol(symbol)
	assert.NoError(t, err)
	assert.Equal(t, mockAsset, asset)

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestGetAssets(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	assetService := services.NewAssetService(db)
	mockAssets := []models.Asset{
		{ID: 1, Symbol: "AAPL", Description: "Apple Inc.", AssetType: "Stock"},
		{ID: 2, Symbol: "MSFT", Description: "Microsoft Corporation", AssetType: "Stock"},
	}

	rows := sqlmock.NewRows([]string{"id", "symbol", "description", "asset_type"})
	for _, asset := range mockAssets {
		rows.AddRow(asset.ID, asset.Symbol, asset.Description, asset.AssetType)
	}
	mock.ExpectQuery(regexp.QuoteMeta("SELECT a.id, a.symbol, a.description, a.asset_type FROM asset a")).WillReturnRows(rows)

	assets, err := assetService.GetAssets()
	assert.NoError(t, err)
	assert.Equal(t, mockAssets, assets)

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("There were unfulfilled expectations: %s", err)
	}
}
