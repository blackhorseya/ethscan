package main

import (
	"fmt"
	"time"

	"github.com/blackhorseya/portto/pkg/adapters"
	"github.com/blackhorseya/portto/pkg/contextx"
	bb "github.com/blackhorseya/portto/pkg/entity/domain/block/biz"
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

	scannedHeight uint64
	biz           bb.IBiz

	taskC chan time.Time
	done  chan bool
}

// NewCronjob create a cronjob instance
func NewCronjob(opts *Options, logger *zap.Logger, initHeight uint64, biz bb.IBiz) adapters.Cronjob {
	return &impl{
		opts:          opts,
		logger:        logger,
		scannedHeight: initHeight,
		biz:           biz,
		taskC:         make(chan time.Time, 1),
		done:          make(chan bool),
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
			i.execute()
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

func (i *impl) execute() {
	ctx := contextx.BackgroundWithLogger(i.logger)

	last, progress, done, _ := i.biz.ScanBlock(ctx, i.scannedHeight)
	if done == nil {
		return
	}

	for {
		select {
		case <-done:
			i.scannedHeight = last
			i.logger.Info("done", zap.Uint64("peak_height", last))
			return
		case record := <-progress:
			i.logger.Info(fmt.Sprintf("progress(%v/%v)", record.Height, last), zap.Uint64("height", record.Height), zap.String("hash", record.Hash))
		case <-time.After(5 * time.Second):
			return
		}
	}
}
