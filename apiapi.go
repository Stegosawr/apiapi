package apiapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

// Item of amiami API
type Item struct {
	GCode                 string `json:"gcode"`
	SCode                 string `json:"scode"`
	GName                 string `json:"gname"`
	SName                 string `json:"sname"`
	GNameSub              string `json:"gname_sub"`
	SNameSimple           string `json:"sname_simple"`
	SNameSimpleJ          string `json:"sname_simple_j"`
	MainImageURL          string `json:"main_image_url"`
	MainImageAlt          string `json:"main_image_alt"`
	MainImageTitle        string `json:"main_image_title"`
	ImageComment          string `json:"image_comment"`
	Youtube               string `json:"youtube"`
	ListPrice             int    `json:"list_price"`
	CPriceTaxed           int    `json:"c_price_taxed"`
	Price                 int    `json:"price"`
	Point                 int    `json:"point"`
	Salestatus            string `json:"salestatus"`
	Releasedate           string `json:"releasedate"`
	PeriodFrom            string `json:"period_from"`
	PeriodTo              string `json:"period_to"`
	CartType              int    `json:"cart_type"`
	MaxCartinCount        int    `json:"max_cartin_count"`
	IncludeInstockOnlyFlg int    `json:"include_instock_only_flg"`
	Remarks               string `json:"remarks"`
	SizeInfo              string `json:"size_info"`
	WatchListAvailable    int    `json:"watch_list_available"`
	Jancode               string `json:"jancode"`
	MakerName             string `json:"maker_name"`
	Modeler               string `json:"modeler"`
	Modelergroup          string `json:"modelergroup"`
	Spec                  string `json:"spec"`
	Memo                  string `json:"memo"`
	Copyright             string `json:"copyright"`
	Saleitem              int    `json:"saleitem"`
	ConditionFlg          int    `json:"condition_flg"`
	Preorderitem          int    `json:"preorderitem"`
	Backorderitem         int    `json:"backorderitem"`
	StoreBonus            int    `json:"store_bonus"`
	AmiamiLimited         int    `json:"amiami_limited"`
	InstockFlg            int    `json:"instock_flg"`
	OrderClosedFlg        int    `json:"order_closed_flg"`
	PreownAttention       int    `json:"preown_attention"`
	Producttypeattention  int    `json:"producttypeattention"`
	Agelimit              int    `json:"agelimit"`
	CustomsWarningFlg     int    `json:"customs_warning_flg"`
	Preorderattention     string `json:"preorderattention"`
	PreorderBonusFlg      int    `json:"preorder_bonus_flg"`
	Domesticitem          int    `json:"domesticitem"`
	Metadescription       string `json:"metadescription"`
	Metawords             string `json:"metawords"`
	ReleasechangeText     string `json:"releasechange_text"`
	Cate1                 []int  `json:"cate1"`
	Cate2                 []int  `json:"cate2"`
	Cate3                 []int  `json:"cate3"`
	Cate4                 []int  `json:"cate4"`
	Cate5                 []int  `json:"cate5"`
	Cate6                 []int  `json:"cate6"`
	Cate7                 []int  `json:"cate7"`
	Salestalk             string `json:"salestalk"`
	BuyFlg                int    `json:"buy_flg"`
	BuyPrice              int    `json:"buy_price"`
	BuyRemarks            string `json:"buy_remarks"`
	EndFlg                int    `json:"end_flg"`
	DispFlg               int    `json:"disp_flg"`
	OnsaleFlg             int    `json:"onsale_flg"`
	HandlingStore         string `json:"handling_store"`
	SalestatusDetail      string `json:"salestatus_detail"`
	Stock                 int    `json:"stock"`
	Newitem               int    `json:"newitem"`
	Saletopitem           int    `json:"saletopitem"`
	ResaleFlg             int    `json:"resale_flg"`
	PreownedSaleFlg       int    `json:"preowned_sale_flg"`
	BigTitleFlg           int    `json:"big_title_flg"`
	SoldoutFlg            int    `json:"soldout_flg"`
	IncTxt1               int    `json:"inc_txt1"`
	IncTxt2               int    `json:"inc_txt2"`
	IncTxt3               int    `json:"inc_txt3"`
	IncTxt4               int    `json:"inc_txt4"`
	IncTxt5               int    `json:"inc_txt5"`
	IncTxt6               int    `json:"inc_txt6"`
	IncTxt7               int    `json:"inc_txt7"`
	IncTxt8               int    `json:"inc_txt8"`
	IncTxt9               int    `json:"inc_txt9"`
	IncTxt10              int    `json:"inc_txt10"`
	ImageOn               int    `json:"image_on"`
	ImageCategory         string `json:"image_category"`
	ImageName             string `json:"image_name"`
	Metaalt               string `json:"metaalt"`
	ImageReviewnumber     int    `json:"image_reviewnumber"`
	ImageReviewcategory   string `json:"image_reviewcategory"`
	Price1                int    `json:"price1"`
	Price2                int    `json:"price2"`
	Price3                int    `json:"price3"`
	Price4                int    `json:"price4"`
	Price5                int    `json:"price5"`
	Discountrate1         int    `json:"discountrate1"`
	Discountrate2         int    `json:"discountrate2"`
	Discountrate3         int    `json:"discountrate3"`
	Discountrate4         int    `json:"discountrate4"`
	Discountrate5         int    `json:"discountrate5"`
	Sizew                 string `json:"sizew"`
	Colorw                string `json:"colorw"`
	ThumbURL              string `json:"thumb_url"`
	ThumbAlt              string `json:"thumb_alt"`
	ThumbTitle            string `json:"thumb_title"`
	ThumbAgelimit         int    `json:"thumb_agelimit"`
}

// ReviewImage struct of amiami API
type ReviewImage struct {
	ImageURL string `json:"image_url"`
	ThumbURL string `json:"thumb_url"`
	Alt      string `json:"alt"`
	Title    string `json:"title"`
}

// Maker struct of amiami API
type Maker struct {
	ID   int
	Name string
}

// Title struct of amiami API
type Title struct {
	ID   int
	Name string
}

// Name struct of amiami API
type Name struct {
	ID   int
	Name string
}

// Embed struct of amiami API actually called _embedded
type Embed struct {
	ReviewImages   []ReviewImage `json:"review_images"`
	BonusImages    []ReviewImage `json:"bonus_images"`
	RelatedItems   []interface{} `json:"related_items"`
	OtherItems     []interface{} `json:"other_items"`
	Makers         []Maker       `json:"makers"`
	SeriesTitles   []Title       `json:"series_titles"`
	OriginalTtles  []Title       `json:"original_titles"`
	CharacterNames []Name        `json:"character_names"`
}

// ProductDetails struct of amiami API
type ProductDetails struct {
	RMessage string
	RSuccess bool
	Item     Item
	Embedded Embed `json:"_embedded"`
}

type ItemList struct {
	RSuccess     bool        `json:"RSuccess"`
	RValue       interface{} `json:"RValue"`
	RMessage     string      `json:"RMessage"`
	SearchResult struct {
		TotalResults int `json:"total_results"`
	} `json:"search_result"`
	Items []struct {
		Gcode                  string      `json:"gcode"`
		Gname                  string      `json:"gname"`
		ThumbURL               string      `json:"thumb_url"`
		ThumbAlt               string      `json:"thumb_alt"`
		ThumbTitle             string      `json:"thumb_title"`
		MinPrice               int         `json:"min_price"`
		MaxPrice               int         `json:"max_price"`
		CPriceTaxed            int         `json:"c_price_taxed"`
		MakerName              string      `json:"maker_name"`
		Saleitem               int         `json:"saleitem"`
		ConditionFlg           int         `json:"condition_flg"`
		ListPreorderAvailable  int         `json:"list_preorder_available"`
		ListBackorderAvailable int         `json:"list_backorder_available"`
		ListStoreBonus         int         `json:"list_store_bonus"`
		ListAmiamiLimited      int         `json:"list_amiami_limited"`
		InstockFlg             int         `json:"instock_flg"`
		OrderClosedFlg         int         `json:"order_closed_flg"`
		ElementID              interface{} `json:"element_id"`
		Salestatus             string      `json:"salestatus"`
		SalestatusDetail       string      `json:"salestatus_detail"`
		Releasedate            string      `json:"releasedate"`
		Jancode                string      `json:"jancode"`
		Preorderitem           int         `json:"preorderitem"`
		Saletopitem            int         `json:"saletopitem"`
		ResaleFlg              int         `json:"resale_flg"`
		PreownedSaleFlg        interface{} `json:"preowned_sale_flg"`
		ForWomenFlg            int         `json:"for_women_flg"`
		GenreMoe               int         `json:"genre_moe"`
		Cate6                  interface{} `json:"cate6"`
		Cate7                  interface{} `json:"cate7"`
		BuyFlg                 int         `json:"buy_flg"`
		BuyPrice               int         `json:"buy_price"`
		BuyRemarks             interface{} `json:"buy_remarks"`
		StockFlg               int         `json:"stock_flg"`
		ImageOn                int         `json:"image_on"`
		ImageCategory          string      `json:"image_category"`
		ImageName              string      `json:"image_name"`
		Metaalt                interface{} `json:"metaalt"`
	} `json:"items"`
}

// CodeType of item code
type CodeType string

const (
	// CodeTypeG indicates the type of item code is G.
	CodeTypeG CodeType = "G"
	// CodeTypeS indicates the type of item code is S.
	CodeTypeS CodeType = "S"
)

const apiItemURL = "https://api.amiami.com/api/v1.0/item?%scode=%s&lang=eng"
const apiItemListURL = "https://api.amiami.com/api/v1.0/items?pagemax=20&lang=eng&s_keywords=%s"

// GetItemByCode from amiami.com
func GetItemByCode(kind CodeType, code string) (ProductDetails, error) {

	data, err := get(fmt.Sprintf(apiItemURL, kind, code))
	if err != nil {
		return ProductDetails{}, err
	}

	details := ProductDetails{}
	err = json.Unmarshal([]byte(data), &details)
	if err != nil {
		return ProductDetails{}, err
	}
	if !details.RSuccess {
		return ProductDetails{}, fmt.Errorf("RSuccess: false for %s", fmt.Sprintf(apiItemURL, kind, code))
	}

	return details, nil
}

func GetItemsByKeywords(keywords string) ([]ProductDetails, error) {

	keywords = strings.ReplaceAll(keywords, "%20", "+")

	data, err := get(fmt.Sprintf(apiItemListURL, keywords))
	if err != nil {
		return nil, err
	}

	iList := ItemList{}
	err = json.Unmarshal(data, &iList)
	if err != nil {
		return nil, err
	}
	if !iList.RSuccess {
		return nil, fmt.Errorf("RSuccess: false for %s", fmt.Sprintf(apiItemListURL, keywords))
	}

	out := []ProductDetails{}
	for _, item := range iList.Items {
		time.Sleep(200 * time.Millisecond)
		details, err := GetItemByCode(CodeTypeG, item.Gcode)
		if err != nil {
			return nil, err
		}

		out = append(out, details)
	}

	return out, nil
}

func get(URL string) ([]byte, error) {

	command := []string{"--header", "Referer: https://www.amiami.com/", "--header", "X-User-Key: amiami_dev", URL}
	cmd := exec.Command("curl_ff117", command...)
	var inb, outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	cmd.Stdin = &inb
	if err := cmd.Run(); err != nil {
		log.Println(errb.String())
	}

	return outb.Bytes(), nil
}
