package blocks

import (
	"context"
	blocks "github.com/ipfs/go-block-format"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/go-pg/pg/v10"
)

type BlockHeader struct {
	Cid             string `pg:",pk,notnull"`
	Miner           string `pg:",notnull"`
	ParentWeight    string `pg:",notnull"`
	ParentBaseFee   string `pg:",notnull"`
	ParentStateRoot string `pg:",notnull"`

	Height        int64  `pg:",use_zero"`
	WinCount      int64  `pg:",use_zero"`
	Timestamp     uint64 `pg:",use_zero"`
	ForkSignaling uint64 `pg:",use_zero"`

	Ticket        []byte
	ElectionProof []byte
}

func NewBlockHeader(bh *types.BlockHeader) *BlockHeader {
	return &BlockHeader{
		Cid:             bh.Cid().String(),
		Miner:           bh.Miner.String(),
		ParentWeight:    bh.ParentWeight.String(),
		ParentBaseFee:   bh.ParentBaseFee.String(),
		ParentStateRoot: bh.ParentStateRoot.String(),
		Height:          int64(bh.Height),
		WinCount:        bh.ElectionProof.WinCount,
		Timestamp:       bh.Timestamp,
		ForkSignaling:   bh.ForkSignaling,
		Ticket:          bh.Ticket.VRFProof,
		ElectionProof:   bh.ElectionProof.VRFProof,
	}
}

//func (bh *BlockHeader) Persist(ctx context.Context, headers map[cid.Cid]*types.BlockHeader) error {
func (bh *BlockHeader) PersistWithTx(ctx context.Context, tx *pg.Tx) error {
	if _, err := tx.ModelContext(ctx, bh).
		OnConflict("do nothing").
		Insert(); err != nil {
		return err
	}
}
