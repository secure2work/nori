package main

import (
	"context"

	"github.com/nori-io/nori-common/config"
	"github.com/nori-io/nori-common/meta"
	pl "github.com/nori-io/nori-common/plugin"
)

var Plugin plugin

type plugin struct{}

func (p plugin) Meta() meta.Meta {
	return meta.Data{}
}

func (p plugin) Instance() interface{} {
	return nil
}

func (p plugin) Init(ctx context.Context, config config.Manager) error {
	return nil
}

func (p plugin) Start(ctx context.Context, registry pl.Registry) error {
	return nil
}

func (p plugin) Stop(ctx context.Context, registry pl.Registry) error {
	return nil
}