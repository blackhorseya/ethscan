package biz

import (
	"fmt"

	"github.com/blackhorseya/portto/internal/app/domain/block/biz/repo"
	"github.com/blackhorseya/portto/internal/pkg/errorx"
	"github.com/blackhorseya/portto/pkg/contextx"
	bb "github.com/blackhorseya/portto/pkg/entity/domain/block/biz"
	bm "github.com/blackhorseya/portto/pkg/entity/domain/block/model"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var ProviderSet = wire.NewSet(repo.ProviderSet, NewImpl)

type impl struct {
	repo repo.IRepo
}

func NewImpl(repo repo.IRepo) bb.IBiz {
	return &impl{repo: repo}
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

	end, err := i.repo.FetchCurrentHeight(ctx)
	if err != nil {
		ctx.Error(errorx.ErrFetchCurrentHeight.LogMessage, zap.Error(err))
		errC <- errorx.ErrFetchCurrentHeight
		return 0, nil, nil, nil
	}
	if start == end {
		return end, nil, nil, nil
	}

	ctx.Info(fmt.Sprintf("start to scan from %v to %v", start, end))

	total := end - start
	completed := uint64(0)
	next := start + 1
	for next <= end {
		go func(height uint64) {
			record, err := i.repo.FetchRecordByHeight(ctx, height)
			if err != nil {
				ctx.Error(errorx.ErrFetchRecord.LogMessage, zap.Error(err), zap.Uint64("height", height))
				errC <- errorx.ErrFetchRecord
				completed++
				return
			}

			err = i.repo.CreateRecord(ctx, record)
			if err != nil {
				ctx.Error(errorx.ErrCreateRecord.LogMessage, zap.Error(err), zap.Any("record", record))
				errC <- errorx.ErrCreateRecord
				completed++
				return
			}

			progress <- record
			completed++
			if total == completed {
				done <- struct{}{}
			}
		}(next)

		next++
	}

	return end, progress, done, errC
}
