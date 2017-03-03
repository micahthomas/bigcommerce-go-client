package bigcommerce

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestJSON(t *testing.T) {
	var v Product
	dec := json.NewDecoder(strings.NewReader(ProductData))
	if err := dec.Decode(&v); err != nil {
		t.Error(err)
	}

	jsonV, err := json.Marshal(&v)
	if err != nil {
		t.Error(err)
	}

	fmt.Print(ProductData)
	fmt.Print("\n\n")
	fmt.Print(string(jsonV))
	fmt.Print("\n\n")

	parsedProduct := &ParsedProduct{
		&v,
		[]BCBrand{
			BCBrand{ID: 1},
		},
	}

	jsonV, err = json.Marshal(&parsedProduct)
	if err != nil {
		t.Error(err)
	}

	fmt.Print(string(jsonV))
}

const ProductData = `{
  "id": 32,
  "keyword_filter": null,
  "name": "[Sample] Tomorrow is today, Red printed scarf",
  "type": "physical",
  "sku": "",
  "description": "Densely pack your descriptions with useful information and watch products fly off the shelf.",
  "search_keywords": null,
  "availability_description": "",
  "price": "89.0000",
  "cost_price": "0.0000",
  "retail_price": "0.0000",
  "sale_price": "0.0000",
  "calculated_price": "89.0000",
  "sort_order": 0,
  "is_visible": true,
  "is_featured": true,
  "related_products": "-1",
  "inventory_level": 0,
  "inventory_warning_level": 0,
  "warranty": null,
  "weight": "0.3000",
  "width": "0.0000",
  "height": "0.0000",
  "depth": "0.0000",
  "fixed_cost_shipping_price": "10.0000",
  "is_free_shipping": false,
  "inventory_tracking": "none",
  "rating_total": 0,
  "rating_count": 0,
  "total_sold": 0,
  "date_created": "Fri, 21 Sep 2012 02:31:01 +0000",
  "brand_id": 17,
  "view_count": 4,
  "page_title": "",
  "meta_keywords": null,
  "meta_description": null,
  "layout_file": "product.html",
  "is_price_hidden": false,
  "price_hidden_label": "",
  "categories": [
    14
  ],
  "date_modified": "Mon, 24 Sep 2012 01:34:57 +0000",
  "event_date_field_name": "Delivery Date",
  "event_date_type": "none",
  "event_date_start": "",
  "event_date_end": "",
  "myob_asset_account": "",
  "myob_income_account": "",
  "myob_expense_account": "",
  "peachtree_gl_account": "",
  "condition": "New",
  "is_condition_shown": false,
  "preorder_release_date": "",
  "is_preorder_only": false,
  "preorder_message": "",
  "order_quantity_minimum": 0,
  "order_quantity_maximum": 0,
  "open_graph_type": "product",
  "open_graph_title": "",
  "open_graph_description": null,
  "is_open_graph_thumbnail": true,
  "upc": null,
  "avalara_product_tax_code": "",
  "date_last_imported": "",
  "option_set_id": null,
  "tax_class_id": 0,
  "option_set_display": "right",
  "bin_picking_number": "",
  "custom_url": "/tomorrow-is-today-red-printed-scarf/",
  "primary_image": {
    "id": 247,
    "zoom_url": "https://cdn.url.path/bcapp/et7xe3pz/products/32/images/247/in_123__14581.1393831046.1280.1280.jpg?c=1",
    "thumbnail_url": "https://cdn.url.path/bcapp/et7xe3pz/products/32/images/247/in_123__14581.1393831046.220.290.jpg?c=1",
    "standard_url": "https://cdn.url.path/bcapp/et7xe3pz/products/32/images/247/in_123__14581.1393831046.386.513.jpg?c=1",
    "tiny_url": "https://cdn.url.path/bcapp/et7xe3pz/products/32/images/247/in_123__14581.1393831046.44.58.jpg?c=1"
  },
  "availability": "available",
  "brand": {
    "url": "https://store-et7xe3pz.mybigcommerce.com/api/v2/brands/17.json",
    "resource": "/brands/17"
  },
  "images": {
    "url": "https://store-et7xe3pz.mybigcommerce.com/api/v2/products/32/images.json",
    "resource": "/products/32/images"
  },
  "discount_rules": {
    "url": "https://store-et7xe3pz.mybigcommerce.com/api/v2/products/32/discountrules.json",
    "resource": "/products/32/discountrules"
  },
  "configurable_fields": {
    "url": "https://store-et7xe3pz.mybigcommerce.com/api/v2/products/32/configurablefields.json",
    "resource": "/products/32/configurablefields"
  },
  "custom_fields": {
    "url": "https://store-et7xe3pz.mybigcommerce.com/api/v2/products/32/customfields.json",
    "resource": "/products/32/customfields"
  },
  "videos": {
    "url": "https://store-et7xe3pz.mybigcommerce.com/api/v2/products/32/videos.json",
    "resource": "/products/32/videos"
  },
  "skus": {
    "url": "https://store-et7xe3pz.mybigcommerce.com/api/v2/products/32/skus.json",
    "resource": "/products/32/skus"
  },
  "rules": {
    "url": "https://store-et7xe3pz.mybigcommerce.com/api/v2/products/32/rules.json",
    "resource": "/products/32/rules"
  },
  "option_set": null,
  "options": {
    "url": "https://store-et7xe3pz.mybigcommerce.com/api/v2/products/32/options.json",
    "resource": "/products/32/options"
  },
  "tax_class": {
    "url": "https://store-et7xe3pz.mybigcommerce.com/api/v2/taxclasses/0.json",
    "resource": "/taxclasses/0"
  }
}`

// ***Experimenting with GO's JSON Marshalling for DateRFC2822 (Not Implemented)***
const DateInput = `"Mon, 02 Jan 2006 15:04:05 -0700"`
const DateOutput = `Mon, 02 Jan 2006 15:04:05 -0700`

func TestDateRFC2822(t *testing.T) {
	dateCreated := DateRFC2822(time.Now())
	if err := dateCreated.UnmarshalJSON([]byte(DateInput)); err != nil {
		t.Error(err)
	}

	output := time.Time(dateCreated).Format("Mon, 02 Jan 2006 15:04:05 -0700")

	if output != DateOutput {
		t.Error("Expected", DateOutput, "to match", output)
	}
}
