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
	"encoding/json"
	"log"
	"strconv"
)

func CreateBuilding(p ProposedBuilding, id int) (bool, error) {
	log.Println("INFO: Creating a panel: " + p.Name)
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

	panelInfo := `INSERT INTO Buildings (Name, City, Region, CreatorId) VALUES (?, ?, ?, ?)`
	q, err := t.Prepare(panelInfo)
	if err != nil {
		log.Println("ERROR: Cannot prepare DB query: " + string(err.Error()))
		return false, err
	}

	_, err = q.Exec(p.Name, p.City, p.Region, id)
	if err != nil {
		log.Println("ERROR: Cannot execute DB transaction: " + string(err.Error()))
		return false, err
	}

	err = t.Commit()
	if err != nil {
		log.Println("ERROR: Cannot commit DB transaction: " + string(err.Error()))
		return false, err
	}

	log.Println("INFO: Panel entry created")
	return true, nil
}

func GetBuildingById(id int) (Building, error) {
	log.Println("INFO: Getting building by Id")
	ent, err := DB.Prepare("SELECT * FROM Buildings WHERE Id = ?")
	if err != nil {
		return Building{}, err
	}
	defer ent.Close()

	building := Building{}
	record, err := ent.Query(id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("ERROR: No such building found in DB: " + string(err.Error()))
			return Building{}, nil
		}
		log.Println("ERROR: Cannot retrieve building from DB: " + string(err.Error()))
		return Building{}, err
	}
	err = record.Scan(
		&building.Id,
		&building.Name,
		&building.City,
		&building.Region,
		&building.CreatorId,
		&building.CreationDate,
	)
	if err != nil {
		log.Println("ERROR: Cannot unmarshal the building object!" + string(err.Error()))
		return Building{}, err
	}
	defer record.Close()

	return building, nil
}

func GetBuildingIdByName(buildingName string) (int, error) {
	log.Println("INFO: Getting building id by name")
	ent, err := DB.Prepare("SELECT Id FROM Buildings WHERE Name = ?")
	if err != nil {
		log.Println("ERROR: Cannot prepare SQL query: " + string(err.Error()))
		return -1, err
	}
	defer ent.Close()

	var id int
	record, err := ent.Query(buildingName)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println("ERROR: No such building found in DB: " + string(err.Error()))
			return -1, sql.ErrNoRows
		}
		log.Println("ERROR: Cannot retrieve building from DB: " + string(err.Error()))
		return -1, err
	}
	err = record.Scan(
		&id,
	)
	if err != nil {
		log.Println("ERROR: Cannot unmarshal the building object!" + string(err.Error()))
		return -1, err
	}
	defer record.Close()

	return id, nil
}

func GetBuildings() ([]Building, error) {
	log.Println("INFO: List of building objects requested")
	rows, err := DB.Query("SELECT * FROM Buildings")
	if err != nil {
		log.Println("ERROR: Could not run the DB query!" + string(err.Error()))
		return nil, err
	}
	defer rows.Close()

	log.Println("INFO: Constructing building list")
	buildings := make([]Building, 0)
	for rows.Next() {
		building := Building{}
		err = rows.Scan(
			&building.Id,
			&building.Name,
			&building.City,
			&building.Region,
			&building.CreatorId,
			&building.CreationDate,
		)
		if err != nil {
			log.Println("ERROR: Cannot marshal the building objects!" + string(err.Error()))
			return nil, err
		}
		buildings = append(buildings, building)
	}

	log.Println("INFO: List of all buildings retrieved")
	return buildings, nil
}

func UpdateBuildingById(id int, b BuildingUpdate) (bool, error) {
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

	q, err := t.Prepare("UPDATE Buildings SET Name = ?, City = ?, Region = ? WHERE Id = ?")
	if err != nil {
		log.Println("ERROR: Could not prepare DB query! " + string(err.Error()))
		return false, err
	}
	log.Println("INFO: Building ID to update: " + strconv.Itoa(id))
	log.Println("INFO: Incoming data: name: " + b.Name + ", city: " + b.City + ", region: " + b.Region)

	building, err := json.Marshal(b)
	if err != nil {
		log.Println("ERROR: Cannot marshal JSON: " + string(err.Error()))
		return false, err
	}
	_, err = q.Exec(building, id)
	if err != nil {
		log.Println("ERROR: Cannot execute DB query: " + string(err.Error()))
		return false, err
	}

	err = t.Commit()
	if err != nil {
		log.Println("ERROR: Cannot commit DB transaction: " + string(err.Error()))
		return false, err
	}

	log.Println("INFO: Building entry updated")
	return true, nil
}

func DeleteBuildingById(id int) (bool, error) {
	log.Println("INFO: User deletion requested for Id: " + strconv.Itoa(id))
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

	q, err := DB.Prepare("DELETE FROM Buildings WHERE Id IS ?")
	if err != nil {
		log.Println("ERROR: Could not prepare the DB query!" + string(err.Error()))
		return false, err
	}

	_, err = q.Exec(id)
	if err != nil {
		log.Println("ERROR: Cannot delete building with Id '" + strconv.Itoa(id) + "': " + string(err.Error()))
		return false, err
	}

	err = t.Commit()
	if err != nil {
		log.Println("ERROR: Cannot commit DB transaction: " + string(err.Error()))
		return false, err
	}

	log.Println("INFO: Building with Id '" + strconv.Itoa(id) + "' has been deleted")
	return true, nil
}
