package badgerkv

import (
	"github.com/dalloriam/ocl/kv"
	badger "github.com/dgraph-io/badger/v4"
)

type BadgerKV struct {
	db *badger.DB
}

func NewBadgerKV(dbPath string) (*BadgerKV, error) {
	opts := badger.DefaultOptions(dbPath).WithLogger(nil)
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}
	return &BadgerKV{db: db}, nil
}

func (b *BadgerKV) Get(key string) ([]byte, error) {
	var valCopy []byte
	err := b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			if err == badger.ErrKeyNotFound {
				return kv.ErrKeyNotFound
			}
			return err
		}
		val, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}
		valCopy = val
		return nil
	})
	if err != nil {
		return nil, err
	}
	return valCopy, nil
}

func (b *BadgerKV) Set(key string, value []byte) error {
	err := b.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), value)
	})
	return err
}

func (b *BadgerKV) Delete(key string) error {
	err := b.db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	})
	return err
}

func (b *BadgerKV) Clear() error {
	err := b.db.DropAll()
	return err
}

func (b *BadgerKV) Close() error {
	return b.db.Close()
}
