// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2019-05-06 03:22:32.35974 +0200 CEST m=+0.002916631
// using data from
// coins.json
package coin

const (
	BTC   = 0
	LTC   = 2
	DOGE  = 3
	DASH  = 5
	DCR   = 42
	ETH   = 60
	ETC   = 61
	ICX   = 74
	XVG   = 77
	GAME  = 101
	ARK   = 111
	ZEC   = 133
	LSK   = 134
	STEEM = 135
	XRP   = 144
	XLM   = 148
	XRB   = 165
	POA   = 178
	EOS   = 194
	TRX   = 195
	NIM   = 242
	ZIL   = 313
	AION  = 425
	THETA = 500
	KTS   = 713
	BNB   = 714
	VET   = 818
	CLO   = 820
	NEO   = 888
	TOMO  = 889
	TT    = 1001
	ONT   = 1024
	XTZ   = 1729
	ADA   = 1815
	KIN   = 2017
	DYN   = 3381
	IOTA  = 4218
	GO    = 6060
	WAN   = 5718350
	WAVES = 5741564
	ATOM  = 118
)

var Coins = map[uint]Coin{
	BTC: {
		Index:   BTC,
		Symbol:  "BTC",
		Title:   "Bitcoin",
		Website: "https://bitcoin.org/",
	},
	LTC: {
		Index:   LTC,
		Symbol:  "LTC",
		Title:   "Litecoin",
		Website: "https://litecoin.org/",
	},
	DOGE: {
		Index:   DOGE,
		Symbol:  "DOGE",
		Title:   "Dogecoin",
		Website: "https://github.com/dogecoin/dogecoin",
	},
	DASH: {
		Index:   DASH,
		Symbol:  "DASH",
		Title:   "Dash",
		Website: "https://github.com/dashpay/dash",
	},
	DCR: {
		Index:   DCR,
		Symbol:  "DCR",
		Title:   "Decred",
		Website: "https://decred.org/",
	},
	ETH: {
		Index:    ETH,
		Symbol:   "ETH",
		Title:    "Ether",
		Website:  "https://ethereum.org/ether",
		Decimals: 18,
	},
	ETC: {
		Index:    ETC,
		Symbol:   "ETC",
		Title:    "Ether Classic",
		Website:  "https://ethereumclassic.github.io",
		Decimals: 18,
	},
	ICX: {
		Index:   ICX,
		Symbol:  "ICX",
		Title:   "ICON",
		Website: "https://icon.foundation/",
	},
	XVG: {
		Index:   XVG,
		Symbol:  "XVG",
		Title:   "Verge",
		Website: "https://github.com/vergecurrency/verge/",
	},
	GAME: {
		Index:   GAME,
		Symbol:  "GAME",
		Title:   "GameCredits",
		Website: "https://github.com/gamecredits-project",
	},
	ARK: {
		Index:   ARK,
		Symbol:  "ARK",
		Title:   "ARK",
		Website: "https://ark.io",
	},
	ZEC: {
		Index:   ZEC,
		Symbol:  "ZEC",
		Title:   "Zcash",
		Website: "https://z.cash",
	},
	LSK: {
		Index:   LSK,
		Symbol:  "LSK",
		Title:   "Lisk",
		Website: "https://lisk.io/",
	},
	STEEM: {
		Index:   STEEM,
		Symbol:  "STEEM",
		Title:   "Steem",
		Website: "http://steem.io",
	},
	XRP: {
		Index:    XRP,
		Symbol:   "XRP",
		Title:    "Ripple",
		Website:  "https://ripple.com",
		Decimals: 6,
	},
	XLM: {
		Index:    XLM,
		Symbol:   "XLM",
		Title:    "Stellar Lumens",
		Website:  "https://www.stellar.org/",
		Decimals: 7,
	},
	XRB: {
		Index:   XRB,
		Symbol:  "XRB",
		Title:   "RaiBlocks",
		Website: "https://raiblocks.com",
	},
	POA: {
		Index:    POA,
		Symbol:   "POA",
		Title:    "Poa",
		Website:  "https://poa.network",
		Decimals: 18,
	},
	EOS: {
		Index:   EOS,
		Symbol:  "EOS",
		Title:   "EOS",
		Website: "https://eos.io",
	},
	TRX: {
		Index:   TRX,
		Symbol:  "TRX",
		Title:   "Tron",
		Website: "https://tron.network/enindex.html",
	},
	NIM: {
		Index:    NIM,
		Symbol:   "NIM",
		Title:    "Nimiq",
		Website:  "https://nimiq.com/",
		Decimals: 5,
	},
	ZIL: {
		Index:   ZIL,
		Symbol:  "ZIL",
		Title:   "Zilliqa",
		Website: "https://zilliqa.com/",
	},
	AION: {
		Index:    AION,
		Symbol:   "AION",
		Title:    "Aion",
		Website:  "https://aion.network",
		Decimals: 18,
	},
	THETA: {
		Index:   THETA,
		Symbol:  "THETA",
		Title:   "Theta",
		Website: "https://www.thetatoken.org/",
	},
	KTS: {
		Index:   KTS,
		Symbol:  "KTS",
		Title:   "Katallassos",
		Website: "https://katallassos.com",
	},
	BNB: {
		Index:    BNB,
		Symbol:   "BNB",
		Title:    "Binance",
		Website:  "https://www.binance.org",
		Decimals: 8,
	},
	VET: {
		Index:   VET,
		Symbol:  "VET",
		Title:   "VeChain Token",
		Website: "https://vechain.com/",
	},
	CLO: {
		Index:    CLO,
		Symbol:   "CLO",
		Title:    "Callisto",
		Website:  "http://callisto.network/",
		Decimals: 18,
	},
	NEO: {
		Index:   NEO,
		Symbol:  "NEO",
		Title:   "NEO",
		Website: "https://neo.org/",
	},
	TOMO: {
		Index:   TOMO,
		Symbol:  "TOMO",
		Title:   "TOMO",
		Website: "https://tomochain.com/",
	},
	TT: {
		Index:    TT,
		Symbol:   "TT",
		Title:    "ThunderCore",
		Website:  "https://thundercore.com/",
		Decimals: 18,
	},
	ONT: {
		Index:   ONT,
		Symbol:  "ONT",
		Title:   "Ontology",
		Website: "https://ont.io",
	},
	XTZ: {
		Index:    XTZ,
		Symbol:   "XTZ",
		Title:    "Tezos",
		Website:  "https://tezos.com",
		Decimals: 6,
	},
	ADA: {
		Index:   ADA,
		Symbol:  "ADA",
		Title:   "Cardano",
		Website: "https://www.cardanohub.org/en/home/",
	},
	KIN: {
		Index:    KIN,
		Symbol:   "KIN",
		Title:    "Kin",
		Website:  "https://www.kinecosystem.org/",
		Decimals: 5,
	},
	DYN: {
		Index:   DYN,
		Symbol:  "DYN",
		Title:   "Dynamic",
		Website: "https://duality.solutions/dynamic/",
	},
	IOTA: {
		Index:   IOTA,
		Symbol:  "IOTA",
		Title:   "IOTA",
		Website: "https://www.iota.org/",
	},
	GO: {
		Index:    GO,
		Symbol:   "GO",
		Title:    "GoChain GO",
		Website:  "https://gochain.io/",
		Decimals: 18,
	},
	WAN: {
		Index:    WAN,
		Symbol:   "WAN",
		Title:    "Wanchain",
		Website:  "https://wanchain.org/",
		Decimals: 18,
	},
	WAVES: {
		Index:   WAVES,
		Symbol:  "WAVES",
		Title:   "Waves",
		Website: "https://wavesplatform.com/",
	},
	ATOM: {
		Index:   ATOM,
		Symbol:  "ATOM",
		Title:   "Cosmos ATOM",
		Website: "https://cosmos.network/",
	},
}
