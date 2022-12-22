package repo

import (
	"context"
	"database/sql"
	"encoding/json"
	"math/big"
	"strings"
	"time"

	"github.com/blackhorseya/portto/pkg/contextx"
	bm "github.com/blackhorseya/portto/pkg/entity/domain/block/model"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	topicNewBlock = "new_block"
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
	opts     *NodeOptions
	rw       *sqlx.DB
	eth      *ethclient.Client
	producer *kafka.Producer
}

func NewImpl(opts *NodeOptions, rw *sqlx.DB, producer *kafka.Producer) (IRepo, error) {
	client, err := ethclient.Dial(opts.BaseURL)
	if err != nil {
		return nil, err
	}

	return &impl{
		opts:     opts,
		rw:       rw,
		eth:      client,
		producer: producer,
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

	txIds := make([]string, len(block.Transactions()))
	for idx, tx := range block.Transactions() {
		txIds[idx] = tx.Hash().String()
	}

	return &bm.BlockRecord{
		Height:         block.NumberU64(),
		Hash:           block.Hash().String(),
		ParentHash:     block.ParentHash().String(),
		TransactionIds: txIds,
		Timestamp:      timestamppb.New(time.Unix(int64(block.Time()), 0)),
		Depth:          1,
		Status:         bm.BlockStatus_BLOCK_STATUS_UNSTABLE,
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

func (i *impl) ListRecord(ctx contextx.Contextx, condition ListRecordCondition) (records []*bm.BlockRecord, err error) {
	timeout, cancelFunc := i.newContextxWithTimeout(ctx)
	defer cancelFunc()

	query := []string{`select hash, height, parent_hash, timestamp from records`}
	var args []interface{}

	query = append(query, `order by timestamp desc`)

	if condition.Limit != 0 {
		query = append(query, `limit ?`)
		args = append(args, condition.Limit)
	}

	if condition.Offset != 0 {
		query = append(query, `offset ?`)
		args = append(args, condition.Offset)
	}

	stmt := strings.Join(query, " ")

	var resp []*blockRecord
	err = i.rw.SelectContext(timeout, &resp, stmt, args...)
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, nil
	}

	ret := make([]*bm.BlockRecord, len(resp))
	for idx, record := range resp {
		ret[idx] = record.ToEntity()
	}

	return ret, nil
}

func (i *impl) CountRecord(ctx contextx.Contextx, condition ListRecordCondition) (total int, err error) {
	timeout, cancelFunc := i.newContextxWithTimeout(ctx)
	defer cancelFunc()

	stmt := `select count(*) from records`

	ret := 0
	err = i.rw.QueryRowxContext(timeout, stmt).Scan(&ret)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}

		return 0, err
	}

	return ret, nil
}

func (i *impl) CreateRecord(ctx contextx.Contextx, record *bm.BlockRecord) error {
	timeout, cancelFunc := i.newContextxWithTimeout(ctx)
	defer cancelFunc()

	stmt := `insert into records (hash, height, parent_hash, timestamp) values (:hash, :height, :parent_hash, :timestamp)`

	_, err := i.rw.NamedExecContext(timeout, stmt, newBlockRecord(record))
	if err != nil {
		return err
	}

	return nil
}

func (i *impl) PublishRecord(ctx contextx.Contextx, record *bm.BlockRecord, delivery chan kafka.Event) error {
	key := record.Hash
	value, err := json.Marshal(record)
	if err != nil {
		return err
	}

	err = i.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topicNewBlock, Partition: kafka.PartitionAny},
		Value:          value,
		Key:            []byte(key),
		Timestamp:      time.Now(),
		TimestampType:  kafka.TimestampCreateTime,
		Opaque:         nil,
		Headers:        nil,
	}, delivery)
	if err != nil {
		return err
	}

	return nil
}

func (i *impl) newContextxWithTimeout(ctx contextx.Contextx) (contextx.Contextx, context.CancelFunc) {
	return contextx.WithTimeout(ctx, i.opts.Timeout)
}
