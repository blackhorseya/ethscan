package biz

import (
	"github.com/blackhorseya/ethscan/internal/app/domain/activity/biz/repo"
	"github.com/blackhorseya/ethscan/internal/pkg/errorx"
	"github.com/blackhorseya/ethscan/pkg/contextx"
	ab "github.com/blackhorseya/ethscan/pkg/entity/domain/activity/biz"
	am "github.com/blackhorseya/ethscan/pkg/entity/domain/activity/model"
	bm "github.com/blackhorseya/ethscan/pkg/entity/domain/block/model"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var ProviderSet = wire.NewSet(repo.ProviderSet, NewImpl)

type impl struct {
	repo repo.IRepo
}

func NewImpl(repo repo.IRepo) ab.IBiz {
	return &impl{
		repo: repo,
	}
}

func (i *impl) GetByHash(ctx contextx.Contextx, hash string) (tx *am.Transaction, err error) {
	ret, err := i.repo.FetchTxByHash(ctx, hash)
	if err != nil {
		ctx.Error(errorx.ErrFetchTx.LogMessage, zap.Error(err), zap.String("hash", hash))
		return nil, errorx.ErrFetchTx
	}

	return ret, nil
}

func (i *impl) HandleNewBlock(ctx contextx.Contextx, record *bm.BlockRecord) (txns []*am.Transaction, err error) {
	var ret []*am.Transaction
	for _, tx := range record.Transactions {
		var events []*am.Event
		for _, event := range tx.Events {
			events = append(events, &am.Event{
				Index: event.Index,
				Data:  event.Data,
			})
		}
		input := &am.Transaction{
			BlockHash: tx.BlockHash,
			Hash:      tx.Hash,
			From:      tx.From,
			To:        tx.To,
			Nonce:     tx.Nonce,
			Data:      tx.Data,
			Value:     tx.Value,
			Events:    events,
		}

		err = i.repo.CreateTx(ctx, input)
		if err != nil {
			ctx.Error(errorx.ErrCreateTx.LogMessage, zap.Error(err), zap.Any("tx", tx))
		}

		ret = append(ret, input)
	}

	return ret, nil
}
