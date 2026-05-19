package client

import (
	"bufio"
	"bytes"
	"cwc/httpcli"
	"cwc/utils"
	"net/http"
	"strings"
)

type MetricSample struct {
	Name   string            `json:"name"`
	Labels map[string]string `json:"labels"`
	Value  string            `json:"value"`
}

type DisplayMetricSample struct {
	Name   string
	Labels string
	Value  string
}

func GetMetrics() ([]MetricSample, error) {
	return fetchAndParseMetrics("")
}

func GetMetricByName(name string) ([]MetricSample, error) {
	return fetchAndParseMetrics(name)
}

func fetchAndParseMetrics(filterName string) ([]MetricSample, error) {
	cli := &http.Client{}
	body, err := httpcli.HttpRequest(cli, "/metrics", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}

	defer body.Close()

	var samples []MetricSample
	scanner := bufio.NewScanner(body)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if utils.IsBlank(line) || strings.HasPrefix(line, "#") {
			continue
		}

		sample, ok := parseMetricLine(line)
		if !ok {
			continue
		}

		if utils.IsNotBlank(filterName) && sample.Name != filterName {
			continue
		}

		samples = append(samples, sample)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return samples, nil
}

func parseMetricLine(line string) (MetricSample, bool) {
	lastSpace := strings.LastIndex(line, " ")
	if lastSpace < 0 {
		return MetricSample{}, false
	}

	value := strings.TrimSpace(line[lastSpace+1:])
	rest := strings.TrimSpace(line[:lastSpace])

	labels := map[string]string{}
	name := rest

	if idx := strings.Index(rest, "{"); idx >= 0 {
		end := strings.LastIndex(rest, "}")
		if end < 0 {
			return MetricSample{}, false
		}
		name = rest[:idx]
		labelsStr := rest[idx+1 : end]
		labels = parsePrometheusLabels(labelsStr)
	}

	return MetricSample{
		Name:   name,
		Labels: labels,
		Value:  value,
	}, true
}

func parsePrometheusLabels(s string) map[string]string {
	result := map[string]string{}
	if utils.IsBlank(s) {
		return result
	}

	parts := strings.Split(strings.TrimSpace(s), ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		eqIdx := strings.Index(part, "=")
		if eqIdx < 0 {
			continue
		}

		k := strings.TrimSpace(part[:eqIdx])
		v := strings.TrimSpace(part[eqIdx+1:])
		v = strings.Trim(v, `"`)
		result[k] = v
	}

	return result
}
