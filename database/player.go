package database

import (
	"fmt"
)

type Player struct {
	Id string `json:"nconst" bson:"nconst"`
	Name string `json:"primaryName" bson:"primaryName"`
	BirthYear int `json:"birthYear" bson:"birthYear,omitempty"`
	DeathYear int `json:"deathYear" bson:"deathYear,omitempty"`
	Profession string `json:"primaryProfession" bson:"primaryProfession,omitempty"`
	Url string `json:"url" bson:"url,omitempty"`
}

func (p Player) GetURL() string {
	return fmt.Sprintf("https://www.imdb.com/name/%s/", p.Id)
}