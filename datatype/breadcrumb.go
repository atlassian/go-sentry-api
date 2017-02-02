package datatype

import (
	"time"
)

//BreadcrumbValue represents a single breadcrumb
type BreadcrumbValue struct {
	Category  *string                 `json:"category,omitempty"`
	Data      *map[string]interface{} `json:"data,omitempty"`
	EventID   *string                 `json:"event_id,omitempty"`
	Level     *string                 `json:"level,omitempty"`
	Message   *string                 `json:"message,omitempty"`
	Timestamp *time.Time              `json:"timestamp,omitempty"`
	Type      *string                 `json:"type,omitempty"`
}

// Breadcrumb represents the breadcrumb interface in sentry
type Breadcrumb struct {
	Values *[]BreadcrumbValue `json:"values,omitempty"`
}
