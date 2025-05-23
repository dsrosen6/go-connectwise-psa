package types

type WorkType struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	UtilizationFlag bool   `json:"utilizationFlag"`
	Info            Info   `json:"_info"`
}

type WorkRole struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info Info   `json:"_info"`
}
