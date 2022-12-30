package repo

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/blackhorseya/ethscan/pkg/contextx"
	am "github.com/blackhorseya/ethscan/pkg/entity/domain/activity/model"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
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
	rw   *sqlx.DB
}

func NewImpl(opts *NodeOptions, rw *sqlx.DB) (IRepo, error) {
	client, err := ethclient.Dial(opts.BaseURL)
	if err != nil {
		return nil, err
	}

	return &impl{
		opts: opts,
		eth:  client,
		rw:   rw,
	}, nil
}

func (i *impl) FetchTxByHash(ctx contextx.Contextx, hash string) (tx *am.Transaction, err error) {
	timeout, cancelFunc := i.newContextxWithTimeout(ctx)
	defer cancelFunc()

	h := common.HexToHash(hash)
	resp, _, err := i.eth.TransactionByHash(timeout, h)
	if err != nil {
		return nil, err
	}
	msg, err := resp.AsMessage(types.LatestSignerForChainID(resp.ChainId()), nil)
	if err != nil {
		return nil, err
	}
	receipt, err := i.eth.TransactionReceipt(timeout, h)
	if err != nil {
		return nil, err
	}

	to := ""
	if resp.To() != nil {
		to = resp.To().String()
	}

	ret := &am.Transaction{
		BlockHash: receipt.BlockHash.String(),
		Hash:      resp.Hash().String(),
		From:      msg.From().String(),
		To:        to,
		Nonce:     resp.Nonce(),
		Data:      common.Bytes2Hex(resp.Data()),
		Value:     resp.Value().String(),
		// todo: 2022/12/23|sean|fill the event log
		Events: nil,
	}

	return ret, nil
}

func (i *impl) CreateTx(ctx contextx.Contextx, tx *am.Transaction) error {
	timeout, cancelFunc := i.newContextxWithTimeout(ctx)
	defer cancelFunc()

	stmt := "insert into txns (hash, `from`, `to`, block_hash) values (:hash, :from, :to, :block_hash)"

	_, err := i.rw.NamedExecContext(timeout, stmt, newTransaction(tx))
	if err != nil {
		return err
	}

	return nil
}

func (i *impl) GetTxByHash(ctx contextx.Contextx, hash string) (tx *am.Transaction, err error) {
	timeout, cancelFunc := i.newContextxWithTimeout(ctx)
	defer cancelFunc()

	stmt := "select hash, `from`, `to`, block_hash from txns where hash = ?"

	var resp transaction
	err = i.rw.GetContext(timeout, &resp, stmt, hash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return resp.ToEntity(), nil
}

func (i *impl) ListTxns(ctx contextx.Contextx, cond ListTxnsCondition) (txns []*am.Transaction, err error) {
	timeout, cancelFunc := i.newContextxWithTimeout(ctx)
	defer cancelFunc()

	var conditions []string
	var args []interface{}
	query := []string{"select hash, `from`, `to`, block_hash from txns"}

	if len(cond.BlockHash) > 0 {
		conditions = append(conditions, "block_hash = ?")
		args = append(args, cond.BlockHash)
	}

	if len(conditions) != 0 {
		query = append(query, "where "+strings.Join(conditions, " and "))
	}

	if cond.Limit != 0 {
		query = append(query, "limit ?")
		args = append(args, cond.Limit)
	}

	if cond.Offset != 0 {
		query = append(query, "offset ?")
		args = append(args, cond.Offset)
	}

	stmt := strings.Join(query, " ")

	var resp []*transaction
	err = i.rw.SelectContext(timeout, &resp, stmt, args...)
	if err != nil {
		return nil, err
	}

	var ret []*am.Transaction
	for _, t := range resp {
		ret = append(ret, t.ToEntity())
	}

	return ret, nil
}

func (i *impl) newContextxWithTimeout(ctx contextx.Contextx) (contextx.Contextx, context.CancelFunc) {
	return contextx.WithTimeout(ctx, i.opts.Timeout)
}
