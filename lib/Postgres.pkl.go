// Code generated from Pkl module `Postgres`. DO NOT EDIT.
package lib

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type Postgres struct {
	Something string `pkl:"something"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Postgres
func LoadFromPath(ctx context.Context, path string) (ret *Postgres, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Postgres
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*Postgres, error) {
	var ret Postgres
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
