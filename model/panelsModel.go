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
	"time"
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

func DeletePanelById(id int) (bool, error) {
	log.Println("INFO: Panel deletion requested: " + strconv.Itoa(id))
	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Could not start DB transaction!" + string(err.Error()))
		return false, err
	}

	q, err := DB.Prepare("DELETE FROM Panels WHERE Id IS ?")
	if err != nil {
		log.Println("ERROR: Could not prepare the DB query!" + string(err.Error()))
		return false, err
	}

	_, err = q.Exec(id)
	if err != nil {
		log.Println("ERROR: Cannot delete panel with id '" + strconv.Itoa(id) + "': " + string(err.Error()))
		return false, err
	}

	t.Commit()

	log.Println("INFO: Panel with id '" + strconv.Itoa(id) + "' has been deleted")
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
			&panel.AgeRestricted,
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

func GetPanelsByLocationId(id int) ([]PanelSQL, error) {
	log.Println("INFO: Panels by location Id requested: Location Id: " + strconv.Itoa(id))
	// first, take the location Id and get back the location name
	rec, err := DB.Prepare("SELECT RoomName FROM Locations WHERE Id = ?")
	if err != nil {
		return nil, err
	}

	var locationName string
	err = rec.QueryRow(id).Scan(
		&locationName,
	)
	if err != nil {
		return nil, err
	}

	// now that we have the location name, do the real query for the panels in that location
	rows, err := DB.Query("SELECT * FROM Panels WHERE Location = ?", locationName)
	if err != nil {
		return nil, err
	}

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
			&panel.AgeRestricted,
			&panel.CreatorId,
			&panel.CreationDateTime,
			&panel.ApprovalStatus,
			&panel.ApprovedById,
			&panel.ApprovalDateTime,
		)
		if err != nil {
			return nil, err
		}

		panels = append(panels, panel)
	}

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
		&panel.AgeRestricted,
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
	log.Println("INFO: Set location for panel Id '" + strconv.Itoa(id) + "'")
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

func SetPanelScheduledTimeById(id int, json PanelScheduledTime) (bool, string, error) {
	log.Println("INFO: Set scheduled time for panel Id '" + strconv.Itoa(id) + "'")

	timeFormat := "2006-01-02 15:04:05"
	panelParsedStartTime, err := time.Parse(timeFormat, json.ScheduledTime)
	panelStartTimeInUnixTime := int(panelParsedStartTime.Unix())
	if err != nil {
		return false, "Could not convert from " + json.ScheduledTime + " to UNIX time", err
	}

	// first, get all panels by location Id
	panels, err := GetPanelsByLocationId(json.LocationId)
	if err != nil {
		return false, "Could not get panels by location Id '" + strconv.Itoa(json.LocationId) + "'", err
	}

	var startTime string
	for _, panel := range panels {
		// now get the start time converted to UNIX time to check against
		if panel.ScheduledTime.Valid || panel.DurationInMinutes != 0 {
			startTime = panel.ScheduledTime.String
			if startTime != "" {
				parsedTime, _ := time.Parse(timeFormat, startTime)
				startTimeUnixTime := int(parsedTime.Unix())
				endTimeUnixTime := startTimeUnixTime + (panel.DurationInMinutes * 60)
				if startTimeUnixTime == panelStartTimeInUnixTime ||
					(panelStartTimeInUnixTime > startTimeUnixTime && panelStartTimeInUnixTime < endTimeUnixTime) {
					return false, "Panel start time conflicts with existing panel in location", new(SchedulingConflict)
				}
			}
		}
	}

	// assume panel time is nil or an empty string, so we should be able to assign the panel
	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Could not start DB transaction: " + string(err.Error()))
		return false, json.ScheduledTime, err
	}
	q, err := DB.Prepare("UPDATE Panels SET ScheduledTime = ? WHERE Id = ?")
	if err != nil {
		log.Println("ERROR: Could not prepare DB query! " + string(err.Error()))
		return false, json.ScheduledTime, err
	}
	result, err := q.Exec(json.ScheduledTime, id)
	if err != nil {
		log.Println("ERROR: Could not execute query for panel Id '" + strconv.Itoa(id) + "': " + string(err.Error()))
		return false, json.ScheduledTime, err
	}
	numberOfRows, err := result.RowsAffected()
	if err != nil {
		return false, json.ScheduledTime, err
	}

	log.Println("INFO: SQL result: Rows: " + strconv.Itoa(int(numberOfRows)))
	t.Commit()

	return true, json.ScheduledTime, nil
}

func SetApprovalStatusPanelById(id int, status PanelApproval, userId int) (bool, error) {
	log.Println("INFO: Set Approval status for panel Id '" + strconv.Itoa(id) + "'")
	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Could not start DB transaction: " + string(err.Error()))
		return false, err
	}
	q, err := DB.Prepare("UPDATE Panels SET ApprovalStatus = ?, ApprovedById = ?, ApprovalDateTime = CURRENT_TIMESTAMP WHERE Id = ?")
	if err != nil {
		log.Println("ERROR: Could not prepare DB query! " + string(err.Error()))
		return false, err
	}
	result, err := q.Exec(status.State, userId, id)
	if err != nil {
		log.Println("ERROR: Could not execute query for panel Id '" + strconv.Itoa(id) + "': " + string(err.Error()))
		return false, err
	}
	numberOfRows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	log.Println("INFO: SQL result: Rows: " + strconv.Itoa(int(numberOfRows)))
	t.Commit()

	return true, nil
}

func SetPanelAgeRestrictionById(id int, status PanelAgeRestrictionState) (bool, error) {
	log.Println("INFO: Set age restriction status for panel Id '" + strconv.Itoa(id) + "'")
	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Could not start DB transaction: " + string(err.Error()))
		return false, err
	}
	q, err := DB.Prepare("UPDATE Panels SET AgeRestricted = ? WHERE Id = ?")
	if err != nil {
		log.Println("ERROR: Could not prepare DB query! " + string(err.Error()))
		return false, err
	}
	result, err := q.Exec(status.RestrictionState, id)
	if err != nil {
		log.Println("ERROR: Could not execute query for panel Id '" + strconv.Itoa(id) + "': " + string(err.Error()))
		return false, err
	}
	numberOfRows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	log.Println("INFO: SQL result: Rows: " + strconv.Itoa(int(numberOfRows)))
	t.Commit()

	return true, nil
}
