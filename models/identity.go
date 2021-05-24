package models

import (
	"errors"

	"github.com/dgraph-io/badger/v3"
	protobuf "github.com/golang/protobuf/proto"
)

func NewIdentityFromBytes(bs []byte) *Identity {
	id := &Identity{}
	err := protobuf.Unmarshal(bs, id)
	if err != nil {
		panic("invalid identity")
	}
	return id
}

// InsertIdentity insert an Identity into db
func (d *DB) InsertIdentity(txn *badger.Txn, i *Identity) (err error) {
	if txn == nil {
		txn = d.db.NewTransaction(true)
		defer func() {
			err = d.CloseTxn(txn, err)
		}()
	}

	bs, err := protobuf.Marshal(i)
	if err != nil {
		return err
	}

	err = txn.Set(IdentityKey(i.Type, i.Name), bs)
	if err != nil {
		return err
	}

	return
}

func (d *DB) GetIdentity(txn *badger.Txn, idType IdentityType, name string) (id *Identity, err error) {
	if txn == nil {
		txn = d.db.NewTransaction(false)
		defer func() {
			err = d.CloseTxn(txn, err)
		}()
	}

	key := IdentityKey(idType, name)
	item, err := txn.Get(key)
	if err != nil {
		// handle not found error manually
		if errors.Is(err, badger.ErrKeyNotFound) {
			return nil, NewNotFoundErr(string(key))
		}
		return nil, err
	}
	err = item.Value(func(val []byte) error {
		id = NewIdentityFromBytes(val)
		return nil
	})
	return
}

// DeleteIdentity delete an Identity by given type and name from DB
func (d *DB) DeleteIdentity(txn *badger.Txn, idType IdentityType, name string) (err error) {
	if txn == nil {
		txn = d.db.NewTransaction(true)
		defer func() {
			err = d.CloseTxn(txn, err)
		}()
	}

	// TODO: we will need to check identity first.
	if err = txn.Delete(IdentityKey(idType, name)); err != nil {
		return err
	}
	return
}

func (d *DB) ListIdentity(idType *IdentityType) ([]*Identity, error) {
	txn := d.db.NewTransaction(false)
	defer txn.Discard()
	it := txn.NewIterator(badger.DefaultIteratorOptions)
	defer it.Close()

	ids := make([]*Identity, 0)

	prefix := IdentityKeyPrefix(idType)
	for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
		item := it.Item()
		err := item.Value(func(v []byte) error {
			ids = append(ids, NewIdentityFromBytes(v))
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	return ids, nil
}
