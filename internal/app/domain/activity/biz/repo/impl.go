package repo

import (
	"time"

	"github.com/blackhorseya/ethscan/pkg/contextx"
	am "github.com/blackhorseya/ethscan/pkg/entity/domain/activity/model"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
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
	eth  *ethclient.Client
}

func NewImpl(opts *NodeOptions) (IRepo, error) {
	client, err := ethclient.Dial(opts.BaseURL)
	if err != nil {
		return nil, err
	}

	return &impl{
		opts: opts,
		eth:  client,
	}, nil
}

func (i *impl) FetchTxByHash(ctx contextx.Contextx, hash string) (tx *am.Transaction, err error) {
	// todo: 2022/12/23|sean|impl me
	panic("implement me")
}
