package model

type Traces struct {
	Data []Trace `json:"data"`
	JaegerRes
}

type Trace struct {
	TraceID string `json:"traceID"`
	Spans   []Span `json:"spans"`
}
