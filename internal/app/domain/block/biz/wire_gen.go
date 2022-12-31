// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package biz

import (
	"github.com/blackhorseya/ethscan/internal/app/domain/block/biz/repo"
	"github.com/blackhorseya/ethscan/pkg/entity/domain/activity/s2s"
	"github.com/blackhorseya/ethscan/pkg/entity/domain/block/biz"
	"github.com/google/wire"
)

// Injectors from wire.go:

func CreateBiz(repo2 repo.IRepo, activity s2s.ServiceClient) biz.IBiz {
	iBiz := NewImpl(repo2, activity)
	return iBiz
}

// wire.go:

var testProviderSet = wire.NewSet(NewImpl)
