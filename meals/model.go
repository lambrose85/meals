package meals

type Meal struct {
	Name        string       `json:"name"`
	Day         string       `json:"day"`
	Ingredients []Ingredient `json:"ingredients"`
}
type Ingredient struct {
	Name   string `json:"name"`
	Amount string `json:"amount"`
}
