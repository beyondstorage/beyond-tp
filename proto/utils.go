package proto

import (
	"errors"

	fs "github.com/aos-dev/go-service-fs/v2"
	qingstor "github.com/aos-dev/go-service-qingstor/v2"
	"github.com/aos-dev/go-storage/v3/types"
	"github.com/google/uuid"
)

// ParseStorager parse endpoint into
func (x *Endpoint) ParseStorager() (types.Storager, error) {
	pairs := make([]types.Pair, 0, len(x.GetPairs()))
	for _, p := range x.GetPairs() {
		pairs = append(pairs, p.transform())
	}

	switch x.GetType() {
	case fs.Type:
		return fs.NewStorager(pairs...)
	case qingstor.Type:
		return qingstor.NewStorager(pairs...)
	default:
		return nil, errors.New("endpoint type unsupported")
	}
}

func (x *Pair) transform() types.Pair {
	return types.Pair{
		Key:   x.Key,
		Value: x.Value,
	}
}

func NewJob() Job {
	return Job{Id: uuid.NewString()}
}
