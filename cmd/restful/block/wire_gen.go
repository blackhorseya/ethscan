// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/blackhorseya/ethscan/internal/entity/domain/block/biz"
	"github.com/blackhorseya/ethscan/internal/entity/domain/block/biz/repo"
	"github.com/blackhorseya/ethscan/internal/pkg/config"
	"github.com/blackhorseya/ethscan/internal/pkg/log"
	"github.com/blackhorseya/ethscan/internal/pkg/storage/mariadb"
	"github.com/blackhorseya/ethscan/internal/pkg/transports/grpcx"
	"github.com/blackhorseya/ethscan/internal/pkg/transports/httpx"
	"github.com/blackhorseya/ethscan/internal/pkg/transports/kafka"
	"github.com/blackhorseya/ethscan/pkg/app"
	"github.com/google/wire"
)

// Injectors from wire.go:

func CreateService(path2 string, id int64) (app.Service, error) {
	viper, err := config.NewConfig(path2)
	if err != nil {
		return nil, err
	}
	options, err := log.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	logger, err := log.NewLogger(options)
	if err != nil {
		return nil, err
	}
	httpxOptions, err := httpx.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	engine := httpx.NewRouter(httpxOptions)
	server := httpx.NewServer(httpxOptions, logger, engine)
	nodeOptions, err := repo.NewNodeOptions(viper)
	if err != nil {
		return nil, err
	}
	mariadbOptions, err := mariadb.NewOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	db, err := mariadb.NewMariadb(mariadbOptions, logger)
	if err != nil {
		return nil, err
	}
	producerOptions, err := kafka.NewProducerOptions(viper)
	if err != nil {
		return nil, err
	}
	producer, err := kafka.NewProducer(producerOptions)
	if err != nil {
		return nil, err
	}
	iRepo, err := repo.NewImpl(nodeOptions, db, producer)
	if err != nil {
		return nil, err
	}
	clientOptions, err := grpcx.NewClientOptions(viper)
	if err != nil {
		return nil, err
	}
	client := grpcx.NewClient(clientOptions)
	serviceClient, err := biz.NewActivityClient(client)
	if err != nil {
		return nil, err
	}
	iBiz := biz.NewImpl(iRepo, serviceClient)
	adaptersRestful := NewRestful(logger, engine, iBiz)
	appService, err := NewService(logger, server, adaptersRestful)
	if err != nil {
		return nil, err
	}
	return appService, nil
}

// wire.go:

var providerSet = wire.NewSet(config.ProviderSet, log.ProviderSet, mariadb.ProviderSet, httpx.ProviderServerSet, kafka.ProviderProducer, grpcx.ProviderClient, biz.ProviderSet, NewService,
	NewRestful,
)
