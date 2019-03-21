// Package ipfslite is a super-lightweight IPFS peer which runs the minimal
// setup to provide an `ipld.DAGService`. It can only do basic DAG
// functionality like retrieving and putting blocks from/to the IPFS
// network.
package ipfslite

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/ipfs/go-bitswap"
	"github.com/ipfs/go-bitswap/network"
	blockservice "github.com/ipfs/go-blockservice"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-datastore"
	blockstore "github.com/ipfs/go-ipfs-blockstore"
	chunker "github.com/ipfs/go-ipfs-chunker"
	offline "github.com/ipfs/go-ipfs-exchange-offline"
	cbor "github.com/ipfs/go-ipld-cbor"
	ipld "github.com/ipfs/go-ipld-format"
	logging "github.com/ipfs/go-log"
	"github.com/ipfs/go-merkledag"
	dag "github.com/ipfs/go-merkledag"
	"github.com/ipfs/go-unixfs/importer/balanced"
	"github.com/ipfs/go-unixfs/importer/helpers"
	"github.com/ipfs/go-unixfs/importer/trickle"
	ufsio "github.com/ipfs/go-unixfs/io"
	host "github.com/libp2p/go-libp2p-host"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	multihash "github.com/multiformats/go-multihash"
)

func init() {
	ipld.Register(cid.DagProtobuf, dag.DecodeProtobufBlock)
	ipld.Register(cid.Raw, dag.DecodeRawBlock)
	ipld.Register(cid.DagCBOR, cbor.DecodeBlock) // need to decode CBOR
}

var logger = logging.Logger("ipfslite")

// Config wraps configuration options for the Peer.
type Config struct {
	// The DAGService will not announce or retrieve blocks from the network
	Offline bool
}

// Peer is an IPFS-Lite peer. It provides a DAG service that can fetch and put
// blocks from/to the IPFS network.
type Peer struct {
	ipld.DAGService
	bstore blockstore.Blockstore
	cfg    *Config
}

// New creates an IPFS-Lite Peer. It uses the given datastore, libp2p Host and
// DHT. The Host and the DHT may be nil if config.Offline is set to true, as
// they are not used in that case. Peer implements the ipld.DAGService
// interface.
func New(
	ctx context.Context,
	store datastore.Batching,
	host host.Host,
	dht *dht.IpfsDHT,
	cfg *Config,
) (*Peer, error) {

	if cfg == nil {
		cfg = &Config{}
	}

	bs := blockstore.NewBlockstore(store)
	bs = blockstore.NewIdStore(bs)
	cachedbs, err := blockstore.CachedBlockstore(ctx, bs, blockstore.DefaultCacheOpts())
	if err != nil {
		return nil, err
	}

	var bserv blockservice.BlockService
	if cfg.Offline {
		bserv = blockservice.New(cachedbs, offline.Exchange(cachedbs))
	} else {
		bswapnet := network.NewFromIpfsHost(host, dht)
		bswap := bitswap.New(ctx, bswapnet, cachedbs)
		bserv = blockservice.New(cachedbs, bswap)
	}

	dags := merkledag.NewDAGService(bserv)
	return &Peer{
		DAGService: dags,
		cfg:        cfg,
		bstore:     cachedbs,
	}, nil
}

// AddParams contains all of the configurable parameters needed to specify the
// importing process of a file.
type AddParams struct {
	Layout    string
	Chunker   string
	RawLeaves bool
	Hidden    bool
	Shard     bool
	HashFun   string
}

// AddFile chunks and adds content to the DAGService from a reader. The content
// is stored as a UnixFS DAG (default for IPFS). It returns the root
// ipld.Node.
func (p *Peer) AddFile(ctx context.Context, r io.Reader, params *AddParams) (ipld.Node, error) {
	if params == nil {
		params = &AddParams{}
	}
	if params.HashFun == "" {
		params.HashFun = "sha2-256"
	}

	prefix, err := merkledag.PrefixForCidVersion(1)
	if err != nil {
		return nil, fmt.Errorf("bad CID Version: %s", err)
	}

	hashFunCode, ok := multihash.Names[strings.ToLower(params.HashFun)]
	if !ok {
		return nil, fmt.Errorf("unrecognized hash function: %s", params.HashFun)
	}
	prefix.MhType = hashFunCode
	prefix.MhLength = -1

	dbp := helpers.DagBuilderParams{
		Dagserv:    p,
		RawLeaves:  params.RawLeaves,
		Maxlinks:   helpers.DefaultLinksPerBlock,
		NoCopy:     false,
		CidBuilder: &prefix,
	}

	chnk, err := chunker.FromString(r, params.Chunker)
	dbh, err := dbp.New(chnk)
	if err != nil {
		return nil, err
	}

	switch params.Layout {
	case "trickle":
		return trickle.Layout(dbh)
	case "balanced", "":
		return balanced.Layout(dbh)
	default:
		return nil, errors.New("invalid Layout")
	}
}

// GetFile returns a reader to a file as identified by its root CID. The file
// must have been added as a UnixFS DAG (default for IPFS).
func (p *Peer) GetFile(ctx context.Context, c cid.Cid) (ufsio.ReadSeekCloser, error) {
	n, err := p.Get(ctx, c)
	if err != nil {
		return nil, err
	}
	return ufsio.NewDagReader(ctx, n, p)
}

// Blockstore offers access to the blockstore underlying the Peer's DAGService.
func (p *Peer) BlockStore() blockstore.Blockstore {
	return p.bstore
}