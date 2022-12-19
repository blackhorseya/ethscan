package biz

import (
	"github.com/blackhorseya/portto/internal/app/domain/biz/repo"
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

func (i *impl) ScanByHeight(ctx contextx.Contextx, height uint64) (record *bm.BlockRecord, next bool, err error) {
	peakHeight, err := i.repo.FetchCurrentHeight(ctx)
	if err != nil {
		ctx.Error(errorx.ErrFetchCurrentHeight.LogMessage, zap.Error(err))
		return nil, false, errorx.ErrFetchCurrentHeight
	}

	ret, err := i.repo.FetchRecordByHeight(ctx, height)
	if err != nil {
		ctx.Error(errorx.ErrFetchRecord.LogMessage, zap.Error(err), zap.Uint64("height", height))
		return nil, false, errorx.ErrFetchRecord
	}

	if ret.Height+1 <= peakHeight {
		next = true
	}

	// todo: 2022/12/19|sean|create record into database

	return ret, next, nil
}
