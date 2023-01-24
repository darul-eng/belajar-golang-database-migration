//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func InitializedService() (*SimpleService, error) {
	wire.Build(NewSimpleService, NewSimpleRepository)
	return nil, nil
}
