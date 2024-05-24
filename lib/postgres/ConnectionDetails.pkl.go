// Code generated from Pkl module `lib.postgres.ConnectionDetails`. DO NOT EDIT.
package postgres

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type ConnectionDetails struct {
	Username string `pkl:"username"`

	Password string `pkl:"password"`

	Host string `pkl:"host"`

	Port int `pkl:"port"`

	DbName string `pkl:"dbName"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a ConnectionDetails
func LoadFromPath(ctx context.Context, path string) (ret *ConnectionDetails, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a ConnectionDetails
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*ConnectionDetails, error) {
	var ret ConnectionDetails
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
