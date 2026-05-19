package store

import (
	"encoding/json"
	"time"
)

type (
	Object struct {
		ID        any               `json:"id"`
		Class     any               `json:"class"`
		CreatedBy any               `json:"createdBy"`
		CreatedAt time.Time         `json:"createdAt"`
		UpdatedBy any               `json:"updatedBy"`
		UpdatedAt time.Time         `json:"updatedAt"`
		DeletedBy any               `json:"deletedBy,omitempty"`
		DeletedAt time.Time         `json:"deletedAt,omitempty"`
		Data      []json.RawMessage `json:"data"`
	}
)
