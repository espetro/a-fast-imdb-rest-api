package database

import (
	"fmt"
)

type Film struct {
	Id string `json:"tconst" bson:"tconst"`
	Title string `json:"primaryTitle" bson:"primaryTitle"`
	StartYear int `json:"startYear" bson:"startYear,omitempty"`
	EndYear int `json:"endYear" bson:"endYear,omitempty"`
	Genres string `json:"genres" bson:"genres,omitempty"`
	Type string `json:"titleType" bson:"titleType"`
	Rating int `json:"averageRating" bson:"averageRating,omitempty"`
	Url string `json:"url" bson:"url,omitempty"`
}

func (f Film) GetURL() string {
	return fmt.Sprintf("https://www.imdb.com/title/%s/", f.Id)
}