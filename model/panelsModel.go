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

func CreatePanel(p ProposedPanel, id int) (bool, error) {
	log.Println("INFO: Creating a panel: " + p.Topic)
	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Cannot start DB transaction: " + string(err.Error()))
		return false, err
	}

	panelInfo := `INSERT INTO Panels (
		Topic, Description, PanelRequestorEmail, CreatorId)
		VALUES (?, ?, ?, ?)`
	q, err := t.Prepare(panelInfo)
	if err != nil {
		log.Println("ERROR: Cannot prepare DB query: " + string(err.Error()))
		return false, err
	}

	_, err = q.Exec(p.Topic, p.Description, p.PanelRequestorEmail, id)
	if err != nil {
		log.Println("ERROR: Cannot execute DB transaction: " + string(err.Error()))
		return false, err
	}

	t.Commit()

	log.Println("INFO: Panel entry created")

	return true, nil
}

func GetPanels() ([]PanelSQL, error) {
	log.Println("INFO: List of panel objects requested")
	rows, err := DB.Query("SELECT * FROM Panels")
	if err != nil {
		log.Println("ERROR: Could not run the DB query!" + string(err.Error()))
		return nil, err
	}

	log.Println("INFO: Building panel list")
	panels := make([]PanelSQL, 0)
	for rows.Next() {
		panel := PanelSQL{}
		err = rows.Scan(
			&panel.Id,
			&panel.Topic,
			&panel.Description,
			&panel.PanelRequestorEmail,
			&panel.Location,
			&panel.ScheduledTime,
			&panel.DurationInMinutes,
			&panel.CreatorId,
			&panel.CreationDateTime,
			&panel.ApprovalStatus,
			&panel.ApprovedById,
			&panel.ApprovalDateTime,
		)
		if err != nil {
			log.Println("ERROR: Cannot marshal the panel objects!" + string(err.Error()))
			return nil, err
		}
		panels = append(panels, panel)
	}

	log.Println("INFO: List of all panels retrieved")
	return panels, nil
}

func GetPanelById(id int) (Panel, error) {
	log.Println("INFO: Panel by Id requested: " + strconv.Itoa(id))
	rec, err := DB.Prepare("SELECT * FROM Panels WHERE Id = ?")
	if err != nil {
		log.Println("ERROR: Could not prepare the DB query!" + string(err.Error()))
		return Panel{}, err
	}

	panel := Panel{}
	err = rec.QueryRow(id).Scan(
		&panel.Id,
		&panel.Topic,
		&panel.Description,
		&panel.PanelRequestorEmail,
		&panel.Location,
		&panel.ScheduledTime,
		&panel.DurationInMinutes,
		&panel.CreatorId,
		&panel.CreationDateTime,
		&panel.ApprovalStatus,
		&panel.ApprovedById,
		&panel.ApprovalDateTime,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("ERROR: No such panel found in DB: " + string(err.Error()))
			return Panel{}, nil
		}
		log.Println("ERROR: Cannot retrieve panel from DB: " + string(err.Error()))
		return Panel{}, err
	}

	return panel, nil
}

func GetPanelLocationByPanelId(id int) (Location, error) {
	log.Println("INFO: Panel location by panel Id requested: " + strconv.Itoa(id))
	rec, err := DB.Prepare("SELECT Location FROM Panels WHERE Id = ?")
	if err != nil {
		return Location{}, err
	}

	location := Location{}
	err = rec.QueryRow(id).Scan(
		&location.Location,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("ERROR: No such panel found in DB: " + string(err.Error()))
			return Location{}, nil
		}
		log.Println("ERROR: Cannot retrieve panel location from DB: " + string(err.Error()))
		return Location{}, err
	}

	return location, nil
}

func GetPanelScheduleByPanelId(id int) (Schedule, error) {
	log.Println("INFO: Panel schedule by panel Id requested: " + strconv.Itoa(id))
	rec, err := DB.Prepare("SELECT ScheduledTime, DurationInMinutes FROM Panels WHERE Id = ?")
	if err != nil {
		return Schedule{}, err
	}

	schedule := Schedule{}
	err = rec.QueryRow(id).Scan(
		&schedule.StartTime,
		&schedule.DurationInMinutes,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("ERROR: No such panel found in DB: " + string(err.Error()))
			return Schedule{}, nil
		}
		log.Println("ERROR: Cannot retrieve panel schedule from DB: " + string(err.Error()))
		return Schedule{}, err
	}

	return schedule, nil
}

func SetPanelLocation(id int, j Location) (bool, error) {
	log.Println("INFO: Set user status for panel Id '" + strconv.Itoa(id) + "'")
	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Could not start DB transaction: " + string(err.Error()))
		return false, err
	}

	q, err := DB.Prepare("UPDATE Panels SET Location = ? WHERE Id = ?")
	if err != nil {
		log.Println("ERROR: Could not prepare DB query! " + string(err.Error()))
		return false, err
	}
	// ensure the UserStatus.Status value is either 'enabled' or 'locked'
	log.Println("INFO: panel Id to set location of: " + strconv.Itoa(id))
	log.Println("INFO: requested location to assign the panel to: " + j.Location)
	// TODO: Add a test for valid locations after we add the location table

	result, err := q.Exec(j.Location, id)
	if err != nil {
		log.Println("ERROR: Could not execute query for panel Id '" + strconv.Itoa(id) + "': " + string(err.Error()))
		return false, err
	}
	numberOfRows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	t.Commit()

	log.Println("INFO: SQL result: Rows: " + strconv.Itoa(int(numberOfRows)))
	return true, nil
}
