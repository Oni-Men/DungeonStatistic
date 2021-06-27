package model

type JaegerRes struct {
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type Ref struct {
	RefType string `json:"refType"`
	TraceID string `json:"traceID"`
	SpanID  string `json:"spanID"`
}

type Tag struct {
	Key   string
	Value string
}

type Log struct {
	Timestamp uint64
	Fields    []Tag
}

func GetValueFromTags(tags []Tag, key string) *string {
	for _, tag := range tags {
		if tag.Key == key {
			return &tag.Value
		}
	}
	return nil
}

func TagContainsKey(tags []Tag, key string) bool {
	return GetValueFromTags(tags, key) != nil
}
