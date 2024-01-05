package models

type Credit struct {
	ID              int `json:"id"`
	RawScore        int `json:"rawScore"`
	ScorePercentage int `json:"scorePercentage"`
}
