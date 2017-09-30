package database

/*
func (db *DB) GetAllPlaces() ([]models.Place, error) {
	rows, err := sq.Select("*").
		From("places").
		RunWith(db.db).
		Query()

	if err != nil{
		return nil, err
	}

	var places []models.Place

	for rows.Next() {
		place := models.Place{}
		rows.Scan(
			&place.Place_id,
			&place.Place_name,
			&place.Importance,
		)

		places = append(places, place)
	}

	return places, nil
}

func (db *DB) InsertPlace(m models.Place) error {
	_, err := sq.Insert("tasks").
		Columns("Place_name", "important").
		Values(m.Place_name, m.Importance).
		RunWith(db.db).
		Exec()

	if err != nil {
		return err
	}

	return nil
}

/*
func (db *DB) DeletePlace(id int) error {
	rows, err := sq.Select("task_id").
		From("places").
		RunWith(db.db).
		Query()


	return nil
}
*/
