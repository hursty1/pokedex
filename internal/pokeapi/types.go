package pokeapi

type LocationAreaResponse struct {
	Count 			int 	`json:"count"`
	Next			*string	`json:"next"`
	Previous		*string	`json:"previous"`
	Results			[]LocationArea `json:"results"`
}
type LocationArea struct {
	Id	 			int 	`json:"id"`
	Name	 		string 	`json:"name"`
	Game_Index	 	int 	`json:"game_index"`
	// encounter_method_rates	 []EncounterMethodRate
	// location	 NamedAPIResource
	// names	 []Name
	// pokemon_encounters	[]PokemonEncounter
}

const (
	baseURL = "https://pokeapi.co/api/v2"
)