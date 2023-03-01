package models

import (
	"context"
	"errors"

	"github.com/Eiliv17/HotelFinder/initializers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Hotel struct {
	Name       string `json:"name" bson:"name"`
	StarRating int    `json:"starRating" bson:"starRating"`
	Address    string `json:"address" bson:"address"`
	State      string `json:"state" bson:"state"`
	Location   Point  `json:"location" bson:"location"`
}

var coll *mongo.Collection = initializers.DB.Database(initializers.DBName).Collection("hotels")

func CreateHotel(name string, starRating int, address string, state string, location Point) (Hotel, error) {

	// check if star rating is valid
	if !(starRating >= 0 && starRating <= 5) {
		return Hotel{}, errors.New("starRating not valid")
	}

	return Hotel{
		Name:       name,
		StarRating: starRating,
		Address:    address,
		State:      state,
		Location:   location,
	}, nil
}

func RetrieveHotel(ctx context.Context, id string) (Hotel, error) {
	// transform id
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Hotel{}, err
	}

	// search document
	filter := bson.D{primitive.E{Key: "_id", Value: objID}}
	result := coll.FindOne(ctx, filter)

	// decode document
	var h Hotel
	err = result.Decode(&h)
	if err != nil {
		return Hotel{}, err
	}

	return h, nil
}

func AddHotel(ctx context.Context, hotel Hotel) error {
	// insert it
	_, err := coll.InsertOne(ctx, hotel)
	if err != nil {
		return err
	}

	return nil
}

func UpdateHotel(ctx context.Context, id string, hotel Hotel) error {
	// transform id
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	// update document
	filter := bson.D{primitive.E{Key: "_id", Value: objID}}
	_, err = coll.ReplaceOne(ctx, filter, hotel)
	if err != nil {
		return err
	}

	return nil
}

func DeleteHotel(ctx context.Context, id string) error {
	// transform id
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	// delete document
	filter := bson.D{primitive.E{Key: "_id", Value: objID}}
	_, err = coll.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func SearchHotel(ctx context.Context, coordinates Point, radius float64, offset int, limit int) ([]Hotel, error) {
	// covnersion to radiants
	radiusRadiants := radius / 6371.01

	// create the stages
	geoStage := bson.D{
		{"$geoNear", bson.D{
			{"near", bson.A{coordinates.Coordinates[0], coordinates.Coordinates[1]}},
			{"distanceField", "distance"},
			{"spherical", true},
			{"maxDistance", radiusRadiants},
		}},
	}
	sortStage := bson.D{{"$sort", bson.D{{"distance", 1}}}}
	skipStage := bson.D{{"$skip", offset}}
	limitStage := bson.D{{"$limit", limit}}

	// pass the pipeline to the Aggregate() method
	cursor, err := coll.Aggregate(ctx, mongo.Pipeline{geoStage, sortStage, skipStage, limitStage})
	if err != nil {
		return nil, err
	}

	// decode the results
	var results []Hotel
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}
