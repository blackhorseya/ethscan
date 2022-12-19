package repo

import (
	"context"
	"database/sql"
	"math/big"
	"time"

	"github.com/blackhorseya/portto/pkg/contextx"
	bm "github.com/blackhorseya/portto/pkg/entity/domain/block/model"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type NodeOptions struct {
	BaseURL string        `json:"base_url" yaml:"baseURL"`
	Timeout time.Duration `json:"timeout" yaml:"timeout"`
}

func NewNodeOptions(v *viper.Viper) (*NodeOptions, error) {
	opts := new(NodeOptions)

	err := v.UnmarshalKey("node", &opts)
	if err != nil {
		return nil, err
	}

	return opts, nil
}

type impl struct {
	opts *NodeOptions
	rw   *sqlx.DB
	eth  *ethclient.Client
}

func NewImpl(opts *NodeOptions, rw *sqlx.DB) (IRepo, error) {
	client, err := ethclient.Dial(opts.BaseURL)
	if err != nil {
		return nil, err
	}

	return &impl{
		opts: opts,
		rw:   rw,
		eth:  client,
	}, nil
}

func (i *impl) FetchCurrentHeight(ctx contextx.Contextx) (height uint64, err error) {
	timeout, cancelFunc := i.newContextxWithTimeout(ctx)
	defer cancelFunc()

	ret, err := i.eth.BlockNumber(timeout)
	if err != nil {
		return 0, err
	}

	return ret, nil
}

func (i *impl) FetchRecordByHeight(ctx contextx.Contextx, height uint64) (record *bm.BlockRecord, err error) {
	timeout, cancelFunc := i.newContextxWithTimeout(ctx)
	defer cancelFunc()

	block, err := i.eth.BlockByNumber(timeout, big.NewInt(int64(height)))
	if err != nil {
		return nil, err
	}

	return &bm.BlockRecord{
		Height:         block.NumberU64(),
		Hash:           block.Hash().String(),
		ParentHash:     block.ParentHash().String(),
		TransactionIds: nil,
		Timestamp:      timestamppb.New(time.Unix(int64(block.Time()), 0)),
		Depth:          0,
		Status:         bm.BlockStatus_BLOCK_STATUS_UNSPECIFIED,
	}, nil
}

func (i *impl) GetRecordByHash(ctx contextx.Contextx, hash string) (record *bm.BlockRecord, err error) {
	timeout, cancelFunc := i.newContextxWithTimeout(ctx)
	defer cancelFunc()

	stmt := `select hash, height, parent_hash, timestamp from records where hash = ?`

	var resp blockRecord
	err = i.rw.GetContext(timeout, &resp, stmt, hash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return resp.ToEntity(), nil
}

func (i *impl) CreateRecord(ctx contextx.Contextx, record *bm.BlockRecord) error {
	// todo: 2022/12/18|sean|impl me
	panic("implement me")
}

func (i *impl) UpdateRecord(ctx contextx.Contextx, record *bm.BlockRecord) error {
	// todo: 2022/12/18|sean|impl me
	panic("implement me")
}

func (i *impl) newContextxWithTimeout(ctx contextx.Contextx) (contextx.Contextx, context.CancelFunc) {
	return contextx.WithTimeout(ctx, i.opts.Timeout)
}
