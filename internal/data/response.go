package data

import (
	"time"
)

type AddProductPayload struct {
	ConfigFk        string  `json:"configFk,omitempty"`
	Discount        float64 `json:"discount"`
	ExtraFee        float64 `json:"extraFee"`
	MainCartFk      string  `json:"mainCartFk"`
	MainCartItemsPk int     `json:"mainCartItemsPk"`
	ProductFk       int     `json:"productFk"`
	Quantity        int     `json:"quantity"`
	SetupForAs      bool    `json:"setupForAs"`
}

type ProductAddedResponse struct {
	Id                          int           `json:"id"`
	Quantity                    int           `json:"quantity"`
	ProductFk                   int           `json:"productFk"`
	PackageFk                   interface{}   `json:"packageFk"`
	Name                        string        `json:"name"`
	ImageUrl                    string        `json:"imageUrl"`
	SetupForAs                  bool          `json:"setupForAs"`
	ConfigFk                    interface{}   `json:"configFk"`
	PriceListFk                 int           `json:"priceListFk"`
	IsPaCItem                   bool          `json:"isPaCItem"`
	IsSignup                    bool          `json:"isSignup"`
	Sku                         string        `json:"sku"`
	PriceType                   string        `json:"priceType"`
	CountryCode                 string        `json:"countryCode"`
	CurrencyPK                  int           `json:"currencyPK"`
	CurrencyCode                string        `json:"currencyCode"`
	CurrencySymbol              string        `json:"currencySymbol"`
	ExtraFee                    interface{}   `json:"extraFee"`
	Discount                    float64       `json:"discount"`
	FormattedDiscount           string        `json:"formattedDiscount"`
	BasePrice                   float64       `json:"basePrice"`
	FormattedBasePrice          string        `json:"formattedBasePrice"`
	UnitPrice                   float64       `json:"unitPrice"`
	FormattedUnitPrice          string        `json:"formattedUnitPrice"`
	RetailTaxablePrice          float64       `json:"retailTaxablePrice"`
	FormattedRetailTaxablePrice string        `json:"formattedRetailTaxablePrice"`
	Cv                          float64       `json:"cv"`
	Sp                          float64       `json:"sp"`
	MaxLimit                    int           `json:"maxLimit"`
	Points                      float64       `json:"points"`
	IsShippable                 bool          `json:"isShippable"`
	IsStarterKit                bool          `json:"isStarterKit"`
	BrandName                   string        `json:"brandName"`
	LineSubTotal                float64       `json:"lineSubTotal"`
	FormattedLineSubTotal       string        `json:"formattedLineSubTotal"`
	LineTotal                   float64       `json:"lineTotal"`
	FormattedLineTotal          string        `json:"formattedLineTotal"`
	OfferAffiliateProgram       bool          `json:"offerAffiliateProgram"`
	IsVolumeBasedRSB            bool          `json:"isVolumeBasedRSB"`
	OfferPreferredCust          bool          `json:"offerPreferredCust"`
	OfferLoyaltyProgram         bool          `json:"offerLoyaltyProgram"`
	ShowSDCheckbox              bool          `json:"showSDCheckbox"`
	JoinMaxLifetimeLimitCatCode string        `json:"joinMaxLifetimeLimitCatCode"`
	JoinMaxLifetimeLimit        int           `json:"joinMaxLifetimeLimit"`
	ChildItems                  []interface{} `json:"childItems"`
}

type ShippedProductResponse struct {
	PackagePk                 int         `json:"packagePk"`
	ProductPk                 int         `json:"productPk"`
	PackageName               string      `json:"packageName"`
	ProductName               string      `json:"productName"`
	IsPackage                 bool        `json:"isPackage"`
	Quantity                  int         `json:"quantity"`
	Cv                        int         `json:"cv"`
	Sp                        int         `json:"sp"`
	Price                     int         `json:"price"`
	FormattedPrice            string      `json:"formattedPrice"`
	CurrencyCode              string      `json:"currencyCode"`
	ShipmentNumber            string      `json:"shipmentNumber"`
	ShipmentStatus            string      `json:"shipmentStatus"`
	ShippedDate               string      `json:"shippedDate"`
	TrackingNumber            string      `json:"trackingNumber"`
	TrackingLink              string      `json:"trackingLink"`
	VideoOrderPackagingInfoPK interface{} `json:"videoOrderPackagingInfoPK"`
}

type PackageItem struct {
	ProductPK   int    `json:"productPK"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImgUrl      string `json:"imgUrl"`
	Qty         int    `json:"qty"`
}

type Pricing struct {
	PriceType      string  `json:"priceType"`
	CurrencySymbol string  `json:"currencySymbol"`
	Price          float64 `json:"price"`
	NoVatPrice     float64 `json:"noVatPrice"`
	FormattedPrice string  `json:"formattedPrice"`
	PriceWarning   string  `json:"priceWarning"`
}

type ImageUrl struct {
	ImageUrl  string `json:"imageUrl"`
	ImageName string `json:"imageName"`
}

type ProductCmsData struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	DataTag string `json:"dataTag"`
}

type ShopifyProduct struct {
	Id             int64     `json:"id"`
	Title          string    `json:"title"`
	BodyHtml       string    `json:"body_html"`
	Vendor         string    `json:"vendor"`
	ProductType    string    `json:"product_type"`
	Handle         string    `json:"handle"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	PublishedAt    time.Time `json:"published_at"`
	PublishedScope string    `json:"published_scope"`
	Tags           string    `json:"tags"`
	Status         string    `json:"status"`
	Options        []struct {
		Id        int64    `json:"id"`
		ProductId int64    `json:"product_id"`
		Name      string   `json:"name"`
		Position  int      `json:"position"`
		Values    []string `json:"values"`
	} `json:"options"`
	Variants []struct {
		Id                   int64     `json:"id"`
		ProductId            int64     `json:"product_id"`
		Title                string    `json:"title"`
		Sku                  string    `json:"sku"`
		Position             int       `json:"position"`
		Grams                int       `json:"grams"`
		InventoryPolicy      string    `json:"inventory_policy"`
		Price                string    `json:"price"`
		FulfillmentService   string    `json:"fulfillment_service"`
		InventoryManagement  string    `json:"inventory_management"`
		InventoryItemId      int64     `json:"inventory_item_id"`
		Option1              string    `json:"option1"`
		CreatedAt            time.Time `json:"created_at"`
		UpdatedAt            time.Time `json:"updated_at"`
		Taxable              bool      `json:"taxable"`
		Barcode              string    `json:"barcode"`
		InventoryQuantity    int       `json:"inventory_quantity"`
		Weight               string    `json:"weight"`
		WeightUnit           string    `json:"weight_unit"`
		OldInventoryQuantity int       `json:"old_inventory_quantity"`
		RequiresShipping     bool      `json:"requires_shipping"`
		AdminGraphqlApiId    string    `json:"admin_graphql_api_id"`
	} `json:"variants"`
	Image struct {
		Id                int64     `json:"id"`
		ProductId         int64     `json:"product_id"`
		Position          int       `json:"position"`
		CreatedAt         time.Time `json:"created_at"`
		UpdatedAt         time.Time `json:"updated_at"`
		Width             int       `json:"width"`
		Height            int       `json:"height"`
		Src               string    `json:"src"`
		Alt               string    `json:"alt"`
		AdminGraphqlApiId string    `json:"admin_graphql_api_id"`
	} `json:"image"`
	Images []struct {
		Id                int64     `json:"id"`
		ProductId         int64     `json:"product_id"`
		Position          int       `json:"position"`
		CreatedAt         time.Time `json:"created_at"`
		UpdatedAt         time.Time `json:"updated_at"`
		Width             int       `json:"width"`
		Height            int       `json:"height"`
		Src               string    `json:"src"`
		Alt               string    `json:"alt"`
		AdminGraphqlApiId string    `json:"admin_graphql_api_id"`
	} `json:"images"`
	TemplateSuffix    string `json:"template_suffix"`
	AdminGraphqlApiId string `json:"admin_graphql_api_id"`
}

type CartItem struct {
	Id                          int           `json:"id"`
	Quantity                    int           `json:"quantity"`
	ProductFk                   int           `json:"productFk"`
	PackageFk                   interface{}   `json:"packageFk"`
	Name                        string        `json:"name"`
	ImageUrl                    string        `json:"imageUrl"`
	SetupForAs                  bool          `json:"setupForAs"`
	ConfigFk                    interface{}   `json:"configFk"`
	PriceListFk                 int           `json:"priceListFk"`
	IsPaCItem                   bool          `json:"isPaCItem"`
	IsSignup                    bool          `json:"isSignup"`
	Sku                         string        `json:"sku"`
	PriceType                   string        `json:"priceType"`
	CountryCode                 string        `json:"countryCode"`
	CurrencyPK                  int           `json:"currencyPK"`
	CurrencyCode                string        `json:"currencyCode"`
	CurrencySymbol              string        `json:"currencySymbol"`
	ExtraFee                    interface{}   `json:"extraFee"`
	Discount                    float64       `json:"discount"`
	FormattedDiscount           string        `json:"formattedDiscount"`
	BasePrice                   float64       `json:"basePrice"`
	FormattedBasePrice          string        `json:"formattedBasePrice"`
	UnitPrice                   float64       `json:"unitPrice"`
	FormattedUnitPrice          string        `json:"formattedUnitPrice"`
	RetailTaxablePrice          float64       `json:"retailTaxablePrice"`
	FormattedRetailTaxablePrice string        `json:"formattedRetailTaxablePrice"`
	Cv                          float64       `json:"cv"`
	Sp                          float64       `json:"sp"`
	MaxLimit                    int           `json:"maxLimit"`
	Points                      float64       `json:"points"`
	IsShippable                 bool          `json:"isShippable"`
	IsStarterKit                bool          `json:"isStarterKit"`
	BrandName                   string        `json:"brandName"`
	LineSubTotal                float64       `json:"lineSubTotal"`
	FormattedLineSubTotal       string        `json:"formattedLineSubTotal"`
	LineTotal                   float64       `json:"lineTotal"`
	FormattedLineTotal          string        `json:"formattedLineTotal"`
	OfferAffiliateProgram       bool          `json:"offerAffiliateProgram"`
	IsVolumeBasedRSB            bool          `json:"isVolumeBasedRSB"`
	OfferPreferredCust          bool          `json:"offerPreferredCust"`
	OfferLoyaltyProgram         bool          `json:"offerLoyaltyProgram"`
	ShowSDCheckbox              bool          `json:"showSDCheckbox"`
	JoinMaxLifetimeLimitCatCode string        `json:"joinMaxLifetimeLimitCatCode"`
	JoinMaxLifetimeLimit        int           `json:"joinMaxLifetimeLimit"`
	ChildItems                  []interface{} `json:"childItems"`
}

type ProductInformation struct {
	ProductPK                    int              `json:"productPK"`
	ProductCode                  string           `json:"productCode"`
	Sku                          interface{}      `json:"sku"`
	ProductCategory              string           `json:"productCategory"`
	BrandId                      int              `json:"brandId"`
	BrandName                    string           `json:"brandName"`
	ProductBrand                 interface{}      `json:"productBrand"`
	Name                         string           `json:"name"`
	ImageUrl                     string           `json:"imageUrl"`
	Weight                       float64          `json:"weight"`
	IsComingSoon                 bool             `json:"isComingSoon"`
	ComingSoonMessage            interface{}      `json:"comingSoonMessage"`
	IsPackage                    bool             `json:"isPackage"`
	PackageItems                 []PackageItem    `json:"packageItems"`
	IsConfigurable               bool             `json:"isConfigurable"`
	IsProductAvailableOnAutoship bool             `json:"isProductAvailableOnAutoship"`
	IsProductOnAutoship          bool             `json:"isProductOnAutoship"`
	AutoshipProductPk            int              `json:"autoshipProductPk"`
	MaxLimit                     int              `json:"maxLimit"`
	Points                       float64          `json:"points"`
	Bv                           float64          `json:"bv"`
	Sp                           float64          `json:"sp"`
	ProductMenuId                int              `json:"productMenuId"`
	ProductMenu                  string           `json:"productMenu"`
	Configurations               []interface{}    `json:"configurations"`
	Pricing                      []Pricing        `json:"pricing"`
	Description                  string           `json:"description"`
	ImageUrls                    []ImageUrl       `json:"imageUrls"`
	AdditionalInfo               []interface{}    `json:"additionalInfo"`
	Documents                    []interface{}    `json:"documents"`
	IsShippable                  bool             `json:"isShippable"`
	IsStarterKit                 bool             `json:"isStarterKit"`
	SeqNo                        int              `json:"seqNo"`
	IsFoodProduct                bool             `json:"isFoodProduct"`
	RankInfo                     interface{}      `json:"rankInfo"`
	IsRetailPackage              bool             `json:"isRetailPackage"`
	IsVolumeBasedRSB             bool             `json:"isVolumeBasedRSB"`
	MainType                     int              `json:"mainType"`
	ActiveSmartDelivery          bool             `json:"activeSmartDelivery"`
	IsRedemption                 bool             `json:"isRedemption"`
	PriceType                    string           `json:"priceType"`
	DoNotSplitPackBV             bool             `json:"doNotSplitPackBV"`
	SdOnlyPackage                bool             `json:"sdOnlyPackage"`
	ShowSDCheckbox               bool             `json:"showSDCheckbox"`
	OfferAffiliateProgram        bool             `json:"offerAffiliateProgram"`
	OfferPreferredCust           bool             `json:"offerPreferredCust"`
	OfferLoyaltyProgram          bool             `json:"offerLoyaltyProgram"`
	OfferSDOnShop                bool             `json:"offerSDOnShop"`
	IsRetailCart                 bool             `json:"isRetailCart"`
	ProductLineId                int              `json:"productLineId"`
	ProductLine                  string           `json:"productLine"`
	ProductFunction              string           `json:"productFunction"`
	MaxLifetimeLimitCatCode      interface{}      `json:"maxLifetimeLimitCatCode"`
	MaxLifetimeLimit             interface{}      `json:"maxLifetimeLimit"`
	JoinMaxLifetimeLimitCatCode  string           `json:"joinMaxLifetimeLimitCatCode"`
	JoinMaxLifetimeLimit         int              `json:"joinMaxLifetimeLimit"`
	ProductCmsData               []ProductCmsData `json:"productCmsData"`
}

type OrderResponse struct {
	OrderDate               string  `json:"orderDate"`
	MainOrdersPK            int     `json:"mainOrdersPK"`
	OrderType               string  `json:"orderType"`
	FinalOrderType          string  `json:"finalOrderType,omitempty"`
	SiteURL                 string  `json:"siteURL"`
	EncOrderNumber          string  `json:"encOrderNumber"`
	CurrencySymbol          string  `json:"currencySymbol"`
	CurrencyCode            string  `json:"currencyCode"`
	PaidStatus              bool    `json:"paidStatus"`
	HasTaxInvoice           bool    `json:"hasTaxInvoice"`
	HasCommercialInvoice    bool    `json:"hasCommercialInvoice"`
	HasCreditNote           bool    `json:"hasCreditNote"`
	IsShippingPending       bool    `json:"isShippingPending"`
	IsPB                    bool    `json:"isPB"`
	IsPA                    bool    `json:"isPA"`
	IsCC                    bool    `json:"isCC"`
	MainFK                  int     `json:"mainFK"`
	MainOrderTypeFK         int     `json:"mainOrderTypeFK"`
	VoucherURL              string  `json:"voucherURL,omitempty"`
	ShipCountry             string  `json:"shipCountry"`
	ShippingStatus          string  `json:"shippingStatus"`
	OrderShippingStatus     string  `json:"orderShippingStatus"`
	OrderTypeValue          string  `json:"orderTypeValue,omitempty"`
	PaidStatusValue         string  `json:"paidStatusValue"`
	Quantity                int     `json:"quantity"`
	Email                   string  `json:"email,omitempty"`
	Phone                   string  `json:"phone,omitempty"`
	ShipFirstName           string  `json:"shipFirstName,omitempty"`
	ShipLastName            string  `json:"shipLastName,omitempty"`
	MarkedPaidDate          string  `json:"markedPaidDate"`
	Total                   float64 `json:"total"`
	ConvTotal               float64 `json:"convTotal"`
	ConvTotalFormat         string  `json:"convTotalFormat"`
	SubTotal                float64 `json:"subTotal"`
	ConvSubtotal            float64 `json:"convSubtotal"`
	ShipTotal               float64 `json:"shipTotal"`
	ConvShipTotal           float64 `json:"convShipTotal"`
	Taxes                   float64 `json:"taxes"`
	TaxLabel                string  `json:"taxLabel"`
	ProductTax              float64 `json:"productTax"`
	ShippingTax             float64 `json:"shippingTax"`
	TotalProductTax         float64 `json:"totalProductTax"`
	AdditionalTaxLabel      string  `json:"additionalTaxLabel"`
	AdditionalTax           string  `json:"additionalTax,omitempty"`
	ConvTaxes               float64 `json:"convTaxes"`
	OrderProcessingFees     string  `json:"orderProcessingFees,omitempty"`
	ConvOrderProcessingFees string  `json:"convOrderProcessingFees,omitempty"`
	Discount                float64 `json:"discount"`
	ConvDiscount            float64 `json:"convDiscount"`
	RefundAmount            float64 `json:"refundAmount"`
	ConvRefund              float64 `json:"convRefund"`
	SalesCampaignFK         int     `json:"salesCampaignFK,omitempty"`
	Paidstatusfk            int     `json:"paidstatusfk"`
	DeliveryDate            string  `json:"deliveryDate,omitempty"`
	ShippingDetails         string  `json:"shippingDetails,omitempty"`
	OrderItems              string  `json:"orderItems,omitempty"`
	Payments                string  `json:"payments,omitempty"`
	IsPrepaidOrder          string  `json:"isPrepaidOrder,omitempty"`
	ShowInvoice             bool    `json:"showInvoice"`
	ShowOrderInvoice        bool    `json:"showOrderInvoice"`
	KrGuaranteeNo           string  `json:"krGuaranteeNo"`
	WeChatOrderNumber       string  `json:"weChatOrderNumber,omitempty"`
	MemberID                string  `json:"memberID,omitempty"`
}

type Cart struct {
	CartKey                   string      `json:"cartKey"`
	CartType                  string      `json:"cartType"`
	CountryCode               string      `json:"countryCode"`
	MainFK                    int         `json:"mainFK"`
	MainReferrerFK            int         `json:"mainReferrerFK"`
	Culture                   string      `json:"culture"`
	LanguageFK                int         `json:"languageFK"`
	MainOrderTypeFK           int         `json:"mainOrderTypeFK"`
	PriceListFK               int         `json:"priceListFK"`
	CampaignCode              string      `json:"campaignCode"`
	SalesCampaignFK           interface{} `json:"salesCampaignFK"`
	Ip                        string      `json:"ip"`
	DateEntered               string      `json:"dateEntered"`
	GaCode                    string      `json:"gaCode"`
	FacebookCode              string      `json:"facebookCode"`
	LuckyOrange               string      `json:"luckyOrange"`
	ReferrerSiteUrl           string      `json:"referrerSiteUrl"`
	ReferrerIsCorporate       bool        `json:"referrerIsCorporate"`
	CustomerReferralId        string      `json:"customerReferralId"`
	MainCreditCardsFK         interface{} `json:"mainCreditCardsFK"`
	MainOrdersFK              uint64      `json:"mainOrdersFK"`
	ShippingTypeFK            int         `json:"shippingTypeFK"`
	CartStatus                interface{} `json:"cartStatus"`
	FirstName                 string      `json:"firstName"`
	LastName                  string      `json:"lastName"`
	Phone                     string      `json:"phone"`
	Email                     string      `json:"email"`
	DateModified              string      `json:"dateModified"`
	SubTotal                  float64     `json:"subTotal"`
	FormattedSubTotal         string      `json:"formattedSubTotal"`
	Tax                       float64     `json:"tax"`
	FormattedTax              string      `json:"formattedTax"`
	Shipping                  float64     `json:"shipping"`
	FormattedShipping         string      `json:"formattedShipping"`
	Discount                  float64     `json:"discount"`
	FormattedDiscount         string      `json:"formattedDiscount"`
	Total                     float64     `json:"total"`
	FormattedTotal            string      `json:"formattedTotal"`
	PointsTotal               float64     `json:"pointsTotal"`
	ShipSignatureRequired     bool        `json:"shipSignatureRequired"`
	ShipSignatureFee          float64     `json:"shipSignatureFee"`
	CurrencyFK                int         `json:"currencyFK"`
	CurrencyCode              string      `json:"currencyCode"`
	MainDiscountCode          string      `json:"mainDiscountCode"`
	ActiveSmartDelivery       bool        `json:"activeSmartDelivery"`
	AllowImportCart           bool        `json:"allowImportCart"`
	OfferPreferredCust        bool        `json:"offerPreferredCust"`
	IsAffiliateOn             bool        `json:"isAffiliateOn"`
	IsVolumeBasedRSB          bool        `json:"isVolumeBasedRSB"`
	OfferLoyaltyProgram       bool        `json:"offerLoyaltyProgram"`
	OfferSmartDelivery        bool        `json:"offerSmartDelivery"`
	IsRetailSignup            bool        `json:"isRetailSignup"`
	HasRetailStarterKit       bool        `json:"hasRetailStarterKit"`
	AllowRetailSignup         bool        `json:"allowRetailSignup"`
	EventMemberID             interface{} `json:"eventMemberID"`
	ShowAbandonedOrderWarning bool        `json:"showAbandonedOrderWarning"`
	ShouldCreateAccount       bool        `json:"shouldCreateAccount"`
	IsCartCreatedFromBag      bool        `json:"isCartCreatedFromBag"`
	IsCartCreatedFromSignup   bool        `json:"isCartCreatedFromSignup"`
	CurrencySymbol            string      `json:"currencySymbol"`
	ShippingAddress           struct {
		MainCartAddressPK int         `json:"mainCartAddressPK"`
		FirstName         interface{} `json:"firstName"`
		LastName          interface{} `json:"lastName"`
		Company           interface{} `json:"company"`
		Address1          interface{} `json:"address1"`
		Address2          interface{} `json:"address2"`
		Address3          interface{} `json:"address3"`
		City              interface{} `json:"city"`
		State             interface{} `json:"state"`
		PostalCode        interface{} `json:"postalCode"`
		Country           interface{} `json:"country"`
		IsPOBox           interface{} `json:"isPOBox"`
		IsResidential     interface{} `json:"isResidential"`
	} `json:"shippingAddress"`
	MailingAddress struct {
		MainCartAddressPK int         `json:"mainCartAddressPK"`
		FirstName         interface{} `json:"firstName"`
		LastName          interface{} `json:"lastName"`
		Company           interface{} `json:"company"`
		Address1          interface{} `json:"address1"`
		Address2          interface{} `json:"address2"`
		Address3          interface{} `json:"address3"`
		City              interface{} `json:"city"`
		State             interface{} `json:"state"`
		PostalCode        interface{} `json:"postalCode"`
		Country           interface{} `json:"country"`
		IsPOBox           interface{} `json:"isPOBox"`
		IsResidential     interface{} `json:"isResidential"`
	} `json:"mailingAddress"`
	BillingAddress struct {
		MainCartAddressPK int         `json:"mainCartAddressPK"`
		FirstName         interface{} `json:"firstName"`
		LastName          interface{} `json:"lastName"`
		Company           interface{} `json:"company"`
		Address1          interface{} `json:"address1"`
		Address2          interface{} `json:"address2"`
		Address3          interface{} `json:"address3"`
		City              interface{} `json:"city"`
		State             interface{} `json:"state"`
		PostalCode        interface{} `json:"postalCode"`
		Country           interface{} `json:"country"`
		IsPOBox           interface{} `json:"isPOBox"`
		IsResidential     interface{} `json:"isResidential"`
	} `json:"billingAddress"`
	FormattedAutoshipSubtotal string     `json:"formattedAutoshipSubtotal"`
	CartItems                 []CartItem `json:"cartItems"`
}

type BillingAddress struct {
	FirstName string      `json:"firstName"`
	LastName  string      `json:"lastName"`
	ShipFName interface{} `json:"shipFName"`
	ShipLName interface{} `json:"shipLName"`
	Address1  string      `json:"address1"`
	Address2  string      `json:"address2"`
	Address3  string      `json:"address3"`
	City      string      `json:"city"`
	CityName  interface{} `json:"cityName"`
	Zip       string      `json:"zip"`
	State     struct {
		Code  string      `json:"code"`
		Name  string      `json:"name"`
		Name2 interface{} `json:"name2"`
	} `json:"state"`
	Phone       string `json:"phone"`
	SecondPhone string `json:"secondPhone"`
	Email       string `json:"email"`
	Country     struct {
		Code2  string      `json:"code2"`
		States interface{} `json:"states"`
	} `json:"country"`
	Ssn                  interface{} `json:"ssn"`
	Area                 string      `json:"area"`
	AreaName             interface{} `json:"areaName"`
	SiteUrl              interface{} `json:"siteUrl"`
	IsUseShippingAddress bool        `json:"isUseShippingAddress"`
}

type TrackingInfo struct {
	PackagePk                 int         `json:"packagePk"`
	ProductPk                 int         `json:"productPk"`
	PackageName               string      `json:"packageName"`
	ProductName               string      `json:"productName"`
	IsPackage                 bool        `json:"isPackage"`
	Quantity                  int         `json:"quantity"`
	Cv                        int         `json:"cv"`
	Sp                        int         `json:"sp"`
	Price                     int         `json:"price"`
	FormattedPrice            string      `json:"formattedPrice"`
	CurrencyCode              string      `json:"currencyCode"`
	ShipmentNumber            string      `json:"shipmentNumber"`
	ShipmentStatus            string      `json:"shipmentStatus"`
	ShippedDate               string      `json:"shippedDate"`
	TrackingNumber            string      `json:"trackingNumber"`
	TrackingLink              string      `json:"trackingLink"`
	VideoOrderPackagingInfoPK interface{} `json:"videoOrderPackagingInfoPK"`
}
