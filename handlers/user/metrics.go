package user

import (
	"cwc/client"
	"cwc/config"
	"cwc/utils"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func HandleListMetrics(samples []client.MetricSample, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		displayMetricsAsTable(samples)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(samples)
	} else {
		display := toDisplaySamples(samples)
		utils.PrintMultiRow(client.DisplayMetricSample{}, display)
	}
}

func HandleGetMetric(name string, samples []client.MetricSample, pretty *bool, valueOnly bool) {
	if valueOnly {
		displayMetricValues(samples)
		return
	}

	if utils.IsEmpty(samples) {
		fmt.Printf("No metrics found for name: %s\n", name)
		return
	}

	if config.IsPrettyFormatExpected(pretty) {
		displayMetricsAsTable(samples)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(samples)
	} else {
		display := toDisplaySamples(samples)
		utils.PrintMultiRow(client.DisplayMetricSample{}, display)
	}
}

func displayMetricValues(samples []client.MetricSample) {
	values := make([]string, 0, len(samples))
	for _, s := range samples {
		values = append(values, s.Value)
	}

	utils.PrintArray(values)
}

func toDisplaySamples(samples []client.MetricSample) []client.DisplayMetricSample {
	display := make([]client.DisplayMetricSample, 0, len(samples))
	for _, s := range samples {
		display = append(display, client.DisplayMetricSample{
			Name:   s.Name,
			Labels: formatMetricLabels(s.Labels),
			Value:  s.Value,
		})
	}
	return display
}

func formatMetricLabels(labels map[string]string) string {
	if utils.IsEmpty(labels) {
		return ""
	}

	keys := make([]string, 0, len(labels))
	for k := range labels {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s=%s", k, labels[k]))
	}

	return strings.Join(parts, ", ")
}

func displayMetricsAsTable(samples []client.MetricSample) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Labels", "Value"})
	table.SetAutoWrapText(false)
	table.SetColWidth(60)

	if utils.IsEmpty(samples) {
		table.Append([]string{"No metrics available", "", ""})
		table.Render()
		return
	}

	for _, s := range samples {
		table.Append([]string{
			s.Name,
			formatMetricLabels(s.Labels),
			s.Value,
		})
	}
	table.Render()
}
