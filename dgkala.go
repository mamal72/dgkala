package dgkala

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	incredibleOffersAPIAddress = "https://service2.digikala.com/api/IncredibleOffer/GetIncredibleOffer"
)

// ImagePath is a struct containing the product images in various sizes
// type ImagePath struct {
// 	Original, Size70, Size110, Size180, Size220 string
// }

// IncredibleOffer is a struct containing
// DGKala incredible offer properties
type IncredibleOffer struct {
	ID         uint
	ProductID  uint
	Title      string
	ImagePaths struct {
		Original, Size70, Size110, Size180, Size220 string
	}
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

type incredibleOffersResponse struct {
	Data   []IncredibleOffer
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

// IncredibleOffers get a slice of DGKala IncredibleOffer items
func IncredibleOffers() ([]IncredibleOffer, error) {
	headers := map[string]string{"ApplicationVersion": "1.3.2"}
	response, err := sendRequest(incredibleOffersAPIAddress, headers)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	offersResponse := &incredibleOffersResponse{}
	err = json.Unmarshal(body, offersResponse)
	if err != nil {
		return nil, err
	}
	IncredibleOffers := offersResponse.Data
	return IncredibleOffers, nil
}
