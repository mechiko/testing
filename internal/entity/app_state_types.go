package entity

type AppStateItem struct {
	Module string `json:"module"`
	Key    string `json:"key"`
	Value  string `json:"value"`
}

type AppStateList struct {
	Total int64          `json:"total"`
	Items []AppStateItem `json:"rows"`
}

type AppState interface {
	Set(module string, key string, value string) error
	Get(module string, key string) (string, error)
}
