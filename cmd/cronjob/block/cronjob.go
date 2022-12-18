package main

import (
	"time"

	"github.com/blackhorseya/portto/pkg/adapters"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Options declare cronjob configuration
type Options struct {
	Enabled   bool          `json:"enabled" yaml:"enabled"`
	Interval  time.Duration `json:"interval" yaml:"interval"`
	NeedDepth int           `json:"need_depth" yaml:"needDepth"`
}

// NewCronjobOptions serve caller to create an Options
func NewCronjobOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	ret := new(Options)

	err := v.UnmarshalKey("cronjob", &ret)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load cronjob options")
	}

	logger.Info("load cronjob options success")

	return ret, nil
}

type impl struct {
	opts   *Options
	logger *zap.Logger

	initHeight uint64

	taskC chan time.Time
	done  chan bool
}

// NewCronjob create a cronjob instance
func NewCronjob(opts *Options, logger *zap.Logger, initHeight uint64) adapters.Cronjob {
	return &impl{
		opts:       opts,
		logger:     logger,
		initHeight: initHeight,
		taskC:      make(chan time.Time, 1),
		done:       make(chan bool),
	}
}

func (i *impl) Start() error {
	if !i.opts.Enabled {
		return nil
	}

	i.logger.Info("starting cronjob engine...")

	go i.worker()

	return nil
}

func (i *impl) Stop() error {
	if !i.opts.Enabled {
		return nil
	}

	i.logger.Info("stopping cronjob engine...")

	i.done <- true

	return nil
}

func (i *impl) worker() {
	ticker := time.NewTicker(i.opts.Interval)

	for {
		select {
		case <-i.done:
			return
		case <-ticker.C:
			i.executeTo()
		case <-i.taskC:
			// todo: 2022/12/18|sean|impl me
		}
	}
}

func (i *impl) executeTo() {
	select {
	case i.taskC <- time.Now():
	case <-time.After(1 * time.Second):
		return
	}
}
