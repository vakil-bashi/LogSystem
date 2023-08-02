package ES

type SearchRES struct {
	Hits     SearchRESHits `json:"hits"`
	TimedOut bool          `json:"timed_out"`
	Took     int64         `json:"took"`
}

type SearchRESHits struct {
	Hits     []SearchREQHitsHits    `json:"hits"`
	MaxScore float64                `json:"max_score"`
	Total    map[string]interface{} `json:"total"`
}

type SearchREQHitsHits struct {
	Id     string                 `json:"_id"`
	Index  string                 `json:"_index"`
	Score  float64                `json:"_score"`
	Source map[string]interface{} `json:"_source"`
}
