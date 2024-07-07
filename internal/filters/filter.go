package filters

type Filter struct {
	Fields    []string         `json:"fields,omitempty"`
	Skip      int              `json:"skip,omitempty"`
	Limit     int              `json:"limit,omitempty"`
	OrderBy   string           `json:"orderBy,omitempty"`
	Order     string           `json:"order,omitempty"`
	Operators []FilterOperator `json:"filterOperator"`
}

type FilterOperator struct {
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
	Field    string      `json:"field"`
}
