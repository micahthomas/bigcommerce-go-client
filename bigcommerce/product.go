package bigcommerce

import (
	"fmt"
	"time"
)

// Product describes a BigCommerce Product Object
type Product struct {
	ID                      int64               `json:"id,omitempty"`                        // The unique numerical ID of the product. Increments sequentially.
	KeywordFilter           string              `json:"keyword_filter,omitempty"`            // (This property is deprecated.)
	Name                    string              `json:"name,omitempty"`                      // The product name.
	Type                    ProductType         `json:"type,omitempty"`                      // The product type.
	SKU                     string              `json:"sku,omitempty"`                       // User-defined product code/stock keeping unit (SKU).
	Description             string              `json:"description,omitempty"`               // Product description, which can include HTML formatting.
	SearchKeywords          string              `json:"search_keywords,omitempty"`           // A comma-separated list of keywords that can be used to locate the product when searching the store.
	AvailabilityDescription string              `json:"availability_description,omitempty"`  // Availability text, displayed on the checkout page under the product title, telling the customer how long it will normally take to ship this product. E.g.: “Usually ships in 24 hours”.
	Price                   string              `json:"price,omitempty"`                     // The product’s price. Should include, or exclude, tax based on the store settings.
	CostPrice               string              `json:"cost_price,omitempty"`                // The product’s cost price. Stored for reference only; not used or displayed anywhere on the store.
	RetailPrice             string              `json:"retail_price,omitempty"`              // The product’s retail cost. If entered, this retail price will be shown on the product page.
	SalePrice               string              `json:"sale_price,omitempty"`                // Sale price. If entered, this will be used instead of value in the price field when calculating the product’s cost.
	CalculatedPrice         string              `json:"calculated_price,omitempty"`          // Price as displayed to guests, adjusted for applicable sales and rules. (Cart price might incorporate further discounts for logged-in customers or customer groups.) Read-only.
	SortOrder               int64               `json:"sort_order,omitempty"`                // Priority to give this product when included in product lists on category pages and in search results. Lower integers will place the product closer to the top of the results.
	IsVisible               bool                `json:"is_visible,omitempty"`                // Flag to determine whether or not the product should be displayed to customers browsing. If true, the product will be displayed. If false, the product will be hidden from view.
	IsFeatured              bool                `json:"is_featured,omitempty"`               // Flag to determine whether the product should be included in the “featured products” panel for shoppers viewing the store.
	RelatedProducts         string              `json:"related_products,omitempty"`          //	Defaults to -1, which causes the store to automatically generate a list of related products. To manually specify the list of related products, include their IDs, separated by commas. For example: 3, 6, 7, 21.
	InventoryLevel          int64               `json:"inventory_level,omitempty"`           //	Current inventory level of the product. Simple inventory tracking must be enabled (see the inventory_tracking field) for this to take effect.
	InventoryWarningLevel   int64               `json:"inventory_warning_level,omitempty"`   //	Inventory Warning level for the product. When the product’s inventory level drops below this warning level, the store owner will be sent a notification. Simple inventory tracking must be enabled (see the inventory_tracking field) for this to take effect.
	Warranty                string              `json:"warranty,omitempty"`                  //	Warranty information displayed on the product page. Can include HTML formatting.
	Weight                  string              `json:"weight,omitempty"`                    //	Weight of the product, which can be used when calculating shipping costs.
	Width                   string              `json:"width,omitempty"`                     //	Width of the product, which can be used when calculating shipping costs.
	Height                  string              `json:"height,omitempty"`                    //	Height of the product, which can be used when calculating shipping costs.
	Depth                   string              `json:"depth,omitempty"`                     //	Depth of the product, which can be used when calculating shipping costs.
	FixedCostShippingPrice  string              `json:"fixed_cost_shipping_price,omitempty"` //	A fixed shipping cost for the product. If defined, this value will be used instead of normal shipping-cost calculation during checkout.
	IsFreeShipping          bool                `json:"is_free_shipping,omitempty"`          //	Flag used to indicate whether or not the product has free shipping. If true, the shipping cost for the product will be zero.
	InventoryTracking       *InventoryType      `json:"inventory_tracking,omitempty"`        //	The type of inventory tracking for the product. One of:
	RatingTotal             int64               `json:"rating_total,omitempty"`              //	The total rating for the product.
	RatingCount             int64               `json:"rating_count,omitempty"`              //	The total number of ratings the product has had.
	TotalSold               int64               `json:"total_sold,omitempty"`                //	Total quantity of this product sold through transactions.
	DateCreated             string              `json:"date_created,omitempty"`              //	The date of which the product was created.
	BrandID                 int64               `json:"brand_id,omitempty"`                  // The product’s brand
	ViewCount               int64               `json:"view_count,omitempty"`                // The number of times the product has been viewed.
	PageTitle               string              `json:"page_title,omitempty"`                // Custom title for the product’s page. If not defined, the product name will be used as the page title.
	MetaKeywords            string              `json:"meta_keywords,omitempty"`             // Custom meta keywords for the product page. If not defined, the store’s default keywords will be used.
	MetaDescription         string              `json:"meta_description,omitempty"`          // Custom meta description for the product page. If not defined, the store’s default meta description will be used.
	LayoutFile              string              `json:"layout_file,omitempty"`               // The layout template file used to render this product category.
	IsPriceHidden           bool                `json:"is_price_hidden,omitempty"`           // The default false value indicates that this product’s price should be shown on the product page. If set to true, the price will be hidden hidden. (NOTE: To successfully set is_price_hidden to true, the availability value must be disabled.)
	PriceHiddenLabel        string              `json:"price_hidden_label,omitempty"`        // By default, an empty string. If is_price_hidden is true, the value of price_hidden_label will be displayed instead of the price. (NOTE: To successfully set a non-empty string value for price_hidden_label, the availability value must be disabled.)
	Categories              []int64             `json:"categories,omitempty"`                // An array of IDs for the categories this product belongs to. When updating a product, if an array of categories is supplied, then all product categories will be overwritten. Does not accept more than 1,000 ID values.
	DateModified            string              `json:"date_modified,omitempty"`             // The date that the product was last modified.
	EventDateFieldName      string              `json:"event_date_field_name,omitempty"`     // Name of the field to be displayed on the product page when selecting the event/delivery date.
	EventDateType           *EventDateFieldType `json:"event_date_type,omitempty"`           // One of the following values:
	EventDateStart          string              `json:"event_date_start,omitempty"`          // When the product requires the customer to select an event/delivery date, this date is used as the “after” date.
	EventDateEnd            string              `json:"event_date_end,omitempty"`            // When the product requires the customer to select an event/delivery date, this date is used as the “before” date.
	MYOBAssetAccount        string              `json:"myob_asset_account,omitempty"`        // MYOB Asset Account.
	MYOBIncomeAccount       string              `json:"myob_income_account,omitempty"`       // MYOB Income Account.
	MYOBExpenseAccount      string              `json:"myob_expense_account,omitempty"`      // MYOB Expense/COS Account.
	PeachtreeGLAccount      string              `json:"peachtree_gl_account,omitempty"`      // Peachtree General Ledger Account.
	Condition               string              `json:"condition,omitempty"`                 // The product’s condition. Will be shown on the product page if the value of the is_condition_shown field is true. Possible values: New, Used, Refurbished.
	IsConditionShown        bool                `json:"is_condition_shown,omitempty"`        // Flag used to determine whether the product’s condition will be shown to the customer on the product page.
	PreorderReleaseDate     string              `json:"preorder_release_date,omitempty"`     // Pre-order release date. See availability field for details on setting a product’s availability to accept pre-orders.
	IsPreorderOnly          bool                `json:"is_preorder_only,omitempty"`          // If set to false, the product will not change its availability from preorder to available on the release date. Otherwise, on the release date the product’s availability/status will change to available.
	PreorderMessage         string              `json:"preorder_message,omitempty"`          // Custom expected-date message to display on the product page. If undefined, the message defaults to the storewide setting. Can contain the %%DATE%% placeholder, which will be replaced with the release date.
	OrderQuantityMinimum    int64               `json:"order_quantity_minimum,omitempty"`    // The minimum quantity an order must contain in order to purchase this product.
	OrderQuantityMaximum    int64               `json:"order_quantity_maximum,omitempty"`    // The maximum quantity an order can contain when purchasing the product.
	OpenGraphType           string              `json:"open_graph_type,omitempty"`           // Type of product. Acceptable values are: product, album, book, drink, food, game, movie, song, tv_show
	OpenGraphTitle          string              `json:"open_graph_title,omitempty"`          // Title of the product. If not specified, the product’s name will be used instead.
	OpenGraphDescription    string              `json:"open_graph_description,omitempty"`    // Description to use for the product. If not specified, the meta_description will be used instead.
	IsOpenGraphThumbnail    bool                `json:"is_open_graph_thumbnail,omitempty"`   // If set to true, the product thumbnail image will be used as the open graph image.
	UPC                     string              `json:"upc,omitempty"`                       // The product UPC code, which is used in feeds for shopping comparison sites.
	DateLastImported        string              `json:"date_last_imported,omitempty"`        // The date on which the product was last imported using the bulk importer.
	OptionSetID             int64               `json:"option_set_id,omitempty"`             // The ID of the option set applied to the product. (NOTE: To remove the option set from the product, set the value to null on update.)
	TaxClassID              int64               `json:"tax_class_id,omitempty"`              // The ID of the tax class applied to the product. (NOTE: Value ignored if automatic tax is enabled.)
	OptionSetDisplay        string              `json:"option_set_display,omitempty"`        // The position on the product page where options from the option set will be displayed.
	BinPickingNumber        string              `json:"bin_picking_number,omitempty"`        // The BIN picking number for the product.
	CustomURL               string              `json:"custom_url,omitempty"`                // Custom URL (if set) overriding the structure dictated in the store’s settings. If no custom URL is set, this will contain the default URL.
	PrimaryImage            *ProductImage       `json:"primary_image,omitempty"`             // An image object, corresponding to the image that is set as the product’s thumbnail. This object includes that image’s id, plus four URL values identifying where to pull the image at different sizes:
	Availability            ProductAvailability `json:"availability,omitempty"`              // Availability of the product
	Brand                   *BCResource         `json:"brand,omitempty"`                     // The product’s brand
	DiscountRules           *BCResource         `json:"discount_rules,omitempty"`            // See the Bulk Pricing/Discount resource for information.
	ConfigurableFields      *BCResource         `json:"configurable_fields,omitempty"`       // See the Configurable Fields resource for information.
	CustomFields            *BCResource         `json:"custom_fields,omitempty"`             // See the Custom Fields resource for information.
	Rules                   *BCResource         `json:"rules,omitempty"`                     // Rules that apply only to this product, based on the product’s option set. See Product Rules resource for information.
	OptionSet               *BCResource         `json:"option_set,omitempty"`                // See the Product Option Sets resource for information.
	Options                 *BCResource         `json:"options,omitempty"`                   // Options from the option set applied to the product. See the Product Options resource for information.
}

// ParsedProduct is a complete product containing parsed brands, discount_rules, custom_fields, etc...
type ParsedProduct struct {
	*Product
	Brand []BCBrand `json:"brand,omitempty"`
}

// BCBrand describes a brand object for BigCommerce
type BCBrand struct {
	ID              int    `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	PageTitle       string `json:"page_title,omitempty"`
	MetaKeywords    string `json:"meta_keywords,omitempty"`
	MetaDescription string `json:"meta_description,omitempty"`
	ImageFile       string `json:"image_file,omitempty"`
	SearchKeywords  string `json:"search_keywords,omitempty"`
}

// BCResource - BigCommerce Resource Endpoint
type BCResource struct {
	URL      string `json:"url,omitempty"`
	Resource string `json:"resource,omitempty"`
}

// ProductType - The product type
type ProductType string

// InventoryType - The type of inventory tracking for the product
type InventoryType string

// EventDateFieldType - One of the following values:
type EventDateFieldType string

// ProductAvailability - Availability of the product.
type ProductAvailability string

// ProductImage describes a BigCommerce Product's image field
type ProductImage struct {
	ID           int64  `json:"id,omitempty"`
	ZoomURL      string `json:"zoom_url,omitempty"`
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
	StandardURL  string `json:"standard_url,omitempty"`
	TinyURL      string `json:"tiny_url,omitempty"`
}

const (
	// PhysicalProduct describes a Physical Product
	PhysicalProduct ProductType = "physical"
	// DigitalProduct describes a Digital Product
	DigitalProduct ProductType = "digital"

	// NoInventory - inventory levels will not be tracked.
	NoInventory InventoryType = "none"
	// SimpleInventory - inventory levels will be tracked using the inventory_level and inventory_warning_level fields
	SimpleInventory InventoryType = "simple"
	// SKUInventory - inventory levels will be tracked based on individual product options, which maintain their own warning levels and inventory levels
	SKUInventory InventoryType = "sku"

	// NoEventDateField - Disables the event/delivery date requirement and field.
	NoEventDateField EventDateFieldType = "none"
	// AfterEventDateField – The selected date must fall either on, or after, the date specified in the event_date_start field.
	AfterEventDateField EventDateFieldType = "after"
	// BeforeEventDateField - The selected date must fall either before, or on, the date specified in the event_date_end field.
	BeforeEventDateField EventDateFieldType = "before"
	// RangeEventDateField - The selected date must fall between the event_date_start and event_date_end dates.
	RangeEventDateField EventDateFieldType = "range"

	// AvailableProduct - the product can be purchased on the storefront.
	AvailableProduct ProductAvailability = "available"
	// DisabledProduct - the product is listed on the storefront, but cannot be purchased.
	DisabledProduct ProductAvailability = "disabled"
	// PreorderProduct – the product is listed for pre-orders.
	PreorderProduct ProductAvailability = "preorder"
)

// DateRFC2822 describes RFC2822 type of Date, used by BigCommerce
// ***Experimenting with GO's JSON Marshalling for DateRFC2822 (Not Implemented)***
type DateRFC2822 time.Time

const rfc2822 = "Mon, 02 Jan 2006 15:04:05 -0700"

// UnmarshalJSON handles the JSON Conversion from RFC2822 to time.Time
func (t *DateRFC2822) UnmarshalJSON(data []byte) error {
	var timeString string
	if len(data) >= 30 && data[0] == '"' {
		timeString = string(data[1 : len(data)-1])
	} else if len(data) < 30 {
		t = nil
		return nil
	}

	parsedTime, err := time.Parse(rfc2822, timeString)
	if err == nil {
		*t = DateRFC2822(parsedTime)
	} else {
		*t = DateRFC2822(time.Now())
	}

	return err
}

// MarshalJSON handles DateRFC2822 to JSON epoch conversion
func (t *DateRFC2822) MarshalJSON() ([]byte, error) {
	epoch := time.Time(*t).Unix()
	if (epoch < 0) || t == nil {
		return []byte(`",omitempty"`), nil
	}

	stamp := fmt.Sprint(epoch)
	return []byte(stamp), nil
}
