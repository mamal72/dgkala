package dgkala

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/buger/jsonparser"
)

const (
	incredibleOffersAPIAddress = "https://service2.digikala.com/api/IncredibleOffer/GetIncredibleOffer"
	searchAPIAddress           = "https://search.digikala.com/api/search?keyword=%s"
	staticFilesPath            = "https://file.digikala.com/digikala/%s"
	productByIDAPIAddress      = "https://service2.digikala.com/api/ProductCache/GetProductById/%d"
)

type requestHeader map[string]string

// ProductExistsStatus is a iota type for product existing status for buying
type ProductExistsStatus int

const (
	_ ProductExistsStatus = iota
	_
	// Available means the product is available to buy
	Available
	// OutOfStock means the product is not available  now and is out of stock
	OutOfStock
	// Discontinued means the product is discontinued
	Discontinued
)

// ImagePaths contains a product images in various sizes
type ImagePaths struct {
	Original, Size70, Size110, Size180, Size220 string
}

// IncredibleOffer is a struct containing
// DGKala incredible offer properties
type IncredibleOffer struct {
	ID                 uint
	ProductID          uint
	Title              string
	ImagePaths         ImagePaths
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

// ProductColor is a struct with properties of products colors
type ProductColor struct {
	Title string
	Hex   string
	Code  string
}

// ProductSearchResult is a struct containing a product details for a search result
type ProductSearchResult struct {
	ID                  int64
	EnglishTitle        string
	PersianTitle        string
	Image               string
	ExistsStatus        ProductExistsStatus
	IsActive            bool
	URL                 string
	Rate                int64
	MinimumPrice        int64
	MaximumPrice        int64
	Likes               int64
	LastPeriodLikes     int64
	Views               int64
	LastPeriodViews     int64
	IsSpecialOffer      bool
	RegisteredDateTime  time.Time
	HasVideo            bool
	Colors              []ProductColor
	UserRatingCount     int64
	Favorites           int64
	LastPeriodFavorites int64
	LastPeriodSales     int64
	HasGift             bool
	HTMLDetails         string
}

// SearchResult returns a struct containing results of the search for a keyword
type SearchResult struct {
	ResponseTime int64
	Count        int64
	Results      []ProductSearchResult
}

// ProductByIDResult returns a struct containing results of the request for product details by ID
type ProductByIDResult struct {
	Data ProductByID
}

// ProductByID is a struct containing a product details when you get it by ID
type ProductByID struct {
	ID                uint `json:"ProductId"`
	EnglishTitle      string
	PersianTitle      string
	Description       string
	ImagePaths        ImagePaths
	IsIncredibleOffer bool
	Strengths         string
	Weaknesses        string
	MinPrice          uint
}

func sendRequest(address string, headers requestHeader) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, address, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		request.Header.Add(key, value)
	}

	var client http.Client
	return client.Do(request)
}

func getStaticResourceAddress(resourcePath string) string {
	return fmt.Sprintf(staticFilesPath, resourcePath)
}

func getSearchAPIAddress(keyword string) string {
	query := url.QueryEscape(keyword)
	return fmt.Sprintf(searchAPIAddress, query)
}

func getProductByIDAPIAddress(productID int) string {
	return fmt.Sprintf(productByIDAPIAddress, productID)
}

func getRequestHeaders() requestHeader {
	return map[string]string{"ApplicationVersion": "1.4.1"}
}

// IncredibleOffers get a slice of DGKala IncredibleOffer items
func IncredibleOffers() ([]IncredibleOffer, error) {
	headers := getRequestHeaders()
	response, err := sendRequest(incredibleOffersAPIAddress, headers)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var offersResponse incredibleOffersResponse
	err = json.Unmarshal(body, &offersResponse)
	if err != nil {
		return nil, err
	}
	incredibleOffers := offersResponse.Data
	return incredibleOffers, nil
}

// Search for a product in DGKala and return a slice of DGKala SearchResult items
func Search(keyword string) (SearchResult, error) {
	searchAddress := getSearchAPIAddress(keyword)
	httpResponse, err := sendRequest(searchAddress, requestHeader{})
	if err != nil {
		return SearchResult{}, err
	}

	responseBody, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return SearchResult{}, err
	}

	responseTime, err := jsonparser.GetInt(responseBody, "took")
	if err != nil {
		return SearchResult{}, err
	}

	count, err := jsonparser.GetInt(responseBody, "hits", "total")
	if err != nil {
		return SearchResult{}, err
	}

	productSearchResults := []ProductSearchResult{}
	realResultsJSONPath := []string{"hits", "hits"}
	parentJSONResultKey := "_source"
	jsonparser.ArrayEach(responseBody, func(value []byte, _ jsonparser.ValueType, _ int, _ error) {
		ID, _ := jsonparser.GetInt(value, parentJSONResultKey, "Id")
		englishTitle, _ := jsonparser.GetString(value, parentJSONResultKey, "EnTitle")
		persianTitle, _ := jsonparser.GetString(value, parentJSONResultKey, "FaTitle")
		imagePath, _ := jsonparser.GetString(value, parentJSONResultKey, "ImagePath")
		image := getStaticResourceAddress(imagePath)
		existsStatusInt, _ := jsonparser.GetInt(value, parentJSONResultKey, "ExistStatus")
		existsStatus := ProductExistsStatus(existsStatusInt)
		isActive, _ := jsonparser.GetBoolean(value, parentJSONResultKey, "IsActive")
		URL, _ := jsonparser.GetString(value, parentJSONResultKey, "UrlCode")
		rate, _ := jsonparser.GetInt(value, parentJSONResultKey, "Rate")
		minimumPrice, _ := jsonparser.GetInt(value, parentJSONResultKey, "MinPrice")
		maximumPrice, _ := jsonparser.GetInt(value, parentJSONResultKey, "MaxPrice")
		likes, _ := jsonparser.GetInt(value, parentJSONResultKey, "LikeCounter")
		lastPeriodLikes, _ := jsonparser.GetInt(value, parentJSONResultKey, "LastPeriodLikeCounter")
		views, _ := jsonparser.GetInt(value, parentJSONResultKey, "ViewCounter")
		lastPeriodViews, _ := jsonparser.GetInt(value, parentJSONResultKey, "LastPeriodViewCounter")
		isSpecialOffer, _ := jsonparser.GetBoolean(value, parentJSONResultKey, "IsSpecialOffer")
		regDateTimeString, _ := jsonparser.GetString(value, parentJSONResultKey, "RegDateTime")
		registeredDateTime, _ := time.Parse("2006-01-02T15:04:05", regDateTimeString)
		hasVideo, _ := jsonparser.GetBoolean(value, parentJSONResultKey, "HasVideo")
		colors := []ProductColor{}
		jsonparser.ArrayEach(value, func(colorsValue []byte, _ jsonparser.ValueType, _ int, _ error) {
			colorTitle, _ := jsonparser.GetString(colorsValue, "ColorTitle")
			colorHex, _ := jsonparser.GetString(colorsValue, "ColorHex")
			colorCode, _ := jsonparser.GetString(colorsValue, "ColorCode")
			currentColor := ProductColor{
				colorTitle,
				colorHex,
				colorCode,
			}
			colors = append(colors, currentColor)
		}, parentJSONResultKey, "ProductColorList")
		userRatingCount, _ := jsonparser.GetInt(value, parentJSONResultKey, "UserRating")
		favorites, _ := jsonparser.GetInt(value, parentJSONResultKey, "FavoriteCounter")
		lastPeriodFavorites, _ := jsonparser.GetInt(value, parentJSONResultKey, "LastPeriodFavoriteCounter")
		lastPeriodSales, _ := jsonparser.GetInt(value, parentJSONResultKey, "LastPeriodSaleCounter")
		hasGift, _ := jsonparser.GetBoolean(value, parentJSONResultKey, "HasGift")
		hTMLDetails, _ := jsonparser.GetString(value, parentJSONResultKey, "DetailSource")

		currentProductSearchResult := ProductSearchResult{
			ID,
			englishTitle,
			persianTitle,
			image,
			existsStatus,
			isActive,
			URL,
			rate,
			minimumPrice,
			maximumPrice,
			likes,
			lastPeriodLikes,
			views,
			lastPeriodViews,
			isSpecialOffer,
			registeredDateTime,
			hasVideo,
			colors,
			userRatingCount,
			favorites,
			lastPeriodFavorites,
			lastPeriodSales,
			hasGift,
			hTMLDetails,
		}
		productSearchResults = append(productSearchResults, currentProductSearchResult)
	}, realResultsJSONPath...)

	result := SearchResult{
		responseTime,
		count,
		productSearchResults,
	}

	return result, nil
}

// GetProductByID returns a product by getting it's ID
func GetProductByID(productID int) (ProductByID, error) {
	headers := getRequestHeaders()
	apiAddress := getProductByIDAPIAddress(productID)

	httpResponse, err := sendRequest(apiAddress, headers)

	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return ProductByID{}, err
	}
	var productByIDResult ProductByIDResult
	err = json.Unmarshal(body, &productByIDResult)
	if err != nil {
		return ProductByID{}, err
	}
	product := productByIDResult.Data
	return product, nil
}
