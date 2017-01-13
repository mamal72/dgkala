package dgkala

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const apiAddress = "https://service2.digikala.com/api/IncredibleOffer/GetIncredibleOffer"

// ImagePath is a struct containing the product images in various sizes
type ImagePath struct {
	Original, Size70, Size110, Size180, Size220 string
}

// SpecialOffer is a struct containing
// DGKala special offer properties
type SpecialOffer struct {
	ID                 uint
	ProductID          uint
	Title              string
	ImagePaths         ImagePath
	BannerPath         string
	BannerPathMobile   string
	BannerPathTablet   string
	Row                uint
	ProductTitleFa     string
	ProductTitleEn     string
	Discount           uint
	Price              uint
	OnlyForApplication bool
	OnlyForMembers     bool
}

type specialOffersResponse struct {
	Data   []SpecialOffer
	Status string
}

func sendRequest(address string, headers map[string]string) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, address, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		request.Header.Add(key, value)
	}

	client := http.Client{}
	response, err := client.Do(request)
	return response, err
}

// SpecialOffers returns a slice of DGKala SpecialOffer
func SpecialOffers() ([]SpecialOffer, error) {
	headers := map[string]string{"ApplicationVersion": "1.3.2"}
	response, err := sendRequest(apiAddress, headers)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var specialOffersResponse specialOffersResponse
	err = json.Unmarshal(body, &specialOffersResponse)
	if err != nil {
		return nil, err
	}
	specialOffers := specialOffersResponse.Data
	return specialOffers, nil
}
