package config

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

// LoadFromHTTP loads the pkl module at the given path and evaluates it into a Formation
func LoadFromHTTP(ctx context.Context, url string) (ret *Config, err error) {
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
	ret, err = Load(ctx, evaluator, pkl.UriSource(url))
	return ret, err
}
