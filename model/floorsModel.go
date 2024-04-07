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
	"log"
)

func CreateFloor(f ProposedFloor, id int) (bool, error) {
	log.Println("INFO: Creating a floor: " + f.Name)
	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Cannot start DB transaction: " + string(err.Error()))
		return false, err
	}

	// get building ID from building name
	buildingId, err := GetBuildingIdByName(f.BuildingName)
	if err != nil {
		log.Println("ERROR: Cannot prepare SQL query: " + string(err.Error()))
		return false, err
	}

	floorQuery := `INSERT INTO BuildingFloors (FloorName, BuildingId, CreatorId) VALUES (?, ?, ?)`
	q, err := t.Prepare(floorQuery)
	if err != nil {
		log.Println("ERROR: Cannot prepare DB query: " + string(err.Error()))
		return false, err
	}
	_, err = q.Exec(f.Name, buildingId, id)
	if err != nil {
		log.Println("ERROR: Cannot execute DB transaction: " + string(err.Error()))
		return false, err
	}

	t.Commit()

	log.Println("INFO: Floor entry created")

	return true, nil
}

func GetAllFloors() ([]BuildingFloor, error) {
	log.Println("INFO: List of floor objects requested")
	rows, err := DB.Query("SELECT * FROM BuildingFloors")
	if err != nil {
		log.Println("ERROR: Could not run the DB query!" + string(err.Error()))
		return nil, err
	}

	floors := make([]BuildingFloor, 0)
	for rows.Next() {
		floor := BuildingFloor{}
		err = rows.Scan(
			&floor.Id,
			&floor.FloorName,
			&floor.BuildingId,
			&floor.CreatorId,
			&floor.CreationDate,
		)
		if err != nil {
			log.Println("ERROR: Cannot marshal the floor objects!" + string(err.Error()))
			return nil, err
		}
		floors = append(floors, floor)
	}

	log.Println("INFO: List of all floors retrieved")
	return floors, nil
}
