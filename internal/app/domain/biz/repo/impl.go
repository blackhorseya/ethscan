package repo

import (
	"github.com/blackhorseya/portto/pkg/contextx"
	"github.com/blackhorseya/portto/pkg/entity/domain/block/model"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type NodeOptions struct {
	BaseURL string `json:"base_url" yaml:"baseURL"`
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
	ret, err := i.eth.BlockNumber(ctx)
	if err != nil {
		return 0, err
	}

	return ret, nil
}

func (i *impl) FetchRecordByHeight(ctx contextx.Contextx, height uint64) (record *model.BlockRecord, err error) {
	// todo: 2022/12/18|sean|impl me
	panic("implement me")
}

func (i *impl) GetRecordByHash(ctx contextx.Contextx, hash string) (record *model.BlockRecord, err error) {
	// todo: 2022/12/18|sean|impl me
	panic("implement me")
}

func (i *impl) CreateRecord(ctx contextx.Contextx, record *model.BlockRecord) error {
	// todo: 2022/12/18|sean|impl me
	panic("implement me")
}

func (i *impl) UpdateRecord(ctx contextx.Contextx, record *model.BlockRecord) error {
	// todo: 2022/12/18|sean|impl me
	panic("implement me")
}
