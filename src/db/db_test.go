package db_test

import (
	"errors"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sauce-service/src/db"
)

// Utilitaire pour créer un gorm.DB mocké et un sqlmock.DB
func setupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, func()) {
	sqlDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	assert.NoError(t, err)
	cleanup := func() { sqlDB.Close() }
	return gormDB, mock, cleanup
}

func TestConnectWith_Success(t *testing.T) {
	gormDB, mock, cleanup := setupMockDB(t)
	defer cleanup()
	mock.ExpectPing()
	handler, err := db.ConnectWith(
		func() (string, error) { return "mock-dsn", nil },
		func(dsn string) (*gorm.DB, error) { return gormDB, nil },
	)
	assert.NoError(t, err)
	assert.NotNil(t, handler)
	assert.NotNil(t, handler.GetDB())
}

func TestConnectWith_FailsOnDSN(t *testing.T) {
	handler, err := db.ConnectWith(
		func() (string, error) { return "", errors.New("dsn error") },
		func(dsn string) (*gorm.DB, error) {
			t.Fatal("opener should not be called")
			return nil, nil
		},
	)
	assert.Nil(t, handler)
	assert.EqualError(t, err, "dsn error")
}

func TestConnectWith_FailsOnOpen(t *testing.T) {
	handler, err := db.ConnectWith(
		func() (string, error) { return "mock-dsn", nil },
		func(dsn string) (*gorm.DB, error) {
			return nil, errors.New("open error")
		},
	)
	assert.Nil(t, handler)
	assert.EqualError(t, err, "open error")
}

func TestDBHandler_Close(t *testing.T) {
	gormDB, mock, cleanup := setupMockDB(t)
	defer cleanup()
	mock.ExpectClose()
	handler, err := db.ConnectWith(
		func() (string, error) { return "mock-dsn", nil },
		func(dsn string) (*gorm.DB, error) { return gormDB, nil },
	)
	assert.NoError(t, err)
	err = handler.Close()
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
