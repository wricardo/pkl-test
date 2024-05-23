// Code generated from Pkl module `formation`. DO NOT EDIT.
package pkltest

import (
	"context"

	"bitbucket.org/zetaactions/pkltest/services/kosmos"
	"bitbucket.org/zetaactions/pkltest/services/zuri"
	"github.com/apple/pkl-go/pkl"
)

type Formation struct {
	Zuri zuri.Zuri `pkl:"zuri"`

	Kosmos kosmos.Kosmos `pkl:"kosmos"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Formation
func LoadFromPath(ctx context.Context, path string) (ret *Formation, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Formation
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*Formation, error) {
	var ret Formation
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
