package valueobject

import "fmt"

const (
	MERCHANT_CATEGORY_SMALL_RESTAURANT = iota
	MERCHANT_CATEGORY_MEDIUM_RESTAURANT
	MERCHANT_CATEGORY_LARGE_RESTAURANT
	MERCHANT_CATEGORY_MERCHANDISE_RESTAURANT
	MERCHANT_CATEGORY_BOOTH_KIOSK
	MERCHANT_CATEGORY_CONVENIENCE_STORE
)

var merchantCategory = map[int]string{
	MERCHANT_CATEGORY_SMALL_RESTAURANT:       "SmallRestaurant",
	MERCHANT_CATEGORY_MEDIUM_RESTAURANT:      "MediumRestaurant",
	MERCHANT_CATEGORY_LARGE_RESTAURANT:       "LargeRestaurant",
	MERCHANT_CATEGORY_MERCHANDISE_RESTAURANT: "MerchandiseRestaurant",
	MERCHANT_CATEGORY_BOOTH_KIOSK:            "BoothKiosk",
	MERCHANT_CATEGORY_CONVENIENCE_STORE:      "ConvenienceStore",
}

type MerchantCategory struct {
	types int
}

func CheckMerchantCategory(types interface{}) error {
	switch t := types.(type) {
	case int:
		if _, ok := merchantCategory[t]; ok {
			return nil
		}
	case string:
		for _, v := range merchantCategory {
			if v == t {
				return nil
			}
		}
	}

	return fmt.Errorf("invalid merchant category '%v'", types)
}

func NewMerchantCategory(types interface{}) (*MerchantCategory, error) {
	switch t := types.(type) {
	case int:
		if _, ok := merchantCategory[t]; ok {
			return &MerchantCategory{
				types: t,
			}, nil
		}
	case string:
		for k, v := range merchantCategory {
			if v == t {
				return &MerchantCategory{
					types: k,
				}, nil
			}
		}
	}

	return nil, fmt.Errorf("invalid merchant category '%v'", types)
}

func (mc *MerchantCategory) GetTypesInt() int {
	return mc.types
}

func (mc *MerchantCategory) GetTypesStr() string {
	return merchantCategory[mc.types]
}
