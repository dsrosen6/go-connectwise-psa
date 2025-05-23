package types

type Opportunity struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info Info   `json:"_info"`
}
