package transformers

import (
	"misteraladin.com/jasmine/rate-structure/models"
)

type (
	TransformerInit struct {
		Data interface{} `json:"data"`
	}
)

type (
	TransformerBOInit struct {
		Data []interface{} `json:"data"`
	}
)

func (res *TransformerInit) TransformAvailableExtranet(rates []*models.RateExtranet) {
	var result =  make(map[string]interface{})
	for _, rate := range rates {
		rateRes := assignRateExtranet(rate)
		result[rate.RoomID] = rateRes
	}
	res.Data = result
}

func (res *TransformerBOInit) TransformAvailableBackoffice(rates []*models.RateBackoffice){
	for _, rate := range rates {
		res.Data = append(res.Data, assignRate(rate))
	}
}

