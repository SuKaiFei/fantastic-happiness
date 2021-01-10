package juhe

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
	"time"
)

const (
	// _apiGetShAllFinanceStock 沪股列表
	// 接口地址：http://web.juhe.cn:8080/finance/stock/shall
	// 返回格式：json
	// 请求方式：get/post
	// 请求示例：http://web.juhe.cn:8080/finance/stock/shall?key=您申请的KEY&page=1
	_apiGetShAllFinanceStock = _baseURL + "/finance/stock/shall?key=%s&page=%d&stock=a&type=4"

	// _apiGetSzAllFinanceStock 深圳股市列表
	// 接口地址：http://web.juhe.cn:8080/finance/stock/szall
	// 返回格式：json
	// 请求方式：get/post
	// 请求示例：http://web.juhe.cn:8080/finance/stock/szall?key=您申请的APPKEY&page=1
	_apiGetSzAllFinanceStock = _baseURL + "/finance/stock/szall?key=%s&page=%d&stock=a&type=4"

	// _apiGetHSFinanceStock 沪深股市
	// 接口地址：http://web.juhe.cn:8080/finance/stock/hs
	// 返回格式：json
	// 请求方式：get
	// 请求示例：http://web.juhe.cn:8080/finance/stock/hs?gid=sh601009&key=您申请的APPKEY
	_apiGetHSFinanceStock = _baseURL + "/finance/stock/hs?key=%s&gid=%s"
)

// GetHSFinanceStock 沪深股市
func GetHSFinanceStock(key, gid string) (resp *HSFinanceStockResp, err error) {
	statusCode, body, err := fasthttp.GetTimeout(nil, fmt.Sprintf(_apiGetHSFinanceStock, key, gid), time.Minute)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if statusCode != fasthttp.StatusOK {
		return nil, errors.New(string(body))
	}

	resp = new(HSFinanceStockResp)
	if err := jsoniter.Unmarshal(body, resp); err != nil {
		return nil, errors.WithStack(err)
	}
	if err := resp.Error(); err != nil {
		return nil, errors.WithStack(err)
	}

	return
}

// GetShAllFinanceStock 沪股列表
func GetShAllFinanceStock(key string, page uint32) (resp *ShAllFinanceStockResp, err error) {
	statusCode, body, err := fasthttp.GetTimeout(nil, fmt.Sprintf(_apiGetShAllFinanceStock, key, page), time.Minute)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if statusCode != fasthttp.StatusOK {
		return nil, errors.New(string(body))
	}

	resp = new(ShAllFinanceStockResp)
	if err := jsoniter.Unmarshal(body, resp); err != nil {
		return nil, errors.WithStack(err)
	}

	return
}

// GetShAllFinanceStock 深圳股市列表
func GetSzAllFinanceStock(key string, page uint32) (resp *SzAllFinanceStockResp, err error) {
	statusCode, body, err := fasthttp.GetTimeout(nil, fmt.Sprintf(_apiGetSzAllFinanceStock, key, page), time.Minute)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if statusCode != fasthttp.StatusOK {
		return nil, errors.New(string(body))
	}

	resp = new(SzAllFinanceStockResp)
	if err := jsoniter.Unmarshal(body, resp); err != nil {
		return nil, errors.WithStack(err)
	}

	return
}

// SzAllFinanceStockResp 深圳股市列表响应
type SzAllFinanceStockResp struct {
	CommonResp
	Result struct {
		TotalCount string `json:"totalCount"`
		Page       string `json:"page"`
		Num        string `json:"num"`
		Data       []struct {
			Symbol        string  `json:"symbol"`
			Name          string  `json:"name"`
			Trade         string  `json:"trade"`
			Pricechange   float64 `json:"pricechange"`
			Changepercent float64 `json:"changepercent"`
			Buy           string  `json:"buy"`
			Sell          string  `json:"sell"`
			Settlement    string  `json:"settlement"`
			Open          string  `json:"open"`
			High          string  `json:"high"`
			Low           string  `json:"low"`
			Volume        int     `json:"volume"`
			Amount        int     `json:"amount"`
			Code          string  `json:"code"`
			Ticktime      string  `json:"ticktime"`
			Per           float64 `json:"per"`
			Pb            float64 `json:"pb"`
			Mktcap        float64 `json:"mktcap"`
			Nmc           float64 `json:"nmc"`
			Turnoverratio float64 `json:"turnoverratio"`
		} `json:"data"`
	} `json:"result"`
}

// ShAllFinanceStockResp 沪股列表响应
type ShAllFinanceStockResp struct {
	CommonResp
	Result struct {
		TotalCount string `json:"totalCount"`
		Page       string `json:"page"`
		Num        string `json:"num"`
		Data       []struct {
			Symbol        string  `json:"symbol"`
			Name          string  `json:"name"`
			Trade         string  `json:"trade"`
			Pricechange   float64 `json:"pricechange"`
			Changepercent float64 `json:"changepercent"`
			Buy           string  `json:"buy"`
			Sell          string  `json:"sell"`
			Settlement    string  `json:"settlement"`
			Open          string  `json:"open"`
			High          string  `json:"high"`
			Low           string  `json:"low"`
			Volume        int     `json:"volume"`
			Amount        int     `json:"amount"`
			Code          string  `json:"code"`
			Ticktime      string  `json:"ticktime"`
			Per           float64 `json:"per"`
			Pb            float64 `json:"pb"`
			Mktcap        float64 `json:"mktcap"`
			Nmc           float64 `json:"nmc"`
			Turnoverratio float64 `json:"turnoverratio"`
		} `json:"data"`
	} `json:"result"`
}

// HSFinanceStockResp 沪深股市响应
type HSFinanceStockResp struct {
	CommonResp
	Result []struct {
		Data struct {
			Gid            string `json:"gid"`
			IncrePer       string `json:"increPer"`
			Increase       string `json:"increase"`
			Name           string `json:"name"`
			TodayStartPri  string `json:"todayStartPri"`
			YestodEndPri   string `json:"yestodEndPri"`
			NowPri         string `json:"nowPri"`
			TodayMax       string `json:"todayMax"`
			TodayMin       string `json:"todayMin"`
			CompetitivePri string `json:"competitivePri"`
			ReservePri     string `json:"reservePri"`
			TraNumber      string `json:"traNumber"`
			TraAmount      string `json:"traAmount"`
			BuyOne         string `json:"buyOne"`
			BuyOnePri      string `json:"buyOnePri"`
			BuyTwo         string `json:"buyTwo"`
			BuyTwoPri      string `json:"buyTwoPri"`
			BuyThree       string `json:"buyThree"`
			BuyThreePri    string `json:"buyThreePri"`
			BuyFour        string `json:"buyFour"`
			BuyFourPri     string `json:"buyFourPri"`
			BuyFive        string `json:"buyFive"`
			BuyFivePri     string `json:"buyFivePri"`
			SellOne        string `json:"sellOne"`
			SellOnePri     string `json:"sellOnePri"`
			SellTwo        string `json:"sellTwo"`
			SellTwoPri     string `json:"sellTwoPri"`
			SellThree      string `json:"sellThree"`
			SellThreePri   string `json:"sellThreePri"`
			SellFour       string `json:"sellFour"`
			SellFourPri    string `json:"sellFourPri"`
			SellFive       string `json:"sellFive"`
			SellFivePri    string `json:"sellFivePri"`
			Date           string `json:"date"`
			Time           string `json:"time"`
		} `json:"data"`
		Dapandata struct {
			Dot       string `json:"dot"`
			Name      string `json:"name"`
			NowPic    string `json:"nowPic"`
			Rate      string `json:"rate"`
			TraAmount string `json:"traAmount"`
			TraNumber string `json:"traNumber"`
		} `json:"dapandata"`
		Gopicture struct {
			Minurl   string `json:"minurl"`
			Dayurl   string `json:"dayurl"`
			Weekurl  string `json:"weekurl"`
			Monthurl string `json:"monthurl"`
		} `json:"gopicture"`
	} `json:"result"`
}
