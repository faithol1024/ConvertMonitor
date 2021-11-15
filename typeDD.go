package main

type DataDog struct {
	Title             string              `json:"title"`
	Description       string              `json:"description"`
	Widgets           []Widgets           `json:"widgets"`
	TemplateVariables []TemplateVariables `json:"template_variables"`
	LayoutType        string              `json:"layout_type"`
	IsReadOnly        bool                `json:"is_read_only"`
	NotifyList        []interface{}       `json:"notify_list"`
	ReflowType        string              `json:"reflow_type"`
	ID                string              `json:"id"`
}
type Style struct {
	Palette   string `json:"palette"`
	LineType  string `json:"line_type"`
	LineWidth string `json:"line_width"`
}

type Requests struct {
	Q                  string               `json:"q"`
	Formulas           []Formulas           `json:"formulas"`
	ConditionalFormats []ConditionalFormats `json:"conditional_formats"`
	ResponseFormat     string               `json:"response_format"`
	Queries            []Queries            `json:"queries"`
	Style              Style                `json:"style"`
	DisplayType        string               `json:"display_type"`
}
type Yaxis struct {
	IncludeZero bool   `json:"include_zero"`
	Scale       string `json:"scale"`
	Label       string `json:"label"`
	Min         string `json:"min"`
	Max         string `json:"max"`
}
type Definition struct {
	Title         string        `json:"title"`
	ShowLegend    bool          `json:"show_legend"`
	LegendSize    string        `json:"legend_size"`
	Type          string        `json:"type"`
	Requests      []Requests    `json:"requests"`
	Yaxis         Yaxis         `json:"yaxis"`
	ShowTitle     bool          `json:"show_title"`
	LayoutType    string        `json:"layout_type"`
	Widgets       []Widgets     `json:"widgets"`
	TitleSize     string        `json:"title_size"`
	AlertID       string        `json:"alert_id"`
	Unit          string        `json:"unit"`
	TextAlign     string        `json:"text_align"`
	Precision     int           `json:"precision"`
	LegendLayout  string        `json:"legend_layout"`
	LegendColumns []string      `json:"legend_columns"`
	Markers       []interface{} `json:"markers"`
	Autoscale     bool          `json:"autoscale"`
	CustomUnit    string        `json:"custom_unit"`
}
type Widgets struct {
	ID         int64      `json:"id"`
	Definition Definition `json:"definition"`
}
type Queries struct {
	Query      string `json:"query"`
	DataSource string `json:"data_source"`
	Name       string `json:"name"`
	Aggregator string `json:"aggregator"`
}
type ConditionalFormats struct {
	Comparator string `json:"comparator"`
	Palette    string `json:"palette"`
	Value      int    `json:"value"`
}
type Limit struct {
	Count int    `json:"count"`
	Order string `json:"order"`
}
type Formulas struct {
	Formula string `json:"formula"`
	Limit   Limit  `json:"limit"`
}

type TemplateVariables struct {
	Name            string        `json:"name"`
	Default         string        `json:"default"`
	Prefix          string        `json:"prefix"`
	AvailableValues []interface{} `json:"available_values"`
}
