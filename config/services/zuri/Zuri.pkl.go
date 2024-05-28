// Code generated from Pkl module `Zuri`. DO NOT EDIT.
package zuri

import (
	"context"

	"bitbucket.org/zetaactions/pklpoc/config/lib/postgres"
	"bitbucket.org/zetaactions/pklpoc/config/lib/redis"
	"github.com/apple/pkl-go/pkl"
)

type Zuri struct {
	Enabled bool `pkl:"enabled"`

	Postgres *postgres.ConnectionDetails `pkl:"postgres"`

	Redis *redis.ConnectionDetails `pkl:"redis"`

	ZuriRealmKey string `pkl:"zuri_realm_key"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Zuri
func LoadFromPath(ctx context.Context, path string) (ret *Zuri, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Zuri
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*Zuri, error) {
	var ret Zuri
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
