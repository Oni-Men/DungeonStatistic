package model

type JaegerRes struct {
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type Tag struct {
	Key   string
	Value string
}

type Log struct {
	Timestamp int64
	Fields    []Tag
}

func (log *Log) GetValue(key string) *string {
	return GetValueFromTags(log.Fields, key)
}

func (log *Log) ContainsKey(key string) bool {
	return TagContainsKey(log.Fields, key)
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
