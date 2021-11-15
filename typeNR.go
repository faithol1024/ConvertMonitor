package main

type NewRelic struct {
	Name        string      `json:"name"`
	Description interface{} `json:"description"`
	Permissions string      `json:"permissions"`
	Pages       []Pages     `json:"pages"`
}
type Visualization struct {
	ID string `json:"id"`
}
type Layout struct {
	Column int `json:"column"`
	Row    int `json:"row"`
	Height int `json:"height"`
	Width  int `json:"width"`
}
type Legend struct {
	Enabled bool `json:"enabled"`
}
type NrqlQueries struct {
	AccountID int    `json:"accountId"`
	Query     string `json:"query"`
}
type YAxisLeft struct {
	Zero bool `json:"zero"`
}
type RawConfiguration struct {
	Legend         Legend        `json:"legend"`
	NrqlQueries    []NrqlQueries `json:"nrqlQueries"`
	YAxisLeft      YAxisLeft     `json:"yAxisLeft"`
	DataFormatters []interface{} `json:"dataFormatters"`
	Facet          Facet         `json:"facet"`
	Thresholds     []Thresholds  `json:"thresholds"`
}
type Thresholds struct {
	AlertSeverity string `json:"alertSeverity"`
	Value         int    `json:"value"`
}
type Facet struct {
	ShowOtherSeries bool `json:"showOtherSeries"`
}
type WidgetsNR struct {
	Visualization     Visualization    `json:"visualization"`
	Layout            Layout           `json:"layout"`
	Title             string           `json:"title"`
	RawConfiguration  RawConfiguration `json:"rawConfiguration,omitempty"`
	LinkedEntityGuids interface{}      `json:"linkedEntityGuids"`
}
type Pages struct {
	Name        string      `json:"name"`
	Description interface{} `json:"description"`
	Widgets     []WidgetsNR `json:"widgets"`
}
