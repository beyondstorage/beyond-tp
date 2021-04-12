package models

import (
	"github.com/dgraph-io/badger/v3"
	"go.uber.org/zap"
)

type DB struct {
	db     *badger.DB
	logger *zap.Logger
}

func NewDB(path string, logger *zap.Logger) (*DB, error) {
	db, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		return nil, err
	}
	return &DB{db: db, logger: logger}, nil
}

func (d *DB) Close() (err error) {
	return d.db.Close()
}

func (d *DB) CloseTxn(txn *badger.Txn, err error) error {
	// Discard all changes and return input error.
	if err != nil {
		txn.Discard()
		return err
	}

	// Commit all changes and return error during commit.
	err = txn.Commit()
	if err != nil {
		d.logger.Error("txn commit", zap.Error(err))
		return err
	}
	return nil
}
