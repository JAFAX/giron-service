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

import "log"

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
