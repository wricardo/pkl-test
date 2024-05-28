// Code generated from Pkl module `Config`. DO NOT EDIT.
package config

import (
	"context"

	"bitbucket.org/zetaactions/pklpoc/config/services/kosmos"
	"bitbucket.org/zetaactions/pklpoc/config/services/zasper"
	"bitbucket.org/zetaactions/pklpoc/config/services/zuri"
	"github.com/apple/pkl-go/pkl"
)

type Config struct {
	Zuri *zuri.Zuri `pkl:"zuri"`

	Kosmos *kosmos.Kosmos `pkl:"kosmos"`

	Zasper *zasper.Zasper `pkl:"zasper"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Config
func LoadFromPath(ctx context.Context, path string) (ret *Config, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Config
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*Config, error) {
	var ret Config
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
