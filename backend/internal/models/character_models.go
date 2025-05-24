package models

type Character struct {
	Name      string         `json:"name"`
	Epithet   string         `json:"epithet"`
	Skills    map[string]int `json:"skills"`
	About     string         `json:"about"`
	Channels  []Channel      `json:"channels"`
	Specials  []Special      `json:"specials"`
	Aspects   []Aspect       `json:"aspects"`
	Wounds    []Wound        `json:"wounds"`
	Inventory []Item         `json:"inventory"`
	Modifiers map[string]int `json:"modifiers"`
}

type Channel struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Effects     []Effect `json:"effects"`
}

type Aspect struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Effects     []Effect `json:"effects"`
}

type Wound struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Effects     []Effect `json:"effects"`
}

type Item struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Effects     []Effect `json:"effects"`
}

type Special struct {
	// Placeholder — define fields when data is known
}

type Effect struct {
	Skills      []string `json:"skills"`
	Actions     []string `json:"actions"`
	Modifier    int      `json:"modifier"`
	Channel     string   `json:"channel,omitempty"`
	Description string   `json:"description,omitempty"`
}
