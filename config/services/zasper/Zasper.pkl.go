// Code generated from Pkl module `Zasper`. DO NOT EDIT.
package zasper

import (
	"context"

	"bitbucket.org/zetaactions/pklpoc/config/lib/postgres"
	"github.com/apple/pkl-go/pkl"
)

type Zasper struct {
	Enabled bool `pkl:"enabled"`

	ZasperSomething string `pkl:"zasper_something"`

	SomeArray []any `pkl:"some_array"`

	Poopf *ZasperXYZ `pkl:"poopf"`

	Redises []*postgres.ConnectionDetails `pkl:"redises"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Zasper
func LoadFromPath(ctx context.Context, path string) (ret *Zasper, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Zasper
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*Zasper, error) {
	var ret Zasper
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
