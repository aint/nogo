package main

type HrResponse struct {
	Status string       `json:"status"`
	Data   DataResponse `json:"data"`
}

type DataResponse struct {
	CAD CurrencyResponse `json:"124"`
	JPY CurrencyResponse `json:"392"`
	RUB CurrencyResponse `json:"643"`
	CHF CurrencyResponse `json:"756"`
	GBP CurrencyResponse `json:"826"`
	USD CurrencyResponse `json:"840"`
	EUR CurrencyResponse `json:"978"`
	UAH CurrencyResponse `json:"980"`
	PLN CurrencyResponse `json:"985"`
}

type CurrencyResponse struct {
	Avg        CurrencyType `json:"avg"`
	Government CurrencyType `json:"government"`
	Commercial CurrencyType `json:"commercial"`
	Interbank  CurrencyType `json:"interbank"`
	Black      CurrencyType `json:"black"`
}

type CurrencyType struct {
	Buy     CurValue `json:"buy"`
	Sale    CurValue `json:"sale"`
	Average CurValue `json:"avg"`
}

type CurValue struct {
	Value float64 `json:"value,string"`
	Diff  float64 `json:"diff"`
}
