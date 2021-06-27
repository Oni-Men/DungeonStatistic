package model

type Traces struct {
	Data []Trace `json:"data"`
	JaegerRes
}

type Trace struct {
	TraceID string `json:"traceID"`
	Spans   []Span `json:"spans"`
}

type Span struct {
	SpanID        string `json:"spanID"`
	OperationName string `json:"operationName"`
	StartTime     int64  `json:"startTime"`
	Duration      int64  `json:"duration"`
	References    []Ref  `json:"references"`
	Tags          []Tag  `json:"tags"`
	Logs          []Log  `json:"logs"`
}
