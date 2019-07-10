package main

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	key          = "AKe62df84bJ3d8e4b1hea2R45j11klsb"
	logisticsURL = "http://202.159.30.42:22230/JandT_ecommerce/api/onlineOrder.action"
	dataString   = `{"detail":[{"username":"TEST","api_key":"tes123","orderid":"ORDER54321","shipper_name":"wain","shipper_contact":"0818510686","shipper_phone":"0818510686","shipper_addr":"Menara Imperium 9th Floor, Jl. HR. Rasuna Said, Jakarta Selatan","origin_code":"JKT","receiver_name":"TEST PARCEL","receiver_phone":"0123456789","receiver_addr":"PremPaChaLake & Park Phra Nakhon Si Ayutthaya","receiver_zip":"23657","destination_code":"JKT","receiver_area":"JKT002","qty":"1","weight":"2","goodsdesc":"","servicetype":"6"}]}`
)

func main() {
	dataSign := "MGYxNTZlZDE5NTYzOTQzMmE1OWE4ZTU5MTRhYzg2MWI="
	test, err := encrypt(dataString)
	if err != nil {
		panic(err)
	}
	//fmt.Println(dataSign)
	fmt.Println(test)
	fmt.Println(test == dataSign)
}

func encrypt(dataString string) (string, error) {
	hasher := md5.New()
	hasher.Write([]byte(dataString + key))
	toHash := hex.EncodeToString(hasher.Sum(nil))
	return base64.StdEncoding.EncodeToString([]byte(toHash)), nil
}

func testPost() {
	//'data_param'=>$data_json,
	//'data_sign'=>base64_encode(md5($data_request.$key))
	//base64 (md5 hash ( data_request + key ))
	httpclient := &http.Client{
		Timeout: 20 * time.Second,
	}
	dataSign := "MGYxNTZlZDE5NTYzOTQzMmE1OWE4ZTU5MTRhYzg2MWI="
	data := url.Values{}
	data.Set("data_param", dataString)
	data.Set("data_sign", dataSign)
	u, _ := url.ParseRequestURI(logisticsURL)
	urlStr := u.String()
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	response, err := httpclient.Do(req)
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf.Bytes()))
}
