package model

/*

  Copyright 2024, JAFAX, Inc.

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

*/

import (
	"database/sql"
	"log"
	"strconv"
)

func CreateLocation(p ProposedLocation, id int) (bool, error) {
	log.Println("INFO: Creating a room: " + p.RoomName)
	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Cannot start DB transaction: " + string(err.Error()))
		return false, err
	}

	locationInfo := `INSERT INTO Locations (
		RoomName, FloorId, BuildingId, CreatorId)
		VALUES (?, ?, ?, ?)`
	q, err := t.Prepare(locationInfo)
	if err != nil {
		log.Println("ERROR: Cannot prepare DB query: " + string(err.Error()))
		return false, err
	}

	_, err = q.Exec(p.RoomName, p.FloorId, p.BuildingId, id)
	if err != nil {
		log.Println("ERROR: Cannot execute DB transaction: " + string(err.Error()))
		return false, err
	}

	t.Commit()

	log.Println("INFO: Location entry created")

	return true, nil
}

func DeleteLocationById(id int) (bool, error) {
	log.Println("INFO: Location deletion requested: " + strconv.Itoa(id))
	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Could not start DB transaction!" + string(err.Error()))
		return false, err
	}

	q, err := DB.Prepare("DELETE FROM Locations WHERE Id IS ?")
	if err != nil {
		log.Println("ERROR: Could not prepare the DB query!" + string(err.Error()))
		return false, err
	}

	_, err = q.Exec(id)
	if err != nil {
		log.Println("ERROR: Cannot delete location with id '" + strconv.Itoa(id) + "': " + string(err.Error()))
		return false, err
	}

	t.Commit()

	log.Println("INFO: Location with id '" + strconv.Itoa(id) + "' has been deleted")
	return true, nil
}

func GetAllLocations() ([]Location, error) {
	log.Println("INFO: List of location objects requested")
	rows, err := DB.Query("SELECT * FROM Locations")
	if err != nil {
		log.Println("ERROR: Could not run the DB query!" + string(err.Error()))
		return nil, err
	}

	locations := make([]Location, 0)
	for rows.Next() {
		location := Location{}
		err = rows.Scan(
			&location.Id,
			&location.Location,
			&location.FloorId,
			&location.BuildingId,
			&location.CreatorId,
			&location.CreationDate,
		)
		if err != nil {
			log.Println("ERROR: Cannot marshal the location objects!" + string(err.Error()))
			return nil, err
		}
		locations = append(locations, location)
	}

	log.Println("INFO: List of all locations retrieved")
	return locations, nil
}

func GetLocationById(id int) (Location, error) {
	log.Println("INFO: Location by Id requested")
	ent, err := DB.Prepare("SELECT * FROM Locations WHERE Id = ?")
	if err != nil {
		log.Println("ERROR: Could not run the DB query!" + string(err.Error()))
		return Location{}, err
	}

	location := Location{}
	err = ent.QueryRow(id).Scan(
		&location.Id,
		&location.Location,
		&location.FloorId,
		&location.BuildingId,
		&location.CreatorId,
		&location.CreationDate,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("ERROR: No such location found in DB: " + string(err.Error()))
			return Location{}, nil
		}
		log.Println("ERROR: Cannot retrieve location from DB: " + string(err.Error()))
		return Location{}, err
	}

	log.Println("INFO: Location by Id '" + strconv.Itoa(id) + "' retrieved")
	return location, nil
}

func GetLocationsByFloorId(id int) ([]Location, error) {
	log.Println("INFO: List of locations on a floor requested")
	rows, err := DB.Query("SELECT * FROM Locations WHERE FloorId = ?", id)
	if err != nil {
		log.Println("ERROR: Could not run the DB query!" + string(err.Error()))
		return nil, err
	}

	locations := make([]Location, 0)
	for rows.Next() {
		location := Location{}
		err = rows.Scan(
			&location.Id,
			&location.Location,
			&location.FloorId,
			&location.BuildingId,
			&location.CreatorId,
			&location.CreationDate,
		)
		if err != nil {
			log.Println("ERROR: Cannot marshal the location objects!" + string(err.Error()))
			return nil, err
		}
		locations = append(locations, location)
	}

	log.Println("INFO: List of locations by floor Id retrieved")
	return locations, nil
}

func GetLocationsByBuildingId(id int) ([]Location, error) {
	log.Println("INFO: List of locations at a building requested")
	rows, err := DB.Query("SELECT * FROM Locations WHERE BuildingId = ?", id)
	if err != nil {
		log.Println("ERROR: Could not run the DB query!" + string(err.Error()))
		return nil, err
	}

	locations := make([]Location, 0)
	for rows.Next() {
		location := Location{}
		err = rows.Scan(
			&location.Id,
			&location.Location,
			&location.FloorId,
			&location.BuildingId,
			&location.CreatorId,
			&location.CreationDate,
		)
		if err != nil {
			log.Println("ERROR: Cannot marshal the location objects!" + string(err.Error()))
			return nil, err
		}
		locations = append(locations, location)
	}

	log.Println("INFO: List of locations by building Id retrieved")
	return locations, nil
}
