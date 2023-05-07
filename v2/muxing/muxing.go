package muxing

type MuxingItem struct {
    Key  string `json:"key,omitempty"`
    Size int64  `json:"size,omitempty"`
}

type ProposeMuxing struct {
}
