package database

import (
	"log"
	"time"
	"context"
	"datapi/config"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	dbConn *mongo.Client
	filmColl, playerColl *mongo.Collection
	timeout = 5 * time.Second
)

func processErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	var err error

	uri := "mongodb://" + config.DBHost + ":" + config.DBPort
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	dbConn, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	processErr(err)

	err = dbConn.Ping(ctx, readpref.Primary())
	processErr(err)

	filmColl = dbConn.Database("imdb").Collection("films")
	playerColl = dbConn.Database("imdb").Collection("crew")
}


// GET /film?name=:name
// Searches a film object by its primary name and returns the result as a JSON.
func GetFilm(c *fiber.Ctx) {
	var film Film

	log.Println("Received a film query from " + c.IP())
	name := c.Query("name")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// The query is done via regex as it has shown to perform better
	// than normal queries or text searches
	searchStmt := bson.D{ { "$regex", name }, { "$options", "i" } }
	query := bson.D{ { "primaryTitle", searchStmt } }

	filmColl.FindOne(ctx, query).Decode(&film)

	// If the query doesn't return a film object, then report it
	if len(film.Title) == 0 {
		c.Status(204).Send("No film found for the given search parameters")
		return
	}

	film.Url = film.GetURL()
	c.JSON(film)
}

// GET /film/random
func GetRandomFilm(c *fiber.Ctx) {
	var results []Film
	
	log.Println("Received a random film query from " + c.IP())

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	sampleStage := bson.D{{ "$sample", bson.D{{ "size", 1 }} }}
	cursor, err := filmColl.Aggregate(ctx, mongo.Pipeline{sampleStage})
	processErr(err)

	if err = cursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}
	
	switch len(results) {
	case 0:
		c.Status(204).Send("No film found for the given search parameters")
		return
	default:
		var film Film = results[0]
		film.Url = film.GetURL()
		c.JSON(film)
	}
}

// GET /player?name=:name
func GetPlayer(c *fiber.Ctx) {
	var player Player
	name := c.Query("name")

	log.Println("Received a player query from " + c.IP())

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	searchStmt := bson.D{ { "$regex", name }, { "$options", "i" } }
	query := bson.D{ { "primaryName", searchStmt } }
	playerColl.FindOne(ctx, query).Decode(&player)

	// If the query doesn't return a player object, then report it
	if len(player.Name) == 0 {
		c.Status(204).Send("No player found for the given search parameters")
		return
	}

	player.Url = player.GetURL()
	c.JSON(player)
}

// GET /player/random
func GetRandomPlayer(c *fiber.Ctx) {
	var results []Player
	
	log.Println("Received a random player query from " + c.IP())

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	sampleStage := bson.D{{ "$sample", bson.D{{ "size", 1 }} }}
	cursor, err := playerColl.Aggregate(ctx, mongo.Pipeline{sampleStage})
	processErr(err)

	if err = cursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}

	switch len(results) {
	case 0:
		c.Status(204).Send("No player found for the given search parameters")
		return
	default:
		var player Player = results[0]
		player.Url = player.GetURL()
		c.JSON(player)
	}
}

func disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err := dbConn.Disconnect(ctx)
	processErr(err)
}

func Close() {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	
	if err := dbConn.Disconnect(ctx); err != nil {
		log.Fatal(err)
	}
}