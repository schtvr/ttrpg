package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type MiniCharacter struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Creator string `json:"creator"`
}

var exampleCharacters = []MiniCharacter{
	{
		ID:      "1",
		Name:    "ExampleCharacter1",
		Creator: "ExampleCreator1",
	},
	{
		ID:      "2",
		Name:    "ExampleCharacter2",
		Creator: "ExampleCreator2",
	},
	{
		ID:      "3",
		Name:    "ExampleCharacter3",
		Creator: "ExampleCreator3",
	},
}

type Handler struct {
	// Database connection or other dependencies can be added here
}

func NewCharacterHandler() *Handler {
	return &Handler{}
}

func WriteStatusandEncode(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	if err := e.Encode(data); err != nil {
		log.Printf("error encoding response: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// GetCharacter fetches a character from the database
func (h *Handler) GetCharacter(w http.ResponseWriter, r *http.Request) {
	log.Println("GetCharacter called")

	params := mux.Vars(r)
	id := params["id"]
	log.Printf("fetching character with id: %s\n", id)
	character := findCharacterByID(id)
	if character == nil {
		log.Printf("character with id %s not found\n", id)
		http.Error(w, "Character not found", http.StatusNotFound)
		return
	}

	log.Printf("character found: %+v\n", character)
	WriteStatusandEncode(w, http.StatusOK, character)
}

func findCharacterByID(id string) *MiniCharacter {
	if id == "" {
		return nil
	}
	for _, c := range exampleCharacters {
		if c.ID == id || strings.EqualFold(c.Name, id) {
			return &c
		}
	}
	return nil
}

// ListCharacters lists all characters in the database
func (h *Handler) GetCharacters(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	e.Encode(exampleCharacters)
}

// AddNewCharacter adds a new character to the database
func (h *Handler) AddNewCharacter(w http.ResponseWriter, r *http.Request) {
	log.Println("AddNewCharacter called")

	// decode the request body
	var c MiniCharacter
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		log.Printf("error decoding character: %v\n", err)
		http.Error(w, "Invalid character data", http.StatusBadRequest)
		return
	}

	// validate the character data
	if c.Name == "" {
		log.Println("character name is required")
		http.Error(w, "character name is required", http.StatusBadRequest)
		return
	}
	if c.Creator == "" {
		c.Creator = "rando"
	}
	if match := findCharacterByID(c.Name); match != nil {
		log.Printf("character with id %s already exists\n", c.ID)
		http.Error(w, "character name already in use", http.StatusConflict)
		return
	}

	// generate a new ID for the character
	c.ID = uuid.New().String()
	exampleCharacters = append(exampleCharacters, c)

	WriteStatusandEncode(w, http.StatusCreated, c)
	log.Printf("new character added: %+v\n", c)
}

// CharacterRegistration registers a new character in the database
// CharacterUpdate updates a character in the database
// CharacterDelete deletes a character from the database
