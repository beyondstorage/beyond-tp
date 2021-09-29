package task

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"path/filepath"
	"sync"

	"github.com/dgraph-io/badger/v3"
)

type Server struct {
	logger *zap.Logger

	db *badger.DB

	// task id -> job channel.
	jobCh   map[string]chan *Job
	jobLock sync.Mutex
}

type ServerConfig struct {
	DataDir string
	Logger  *zap.Logger
}

func NewServer(sc *ServerConfig) (s *Server, err error) {
	s = &Server{
		logger: sc.Logger,
	}

	// Init badger db.
	dbDir := filepath.Join(sc.DataDir, "db")
	ops := badger.
		DefaultOptions(dbDir).
		WithLoggingLevel(badger.ERROR)
	db, err := badger.Open(ops)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}
	s.db = db

	s.jobCh = make(map[string]chan *Job)
	return
}

func (s *Server) Serve(ctx context.Context) (err error) {
	return nil
}

func (s *Server) GetTask(ctx context.Context, taskId string) (t *Task, err error) {
	txn := s.db.NewTransaction(false)
	defer func() {
		err = s.CloseTxn(txn, err)
	}()

	item, err := txn.Get(KeyTask(taskId))
	if err != nil && errors.Is(err, badger.ErrKeyNotFound) {
		s.logger.Debug("task not found",
			zap.String("task_id", taskId))
		return nil, nil
	}
	err = item.Value(func(val []byte) error {
		t = ParseTask(val)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("db value: %w", err)
	}
	s.logger.Debug("task found",
		zap.String("task_id", taskId))
	return
}

func (s *Server) InsertTask(ctx context.Context, t *Task) (err error) {
	txn := s.db.NewTransaction(true)
	defer func() {
		err = s.CloseTxn(txn, err)
	}()

	err = txn.Set(KeyTask(t.Id), FormatTask(t))
	if err != nil {
		return fmt.Errorf("db set: %w", err)
	}
	return nil
}

func (s *Server) NextTask(ctx context.Context) (t *Task, err error) {
	err = s.nextValue(PrefixTask, func(val []byte) error {
		t = ParseTask(val)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("db next value: %w", err)
	}
	return
}

func (s *Server) HasTask(ctx context.Context) (has bool, err error) {
	err = s.nextValue(PrefixTask, func(val []byte) error {
		has = true
		return nil
	})
	if err != nil {
		return false, fmt.Errorf("db next value: %w", err)
	}
	return has, nil
}

func (s *Server) DeleteTask(ctx context.Context, taskId string) (err error) {
	txn := s.db.NewTransaction(true)
	defer func() {
		err = s.CloseTxn(txn, err)
	}()

	err = txn.Delete(KeyTask(taskId))
	if err != nil {
		return
	}
	return
}

func (s *Server) NextJob(ctx context.Context, taskId string) (j *Job, err error) {
	err = s.nextValue(PrefixJob(taskId), func(val []byte) error {
		j = ParseJob(val)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("db next value: %w", err)
	}
	return
}

func (s *Server) InsertJob(ctx context.Context, j *Job) (err error) {
	txn := s.db.NewTransaction(true)
	defer func() {
		err = s.CloseTxn(txn, err)
	}()

	err = txn.Set(KeyJob(j.TaskId, j.Id), FormatJob(j))
	if err != nil {
		return fmt.Errorf("db set: %w", err)
	}
	return nil
}

func (s *Server) GetJob(ctx context.Context, taskId, jobId string) (j *Job, err error) {
	txn := s.db.NewTransaction(false)
	defer func() {
		err = s.CloseTxn(txn, err)
	}()

	item, err := txn.Get(KeyJob(taskId, jobId))
	if err != nil && errors.Is(err, badger.ErrKeyNotFound) {
		s.logger.Debug("job not found",
			zap.String("task_id", taskId),
			zap.String("job_id", jobId))
		return nil, nil
	}
	err = item.Value(func(val []byte) error {
		j = ParseJob(val)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("db value: %w", err)
	}
	s.logger.Debug("job found",
		zap.String("job_id", jobId))
	return
}

func (s *Server) WaitJob(ctx context.Context, taskId, jobId string) (err error) {
	j, err := s.GetJob(ctx, taskId, jobId)
	if err != nil {
		return fmt.Errorf("get job: %w", err)
	}
	// If job doesn't exist, we can return directly.
	if j == nil {
		return nil
	}

	prefix := KeyJob(taskId, jobId)
	isDeleted := false

	err = s.db.Subscribe(ctx, func(kv *badger.KVList) error {
		for _, v := range kv.Kv {
			// Check if prefix valid.
			if bytes.Compare(v.Key, KeyJob(taskId, jobId)) != 0 {
				panic(fmt.Errorf("prefix invalid, expected: %s, got: %s", prefix, v.Key))
			}
			if v.Value == nil {
				isDeleted = true
				return nil
			}
		}
		return nil
	}, prefix)
	if err != nil {
		return fmt.Errorf("db subscribe: %w", err)
	}
	if !isDeleted {
		s.logger.Warn("wait job exited without delete",
			zap.String("task_id", taskId),
			zap.String("job_id", jobId))
	}
	return
}

func (s *Server) DeleteJob(ctx context.Context, taskId, jobId string) (err error) {
	txn := s.db.NewTransaction(true)
	defer func() {
		err = s.CloseTxn(txn, err)
	}()

	err = txn.Delete(KeyJob(taskId, jobId))
	if err != nil {
		return
	}
	return
}

func (s *Server) SetMeta(ctx context.Context, taskId, jobId, meteKey, metaValue string) (err error) {
	txn := s.db.NewTransaction(true)
	defer func() {
		err = s.CloseTxn(txn, err)
	}()

	err = txn.Set(KeyMeta(taskId, jobId, meteKey), []byte(metaValue))
	if err != nil {
		return
	}
	return
}
func (s *Server) GetMeta(ctx context.Context, taskId, jobId, meteKey string) (metaValue string, err error) {
	txn := s.db.NewTransaction(false)
	defer func() {
		err = s.CloseTxn(txn, err)
	}()

	item, err := txn.Get(KeyMeta(taskId, jobId, meteKey))
	if err != nil {
		return
	}
	err = item.Value(func(val []byte) error {
		metaValue = string(val)
		return nil
	})
	if err != nil {
		return
	}
	return
}
func (s *Server) DeleteMeta(ctx context.Context, taskId, jobId, meteKey string) (err error) {
	txn := s.db.NewTransaction(true)
	defer func() {
		err = s.CloseTxn(txn, err)
	}()

	err = txn.Delete(KeyMeta(taskId, jobId, meteKey))
	if err != nil {
		return
	}
	return
}

// NewTxn export db.NewTransaction method
func (s *Server) NewTxn(update bool) *badger.Txn {
	return s.db.NewTransaction(update)
}

func (s *Server) CloseTxn(txn *badger.Txn, err error) error {
	// Discard all changes and return input error.
	if err != nil {
		txn.Discard()
		return fmt.Errorf("discard txn: %w", err)
	}

	// Commit all changes and return error during commit.
	err = txn.Commit()
	if err != nil {
		s.logger.Error("txn commit", zap.Error(err))
		return fmt.Errorf("commit txn: %w", err)
	}
	return nil
}

func (s *Server) nextValue(prefix []byte, fn func(val []byte) error) (err error) {
	txn := s.db.NewTransaction(false)
	defer func() {
		err = s.CloseTxn(txn, err)
	}()

	opt := badger.IteratorOptions{
		// Only next key to avoid over read.
		PrefetchSize:   1,
		PrefetchValues: true,
		Prefix:         prefix,
	}

	it := txn.NewIterator(opt)
	defer it.Close()

	it.Seek(prefix)
	if !it.ValidForPrefix(prefix) {
		return nil
	}

	err = it.Item().Value(fn)
	if err != nil {
		return fmt.Errorf("db value: %w", err)
	}
	return nil
}
