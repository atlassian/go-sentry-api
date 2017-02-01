package datatype

//User implements the sentry interface for a authenticated user
type User struct {
	ID        *string                 `json:"id,omitempty"`
	Email     *string                 `json:"email,omitempty"`
	Username  *string                 `json:"username,omitempty"`
	IPAddress *string                 `json:"ipAddress,omitempty"`
	Data      *map[string]interface{} `json:"data,omitempty"`
}
