package types

type SystemLocation struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info Info   `json:"_info"`
}

type TimeZone struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info Info   `json:"_info"`
}

type Department struct {
	Id         int    `json:"id"`
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Info       Info   `json:"_info"`
}

type Member struct {
	Id            int    `json:"id"`
	Identifier    string `json:"identifier"`
	Name          string `json:"name"`
	DailyCapacity int    `json:"dailyCapacity"`
	Info          Info   `json:"_info"`
}

type Document struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info Info   `json:"_info"`
}
