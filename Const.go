package coinapi

type CurrencyPair int

func (c CurrencyPair) String() string {
    if c == 0 {
        return "nil"
    }
    return currencyPairSymbol[c-1]
}

type Currency int

func (c Currency) String() string {
    if c == 0 {
        return "nil"
    }
    return currencySymbol[c]
}

type TradeSide int

func (ts TradeSide) String() string {
    switch ts {
    case 1:
        return "BUY"
    case 2:
        return "SELL"
    case 3:
        return "BUY_MARKET"
    case 4:
        return "SELL_MARKET"
    default:
        return "UNKNOWN"
    }
}

type TradeStatus int

func (ts TradeStatus) String() string {
    return orderStatusSymbol[ts]
}

var currencySymbol = map[Currency]string{
    CNY: "cny",
    USD: "usd",
    BTC: "btc",

    BLK:   "blk",
    BTCD:  "btcd",
    BTS:   "bts",
    BURST: "burst",
    CLAM:  "clam",
    DASH:  "dash",
    DGB:   "dgb",
    DOGE:  "doge",
    EMC2:  "emc2",
    FLDC:  "fldc",
    FLO:   "flo",
    GAME:  "game",
    GRC:   "grc",
    LTC:   "ltc",
    MAID:  "maid",
    OMNI:  "omni",
    NAUT:  "naut",
    NAV:   "nav",
    NEOS:  "neos",
    NXT:   "nxt",
    PINK:  "pink",
    POT:   "pot",
    PPC:   "ppc",
    SJCX:  "sjcx",
    SYS:   "sys",
    VIA:   "via",
    XVC:   "xvc",
    VRC:   "vrc",
    VTC:   "vtc",
    XCP:   "xcp",
    XEM:   "xem",
    XMR:   "xmr",
    XRP:   "xrp",
    ETH:   "eth",
    SC:    "sc",
    BCY:   "bcy",
    EXP:   "exp",
    FCT:   "fct",
    RADS:  "rads",
    AMP:   "amp",
    DCR:   "dcr",
    LSK:   "lsk",
    LBC:   "lbc",
    STEEM: "steem",
    SBD:   "sbd",
    ETC:   "etc",
    REP:   "rep",
    ARDR:  "ardr",
    ZEC:   "zec",
    STRAT: "strat",
    NXC:   "nxc",
    GNT:   "gnt",
    GNO:   "gno",
}
var SymbolCurrency = map[string]Currency{
    "CNY": CNY,
    "USD": USD,
    "BTC": BTC,

    "BLK":   BLK,
    "BTCD":  BTCD,
    "BTS":   BTS,
    "BURST": BURST,
    "CLAM":  CLAM,
    "DASH":  DASH,
    "DGB":   DGB,
    "DOGE":  DOGE,
    "EMC2":  EMC2,
    "FLDC":  FLDC,
    "FLO":   FLO,
    "GAME":  GAME,
    "GRC":   GRC,
    "LTC":   LTC,
    "MAID":  MAID,
    "OMNI":  OMNI,
    "NAUT":  NAUT,
    "NAV":   NAV,
    "NEOS":  NEOS,
    "NXT":   NXT,
    "PINK":  PINK,
    "POT":   POT,
    "PPC":   PPC,
    "SJCX":  SJCX,
    "SYS":   SYS,
    "VIA":   VIA,
    "XVC":   XVC,
    "VRC":   VRC,
    "VTC":   VTC,
    "XCP":   XCP,
    "XEM":   XEM,
    "XMR":   XMR,
    "XRP":   XRP,
    "ETH":   ETH,
    "SC":    SC,
    "BCY":   BCY,
    "EXP":   EXP,
    "FCT":   FCT,
    "RADS":  RADS,
    "AMP":   AMP,
    "DCR":   DCR,
    "LSK":   LSK,
    "LBC":   LBC,
    "STEEM": STEEM,
    "SBD":   SBD,
    "ETC":   ETC,
    "REP":   REP,
    "ARDR":  ARDR,
    "ZEC":   ZEC,
    "STRAT": STRAT,
    "NXC":   NXC,
    "GNT":   GNT,
    "GNO":   GNO,
    "QTUM":  QTUM,
}

const (
    CNY = 1 + iota
    USD
    BTC
    LTC
    ETH
    ETC
    ZEC
    SC
    REP
    BTS
    GNT

    XPM
    XRP
    ZCC
    MEC
    ANC
    BEC
    PPC
    SRC
    TAG
    WDC
    XLM
    DGC
    QRK
    DOGE
    YBC
    RIC
    BOST
    NXT
    BLK
    NRS
    MED
    NCS
    EAC
    XCN
    SYS
    XEM
    VASH
    DASH
    EMC
    HLB
    ARDR
    XZC
    MGC
    TMC
    BNS
    //	BTS
    CORG
    NEOS
    XST
    OneCR
    BDC
    DRKC
    FRAC
    SRCC
    CC
    DAO
    eTOK
    NAV
    TRUST
    AUR
    DIME
    EXP
    GAME
    IOC
    BLU
    FAC
    GEMZ
    CYC
    EMO
    JLH
    XBC
    XDP
    //	DASH
    GAP
    SMC
    XHC
    BTCD
    GRCX
    XUSD
    MIL
    LGC
    PIGGY
    XCP
    BURST
    GNS
    HIRO
    HUGE
    LC
    FLDC
    INDEX
    LEAF
    MYR
    SPA
    CURE
    FLO
    NAUT
    SJCX
    TWE
    MON
    BLOCK
    CHA
    GIAR
    HZ
    IFC
    DGB
    MRC
    VIA
    BTCS
    GOLD
    MMNXT
    XLB
    XMG
    BALLS
    HOT
    NOTE
    SYNC
    //	ARDR
    DNS
    //	ETC
    FRK
    MAST
    CLAM
    NOBL
    XXC
    C2
    NXC
    Q2C
    WIKI
    XSV
    AERO
    FZN
    MINT
    QBK
    VOX
    BURN
    LTBC
    QCN
    //	XEM
    DICE
    FLT
    OMNI
    AC
    APH
    BDG
    BITCNY
    CRYPT
    //	NXT
    OPAL
    RZR
    SHIBE
    SQL
    SUM
    BANK
    CON
    JUG
    METH
    //	SC
    UTIL
    VTC
    LOVE
    MCN
    POT
    CINNI
    ECC
    GDN
    GRS
    KEY
    SHOPX
    XAP
    STEEM
    YIN
    AMP
    NMC
    SRG
    XDN
    YANG
    XAI
    CCN
    CGA
    MAID
    URO
    X13
    VRC
    XCH
    HYP
    MRS
    PLX
    QORA
    USDT
    SLR
    COMM
    DIS
    FVZ
    IXC
    LBC
    GML
    LTCX
    NAS
    AXIS
    CNL
    //	ETH
    FLAP
    FOX
    QTL
    RADS
    //	RIC
    SBD
    BCC
    CNOTE
    FZ
    GPC
    //	MEC
    SXC
    VOOT
    BITUSD
    CAI
    DIEM
    XSI
    ACH
    CNMT
    MAX
    NBT
    NSR
    XMR
    EMC2
    PAWN
    //	SYS
    //	BOST
    EFL
    GRC
    RDD
    STRAT
    TAC
    BTM
    JPC
    KDC
    MTS
    N5X
    //	BTC
    PRC
    UNITY
    BONES
    //	EAC
    FCT
    SILK
    GPUC
    SUN
    //	XCN
    BCN
    MZC
    UIS
    //	XRP
    GEO
    LOL
    DCR
    NTX
    //	ZEC
    PMC
    DVK
    //	LTC
    PAND
    YC
    GUE
    LCL
    BBR
    NL
    PRT
    //	XPM
    DSH
    PTS
    ULTC
    WC
    XCR
    NOXT
    UTC
    AIR
    BCY
    ENC
    LSK
    MMXIV
    SDC
    SOC
    TOR
    SSD
    UVC
    WOLF
    BBL
    GLB
    MMC
    MNTA
    RBY
    ADN
    BELA
    //	DOGE
    GNO
    SWARM
    BITS
    HVC
    ITC
    USDE
    AEON
    EXE
    XC
    ABY
    CACH
    EBT
    MIN
    NXTI
    FCN
    LQD
    MUN
    //	WDC
    XVC
    ARCH
    H2O
    DRM
    STR
    YACC
    //	BLK
    FIBRE
    HUC
    //	NRS
    PASC
    FRQ
    PINK
    //	PPC
    XPB
    QTUM
    STORJ
    OMG
    GAS
    //BCC
    ZRX
)

var currencyPairSymbol = [...]string{"btc_cny", "btc_usd", "btc_jpy", "fx_btc_jpy", "ltc_cny", "ltc_usd", "ltc_btc", "eth_cny",
    "eth_usd", "eth_btc", "etc_cny", "etc_usd", "etc_btc", "etc_eth", "xcn_btc", "sys_btc", "zec_cny", "zec_usd", "zec_btc"}

var CurrencyPairMap = map[string]CurrencyPair{
    "BTC_CNY":    BTC_CNY,
    "BTC_USD":    BTC_USD,
    "BTC_JPY":    BTC_JPY,
    "FX_BTC_JPY": FX_BTC_JPY,
    "LTC_CNY":    LTC_CNY,
    "LTC_USD":    LTC_USD,
    "ETH_CNY":    ETH_CNY,
    "ETH_USD":    ETH_USD,
    "ETC_CNY":    ETC_CNY,
    "ETC_USD":    ETC_USD,
    "ETC_ETH":    ETC_ETH,
    "ZEC_CNY":    ZEC_CNY,
    "ZEC_USD":    ZEC_USD,
    "REP_CNY":    REP_CNY,
    "REP_ETH":    REP_ETH,
    "XRP_CNY":    XRP_CNY,
    "XRP_USD":    XRP_USD,
    "DOGE_CNY":   DOGE_CNY,
    "DOGE_USD":   DOGE_USD,
    "BLK_CNY":    BLK_CNY,
    "BLK_USD":    BLK_USD,
    "LSK_CNY":    LSK_CNY,
    "LSK_USD":    LSK_USD,
    "GAME_CNY":   GAME_CNY,
    "GAME_USD":   GAME_USD,
    "SC_CNY":     SC_CNY,
    "SC_USD":     SC_USD,
    "GNT_CNY":    GNT_CNY,
    "BTS_CNY":    BTS_CNY,
    "BTS_USD":    BTS_USD,
    "HLB_CNY":    HLB_CNY,
    "HLB_USD":    HLB_USD,
    "HLB_BTC":    HLB_BTC,
    "XPM_CNY":    XPM_CNY,
    "XPM_USD":    XPM_USD,
    "XPM_BTC":    XPM_BTC,
    "RIC_CNY":    RIC_CNY,
    "RIC_USD":    RIC_USD,
    "RIC_BTC":    RIC_BTC,
    "XEM_CNY":    XEM_CNY,
    "XEM_USD":    XEM_USD,
    "EAC_CNY":    EAC_CNY,
    "EAC_USD":    EAC_USD,
    "EAC_BTC":    EAC_BTC,
    "PPC_CNY":    PPC_CNY,
    "PPC_USD":    PPC_USD,
    "PLC_CNY":    PLC_CNY,
    "PLC_USD":    PLC_USD,
    "VTC_CNY":    VTC_CNY,
    "VTC_USD":    VTC_USD,
    "VRC_CNY":    VRC_CNY,
    "VRC_USD":    VRC_USD,
    "NXT_CNY":    NXT_CNY,
    "NXT_USD":    NXT_USD,
    "ZCC_CNY":    ZCC_CNY,
    "ZCC_USD":    ZCC_USD,
    "WDC_CNY":    WDC_CNY,
    "WDC_USD":    WDC_USD,
    "WDC_BTC":    WDC_BTC,
    "SYS_CNY":    SYS_CNY,
    "SYS_USD":    SYS_USD,
    "DASH_CNY":   DASH_CNY,
    "DASH_USD":   DASH_USD,
    "YBC_CNY":    YBC_CNY,
    "YBC_USD":    YBC_USD,
    "XCN_BTC":    XCN_BTC,
    "BLK_BTC":    BLK_BTC,
    "BURST_BTC":  BURST_BTC,
    "BTCD_BTC":   BTCD_BTC,
    "CLAM_BTC":   CLAM_BTC,
    "DASH_BTC":   DASH_BTC,
    "DOGE_BTC":   DOGE_BTC,
    "EMC2_BTC":   EMC2_BTC,
    "FLDC_BTC":   FLDC_BTC,
    "GAME_BTC":   GAME_BTC,
    "MAID_BTC":   MAID_BTC,
    "OMNI_BTC":   OMNI_BTC,
    "NAUT_BTC":   NAUT_BTC,
    "NEOS_BTC":   NEOS_BTC,
    "PINK_BTC":   PINK_BTC,
    "SJCX_BTC":   SJCX_BTC,
    "DGB_BTC":    DGB_BTC,
    "FLO_BTC":    FLO_BTC,
    "GRC_BTC":    GRC_BTC,
    "LTC_BTC":    LTC_BTC,
    "NAV_BTC":    NAV_BTC,
    "POT_BTC":    POT_BTC,
    "PPC_BTC":    PPC_BTC,
    "NXT_BTC":    NXT_BTC,
    "BTS_BTC":    BTS_BTC,
    "SYS_BTC":    SYS_BTC,
    "VIA_BTC":    VIA_BTC,
    "XVC_BTC":    XVC_BTC,
    "VRC_BTC":    VRC_BTC,
    "VTC_BTC":    VTC_BTC,
    "XCP_BTC":    XCP_BTC,
    "XEM_BTC":    XEM_BTC,
    "XMR_BTC":    XMR_BTC,
    "XRP_BTC":    XRP_BTC,
    "ETH_BTC":    ETH_BTC,
    "SC_BTC":     SC_BTC,
    "BCY_BTC":    BCY_BTC,
    "EXP_BTC":    EXP_BTC,
    "FCT_BTC":    FCT_BTC,
    "AMP_BTC":    AMP_BTC,
    "DCR_BTC":    DCR_BTC,
    "LSK_BTC":    LSK_BTC,
    "LBC_BTC":    LBC_BTC,
    "SBD_BTC":    SBD_BTC,
    "ETC_BTC":    ETC_BTC,
    "REP_BTC":    REP_BTC,
    "ZEC_BTC":    ZEC_BTC,
    "NXC_BTC":    NXC_BTC,
    "GNT_BTC":    GNT_BTC,
    "GNO_BTC":    GNO_BTC,
    "QTUM_BTC":   QTUM_BTC,
    "QTUM_CNY":   QTUM_CNY,
    "RADS_BTC":   RADS_BTC,
    "ARDR_BTC":   ARDR_BTC,
    "STEEM_BTC":  STEEM_BTC,
    "STRAT_BTC":  STRAT_BTC,
    "STORJ_BTC":  STORJ_BTC,
    "OMG_BTC":    OMG_BTC,
    "GAS_BTC":    GAS_BTC,
    "BCC_BTC":    BCC_BTC,
    "ZRX_BTC":    ZRX_BTC,

}

const (
    BTC_CNY    = 1 + iota
    BTC_USD
    BTC_JPY
    FX_BTC_JPY

    LTC_CNY
    LTC_USD

    ETH_CNY
    ETH_USD

    ETC_CNY
    ETC_USD
    ETC_ETH

    ZEC_CNY
    ZEC_USD

    REP_CNY
    REP_ETH

    XRP_CNY
    XRP_USD

    DOGE_CNY
    DOGE_USD

    BLK_CNY
    BLK_USD

    LSK_CNY
    LSK_USD

    GAME_CNY
    GAME_USD

    SC_CNY
    SC_USD

    GNT_CNY

    BTS_CNY
    BTS_USD

    HLB_CNY
    HLB_USD
    HLB_BTC

    XPM_CNY
    XPM_USD
    XPM_BTC

    RIC_CNY
    RIC_USD
    RIC_BTC

    XEM_CNY
    XEM_USD

    EAC_CNY
    EAC_USD
    EAC_BTC

    PPC_CNY
    PPC_USD

    PLC_CNY
    PLC_USD

    VTC_CNY
    VTC_USD

    VRC_CNY
    VRC_USD

    NXT_CNY
    NXT_USD

    ZCC_CNY
    ZCC_USD

    WDC_CNY
    WDC_USD
    WDC_BTC

    SYS_CNY
    SYS_USD

    DASH_CNY
    DASH_USD

    YBC_CNY
    YBC_USD

    XCN_BTC
    BLK_BTC
    BTCD_BTC
    BTS_BTC
    BURST_BTC
    CLAM_BTC
    DASH_BTC
    DGB_BTC
    DOGE_BTC
    EMC2_BTC
    FLDC_BTC
    FLO_BTC
    GAME_BTC
    GRC_BTC
    LTC_BTC
    MAID_BTC
    OMNI_BTC
    NAUT_BTC
    NAV_BTC
    NEOS_BTC
    NXT_BTC
    PINK_BTC
    POT_BTC
    PPC_BTC
    SJCX_BTC
    SYS_BTC
    VIA_BTC
    XVC_BTC
    VRC_BTC
    VTC_BTC
    XCP_BTC
    XEM_BTC
    XMR_BTC
    XRP_BTC
    ETH_BTC
    SC_BTC
    BCY_BTC
    EXP_BTC
    FCT_BTC
    RADS_BTC
    AMP_BTC
    DCR_BTC
    LSK_BTC
    LBC_BTC
    STEEM_BTC
    SBD_BTC
    ETC_BTC
    REP_BTC
    ARDR_BTC
    ZEC_BTC
    STRAT_BTC
    NXC_BTC
    GNT_BTC
    GNO_BTC
    QTUM_BTC
    QTUM_CNY
    STORJ_BTC
    OMG_BTC
    GAS_BTC
    BCC_BTC
    ZRX_BTC
)

const (
    BUY         = 1 + iota
    SELL
    BUY_MARKET
    SELL_MARKET
)

var orderStatusSymbol = [...]string{"UNFINISH", "PART_FINISH", "FINISH", "CANCEL", "REJECT", "CANCEL_ING"}

const (
    ORDER_UNFINISH    = iota
    ORDER_PART_FINISH
    ORDER_FINISH
    ORDER_CANCEL
    ORDER_REJECT
    ORDER_CANCEL_ING
)

const (
    OPEN_BUY   = 1 + iota //开多
    OPEN_SELL             //开空
    CLOSE_BUY             //平多
    CLOSE_SELL            //平空
)

var CurrencyPairSymbol = map[CurrencyPair]string{
    BTC_CNY:  "btc_cny",
    BTC_USD:  "btc_usd",
    LTC_CNY:  "ltc_cny",
    LTC_USD:  "ltc_usd",
    LTC_BTC:  "ltc_btc",
    ETH_CNY:  "eth_cny",
    ETH_USD:  "eth_usd",
    ETH_BTC:  "eth_btc",
    ETC_CNY:  "etc_cny",
    ETC_USD:  "etc_usd",
    ETC_BTC:  "etc_btc",
    ETC_ETH:  "etc_eth",
    QTUM_CNY: "qtum_cny",
    QTUM_BTC: "qtum_btc",
    BTS_BTC:  "bts_btc",
    BTS_CNY:  "bts_cny"}

var (
    THIS_WEEK_CONTRACT = "this_week" //周合约
    NEXT_WEEK_CONTRACT = "next_week" //次周合约
    QUARTER_CONTRACT   = "quarter"   //季度合约
)

func SymbolPairCurrency(sss string) int {
    return 0
}
