package models

type SimilarRequest struct {
	SourceCode string `json:"source_code"`
	ProblemUrl string `json:"problem_url"`
	Rating     int    `json:"rating"`
	Difficulty int    `json:"difficulty"`
	NRecs      int    `json:"n_recs"`
}

type MlTaskResponse struct {
	ProblemUrl string  `json:"problem_url"`
	Rating     float64 `json:"rating"`
	Tags       []int   `json:"tags"`
}

type MlResponse struct {
	Tasks []MlTaskResponse
}

type StoryItem1 struct {
	ProblemUrl string `json:"problem_url"`
	Rating     int    `json:"rating"`
	Tags       []int  `json:"tags"`
	NAttempts  int    `json:"n_attempts"`
}
type Story1 struct {
	Solved  []StoryItem1 `json:"solved"`
	TooEasy []StoryItem1 `json:"too_easy"`
	TooHard []StoryItem1 `json:"too_hard"`
}

type Problems struct {
	ProblemUrl string `json:"problem_url"`
	Rating     int    `json:"rating"`
	Tags       []int  `json:"tags"`
}

type Recommended struct {
	RecommendedTag int        `json:"recommended_tag"`
	Priority       int        `json:"priority"`
	Problems       []Problems `json:"problems"`
}

type Rec struct {
	Rec []Recommended `json:"rec"`
}

type RecommendedResponse struct {
	RecommendedTag string     `json:"recommended_tag"`
	Priority       int        `json:"priority"`
	Problems       []TaskJSON `json:"problems"`
}

type RecResponse struct {
	Rec []RecommendedResponse `json:"rec"`
}

type Progress struct {
	Tag  int  `json:"tag"`
	Done bool `json:"done"`
}

type ColdStartML struct {
	Finished    bool       `json:"finished"`
	ProblemUrl  string     `json:"problem_url"`
	Tag         int        `json:"tag"`
	Progress    []Progress `json:"progress"`
	ProblemTags []int      `json:"problem_tags"`
	Rating      int        `json:"rating"`
}

type ColdStartResponse struct {
	Finished bool       `json:"finished"`
	Progress []Progress `json:"progress"`
	Task     TaskJSON   `json:"task"`
}
