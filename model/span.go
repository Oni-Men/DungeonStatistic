package model

type Span struct {
	SpanID        string `json:"spanID"`
	OperationName string `json:"operationName"`
	StartTime     int64  `json:"startTime"`
	Duration      int64  `json:"duration"`
	Tags          []Tag  `json:"tags"`
	Logs          []Log  `json:"logs"`
}

func (s *Span) FindLog(key string) *Log {
	for _, log := range s.Logs {
		for _, field := range log.Fields {
			if field.Key == key {
				return &log
			}
		}
	}
	return nil
}

func (s *Span) GetTagValue(key string) *string {
	return GetValueFromTags(s.Tags, key)
}
