package models

type TextInput struct {
	Text string `json:"text"`
}

type TextResult struct {
	Text     string `json:"text"`
	Category string `json:"category"`
}
