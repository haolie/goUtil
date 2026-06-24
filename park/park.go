package park

import (
	"context"
	"fmt"
)

type parkFun func(ctx context.Context) error

var (
	beforeStartMap = make(map[string]parkFun)
	startMap       = make(map[string]parkFun)
	stopMap        = make(map[string]parkFun)
)

func RegisterBeforeStart(k string, fn parkFun) {
	if _, exists := beforeStartMap[k]; exists {
		panic(fmt.Sprintf("repeated register before starting %s", k))
	}

	beforeStartMap[k] = fn
}

func RegisterStart(k string, fn parkFun) {
	if _, exists := startMap[k]; exists {
		panic(fmt.Sprintf("repeated register start %s", k))
	}

	startMap[k] = fn
}

func RegisterStop(k string, fn parkFun) {
	if _, exists := stopMap[k]; exists {
		panic(fmt.Sprintf("repeated register stop %s", k))
	}

	stopMap[k] = fn
}

func RunBeforeStart(ctx context.Context) []error {
	errs := make([]error, 0, 4)
	for _, fn := range beforeStartMap {
		errs = append(errs, fn(ctx))
	}

	return errs
}

func RunStart(ctx context.Context) []error {
	errs := make([]error, 0, 4)
	for _, fn := range startMap {
		errs = append(errs, fn(ctx))
	}

	return errs
}

func RunStop(ctx context.Context) []error {
	errs := make([]error, 0, 4)
	for _, fn := range stopMap {
		errs = append(errs, fn(ctx))
	}

	return errs
}
