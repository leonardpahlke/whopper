package util

import "encoding/json"

// This structure is used to send a query request over dapr to a statestore
// see docs: https://docs.dapr.io/reference/api/state_api/#state-query

type StatestoreQuery []StatestoreQueryElement

func UnmarshalStatestoreQuery(data []byte) (StatestoreQuery, error) {
	var r StatestoreQuery
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StatestoreQuery) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type StatestoreQueryElement struct {
	Filter Filter `json:"filter"`
	Sort   []Sort `json:"sort,omitempty"`
	Page   *Page  `json:"page,omitempty"`
}

type Filter struct {
	Or []Or      `json:"OR,omitempty"`
	In *FilterIN `json:"IN,omitempty"`
	Eq *FilterEQ `json:"EQ,omitempty"`
}

type FilterEQ struct {
	ValueState string `json:"value.state"`
}

type FilterIN struct {
	ValuePersonOrg []string `json:"value.person.org"`
}

type Or struct {
	Eq  *AndEq `json:"EQ,omitempty"`
	And []And  `json:"AND,omitempty"`
}

type And struct {
	Eq *AndEq `json:"EQ,omitempty"`
	In *AndIn `json:"IN,omitempty"`
}

type AndEq struct {
	ValuePersonOrg string `json:"value.person.org"`
}

type AndIn struct {
	ValueState []string `json:"value.state"`
}

type Page struct {
	Limit int64   `json:"limit"`
	Token *string `json:"token,omitempty"`
}

type Sort struct {
	Key   string  `json:"key"`
	Order *string `json:"order,omitempty"`
}
