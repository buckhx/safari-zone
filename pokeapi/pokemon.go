package pokeapi

type Pokemon struct {
	ID                     int           `json:"id"`
	Name                   string        `json:"name"`
	Order                  int           `json:"order"`
	Height                 int           `json:"height"`
	Weight                 int           `json:"weight"`
	IsDefault              bool          `json:"is_default"`
	BaseExperience         int           `json:"base_experience"`
	LocationAreaEncounters string        `json:"location_area_encounters"`
	HeldItems              []interface{} `json:"held_items"`
	Forms                  []struct {
		URL  string `json:"url"`
		Name string `json:"name"`
	} `json:"forms"`
	Abilities []struct {
		Slot     int  `json:"slot"`
		IsHidden bool `json:"is_hidden"`
		Ability  struct {
			URL  string `json:"url"`
			Name string `json:"name"`
		} `json:"ability"`
	} `json:"abilities"`
	Stats []struct {
		Stat struct {
			URL  string `json:"url"`
			Name string `json:"name"`
		} `json:"stat"`
		Effort   int `json:"effort"`
		BaseStat int `json:"base_stat"`
	} `json:"stats"`
	Moves []struct {
		VersionGroupDetails []struct {
			MoveLearnMethod struct {
				URL  string `json:"url"`
				Name string `json:"name"`
			} `json:"move_learn_method"`
			LevelLearnedAt int `json:"level_learned_at"`
			VersionGroup   struct {
				URL  string `json:"url"`
				Name string `json:"name"`
			} `json:"version_group"`
		} `json:"version_group_details"`
		Move struct {
			URL  string `json:"url"`
			Name string `json:"name"`
		} `json:"move"`
	} `json:"moves"`
	Sprites struct {
		BackFemale       interface{} `json:"back_female"`
		BackShinyFemale  interface{} `json:"back_shiny_female"`
		BackDefault      string      `json:"back_default"`
		FrontFemale      interface{} `json:"front_female"`
		FrontShinyFemale interface{} `json:"front_shiny_female"`
		BackShiny        string      `json:"back_shiny"`
		FrontDefault     string      `json:"front_default"`
		FrontShiny       string      `json:"front_shiny"`
	} `json:"sprites"`
	Species struct {
		URL  string `json:"url"`
		Name string `json:"name"`
	} `json:"species"`
	GameIndices []struct {
		Version struct {
			URL  string `json:"url"`
			Name string `json:"name"`
		} `json:"version"`
		GameIndex int `json:"game_index"`
	} `json:"game_indices"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			URL  string `json:"url"`
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}
