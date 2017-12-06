package chbtc

import (
	"encoding/json"
	"errors"
	"fmt"
	. "github.com/qct/crypto_coin_api"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	MARKET_URL = "http://api.chbtc.com/data/v1/"
	TICKER_API = "ticker?currency=%s"
	DEPTH_API  = "depth?currency=%s&size=%d"

	TRADE_URL                 = "https://trade.chbtc.com/api/"
	GET_ACCOUNT_API           = "getAccountInfo"
	GET_ORDER_API             = "getOrder"
	GET_UNFINISHED_ORDERS_API = "getUnfinishedOrdersIgnoreTradeType"
	CANCEL_ORDER_API          = "cancelOrder"
	PLACE_ORDER_API           = "order"
	WITHDRAW_API              = "withdraw"
	CANCELWITHDRAW_API        = "cancelWithdraw"
	GET_WITHDRAWAL            = "getWithdrawRecord"
	GET_DEPOSIT               = "getChargeRecord"
)

type Chbtc struct {
	httpClient *http.Client
	accessKey,
	secretKey string
}

func New(httpClient *http.Client, accessKey, secretKey string) *Chbtc {
	return &Chbtc{httpClient, accessKey, secretKey}
}

func (chbtc *Chbtc) GetExchangeName() string {
	return "chbtc.com"
}

func (chbtc *Chbtc) GetTicker(currency CurrencyPair) (*Ticker, error) {
	resp, err := HttpGet(chbtc.httpClient, MARKET_URL+fmt.Sprintf(TICKER_API, CurrencyPairSymbol[currency]))
	if err != nil {
		return nil, err
	}
	//log.Println(resp)
	if _, ok := resp["ticker"]; !ok {
		return nil, fmt.Errorf("resp is not valid : %v", resp)
	}
	tickermap := resp["ticker"].(map[string]interface{})

	ticker := new(Ticker)
	ticker.Date, _ = strconv.ParseUint(resp["date"].(string), 10, 64)
	ticker.Buy, _ = strconv.ParseFloat(tickermap["buy"].(string), 64)
	ticker.Sell, _ = strconv.ParseFloat(tickermap["sell"].(string), 64)
	ticker.Last, _ = strconv.ParseFloat(tickermap["last"].(string), 64)
	ticker.High, _ = strconv.ParseFloat(tickermap["high"].(string), 64)
	ticker.Low, _ = strconv.ParseFloat(tickermap["low"].(string), 64)
	ticker.Vol, _ = strconv.ParseFloat(tickermap["vol"].(string), 64)

	return ticker, nil
}

func (chbtc *Chbtc) GetDepth(size int, currency CurrencyPair) (*Depth, error) {
	resp, err := HttpGet(chbtc.httpClient, MARKET_URL+fmt.Sprintf(DEPTH_API, CurrencyPairSymbol[currency], size))
	if err != nil {
		return nil, err
	}

	//log.Println(resp);

	asks := resp["asks"].([]interface{})
	bids := resp["bids"].([]interface{})

	//log.Println(asks)
	//log.Println(bids)

	depth := new(Depth)

	for _, e := range bids {
		var r DepthRecord
		ee := e.([]interface{})
		r.Amount = ee[1].(float64)
		r.Price = ee[0].(float64)

		depth.BidList = append(depth.BidList, r)
	}

	for _, e := range asks {
		var r DepthRecord
		ee := e.([]interface{})
		r.Amount = ee[1].(float64)
		r.Price = ee[0].(float64)

		depth.AskList = append(depth.AskList, r)
	}
	sort.Sort(DepthRecords(depth.AskList))
	return depth, nil

}

func (chbtc *Chbtc) buildPostForm(postForm *url.Values) error {
	postForm.Set("accesskey", chbtc.accessKey)

	payload := postForm.Encode()
	secretkeySha, _ := GetSHA(chbtc.secretKey)

	sign, err := GetParamHmacMD5Sign(secretkeySha, payload)
	if err != nil {
		return err
	}

	postForm.Set("sign", sign)
	//postForm.Del("secret_key")
	postForm.Set("reqTime", fmt.Sprintf("%d", time.Now().UnixNano()/1000000))
	return nil
}

func (chbtc *Chbtc) GetAccount() (*Account, error) {
	params := url.Values{}
	params.Set("method", "getAccountInfo")
	chbtc.buildPostForm(&params)
	//log.Println(params.Encode())
	resp, err := HttpPostForm(chbtc.httpClient, TRADE_URL+GET_ACCOUNT_API, params)
	if err != nil {
		return nil, err
	}

	var respmap map[string]interface{}
	err = json.Unmarshal(resp, &respmap)
	if err != nil {
		log.Println("json unmarshal error")
		return nil, err
	}

	if respmap["code"] != nil && respmap["code"].(float64) != 1000 {
		return nil, errors.New(string(resp))
	}

	acc := new(Account)
	acc.Exchange = "chbtc"
	acc.SubAccounts = make(map[Currency]SubAccount)

	resultmap := respmap["result"].(map[string]interface{})
	balancemap := resultmap["balance"].(map[string]interface{})
	frozenmap := resultmap["frozen"].(map[string]interface{})
	p2pmap := resultmap["p2p"].(map[string]interface{})
	netAssets, _ := strconv.ParseFloat(resultmap["netAssets"].(string), 64)
	asset, _ := strconv.ParseFloat(resultmap["totalAssets"].(string), 64)

	acc.NetAsset = netAssets
	acc.Asset = asset

	for t, v := range balancemap {
		vv := v.(map[string]interface{})
		subAcc := SubAccount{}
		subAcc.Amount, _ = strconv.ParseFloat(vv["amount"].(string), 64)

		switch t {
		case "CNY":
			subAcc.Currency = CNY
			cnyfrozen := frozenmap["CNY"].(map[string]interface{})
			subAcc.ForzenAmount, _ = strconv.ParseFloat(cnyfrozen["amount"].(string), 64)
			subAcc.LoanAmount = p2pmap["inCNY"].(float64)
		case "BTC":
			subAcc.Currency = BTC
			btcfrozen := frozenmap["BTC"].(map[string]interface{})
			subAcc.ForzenAmount, _ = strconv.ParseFloat(btcfrozen["amount"].(string), 64)
			subAcc.LoanAmount = p2pmap["inBTC"].(float64)
		case "LTC":
			subAcc.Currency = LTC
			ltcfrozen := frozenmap["LTC"].(map[string]interface{})
			subAcc.ForzenAmount, _ = strconv.ParseFloat(ltcfrozen["amount"].(string), 64)
			subAcc.LoanAmount = p2pmap["inLTC"].(float64)
		case "ETH":
			subAcc.Currency = ETH
			ethfrozen := frozenmap["ETH"].(map[string]interface{})
			subAcc.ForzenAmount, _ = strconv.ParseFloat(ethfrozen["amount"].(string), 64)
			subAcc.LoanAmount = p2pmap["inETH"].(float64)
		case "ETC":
			subAcc.Currency = ETC
			etcfrozen := frozenmap["ETC"].(map[string]interface{})
			subAcc.ForzenAmount, _ = strconv.ParseFloat(etcfrozen["amount"].(string), 64)
			subAcc.LoanAmount = p2pmap["inETC"].(float64)
		case "BTS":
			subAcc.Currency = BTS
			btsfrozen := frozenmap["BTS"].(map[string]interface{})
			subAcc.ForzenAmount, _ = strconv.ParseFloat(btsfrozen["amount"].(string), 64)
			subAcc.LoanAmount = p2pmap["inBTS"].(float64)
		/*case "EOS":
			subAcc.Currency = EOS
			btsfrozen := frozenmap["EOS"].(map[string]interface{})
			subAcc.ForzenAmount,_ = strconv.ParseFloat(btsfrozen["amount"].(string),64)
			subAcc.LoanAmount = p2pmap["inEOS"].(float64)
		case "QTUM":
			subAcc.Currency = QTUM
			qtumfrozen := frozenmap["QTUM"].(map[string]interface{})
			subAcc.ForzenAmount,_ = strconv.ParseFloat(qtumfrozen["amount"].(string),64)
			//subAcc.LoanAmount = p2pmap["inQTUM"].(float64)
		*/
		default:
			//log.Println("unknown ", t)

		}
		acc.SubAccounts[subAcc.Currency] = subAcc
	}

	//log.Println(string(resp))
	//log.Println(acc)

	return acc, nil
}

func (chbtc *Chbtc) placeOrder(amount, price string, currency CurrencyPair, tradeType int) (*Order, error) {
	params := url.Values{}
	params.Set("method", "order")
	params.Set("price", price)
	params.Set("amount", amount)
	params.Set("currency", CurrencyPairSymbol[currency])
	params.Set("tradeType", fmt.Sprintf("%d", tradeType))
	chbtc.buildPostForm(&params)

	resp, err := HttpPostForm(chbtc.httpClient, TRADE_URL+PLACE_ORDER_API, params)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	//log.Println(string(resp));

	respmap := make(map[string]interface{})
	err = json.Unmarshal(resp, &respmap)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	code := respmap["code"].(float64)
	if code != 1000 {
		log.Println(string(resp))
		return nil, errors.New(fmt.Sprintf("%.0f", code))
	}

	orid := respmap["id"].(string)

	order := new(Order)
	order.Amount, _ = strconv.ParseFloat(amount, 64)
	order.Price, _ = strconv.ParseFloat(price, 64)
	order.Status = ORDER_UNFINISH
	order.Currency = currency
	order.OrderTime = int(time.Now().UnixNano() / 1000000)
	order.OrderID, _ = strconv.Atoi(orid)

	switch tradeType {
	case 0:
		order.Side = SELL
	case 1:
		order.Side = BUY
	}

	return order, nil
}

func (chbtc *Chbtc) LimitBuy(amount, price string, currency CurrencyPair) (*Order, error) {
	return chbtc.placeOrder(amount, price, currency, 1)
}

func (chbtc *Chbtc) LimitSell(amount, price string, currency CurrencyPair) (*Order, error) {
	return chbtc.placeOrder(amount, price, currency, 0)
}

func (chbtc *Chbtc) CancelOrder(orderId string, currency CurrencyPair) (bool, error) {
	params := url.Values{}
	params.Set("method", "cancelOrder")
	params.Set("id", orderId)
	params.Set("currency", CurrencyPairSymbol[currency])
	chbtc.buildPostForm(&params)

	resp, err := HttpPostForm(chbtc.httpClient, TRADE_URL+CANCEL_ORDER_API, params)
	if err != nil {
		log.Println(err)
		return false, err
	}

	respmap := make(map[string]interface{})
	err = json.Unmarshal(resp, &respmap)
	if err != nil {
		log.Println(err)
		return false, err
	}

	code := respmap["code"].(float64)

	if code == 1000 {
		return true, nil
	}

	//log.Println(respmap)
	return false, errors.New(fmt.Sprintf("%.0f", code))
}

func parseOrder(order *Order, ordermap map[string]interface{}) {
	//order.Currency = currency;
	order.OrderID, _ = strconv.Atoi(ordermap["id"].(string))
	order.Amount = ordermap["total_amount"].(float64)
	order.DealAmount = ordermap["trade_amount"].(float64)
	order.Price = ordermap["price"].(float64)
	order.Fee = 0.0 //ordermap["fees"].(float64)
	if order.DealAmount > 0 {
		order.AvgPrice = ordermap["trade_money"].(float64) / order.DealAmount
	} else {
		order.AvgPrice = 0
	}

	order.OrderTime = int(ordermap["trade_date"].(float64))

	orType := ordermap["type"].(float64)
	switch orType {
	case 0:
		order.Side = SELL
	case 1:
		order.Side = BUY
	default:
		log.Printf("unknown order type %f", orType)
	}

	_status := TradeStatus(ordermap["status"].(float64))
	switch _status {
	case 0:
		order.Status = ORDER_UNFINISH
	case 1:
		order.Status = ORDER_CANCEL
	case 2:
		order.Status = ORDER_FINISH
	case 3:
		order.Status = ORDER_PART_FINISH

	}

}

func (chbtc *Chbtc) GetOneOrder(orderId string, currency CurrencyPair) (*Order, error) {
	params := url.Values{}
	params.Set("method", "getOrder")
	params.Set("id", orderId)
	params.Set("currency", CurrencyPairSymbol[currency])
	chbtc.buildPostForm(&params)

	resp, err := HttpPostForm(chbtc.httpClient, TRADE_URL+GET_ORDER_API, params)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//println(string(resp))
	ordermap := make(map[string]interface{})
	err = json.Unmarshal(resp, &ordermap)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	order := new(Order)
	order.Currency = currency

	parseOrder(order, ordermap)

	return order, nil
}

func (chbtc *Chbtc) GetUnfinishOrders(currency CurrencyPair) ([]Order, error) {
	params := url.Values{}
	params.Set("method", "getUnfinishedOrdersIgnoreTradeType")
	params.Set("currency", CurrencyPairSymbol[currency])
	params.Set("pageIndex", "1")
	params.Set("pageSize", "100")
	chbtc.buildPostForm(&params)

	resp, err := HttpPostForm(chbtc.httpClient, TRADE_URL+GET_UNFINISHED_ORDERS_API, params)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	respstr := string(resp)
	if strings.Contains(respstr, "\"code\":3001") {
		log.Println(respstr)
		return nil, nil
	}

	var resps []interface{}
	err = json.Unmarshal(resp, &resps)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var orders []Order
	for _, v := range resps {
		ordermap := v.(map[string]interface{})
		order := Order{}
		order.Currency = currency
		parseOrder(&order, ordermap)
		orders = append(orders, order)
	}

	return orders, nil
}

func (chbtc *Chbtc) GetOrderHistorys(currency CurrencyPair, currentPage, pageSize int) ([]Order, error) {
	return nil, nil
}

func (chbtc *Chbtc) GetKlineRecords(currency CurrencyPair, period string, size, since int) ([]Kline, error) {
	return nil, nil
}

/*
GET https://trade.chbtc.com/api/withdraw?accesskey=your_access_key
	&amount=0.01&currency=btc&fees=0.001&itransfer=0&method=withdraw
	&receiveAddr=14fxEPirL9fyfw1i9EF439Pq6gQ5xijUmp&safePwd=资金安全密码
	&sign=请求加密签名串&reqTime=当前时间毫秒数
*/
func (chbtc *Chbtc) Withdraw(amount string, currency Currency, fees, receiveAddr, safePwd string) (string, error) {
	params := url.Values{}
	params.Set("method", "withdraw")
	params.Set("currency", strings.ToLower(currency.String()))
	params.Set("amount", amount)
	params.Set("fees", fees)
	params.Set("receiveAddr", receiveAddr)
	params.Set("safePwd", safePwd)
	params.Set("itransfer", "0")
	chbtc.buildPostForm(&params)

	resp, err := HttpPostForm(chbtc.httpClient, TRADE_URL+WITHDRAW_API, params)
	if err != nil {
		log.Println("withdraw fail.", err)
		return "", err
	}

	respMap := make(map[string]interface{})
	err = json.Unmarshal(resp, &respMap)
	if err != nil {
		log.Println(err, string(resp))
		return "", err
	}

	if respMap["code"].(float64) == 1000 {
		return respMap["id"].(string), nil
	}

	return "", errors.New(string(resp))
}

/*
//# Request
GET https://trade.chbtc.com/api/getWithdrawRecord?method=getWithdrawRecord
	&accesskey=your_access_key&currency=btc&pageIndex=1&pageSize=10
	&sign=请求加密签名串&reqTime=当前时间毫秒数
//# Response
{
    "code": 1000,
    "message": {
        "des": "success",
        "isSuc": true,
        "datas": {
            "list": [
                {
                    "amount": 0.01,
                    "fees": 0.001,
                    "id": 2016042556231,
                    "manageTime": 1461579340000,
                    "status": 3,
                    "submitTime": 1461579288000,
                    "toAddress": "14fxEPirL9fyfw1i9EF439Pq6gQ5xijUmp"
                }...
            ],
            "pageIndex": 1,
            "pageSize": 10,
            "totalCount": 4,
            "totalPage": 1
        }
    }
}

//# Request
GET https://trade.chbtc.com/api/getChargeRecord?method=getChargeRecord
	&accesskey=your_access_key&currency=btc&pageIndex=1&pageSize=10
	&sign=请求加密签名串&reqTime=当前时间毫秒数
//# Response
{
    "code": 1000,
    "message": {
        "des": "success",
        "isSuc": true,
        "datas": {
            "list": [
                {
                    "address": "1FKN1DZqCm8HaTujDioRL2Aezdh7Qj7xxx",
                    "amount": "1.00000000",
                    "confirmTimes": 1,
                    "currency": "BTC",
                    "description": "确认成功",
                    "hash": "7ce842de187c379abafadd64a5fe66c5c61c8a21fb04edff9532234a1dae6xxx",
                    "id": 558,
                    "itransfer": 1,
                    "status": 2,
                    "submit_time": "2016-12-07 18:51:57"
                }...
            ],
            "pageIndex": 1,
            "pageSize": 10,
            "total": 8
        }
    }
}


code : 返回代码
message : 提示信息
amount : 充值金额
confirmTimes : 充值确认次数
currency : 充值货币类型(大写)
description : 充值记录状态描述
hash : 充值交易号
id : 充值记录id
itransfer : 是否内部转账，1是0否
status : 状态(0等待确认，1充值失败，2充值成功)
submit_time : 充值时间
address : 充值地址
*/

type DepositOneRecord struct {
	Amount       string `json:"amount"`
	Id           int    `json:"id"`
	Status       int    `json:"status"`
	Address      string `json:"address"`
	ConfirmTimes int    `json:"confirmTimes"`
	Currency     string `json:"currency"`
	Description  string `json:"description"`
	Hash         string `json:"hash"`
	Itransfer    int    `json:"itransfer"`
	Submit_Time  string `json:"submit_time"`
}
type WithdrawOneRecord struct {
	Amount     float64 `json:"amount"`
	Id         int     `json:"id"`
	Status     int     `json:"status"`
	Fees       float64 `json:"fees"`
	ManageTime int64   `json:"manageTime"`
	SubmitTime int64   `json:"submitTime"`
	ToAddress  string  `json:"toAddress"`
}

type WithrawDatasRecord struct {
	List       []WithdrawOneRecord `json:"list"`
	PageIndex  int                 `json:"pageIndex"`
	PageSize   int                 `json:"pageSize"`
	TotalCount int                 `json:"totalCount"`
	Total      int                 `json:"total"`
	TotalPage  int                 `json:"totalPage"`
}
type WithdrawMessageRecord struct {
	Des   string             `json:"des"`
	IsSuc bool               `json:"isSuc"`
	Datas WithrawDatasRecord `json:"datas"`
}
type WithdrawRecord struct {
	Code    int                   `JSON:"code"`
	Message WithdrawMessageRecord `json:"message"`
}

type DepositDatasRecord struct {
	List       []DepositOneRecord `json:"list"`
	PageIndex  int                `json:"pageIndex"`
	PageSize   int                `json:"pageSize"`
	TotalCount int                `json:"totalCount"`
	Total      int                `json:"total"`
	TotalPage  int                `json:"totalPage"`
}
type DepositMessageRecord struct {
	Des   string             `json:"des"`
	IsSuc bool               `json:"isSuc"`
	Datas DepositDatasRecord `json:"datas"`
}
type DepositRecord struct {
	Code    int                  `JSON:"code"`
	Message DepositMessageRecord `json:"message"`
}

func (chbtc *Chbtc) GetWithdrawal(currency Currency) (*WithdrawRecord, error) {
	//&currency=btc&pageIndex=1&pageSize=10
	params := url.Values{}
	params.Set("method", "getWithdrawRecord")
	params.Set("pageIndex", "1")
	params.Set("pageSize", "100")
	params.Set("currency", currency.String())
	chbtc.buildPostForm(&params)

	resp, err := HttpPostForm(chbtc.httpClient, TRADE_URL+GET_WITHDRAWAL, params)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//log.Println(string(resp))
	var r WithdrawRecord
	err = json.Unmarshal(resp, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (chbtc *Chbtc) GetDeposit(currency Currency) (*DepositRecord, error) {
	//&currency=btc&pageIndex=1&pageSize=10
	params := url.Values{}
	params.Set("method", "getChargeRecord")
	params.Set("pageIndex", "1")
	params.Set("pageSize", "100")
	params.Set("currency", currency.String())
	chbtc.buildPostForm(&params)

	resp, err := HttpPostForm(chbtc.httpClient, TRADE_URL+GET_DEPOSIT, params)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//log.Println(string(resp))
	var r DepositRecord
	err = json.Unmarshal(resp, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (chbtc *Chbtc) CancelWithdraw(id string, currency Currency, safePwd string) (bool, error) {
	params := url.Values{}
	params.Set("method", "cancelWithdraw")
	params.Set("currency", strings.ToLower(currency.String()))
	params.Set("downloadId", id)
	params.Set("safePwd", safePwd)
	chbtc.buildPostForm(&params)

	resp, err := HttpPostForm(chbtc.httpClient, TRADE_URL+CANCELWITHDRAW_API, params)
	if err != nil {
		log.Println("cancel withdraw fail.", err)
		return false, err
	}

	respMap := make(map[string]interface{})
	err = json.Unmarshal(resp, &respMap)
	if err != nil {
		log.Println(err, string(resp))
		return false, err
	}

	if respMap["code"].(float64) == 1000 {
		return true, nil
	}

	return false, errors.New(string(resp))
}

func (chbtc *Chbtc) GetTrades(currencyPair CurrencyPair, since int64) ([]Trade, error) {
	panic("unimplements")
}

func (chbtc *Chbtc) MarketBuy(amount, price string, currency CurrencyPair) (*Order, error) {
	panic("unsupport the market order")
}

func (chbtc *Chbtc) MarketSell(amount, price string, currency CurrencyPair) (*Order, error) {
	panic("unsupport the market order")
}
