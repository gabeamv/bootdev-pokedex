package pokeapi

type LocationArea struct {
	EncounterMethodRates []map[string]any  `json:"encounter_method_rates"`
	Location             map[string]string `json:"location"`
	Names                []map[string]any  `json:"names"`
	PokemonEncounters    []struct {
		Pokemon struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []map[string]any `json:"version_details"`
	} `json:"pokemon_encounters"`
}
