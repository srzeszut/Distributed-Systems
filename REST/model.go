package main

type News struct {
	Title       string
	Description string
	Url         string
}

type Driver struct {
	Id                float64
	Name              string
	Country           string
	Teams             []string
	Wins              int
	Races             int
	Championships     int
	Points            string
	Podiums           int
	Number            int
	WinsPercentage    string
	PodiumsPercentage string
}
type DriverResponse struct {
	Get        string `json:"get"`
	Parameters struct {
		Search string `json:"search"`
	} `json:"parameters"`
	Errors   interface{} `json:"errors"`
	Results  int         `json:"results"`
	Response []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Abbr        string `json:"abbr"`
		Image       string `json:"image"`
		Nationality string `json:"nationality"`
		Country     struct {
			Name string `json:"name"`
			Code string `json:"code"`
		} `json:"country"`
		Birthdate          string `json:"birthdate"`
		Birthplace         string `json:"birthplace"`
		Number             int    `json:"number"`
		GrandsPrixEntered  int    `json:"grands_prix_entered"`
		WorldChampionships int    `json:"world_championships"`
		Podiums            int    `json:"podiums"`
		HighestRaceFinish  struct {
			Position int `json:"position"`
			Number   int `json:"number"`
		} `json:"highest_race_finish"`
		HighestGridPosition int    `json:"highest_grid_position"`
		CareerPoints        string `json:"career_points"`
		Teams               []struct {
			Season int `json:"season"`
			Team   struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Logo string `json:"logo"`
			} `json:"team"`
		} `json:"teams"`
	} `json:"response"`
}

type Result struct { //goroutines
	Data interface{}
	Err  error
}
