package domain


import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// RPCRequest is the standard Geotab envelope
type RPCRequest struct {
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

// AuthParams for the Authenticate method
type AuthParams struct {
	Database string `json:"database"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

// GetFeedParams for ExceptionEvents
type GetFeedParams struct {
	TypeName     string      `json:"typeName"`
	FromVersion  *string     `json:"fromVersion,omitempty"`
	ResultsLimit int         `json:"resultsLimit,omitempty"`
	Credentials  Credentials `json:"credentials"`
}

type Credentials struct {
	Database string `json:"database"`
	UserName string `json:"userName"`
	SessionId string `json:"sessionId"`
}

// FeedResult reflects the structure you provided in the previous turn
type ExceptionEvent struct {
	ID           string  `json:"id"`
	Distance     float64 `json:"distance"`
	Duration     string  `json:"duration"` // "00:00:42"
	ActiveFrom   string  `json:"activeFrom"`
	Device       struct{ ID string `json:"id"` } `json:"device"`
	Version      string  `json:"version"`
}

type FeedResponse struct {
	Result struct {
		Data        []ExceptionEvent `json:"data"`
		ToVersion   string           `json:"toVersion"`
	} `json:"result"`
}
