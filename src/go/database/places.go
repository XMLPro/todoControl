package database

import (
	"github.com/XMLPro/todoControl/src/go/models"
	sq "github.com/Masterminds/squirrel"
	"log"
)

func (db *DB) GetPlaces() []models.Place {
	rows, err := sq.Select("*").
		From("places").
		RunWith(db.db).
		Query()

	if err != nil {
		log.Fatal(err)
	}

	var places []models.Place
	for rows.Next() {
		var place models.Place
		err = rows.Scan(
			&place.Id,
			&place.PlaceName,
			&place.Importance,
		)

		if err != nil {
			log.Fatal(err)
		}

		places = append(places, place)
	}

	return places
}
