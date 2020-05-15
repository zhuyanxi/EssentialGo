package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// IPLocation -
type IPLocation struct {
	IP           string `json:"ip"`
	CountryName  string `json:"countryName"`
	CountryCode  string `json:"countryCode"`
	RegionName   string `json:"regionName"`
	RegionCode   string `json:"regionCode"`
	City         string `json:"city"`
	ZipCode      string `json:"zipCode"`
	TimeZone     string `json:"timeZone"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	ISP          string `json:"isp"`
	Organization string `json:"organization"`
}

// IPLocations :
type IPLocations struct {
	Locations []IPLocation `json:"ipLocations"`
}

type Item struct {
	ItemID         string `json:"itemid"`
	Quantity       int    `json:"quantity"`
	QuantityBefore int    `json:"quantitybefore"`
	QuantityAfter  int    `json:"quantityafter"`
}

func testAnonymousStructs() {
	// ipJSON := []byte(`
	// 		{
	// 			"ip":"112.140.2.90"
	// 		}`)
	// var location IPLocation
	// err := json.Unmarshal(ipJSON, &location)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("%+v\n", location)

	// ipsJSON := []byte(`{"ipLocations":[{"ip":"112.140.2.90","countryName":"Japan","countryCode":"JP","regionName":"Shizuoka","regionCode":"22","city":"Gotenba Shi","zipCode":"412-0048","timeZone":"Asia/Tokyo","latitude":"35.2989","longitude":"138.879","isp":"TOKAI","organization":"TOKAI"},{"ip":"79.40.58.200","countryName":"Italy","countryCode":"IT","regionName":"Abruzzo","regionCode":"65","city":"Pescara","zipCode":"65129","timeZone":"Europe/Rome","latitude":"42.4584","longitude":"14.2028","isp":"Telecom Italia","organization":"Telecom Italia"}]}`)
	// var locations IPLocations
	// err = json.Unmarshal(ipsJSON, &locations)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("%+v\n", locations)

	// jsonData := []byte(`{
	// 	"Fruits": [
	// 		{
	// 			"ip":"112.140.2.90",
	// 			"Name": "Apple"
	// 		},
	// 		{
	// 			"ip":"79.40.58.200",
	// 			"Name": "Pear",
	// 			"PriceTag": "$1.5"
	// 		}
	// 	]
	// 	}`)

	// var basket fruitBasket
	// err = json.Unmarshal(jsonData, &basket)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("%+v\n", basket)

	// itemJSON := []byte(`[{"itemid":"aa7a9c26-3b7f-44df-9c60-b4937395683a","quantity":-1,"quantitybefore":1,"quantityafter":0}, {"itemid":"fcf0736b-0ec0-4ede-b073-9e3de2ff6376","quantity":1,"quantitybefore":0,"quantityafter":1}]`)
	// fmt.Println(string(itemJSON))
	// var items []Item
	// err := json.Unmarshal(itemJSON, &items)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("%+v\n", items)

	// itemJSON2 := []byte(`[{"itemid":aa7a9c26-3b7f-44df-9c60-b4937395683a,"quantity":-1,"quantitybefore":1,"quantityafter":0}, {"itemid":"fcf0736b-0ec0-4ede-b073-9e3de2ff6376","quantity":1,"quantitybefore":0,"quantityafter":1}]`)
	// s := NormalizeJSON(string(itemJSON2))
	// fmt.Println(s)
	// var items2 []Item
	// err2 := json.Unmarshal([]byte(s), &items2)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }
	// fmt.Printf("%+v\n", items2)

	// pattern := "^-?[0-9]\\d*?(}|}])?$" //反斜杠要转义
	// str := "-1}"
	// result, _ := regexp.MatchString(pattern, str)
	// fmt.Println(result)

	i := 20200217
	fmt.Println(strconv.FormatInt(int64(i), 10))
}

type fruit struct {
	Name     string `json:"Name"`
	PriceTag string `json:"PriceTag"`
	IP       string `json:"ip"`
}

type fruitBasket struct {
	Fruit   []fruit `json:"Fruits"`
	ID      int64   `json:"ref"` // 声明对应的json key
	Created time.Time
}

// NormalizeJSON :
func NormalizeJSON(s string) string {
	var result []string
	arr := strings.Split(s, ",")
	for _, val := range arr {
		valArr := strings.Split(val, ":")
		//appendVal := valArr[0] + ":"
		match, _ := regexp.MatchString("^-?[0-9]\\d*?(}|}])?$", valArr[1])
		if !match {
			if strings.HasPrefix(valArr[1], "\"") {
				val = valArr[0] + ":" + valArr[1]
			} else {
				val = valArr[0] + ":\"" + valArr[1]
			}
			if !strings.HasSuffix(valArr[1], "\"") {
				val = val + "\""
			}
		}
		result = append(result, val)
	}
	fmt.Println(result)
	return strings.Join(result, ",")
}
