package transformers

import "misteraladin.com/jasmine/rate-structure/models"

type (
	CollectionTransformer struct {
		Data []interface{} `json:"data"`
		Meta models.Meta   `json:"meta"`
	}

	Transformer struct {
		Data interface{} `json:"data"`
	}
)
