package pokeapi

type Pokemon struct {
	Id             int               `json:"id"`
	Name           string            `json:"name"`
	BaseExperience int               `json:"base_experience"`
	Height         int               `json:"height"`
	IsDefault      bool              `json:"is_default"`
	Order          int               `json:"order"`
	Weight         int               `json:"weight"`
	Abilities      []map[string]any  `json:"abilities"`
	Forms          []map[string]any  `json:"forms"`
	GameIndices    []map[string]any  `json:"game_indices"`
	HeldItems      []map[string]any  `json:"held_items"`
	Moves          []map[string]any  `json:"moves"`
	Species        map[string]string `json:"species"`
	Sprites        map[string]any    `json:"sprites"`
	Cries          map[string]string `json:"cries"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	PastTypes     []map[string]any `json:"past_types"`
	PastAbilities []map[string]any `json:"past_abilities"`
}
