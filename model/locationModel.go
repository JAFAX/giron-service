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
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: DB transaction failed: " + string(err.Error()))
			t.Rollback()
		}
		if err != nil {
			log.Println("ERROR: DB transaction failed: " + string(err.Error()))
			t.Rollback()
		}
	}()

	locationInfo := `INSERT INTO Locations (RoomName, FloorId, BuildingId, CreatorId) VALUES (?, ?, ?, ?)`
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

	err = t.Commit()
	if err != nil {
		log.Println("ERROR: Cannot commit DB transaction: " + string(err.Error()))
		return false, err
	}

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
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: DB transaction failed: " + string(err.Error()))
			t.Rollback()
		}
		if err != nil {
			log.Println("ERROR: DB transaction failed: " + string(err.Error()))
			t.Rollback()
		}
	}()

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

	err = t.Commit()
	if err != nil {
		log.Println("ERROR: Cannot commit DB transaction: " + string(err.Error()))
		return false, err
	}

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
	defer rows.Close()

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
	defer ent.Close()

	location := Location{}
	record, err := ent.Query(id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("ERROR: No such location found in DB: " + string(err.Error()))
			return Location{}, nil
		}
		log.Println("ERROR: Cannot retrieve location from DB: " + string(err.Error()))
		return Location{}, err
	}
	defer record.Close()
	err = record.Scan(
		&location.Id,
		&location.Location,
		&location.FloorId,
		&location.BuildingId,
		&location.CreatorId,
		&location.CreationDate,
	)
	if err != nil {
		log.Println("ERROR: Cannot unmarshal the location object!" + string(err.Error()))
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
	defer rows.Close()

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
	defer rows.Close()

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

func UpdateLocationById(id int, l LocationUpdate) (bool, error) {
	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Could not start DB transaction: " + string(err.Error()))
		return false, err
	}
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: DB transaction failed: " + string(err.Error()))
			t.Rollback()
		}
		if err != nil {
			log.Println("ERROR: DB transaction failed: " + string(err.Error()))
			t.Rollback()
		}
	}()

	q, err := t.Prepare("UPDATE Locations SET FloorId = ?, BuildingId = ? WHERE Id = ?")
	if err != nil {
		log.Println("ERROR: Could not prepare DB query! " + string(err.Error()))
		return false, err
	}
	log.Println("INFO: Location ID to update: " + strconv.Itoa(id))
	log.Println("INFO: Incoming data: Floor Id: " + strconv.Itoa(l.FloorId) + ", Building Id: " + strconv.Itoa(l.BuildingId))

	_, err = q.Exec(l.FloorId, l.BuildingId, id)
	if err != nil {
		log.Println("ERROR: Cannot execute DB query: " + string(err.Error()))
		return false, err
	}

	err = t.Commit()
	if err != nil {
		log.Println("ERROR: Cannot commit DB transaction: " + string(err.Error()))
		return false, err
	}

	log.Println("INFO: Location entry updated")
	return true, nil
}
