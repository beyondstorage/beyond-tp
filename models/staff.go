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

	err = txn.Set(StaffKey(id), bs)
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

	item, err := txn.Get(StaffKey(id))
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

func (d *DB) InsertStaffTask(txn *badger.Txn, staffId, taskId string) (err error) {
	if txn == nil {
		txn = d.db.NewTransaction(true)
		defer func() {
			err = d.CloseTxn(txn, err)
		}()
	}

	err = txn.Set(StaffTaskKey(staffId, taskId), []byte(taskId))
	if err != nil {
		return
	}
	return
}

func (d *DB) NextStaffTask(txn *badger.Txn, staffId string) (taskId string, err error) {
	if txn == nil {
		txn = d.db.NewTransaction(false)
		defer func() {
			err = d.CloseTxn(txn, err)
		}()
	}

	it := txn.NewIterator(badger.DefaultIteratorOptions)
	defer it.Close()

	prefix := StaffTaskPrefix(staffId)

	for it.Seek(prefix); it.ValidForPrefix(prefix); {
		item := it.Item()
		err := item.Value(func(v []byte) error {
			taskId = string(v)
			return nil
		})
		if err != nil {
			return "", err
		} else {
			return taskId, nil
		}
	}
	return "", err
}

func (d *DB) ListStaffTasks(txn *badger.Txn, staffId string) (taskIds []string, err error) {
	if txn == nil {
		txn = d.db.NewTransaction(false)
		defer func() {
			err = d.CloseTxn(txn, err)
		}()
	}

	it := txn.NewIterator(badger.DefaultIteratorOptions)
	defer it.Close()

	prefix := StaffTaskPrefix(staffId)

	taskIds = make([]string, 0)
	for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
		item := it.Item()
		err := item.Value(func(v []byte) error {
			taskIds = append(taskIds, string(v))
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	return taskIds, err
}

func (d *DB) DeleteStaffTask(txn *badger.Txn, staffId, taskId string) (err error) {
	if txn == nil {
		txn = d.db.NewTransaction(false)
		defer func() {
			err = d.CloseTxn(txn, err)
		}()
	}

	err = txn.Delete(StaffTaskKey(staffId, taskId))
	if err != nil {
		return
	}
	return
}
