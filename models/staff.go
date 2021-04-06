package models

import (
	"errors"

	"github.com/dgraph-io/badger/v3"
	protobuf "github.com/golang/protobuf/proto"
)

func NewStaff(id string) *Staff {
	return &Staff{
		Id: id,
	}
}

func NewStaffFromBytes(bs []byte) *Staff {
	s := &Staff{}
	err := protobuf.Unmarshal(bs, s)
	if err != nil {
		panic("invalid staff")
	}
	return s
}

func (d *DB) CreateStaff(id string) (s *Staff, err error) {
	txn := d.db.NewTransaction(true)
	defer txn.Discard()

	s = &Staff{Id: id}

	bs, err := protobuf.Marshal(s)
	if err != nil {
		return
	}

	err = txn.Set(FormatStaffKey(id), bs)
	if err != nil {
		return
	}

	err = txn.Commit()
	if err != nil {
		return
	}
	return
}

func (d *DB) GetStaff(id string) (s *Staff, err error) {
	txn := d.db.NewTransaction(false)
	defer txn.Discard()

	item, err := txn.Get(FormatStaffKey(id))
	if err != nil {
		// handle not found error manually
		if errors.Is(err, badger.ErrKeyNotFound) {
			return nil, NewNotFoundErr(id)
		}
		return nil, err
	}
	err = item.Value(func(val []byte) error {
		s = NewStaffFromBytes(val)
		return nil
	})
	return
}

func (d *DB) ListStaffs() ([]*Staff, error) {
	panic("implement me")
}
