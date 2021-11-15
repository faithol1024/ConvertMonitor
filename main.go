package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var CurrCol = 1
var CurrRow = 1
var RowContent = 1

const (
	Height    int = 3
	Width     int = 8
	MaxWidth  int = 10
	AccountID int = 3044908
)

func main() {
	jsonFile, err := os.Open("datadog.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully Opened datadog.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var dataDog DataDog

	_ = json.Unmarshal(byteValue, &dataDog)

	visual := map[string]string{
		"timeseries":  "viz.line",
		"query_value": "viz.billboard",
		"toplist":     "viz.bar",
	}

	var newRelic NewRelic
	var widgetsNRs []WidgetsNR

	thresholds := []Thresholds{
		{
			AlertSeverity: "WARNING",
			Value:         90,
		},
		{
			AlertSeverity: "CRITICAL",
			Value:         95,
		},
	}
	x := 1
	for _, widgetDDHead := range dataDog.Widgets {
		for _, widgetDD := range widgetDDHead.Definition.Widgets {
			if visual[widgetDD.Definition.Type] == "" {
				continue
			}
			widget := WidgetsNR{
				Visualization: Visualization{ID: visual[widgetDD.Definition.Type]},
				Layout: Layout{
					Column: GetCol(),
					Row:    GetRow(),
					Height: Height,
					Width:  Width,
				},
				Title: widgetDD.Definition.Title,
				RawConfiguration: RawConfiguration{
					Legend: Legend{
						Enabled: true,
					},
				},
			}

			var queries []NrqlQueries
			query := NrqlQueries{
				AccountID: AccountID,
				Query:     queryBuild(widgetDD.Definition.Requests, widget.Visualization.ID),
			}
			queries = append(queries, query)

			widget.RawConfiguration.NrqlQueries = queries
			if widget.Visualization.ID == "viz.billboard" {
				widget.RawConfiguration.Thresholds = thresholds
			}
			widget.RawConfiguration.YAxisLeft.Zero = true
			widgetsNRs = append(widgetsNRs, widget)
			x++
			if x == 999 {
				break
			}
		}
		if x == 999 {
			break
		}
	}

	var pages []Pages

	for i := 1; i <= len(widgetsNRs)/99+1; i++ {
		page := Pages{}
		page.Name = fmt.Sprintf("Page %d", i)

		lowerBound := 0
		if i > 1 {
			lowerBound = (i - 1) * 99
			if lowerBound > len(widgetsNRs) {
				break
			}
		}
		upperBound := i * 99
		if upperBound > len(widgetsNRs) {
			upperBound = len(widgetsNRs)
		}
		println(upperBound)
		println(lowerBound)
		page.Widgets = widgetsNRs[lowerBound:upperBound]
		println(len(page.Widgets))
		pages = append(pages, page)
	}

	newRelic.Name = dataDog.Title
	newRelic.Permissions = "PUBLIC_READ_WRITE"
	newRelic.Pages = pages

	file, _ := json.MarshalIndent(newRelic, "", " ")

	_ = ioutil.WriteFile("newrelic.json", file, 0777)
	fmt.Println("[FINISH]")

}

func GetCol() int {

	col := CurrCol
	if CurrCol >= MaxWidth {
		CurrCol = 1
		return col
	}
	CurrCol += Width
	return col
}

func GetRow() int {

	row := CurrRow
	RowContent++
	if RowContent > 3 {
		CurrRow += Height
		RowContent = 1
	}
	return row
}

func queryBuild(reqs []Requests, visual string) string {
	operation := map[string]string{
		"viz.bar":       "from Metric SELECT percentage(sum(%s), where is_success = 'true')  facet %s SINCE 1 hour ago ",
		"viz.billboard": "from Metric SELECT percentage(sum(%s), where is_success = 'true')  since 1 hour ago",
	}
	operation2 := map[string]string{
		"line_time": "from Metric SELECT max(%s.summary) as 'max_time', average(%s.summary) as '95_perc?' SINCE 1 hour ago TIMESERIES ",
		"line":      "from Metric SELECT percentage(sum(%s), where is_success = 'true')  TIMESERIES since 1 hour ago ",
	}

	var req Requests
	if len(reqs) > 0 {
		req = reqs[0]
	}
	queryRaw := req.Q
	if queryRaw == "" {
		queryRaw = req.Queries[0].Query
	}

	facetRight := strings.Join(strings.Split(queryRaw, "by {")[1:], ":")
	facet := strings.Split(facetRight, "}")[0]

	metric := convertMetric(queryRaw)

	if visual == "viz.bar" {
		return fmt.Sprintf(operation[visual], metric, facet)
	}

	if visual == "viz.billboard" {
		return fmt.Sprintf(operation[visual], metric)
	}

	if visual == "viz.line" && strings.Contains(metric, "process_time") {
		return fmt.Sprintf(operation2["line_time"], metric, metric)
	}

	queryFull := fmt.Sprintf(operation2["line"], metric)
	fmt.Println(queryFull)
	return queryFull
}

func convertMetric(q string) string {
	left := strings.Split(q, "{")[0]
	right := strings.Join(strings.Split(left, ":")[1:], ":")
	query := strings.Replace(right, ".", "_", -1)
	query = strings.Replace(query, "process_time_max", "process_time", -1)
	return query
}
