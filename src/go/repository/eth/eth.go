package eth

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

const etherscanHost = "https://api.etherscan.io/api"

//FIXME:設定ファイルに
const etherscanToken = "myToken"

func GetTransactionList(id, page string) (*transaction, error) {
	values := url.Values{}
	values.Add("module", "account")
	values.Add("action", "txlist")
	values.Add("address", id)
	values.Add("startblock", "0")
	//TODO:latestに変える
	values.Add("startblock", "99999999")
	//TODO:可変に
	values.Add("page", page)
	values.Add("offset", "30")
	values.Add("sort", "asc")
	//TODO:自分のに変える
	values.Add("apikey", etherscanToken)
	url := etherscanHost + "?" + values.Encode()
	fmt.Printf("%v", url)
	resp, e := http.Get(url)

	if e != nil {
		fmt.Println(e)
		return nil, e
	}

	defer resp.Body.Close()

	var r transaction
	e = decode(resp, &r)
	if e != nil {
		fmt.Println(e)
		return nil, e
	}
	fmt.Printf("結果 %#v", r)

	return &r, nil
}

func GetBalance(id string) (float64, error) {
	values := url.Values{}
	values.Add("module", "account")
	values.Add("action", "balance")
	values.Add("address", id)
	values.Add("tag", "latest")
	//TODO:自分のに変える
	values.Add("apikey", etherscanToken)
	url := etherscanHost + "?" + values.Encode()
	fmt.Printf("%v", url)
	resp, e := http.Get(url)

	if e != nil {
		fmt.Println(e)
		return 0, e
	}

	defer resp.Body.Close()

	var r balance
	e = decode(resp, &r)
	if e != nil {
		fmt.Println(e)
		return 0, e
	}
	fmt.Printf("結果 %#v", r)
	f, e := strconv.ParseFloat(r.Result, 64)
	if e != nil {
		fmt.Println(e)
		return 0, e
	}
	eth := weiToETH(f)
	fmt.Printf("eth =  %v", eth)
	return eth, nil
}

func decode(response *http.Response, result interface{}) error {
	err := json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return err
	}

	return nil
}

func weiToETH(wei float64) (eth float64) {
	return wei / math.Pow(10, 18.0)
}

//APIResult is
type APIResult struct {
	Status  string
	Message string
}

type transaction struct {
	APIResult
	Result []transactionResult
}

type transactionResult struct {
	TimeStamp string
	Hash      string
	From      string
	To        string
	Value     string
	Gas       string
	GasPrice  string
}

type balance struct {
	APIResult
	Result string
}
