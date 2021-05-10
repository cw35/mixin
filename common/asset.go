package common

import (
	"fmt"

	"github.com/MixinNetwork/mixin/crypto"
	"github.com/MixinNetwork/mixin/domains/arweave"
	"github.com/MixinNetwork/mixin/domains/bch"
	"github.com/MixinNetwork/mixin/domains/binance"
	"github.com/MixinNetwork/mixin/domains/bitcoin"
	"github.com/MixinNetwork/mixin/domains/cosmos"
	"github.com/MixinNetwork/mixin/domains/dash"
	"github.com/MixinNetwork/mixin/domains/decred"
	"github.com/MixinNetwork/mixin/domains/dogecoin"
	"github.com/MixinNetwork/mixin/domains/eos"
	"github.com/MixinNetwork/mixin/domains/ethereum"
	"github.com/MixinNetwork/mixin/domains/filecoin"
	"github.com/MixinNetwork/mixin/domains/handshake"
	"github.com/MixinNetwork/mixin/domains/horizen"
	"github.com/MixinNetwork/mixin/domains/kusama"
	"github.com/MixinNetwork/mixin/domains/mobilecoin"
	"github.com/MixinNetwork/mixin/domains/monero"
	"github.com/MixinNetwork/mixin/domains/polkadot"
	"github.com/MixinNetwork/mixin/domains/siacoin"
	"github.com/MixinNetwork/mixin/domains/solana"
	"github.com/MixinNetwork/mixin/domains/tron"
	"github.com/MixinNetwork/mixin/domains/zcash"
)

var (
	XINAssetId crypto.Hash
)

type Asset struct {
	ChainId  crypto.Hash
	AssetKey string
}

func init() {
	XINAssetId = crypto.NewHash([]byte("c94ac88f-4671-3976-b60a-09064f1811e8"))
}

func (a *Asset) Verify() error {
	switch a.ChainId {
	case ethereum.EthereumChainId:
		return ethereum.VerifyAssetKey(a.AssetKey)
	case bitcoin.BitcoinChainId:
		return bitcoin.VerifyAssetKey(a.AssetKey)
	case monero.MoneroChainId:
		return monero.VerifyAssetKey(a.AssetKey)
	case zcash.ZcashChainId:
		return zcash.VerifyAssetKey(a.AssetKey)
	case horizen.HorizenChainId:
		return horizen.VerifyAssetKey(a.AssetKey)
	case dogecoin.DogecoinChainId:
		return dogecoin.VerifyAssetKey(a.AssetKey)
	case dash.DashChainId:
		return dash.VerifyAssetKey(a.AssetKey)
	case decred.DecredChainId:
		return decred.VerifyAssetKey(a.AssetKey)
	case bch.BitcoinCashChainId:
		return bch.VerifyAssetKey(a.AssetKey)
	case handshake.HandshakenChainId:
		return handshake.VerifyAssetKey(a.AssetKey)
	case siacoin.SiacoinChainId:
		return siacoin.VerifyAssetKey(a.AssetKey)
	case filecoin.FilecoinChainId:
		return filecoin.VerifyAssetKey(a.AssetKey)
	case solana.SolanaChainId:
		return solana.VerifyAssetKey(a.AssetKey)
	case polkadot.PolkadotChainId:
		return polkadot.VerifyAssetKey(a.AssetKey)
	case kusama.KusamaChainId:
		return kusama.VerifyAddress(a.AssetKey)
	case eos.EOSChainId:
		return eos.VerifyAssetKey(a.AssetKey)
	case tron.TronChainId:
		return tron.VerifyAssetKey(a.AssetKey)
	case mobilecoin.MobileCoinChainId:
		return mobilecoin.VerifyAssetKey(a.AssetKey)
	case cosmos.CosmosChainId:
		return cosmos.VerifyAssetKey(a.AssetKey)
	case binance.BinanceChainId:
		return binance.VerifyAssetKey(a.AssetKey)
	case arweave.ArweaveChainId:
		return arweave.VerifyAssetKey(a.AssetKey)
	default:
		return fmt.Errorf("invalid chain id %s", a.ChainId)
	}
}

func (a *Asset) AssetId() crypto.Hash {
	switch a.ChainId {
	case ethereum.EthereumChainId:
		return ethereum.GenerateAssetId(a.AssetKey)
	case bitcoin.BitcoinChainId:
		return bitcoin.GenerateAssetId(a.AssetKey)
	case monero.MoneroChainId:
		return monero.GenerateAssetId(a.AssetKey)
	case zcash.ZcashChainId:
		return zcash.GenerateAssetId(a.AssetKey)
	case horizen.HorizenChainId:
		return horizen.GenerateAssetId(a.AssetKey)
	case dogecoin.DogecoinChainId:
		return dogecoin.GenerateAssetId(a.AssetKey)
	case dash.DashChainId:
		return dash.GenerateAssetId(a.AssetKey)
	case decred.DecredChainId:
		return decred.GenerateAssetId(a.AssetKey)
	case bch.BitcoinCashChainId:
		return bch.GenerateAssetId(a.AssetKey)
	case handshake.HandshakenChainId:
		return handshake.GenerateAssetId(a.AssetKey)
	case siacoin.SiacoinChainId:
		return siacoin.GenerateAssetId(a.AssetKey)
	case filecoin.FilecoinChainId:
		return filecoin.GenerateAssetId(a.AssetKey)
	case solana.SolanaChainId:
		return solana.GenerateAssetId(a.AssetKey)
	case polkadot.PolkadotChainId:
		return polkadot.GenerateAssetId(a.AssetKey)
	case kusama.KusamaChainId:
		return kusama.GenerateAssetId(a.AssetKey)
	case eos.EOSChainId:
		return eos.GenerateAssetId(a.AssetKey)
	case tron.TronChainId:
		return tron.GenerateAssetId(a.AssetKey)
	case mobilecoin.MobileCoinChainId:
		return mobilecoin.GenerateAssetId(a.AssetKey)
	case cosmos.CosmosChainId:
		return cosmos.GenerateAssetId(a.AssetKey)
	case binance.BinanceChainId:
		return binance.GenerateAssetId(a.AssetKey)
	case arweave.ArweaveChainId:
		return arweave.GenerateAssetId(a.AssetKey)
	default:
		return crypto.Hash{}
	}
}

func (a *Asset) FeeAssetId() crypto.Hash {
	switch a.ChainId {
	case ethereum.EthereumChainId:
		return ethereum.EthereumChainId
	case bitcoin.BitcoinChainId:
		return bitcoin.BitcoinChainId
	case monero.MoneroChainId:
		return monero.MoneroChainId
	case zcash.ZcashChainId:
		return zcash.ZcashChainId
	case horizen.HorizenChainId:
		return horizen.HorizenChainId
	case dogecoin.DogecoinChainId:
		return dogecoin.DogecoinChainId
	case dash.DashChainId:
		return dash.DashChainId
	case decred.DecredChainId:
		return decred.DecredChainId
	case bch.BitcoinCashChainId:
		return bch.BitcoinCashChainId
	case handshake.HandshakenChainId:
		return handshake.HandshakenChainId
	case siacoin.SiacoinChainId:
		return siacoin.SiacoinChainId
	case filecoin.FilecoinChainId:
		return filecoin.FilecoinChainId
	case solana.SolanaChainId:
		return solana.SolanaChainId
	case polkadot.PolkadotChainId:
		return polkadot.PolkadotChainId
	case kusama.KusamaChainId:
		return kusama.KusamaChainId
	case eos.EOSChainId:
		return eos.EOSChainId
	case tron.TronChainId:
		return tron.TronChainId
	case mobilecoin.MobileCoinChainId:
		return mobilecoin.MobileCoinChainId
	case cosmos.CosmosChainId:
		return cosmos.CosmosChainId
	case binance.BinanceChainId:
		return binance.BinanceChainId
	case arweave.ArweaveChainId:
		return arweave.ArweaveChainId
	}
	return crypto.Hash{}
}
