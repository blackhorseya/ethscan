package biz

import (
	"fmt"
	"sync/atomic"

	"github.com/blackhorseya/ethscan/internal/app/domain/block/biz/repo"
	"github.com/blackhorseya/ethscan/internal/pkg/errorx"
	"github.com/blackhorseya/ethscan/pkg/contextx"
	"github.com/blackhorseya/ethscan/pkg/entity/domain/activity/s2s"
	bb "github.com/blackhorseya/ethscan/pkg/entity/domain/block/biz"
	bm "github.com/blackhorseya/ethscan/pkg/entity/domain/block/model"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var ProviderSet = wire.NewSet(repo.ProviderSet, NewImpl, NewActivityClient)

type impl struct {
	repo     repo.IRepo
	activity s2s.ServiceClient
}

func NewImpl(repo repo.IRepo, activity s2s.ServiceClient) bb.IBiz {
	return &impl{
		repo:     repo,
		activity: activity,
	}
}

func (i *impl) GetByHash(ctx contextx.Contextx, hash string) (record *bm.BlockRecord, err error) {
	ret, err := i.repo.GetRecordByHash(ctx, hash)
	if err != nil {
		ctx.Error(errorx.ErrGetRecord.LogMessage, zap.Error(err), zap.String("hash", hash))
		return nil, errorx.ErrGetRecord
	}

	// todo: 2022/12/23|sean|get txns id by hash of block from activity grpc

	return ret, nil
}

func (i *impl) List(ctx contextx.Contextx, cond bb.ListCondition) (records []*bm.BlockRecord, total int, err error) {
	if cond.Page <= 0 {
		ctx.Error(errorx.ErrInvalidPage.LogMessage, zap.Int("page", cond.Page))
		return nil, 0, errorx.ErrInvalidPage
	}

	if cond.Size <= 0 {
		ctx.Error(errorx.ErrInvalidSize.LogMessage, zap.Int("size", cond.Size))
		return nil, 0, errorx.ErrInvalidSize
	}

	condition := repo.ListRecordCondition{
		Limit:  cond.Size,
		Offset: (cond.Page - 1) * cond.Size,
	}
	ret, err := i.repo.ListRecord(ctx, condition)
	if err != nil {
		ctx.Error(errorx.ErrGetRecord.LogMessage, zap.Error(err), zap.Any("condition", condition))
		return nil, 0, errorx.ErrGetRecord
	}

	count, err := i.repo.CountRecord(ctx, condition)
	if err != nil {
		ctx.Error(errorx.ErrCountRecord.LogMessage, zap.Error(err), zap.Any("condition", condition))
		return nil, 0, errorx.ErrCountRecord
	}

	return ret, count, nil
}

func (i *impl) ScanBlock(ctx contextx.Contextx, start uint64) (last uint64, progress chan *bm.BlockRecord, done chan struct{}, errC chan error) {
	progress = make(chan *bm.BlockRecord)
	done = make(chan struct{})
	errC = make(chan error)

	peakHeight, err := i.repo.FetchCurrentHeight(ctx)
	if err != nil {
		ctx.Error(errorx.ErrFetchCurrentHeight.LogMessage, zap.Error(err))
		errC <- errorx.ErrFetchCurrentHeight
		return 0, nil, nil, nil
	}

	scanned, err := i.repo.GetLatestRecord(ctx)
	if err != nil {
		ctx.Error(errorx.ErrGetRecord.LogMessage, zap.Error(err))
		errC <- errorx.ErrGetRecord
		return 0, nil, nil, nil
	}
	scannedHeight := scanned.Height

	if scannedHeight == peakHeight {
		return peakHeight, nil, nil, nil
	}

	total := peakHeight - scannedHeight
	nextHeight := scannedHeight + 1
	completed := uint64(0)

	ctx.Info(fmt.Sprintf("start to scan from %v to %v", nextHeight, peakHeight))

	for nextHeight <= peakHeight {
		go func(height uint64) {
			defer func() {
				atomic.AddUint64(&completed, 1)
				if total == atomic.LoadUint64(&completed) {
					done <- struct{}{}
				}
			}()

			ctx := contextx.BackgroundWithLogger(ctx.GetLogger())

			record, err := i.repo.FetchRecordByHeight(ctx, height)
			if err != nil {
				ctx.Error(errorx.ErrFetchRecord.LogMessage, zap.Error(err), zap.Uint64("height", height))
				errC <- errorx.ErrFetchRecord
				return
			}

			exists, err := i.repo.GetRecordByHash(ctx, record.Hash)
			if err != nil {
				ctx.Error(errorx.ErrGetRecord.LogMessage, zap.Error(err), zap.String("hash", record.Hash))
				errC <- errorx.ErrGetRecord
				return
			}

			if exists == nil {
				err = i.repo.ProduceRecord(ctx, record, nil)
				if err != nil {
					ctx.Error(errorx.ErrProduceRecord.LogMessage, zap.Error(err), zap.Any("record", record))
					return
				}

				err = i.repo.CreateRecord(ctx, record)
				if err != nil {
					ctx.Error(errorx.ErrCreateRecord.LogMessage, zap.Error(err), zap.Any("record", record))
					errC <- errorx.ErrCreateRecord
					return
				}
			}

			progress <- record
		}(nextHeight)

		nextHeight++
	}

	return peakHeight, progress, done, errC
}
