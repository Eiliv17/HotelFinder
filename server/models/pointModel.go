package models

import "errors"

type Point struct {
	Type        string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

func CreatePoint(lat float64, lon float64) (Point, error) {

	if !(lon >= -180 && lon <= 180) {
		return Point{}, errors.New("longitude not valid")
	}
	if !(lat >= -90 && lat <= 90) {
		return Point{}, errors.New("latitude not valid")
	}

	return Point{
		Type:        "Point",
		Coordinates: []float64{lon, lat},
	}, nil
}
