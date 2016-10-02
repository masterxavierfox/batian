package engines

import (
	"time"
	"github.com/ishuah/batian/models"
)

type Report struct {
	ResponseTimes		[]map[string]float64
	RequestsPerMinute	[]map[string]int
}

func AppAnalysis(events models.Events) (Report, error) {
	var report Report
	collection := make(map[string][]float64)
	for _, event := range events {
		if event.Measurement == "requests" {
			rounded := time.Date(event.Timestamp.Year(), event.Timestamp.Month(), event.Timestamp.Day(), event.Timestamp.Hour(), event.Timestamp.Minute(), 0, 0, event.Timestamp.Location())
			
			if values, ok := collection[rounded.String()]; ok {
			    collection[rounded.String()] = append(values, event.Data["response_time"].(float64))
			} else {
				collection[rounded.String()] = append(collection[rounded.String()], event.Data["response_time"].(float64))
			}
		}
	}

	for timestamp, values := range collection {
		rpm := make(map[string]int)
		rpm[timestamp] = len(values)
		report.RequestsPerMinute = append(report.RequestsPerMinute, rpm)
		sumResponseTimes := 0.0
		for _, responseTime := range values {
			sumResponseTimes += responseTime
		}
		avgResponseTime := make(map[string]float64)
		avgResponseTime[timestamp] = sumResponseTimes/float64(len(values))
		report.ResponseTimes = append(report.ResponseTimes, avgResponseTime)
	}
	return report, nil
}