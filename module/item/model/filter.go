package todomodel

type Filter struct {
	Status string `json:"status,omitempty" form:"status"`
}
