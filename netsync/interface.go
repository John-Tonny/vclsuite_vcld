// Copyright (c) 2017 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package netsync

import (
	"github.com/John-Tonny/vclsuite_vcld/blockchain"
	"github.com/John-Tonny/vclsuite_vcld/chaincfg"
	"github.com/John-Tonny/vclsuite_vcld/chaincfg/chainhash"
	"github.com/John-Tonny/vclsuite_vcld/mempool"
	"github.com/John-Tonny/vclsuite_vcld/peer"
	"github.com/John-Tonny/vclsuite_vcld/wire"
	vclutil "github.com/John-Tonny/vclsuite_vclutil"
)

// PeerNotifier exposes methods to notify peers of status changes to
// transactions, blocks, etc. Currently server (in the main package) implements
// this interface.
type PeerNotifier interface {
	AnnounceNewTransactions(newTxs []*mempool.TxDesc)

	UpdatePeerHeights(latestBlkHash *chainhash.Hash, latestHeight int32, updateSource *peer.Peer)

	RelayInventory(invVect *wire.InvVect, data interface{})

	TransactionConfirmed(tx *vclutil.Tx)
}

// Config is a configuration struct used to initialize a new SyncManager.
type Config struct {
	PeerNotifier PeerNotifier
	Chain        *blockchain.BlockChain
	TxMemPool    *mempool.TxPool
	ChainParams  *chaincfg.Params

	DisableCheckpoints bool
	MaxPeers           int

	FeeEstimator *mempool.FeeEstimator
}
