// Code generated from Pkl module `Kosmos`. DO NOT EDIT.
package kosmos

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type Kosmos interface {
	GetEnabled() bool

	GetMykey() string

	GetPostgres() Postgres

	GetSomething() any
}

var _ Kosmos = (*KosmosImpl)(nil)

type KosmosImpl struct {
	Enabled bool `pkl:"enabled"`

	Mykey string `pkl:"mykey"`

	Postgres Postgres `pkl:"postgres"`

	Something any `pkl:"something"`
}

func (rcv *KosmosImpl) GetEnabled() bool {
	return rcv.Enabled
}

func (rcv *KosmosImpl) GetMykey() string {
	return rcv.Mykey
}

func (rcv *KosmosImpl) GetPostgres() Postgres {
	return rcv.Postgres
}

func (rcv *KosmosImpl) GetSomething() any {
	return rcv.Something
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Kosmos
func LoadFromPath(ctx context.Context, path string) (ret Kosmos, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Kosmos
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (Kosmos, error) {
	var ret KosmosImpl
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}