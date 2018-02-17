package poloniex

import (
    "encoding/json"
    "errors"
    "fmt"
    . "github.com/oiovoyo/crypto_coin_api"
    "log"
    "net/http"
    "net/url"
    "strconv"
    "strings"
    "time"
)

const EXCHANGE_NAME = "poloniex.com"

const (
    BASE_URL       = "https://poloniex.com/"
    TRADE_API      = BASE_URL + "tradingApi"
    PUBLIC_URL     = BASE_URL + "public"
    TICKER_API     = "?command=returnTicker"
    CURRENCIES_API = "?command=returnCurrencies"
    ORDER_BOOK_API = "?command=returnOrderBook&currencyPair=%s&depth=%d"
)

var _CURRENCYPAIR_TO_SYMBOL = map[CurrencyPair]string{
    //BTC_LTC : "BTC_LTC",
    ETC_ETH: "ETH_ETC",
    REP_ETH: "ETH_REP",

    BLK_BTC:   "BTC_BLK",
    BTCD_BTC:  "BTC_BTCD",
    BTS_BTC:   "BTC_BTS",
    BURST_BTC: "BTC_BURST",
    CLAM_BTC:  "BTC_CLAM",
    DASH_BTC:  "BTC_DASH",
    DGB_BTC:   "BTC_DGB",
    DOGE_BTC:  "BTC_DOGE",
    EMC2_BTC:  "BTC_EMC2",
    FLDC_BTC:  "BTC_FLDC",
    FLO_BTC:   "BTC_FLO",
    GAME_BTC:  "BTC_GAME",
    GRC_BTC:   "BTC_GRC",
    LTC_BTC:   "BTC_LTC",
    MAID_BTC:  "BTC_MAID",
    OMNI_BTC:  "BTC_OMNI",
    NAUT_BTC:  "BTC_NAUT",
    NAV_BTC:   "BTC_NAV",
    NEOS_BTC:  "BTC_NEOS",
    NXT_BTC:   "BTC_NXT",
    PINK_BTC:  "BTC_PINK",
    POT_BTC:   "BTC_POT",
    PPC_BTC:   "BTC_PPC",
    SJCX_BTC:  "BTC_SJCX",
    SYS_BTC:   "BTC_SYS",
    VIA_BTC:   "BTC_VIA",
    XVC_BTC:   "BTC_XVC",
    VRC_BTC:   "BTC_VRC",
    VTC_BTC:   "BTC_VTC",
    XCP_BTC:   "BTC_XCP",
    XEM_BTC:   "BTC_XEM",
    XMR_BTC:   "BTC_XMR",
    XRP_BTC:   "BTC_XRP",
    ETH_BTC:   "BTC_ETH",
    SC_BTC:    "BTC_SC",
    BCY_BTC:   "BTC_BCY",
    EXP_BTC:   "BTC_EXP",
    FCT_BTC:   "BTC_FCT",
    RADS_BTC:  "BTC_RADS",
    AMP_BTC:   "BTC_AMP",
    DCR_BTC:   "BTC_DCR",
    LSK_BTC:   "BTC_LSK",
    LBC_BTC:   "BTC_LBC",
    STEEM_BTC: "BTC_STEEM",
    SBD_BTC:   "BTC_SBD",
    ETC_BTC:   "BTC_ETC",
    REP_BTC:   "BTC_REP",
    ARDR_BTC:  "BTC_ARDR",
    ZEC_BTC:   "BTC_ZEC",
    STRAT_BTC: "BTC_STRAT",
    NXC_BTC:   "BTC_NXC",
    GNT_BTC:   "BTC_GNT",
    GNO_BTC:   "BTC_GNO",
    QTUM_BTC:  "BTC_QTUM",

    STORJ_BTC: "BTC_STORJ",
    OMG_BTC:   "BTC_OMG",
    GAS_BTC:   "BTC_GAS",
    BCC_BTC:   "BTC_BCC",
    ZRX_BTC:   "BTC_ZRX",
}

type Poloniex struct {
    accessKey,
    secretKey string
    client *http.Client
}

func New(client *http.Client, accessKey, secretKey string) *Poloniex {
    return &Poloniex{accessKey, secretKey, client}
}

func (poloniex *Poloniex) GetExchangeName() string {
    return EXCHANGE_NAME
}
func (poloniex *Poloniex)GetTickers() (map[string]Ticker, error) {
    respmap, err := HttpGet(poloniex.client, PUBLIC_URL+TICKER_API)
    if err != nil {
        log.Println(err)
        return nil, err
    }

    tickers := make(map[string]Ticker)
    for k, v := range respmap {
        tickermap := v.(map[string]interface{})

        ticker := new(Ticker)
        ticker.High, _ = strconv.ParseFloat(tickermap["high24hr"].(string), 64)
        ticker.Low, _ = strconv.ParseFloat(tickermap["low24hr"].(string), 64)
        ticker.Last, _ = strconv.ParseFloat(tickermap["last"].(string), 64)
        ticker.Buy, _ = strconv.ParseFloat(tickermap["highestBid"].(string), 64)
        ticker.Sell, _ = strconv.ParseFloat(tickermap["lowestAsk"].(string), 64)
        ticker.Vol, _ = strconv.ParseFloat(tickermap["quoteVolume"].(string), 64)
        tickers[k] = *ticker
    }


    return tickers, nil
}
func (poloniex *Poloniex) GetTicker(currency CurrencyPair) (*Ticker, error) {
    respmap, err := HttpGet(poloniex.client, PUBLIC_URL+TICKER_API)
    if err != nil {
        log.Println(err)
        return nil, err
    }

    tickermap := respmap[_CURRENCYPAIR_TO_SYMBOL[currency]].(map[string]interface{})

    ticker := new(Ticker)
    ticker.High, _ = strconv.ParseFloat(tickermap["high24hr"].(string), 64)
    ticker.Low, _ = strconv.ParseFloat(tickermap["low24hr"].(string), 64)
    ticker.Last, _ = strconv.ParseFloat(tickermap["last"].(string), 64)
    ticker.Buy, _ = strconv.ParseFloat(tickermap["highestBid"].(string), 64)
    ticker.Sell, _ = strconv.ParseFloat(tickermap["lowestAsk"].(string), 64)
    ticker.Vol, _ = strconv.ParseFloat(tickermap["quoteVolume"].(string), 64)

    log.Println(tickermap)

    return ticker, nil
}
func (poloniex *Poloniex) GetDepth(size int, currency CurrencyPair) (*Depth, error) {
    respmap, err := HttpGet(poloniex.client, PUBLIC_URL+fmt.Sprintf(ORDER_BOOK_API, _CURRENCYPAIR_TO_SYMBOL[currency], size))
    if err != nil {
        log.Println(err)
        return nil, err
    }

    if respmap["asks"] == nil {
        log.Println(respmap)
        return nil, errors.New(fmt.Sprintf("%+v", respmap))
    }

    _, isOK := respmap["asks"].([]interface{})
    if !isOK {
        log.Println(respmap)
        return nil, errors.New(fmt.Sprintf("%+v", respmap))
    }

    var depth Depth

    for _, v := range respmap["asks"].([]interface{}) {
        var dr DepthRecord
        for i, vv := range v.([]interface{}) {
            switch i {
            case 0:
                dr.Price, _ = strconv.ParseFloat(vv.(string), 64)
            case 1:
                dr.Amount = vv.(float64)
            }
        }
        depth.AskList = append(depth.AskList, dr)
    }

    for _, v := range respmap["bids"].([]interface{}) {
        var dr DepthRecord
        for i, vv := range v.([]interface{}) {
            switch i {
            case 0:
                dr.Price, _ = strconv.ParseFloat(vv.(string), 64)
            case 1:
                dr.Amount = vv.(float64)
            }
        }
        depth.BidList = append(depth.BidList, dr)
    }

    return &depth, nil
}
func (Poloniex *Poloniex) GetKlineRecords(currency CurrencyPair, period string, size, since int) ([]Kline, error) {
    return nil, nil
}

func (poloniex *Poloniex) placeLimitOrder(command, amount, price string, currency CurrencyPair) (*Order, error) {
    postData := url.Values{}
    postData.Set("command", command)
    postData.Set("currencyPair", _CURRENCYPAIR_TO_SYMBOL[currency])
    postData.Set("rate", price)
    postData.Set("amount", amount)

    sign, _ := poloniex.buildPostForm(&postData)

    headers := map[string]string{
        "Key":  poloniex.accessKey,
        "Sign": sign}

    resp, err := HttpPostForm2(poloniex.client, TRADE_API, postData, headers)
    if err != nil {
        log.Println(err)
        return nil, err
    }

    respmap := make(map[string]interface{})
    err = json.Unmarshal(resp, &respmap)
    if err != nil || respmap["error"] != nil {
        log.Println(err, string(resp))
        return nil, err
    }

    orderNumber := respmap["orderNumber"].(string)
    order := new(Order)
    order.OrderTime = int(time.Now().Unix() * 1000)
    order.OrderID, _ = strconv.Atoi(orderNumber)
    order.Amount, _ = strconv.ParseFloat(amount, 64)
    order.Price, _ = strconv.ParseFloat(price, 64)
    order.Status = ORDER_UNFINISH
    order.Currency = currency

    switch command {
    case "sell":
        order.Side = SELL
    case "buy":
        order.Side = BUY
    }

    log.Println(string(resp))
    return order, nil
}

func (poloniex *Poloniex) LimitBuy(amount, price string, currency CurrencyPair) (*Order, error) {
    return poloniex.placeLimitOrder("buy", amount, price, currency)
}

func (poloniex *Poloniex) LimitSell(amount, price string, currency CurrencyPair) (*Order, error) {
    return poloniex.placeLimitOrder("sell", amount, price, currency)
}

func (poloniex *Poloniex) CancelOrder(orderId string, currency CurrencyPair) (bool, error) {
    postData := url.Values{}
    postData.Set("command", "cancelOrder")
    postData.Set("orderNumber", orderId)

    sign, err := poloniex.buildPostForm(&postData)
    if err != nil {
        log.Println(err)
        return false, err
    }

    headers := map[string]string{
        "Key":  poloniex.accessKey,
        "Sign": sign}
    resp, err := HttpPostForm2(poloniex.client, TRADE_API, postData, headers)
    if err != nil {
        log.Println(err)
        return false, err
    }

    //log.Println(string(resp));

    respmap := make(map[string]interface{})
    err = json.Unmarshal(resp, &respmap)
    if err != nil || respmap["error"] != nil {
        log.Println(err, string(resp))
        return false, err
    }

    success := int(respmap["success"].(float64))
    if success != 1 {
        log.Println(respmap)
        return false, nil
    }

    return true, nil
}

func (poloniex *Poloniex) GetOneOrder(orderId string, currency CurrencyPair) (*Order, error) {
    postData := url.Values{}
    postData.Set("command", "returnOrderTrades")
    postData.Set("orderNumber", orderId)

    sign, _ := poloniex.buildPostForm(&postData)

    headers := map[string]string{
        "Key":  poloniex.accessKey,
        "Sign": sign}

    resp, err := HttpPostForm2(poloniex.client, TRADE_API, postData, headers)
    if err != nil {
        log.Println(err)
        return nil, err
    }
    //println(string(resp))
    if strings.Contains(string(resp), "error") {
        ords, err1 := poloniex.GetUnfinishOrders(currency)
        if err1 != nil {
            log.Println(err1)
        } else {
            _ordId, _ := strconv.Atoi(orderId)

            for _, ord := range ords {
                if ord.OrderID == _ordId {
                    return &ord, nil
                }
            }
        }
        //log.Println(string(resp))
        return nil, errors.New(string(resp))
    }

    respmap := make([]interface{}, 0)
    err = json.Unmarshal(resp, &respmap)
    if err != nil {
        log.Println(err, string(resp))
        return nil, err
    }

    order := new(Order)
    order.OrderID, _ = strconv.Atoi(orderId)
    order.Currency = currency

    total := 0.0

    for _, v := range respmap {
        vv := v.(map[string]interface{})
        _amount, _ := strconv.ParseFloat(vv["amount"].(string), 64)
        _rate, _ := strconv.ParseFloat(vv["rate"].(string), 64)
        _fee, _ := strconv.ParseFloat(vv["fee"].(string), 64)

        order.DealAmount += _amount
        total += (_amount * _rate)
        order.Fee = _fee

        if strings.Compare("sell", vv["type"].(string)) == 0 {
            order.Side = TradeSide(SELL)
        } else {
            order.Side = TradeSide(BUY)
        }
    }

    order.AvgPrice = total / order.DealAmount

    return order, nil
}

func (poloniex *Poloniex) GetUnfinishOrders(currency CurrencyPair) ([]Order, error) {
    postData := url.Values{}
    postData.Set("command", "returnOpenOrders")
    postData.Set("currencyPair", _CURRENCYPAIR_TO_SYMBOL[currency])

    sign, err := poloniex.buildPostForm(&postData)
    if err != nil {
        log.Println(err)
        return nil, err
    }

    headers := map[string]string{
        "Key":  poloniex.accessKey,
        "Sign": sign}
    resp, err := HttpPostForm2(poloniex.client, TRADE_API, postData, headers)
    if err != nil {
        log.Println(err)
        return nil, err
    }

    orderAr := make([]interface{}, 1)
    err = json.Unmarshal(resp, &orderAr)
    if err != nil {
        log.Println(err, string(resp))
        return nil, err
    }

    orders := make([]Order, 0)
    for _, v := range orderAr {
        vv := v.(map[string]interface{})
        order := Order{}
        order.Currency = currency
        order.OrderID, _ = strconv.Atoi(vv["orderNumber"].(string))
        order.Amount, _ = strconv.ParseFloat(vv["amount"].(string), 64)
        order.Price, _ = strconv.ParseFloat(vv["rate"].(string), 64)
        order.Status = ORDER_UNFINISH

        side := vv["type"].(string)
        switch side {
        case "buy":
            order.Side = TradeSide(BUY)
        case "sell":
            order.Side = TradeSide(SELL)
        }

        orders = append(orders, order)
    }

    //log.Println(orders)
    return orders, nil
}
func (Poloniex *Poloniex) GetOrderHistorys(currency CurrencyPair, currentPage, pageSize int) ([]Order, error) {
    return nil, nil
}

func (poloniex *Poloniex) GetAccount() (*Account, error) {
    postData := url.Values{}
    postData.Add("command", "returnCompleteBalances")
    sign, err := poloniex.buildPostForm(&postData)
    if err != nil {
        return nil, err
    }

    headers := map[string]string{
        "Key":  poloniex.accessKey,
        "Sign": sign}
    resp, err := HttpPostForm2(poloniex.client, TRADE_API, postData, headers)

    if err != nil {
        log.Println(err)
        return nil, err
    }

    respmap := make(map[string]interface{})
    err = json.Unmarshal(resp, &respmap)

    if err != nil || respmap["error"] != nil {
        log.Println(err)
        return nil, err
    }

    acc := new(Account)
    acc.Exchange = EXCHANGE_NAME
    acc.SubAccounts = make(map[Currency]SubAccount)

    for k, v := range respmap {
        var currency Currency

        switch k {
        case "BTC":
            currency = BTC
        case "LTC":
            currency = LTC
        case "ETH":
            currency = ETH
        case "ETC":
            currency = ETC
        case "USD":
            currency = USD
        case "BTS":
            currency = BTS
        default:
            currency = -1
        }

        if currency > 0 {
            vv := v.(map[string]interface{})
            subAcc := SubAccount{}
            subAcc.Currency = currency
            subAcc.Amount, _ = strconv.ParseFloat(vv["available"].(string), 64)
            subAcc.ForzenAmount, _ = strconv.ParseFloat(vv["onOrders"].(string), 64)
            acc.SubAccounts[subAcc.Currency] = subAcc
        }
    }

    return acc, nil
}

func (p *Poloniex) Withdraw(amount string, currency Currency, fees, receiveAddr, safePwd string) (string, error) {
    params := url.Values{}
    params.Add("command", "withdraw")
    params.Add("address", receiveAddr)
    params.Add("amount", amount)
    params.Add("currency", strings.ToUpper(currency.String()))

    sign, err := p.buildPostForm(&params)
    if err != nil {
        return "", err
    }

    headers := map[string]string{
        "Key":  p.accessKey,
        "Sign": sign}

    resp, err := HttpPostForm2(p.client, TRADE_API, params, headers)

    if err != nil {
        log.Println(err)
        return "", err
    }
    println(string(resp))

    respMap := make(map[string]interface{})

    err = json.Unmarshal(resp, &respMap)
    if err != nil {
        log.Println(err)
        return "", err
    }

    if respMap["error"] == nil {
        return string(resp), nil
    }

    return "", errors.New(string(resp))
}

func (p *Poloniex) WithdrawWithMemo(amount string, currency Currency, paymentId, receiveAddr, safePwd string) (string, error) {
    params := url.Values{}
    params.Add("command", "withdraw")
    params.Add("address", receiveAddr)
    params.Add("amount", amount)
    params.Add("currency", strings.ToUpper(currency.String()))
    params.Add("paymentId", paymentId)

    sign, err := p.buildPostForm(&params)
    if err != nil {
        return "", err
    }

    headers := map[string]string{
        "Key":  p.accessKey,
        "Sign": sign}

    resp, err := HttpPostForm2(p.client, TRADE_API, params, headers)

    if err != nil {
        log.Println(err)
        return "", err
    }
    println(string(resp))

    respMap := make(map[string]interface{})

    err = json.Unmarshal(resp, &respMap)
    if err != nil {
        log.Println(err)
        return "", err
    }

    if respMap["error"] == nil {
        return string(resp), nil
    }

    return "", errors.New(string(resp))
}

type PoloniexDepositsWithdrawals struct {
    Deposits []struct {
        Currency      string  `json:"currency"`
        Address       string  `json:"address"`
        Amount        float64 `json:"amount,string"`
        Confirmations int     `json:"confirmations"`
        TransactionID string  `json:"txid"`
        Timestamp     int     `json:"timestamp"`
        //Timestamp     time.Time `json:"timestamp"`
        Status string `json:"status"`
    } `json:"deposits"`
    Withdrawals []struct {
        WithdrawalNumber int64   `json:"withdrawalNumber"`
        Currency         string  `json:"currency"`
        Address          string  `json:"address"`
        Amount           float64 `json:"amount,string"`
        Confirmations    int     `json:"confirmations"`
        TransactionID    string  `json:"txid"`
        Timestamp        int     `json:"timestamp"`
        //Timestamp        time.Time `json:"timestamp"`
        Status    string `json:"status"`
        IPAddress string `json:"ipAddress"`
    } `json:"withdrawals"`
}

func (poloniex *Poloniex) GetDepositsWithdrawals(start, end string) (*PoloniexDepositsWithdrawals, error) {
    params := url.Values{}
    params.Set("command", "returnDepositsWithdrawals")
    println(start)
    if start != "" {
        params.Set("start", start)
    } else {
        params.Set("start", "0")
    }

    if end != "" {
        params.Set("end", end)
    } else {
        params.Set("end", strconv.FormatInt(time.Now().Unix(), 10))
    }

    sign, err := poloniex.buildPostForm(&params)
    if err != nil {
        return nil, err
    }

    headers := map[string]string{
        "Key":  poloniex.accessKey,
        "Sign": sign}

    resp, err := HttpPostForm2(poloniex.client, TRADE_API, params, headers)
    if err != nil {
        log.Println(err)
        return nil, err
    }

    println(string(resp))

    records := new(PoloniexDepositsWithdrawals)
    err = json.Unmarshal(resp, records)

    return records, err
}

func (poloniex *Poloniex) buildPostForm(postForm *url.Values) (string, error) {
    postForm.Add("nonce", fmt.Sprintf("%d", time.Now().UnixNano()+500000000000))
    payload := postForm.Encode()
    //println(payload)
    sign, err := GetParamHmacSHA512Sign(poloniex.secretKey, payload)
    if err != nil {
        return "", err
    }
    //log.Println(sign)
    return sign, nil
}

func (poloniex *Poloniex) GetTrades(currencyPair CurrencyPair, since int64) ([]Trade, error) {
    panic("unimplements")
}

func (poloniex *Poloniex) MarketBuy(amount, price string, currency CurrencyPair) (*Order, error) {
    panic("unsupport the market order")
}

func (poloniex *Poloniex) MarketSell(amount, price string, currency CurrencyPair) (*Order, error) {
    panic("unsupport the market order")
}

func (poloniex *Poloniex) GetCurrency(currency Currency) (*PoloniexCurrency, error) {
    respmap, err := HttpGet(poloniex.client, PUBLIC_URL+CURRENCIES_API)

    if err != nil || respmap["error"] != nil {
        log.Println(err)
        return nil, err
    }

    currencyMap := respmap[strings.ToUpper(currency.String())].(map[string]interface{})

    poloniexCurrency := new(PoloniexCurrency)
    poloniexCurrency.ID = int(currencyMap["id"].(float64))
    poloniexCurrency.Name, _ = currencyMap["name"].(string)
    poloniexCurrency.TxFee, _ = strconv.ParseFloat(currencyMap["txFee"].(string), 64)
    poloniexCurrency.MinConf = int(currencyMap["minConf"].(float64))
    poloniexCurrency.DepositAddress, _ = currencyMap["depositAddress"].(string)
    poloniexCurrency.Disabled = int(currencyMap["disabled"].(float64))
    poloniexCurrency.Delisted = int(currencyMap["delisted"].(float64))
    poloniexCurrency.Frozen = int(currencyMap["frozen"].(float64))

    return poloniexCurrency, nil
}

func (poloniex *Poloniex) GetAllCurrencies() (map[string]*PoloniexCurrency, error) {
    respmap, err := HttpGet(poloniex.client, PUBLIC_URL+CURRENCIES_API)

    if err != nil || respmap["error"] != nil {
        log.Println(err)
        return nil, err
    }

    result := map[string]*PoloniexCurrency{}
    for k, v := range respmap {
        currencyMap := v.(map[string]interface{})
        poloniexCurrency := new(PoloniexCurrency)
        poloniexCurrency.ID = int(currencyMap["id"].(float64))
        poloniexCurrency.Name, _ = currencyMap["name"].(string)
        poloniexCurrency.TxFee, _ = strconv.ParseFloat(currencyMap["txFee"].(string), 64)
        poloniexCurrency.MinConf = int(currencyMap["minConf"].(float64))
        poloniexCurrency.DepositAddress, _ = currencyMap["depositAddress"].(string)
        poloniexCurrency.Disabled = int(currencyMap["disabled"].(float64))
        poloniexCurrency.Delisted = int(currencyMap["delisted"].(float64))
        poloniexCurrency.Frozen = int(currencyMap["frozen"].(float64))

        result[k] = poloniexCurrency
    }
    return result, nil
}

func (poloniex *Poloniex) GetDepositAdresses() (map[string]string, error) {
    params := url.Values{}
    params.Set("command", "returnDepositAddresses")
    sign, err := poloniex.buildPostForm(&params)
    if err != nil {
        return nil, err
    }

    headers := map[string]string{
        "Key":  poloniex.accessKey,
        "Sign": sign}

    resp, err := HttpPostForm2(poloniex.client, TRADE_API, params, headers)
    if err != nil {
        log.Println(err)
        return nil, err
    }

    result := map[string]string{}
    err = json.Unmarshal(resp, &result)
    return result, nil
}

func (poloniex *Poloniex) GenerateNewAddress(currency string) (string, error) {

    params := url.Values{}
    params.Set("command", "generateNewAddress")
    params.Set("currency", currency)
    sign, err := poloniex.buildPostForm(&params)
    if err != nil {
        return "", err
    }

    headers := map[string]string{
        "Key":  poloniex.accessKey,
        "Sign": sign}

    resp, err := HttpPostForm2(poloniex.client, TRADE_API, params, headers)
    if err != nil {
        log.Println(err)
        return "", err
    }
    type Result struct {
        Success  int    `json:"success"`
        Response string `json:"response"`
    }
    result := Result{}
    err = json.Unmarshal(resp, &result)
    if err != nil {
        return "", err
    }
    if result.Success == 1 {
        return result.Response, nil
    }
    return "", fmt.Errorf("response not correct %+v", string(resp))

}
