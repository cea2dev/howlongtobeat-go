package hltb

type SearchArgs struct {
	Term     string
	Page     int
	PageSize int
}

type searchRequest struct {
	SearchType    string        `json:"searchType"`
	SearchTerms   []string      `json:"searchTerms"`
	SearchPage    int           `json:"searchPage"`
	Size          int           `json:"size"`
	SearchOptions searchOptions `json:"searchOptions"`
	Users         searchUsers   `json:"users"`
	Filter        string        `json:"filter"`
	Sort          int           `json:"sort"`
	Randomizer    int           `json:"randomizer"`
}

type searchOptions struct {
	Games    searchOptionsGames    `json:"games"`
	Gameplay searchOptionsGameplay `json:"gameplay"`
	Modifier string                `json:"modifier"`
}

type searchOptionsGames struct {
	UserId        int                    `json:"userId"`
	Platform      string                 `json:"platform"`
	SortCategory  string                 `json:"sortCategory"`
	RangeCategory string                 `json:"rangeCategory"`
	RangeTime     searchOptionsRangeTime `json:"rangeTime"`
}

type searchOptionsRangeTime struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type searchOptionsGameplay struct {
	Perspective string `json:"perspective"`
	Flow        string `json:"flow"`
	Genre       string `json:"genre"`
}

type searchUsers struct {
	SortCategory string `json:"sortCategory"`
}

func getSearchRequestPayload(args SearchArgs) searchRequest {
	r := searchRequest{
		SearchType:  "games",
		SearchTerms: splitStrTerms(args.Term, " "),
		SearchPage:  args.Page,
		Size:        args.PageSize,
		SearchOptions: searchOptions{
			Games: searchOptionsGames{
				SortCategory:  "popular",
				RangeCategory: "main",
			},
		},
		Users: searchUsers{
			SortCategory: "postcount",
		},
	}

	return r
}
