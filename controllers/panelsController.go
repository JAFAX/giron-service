package controllers

/*

  Copyright 2024, YggdrasilSoft, LLC.

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
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/JAFAX/giron-service/helpers"
	"github.com/JAFAX/giron-service/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CreatePanel Add a panel event
//
//	@Summary		Create a new panel event
//	@Description	Create a new panel event
//	@Tags			panels
//	@Accept			json
//	@Produce		json
//	@Param			panel	body	model.Panel	true	"Panel data"
//	@Security		BasicAuth
//	@Success		200	{object}	model.SuccessMsg
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/panel [post]
func (g *GironService) CreatePanel(c *gin.Context) {
	var json model.ProposedPanel
	if err := c.ShouldBindJSON(&json); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// need to get our current user context to get the CreatorId
	session := sessions.Default(c)
	user := session.Get("user")
	// if nil, we have an issue
	if user == nil {
		c.IndentedJSON(http.StatusForbidden, gin.H{"error": "Insufficient access. Access denied!"})
		return
	}

	// convert user interface to a string
	username := fmt.Sprintf("%v", user)
	// lets output our session user
	log.Println("INFO: Session user: " + username)
	// get our user id
	userObject, err := model.GetUserByUserName(username)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// what is our user Id
	log.Println("INFO: Session user's ID: " + strconv.Itoa(userObject.Id))

	s, err := model.CreatePanel(json, userObject.Id)
	if s {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Panel has been added to system"})
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// GetPanels Retrieve list of all panels
//
//	@Summary		Retrieve list of all panels
//	@Description	Retrieve list of all panels
//	@Tags			panels
//	@Produce		json
//	@Success		200	{object}	model.PanelList
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/panels [get]
func (g *GironService) GetPanels(c *gin.Context) {
	panels, err := model.GetPanels()
	helpers.CheckError(err)

	panelSlice := make([]model.Panel, 0)
	for _, panel := range panels {
		panelEnt := model.Panel{}
		panelEnt.Id = panel.Id
		panelEnt.Topic = panel.Topic
		panelEnt.Description = panel.Description
		panelEnt.PanelRequestorEmail = panel.PanelRequestorEmail
		panelEnt.Location = panel.Location
		if panel.ScheduledTime.Valid {
			panelEnt.ScheduledTime = panel.ScheduledTime.String
		} else {
			panelEnt.ScheduledTime = ""
		}
		panelEnt.CreatorId = panel.CreatorId
		panelEnt.CreationDateTime = panel.CreationDateTime
		panelEnt.ApprovalStatus = panel.ApprovalStatus
		if panel.ApprovedById.Valid {
			panelEnt.ApprovedById = int(panel.ApprovedById.Int64)
		} else {
			panelEnt.ApprovedById = 0
		}
		if panel.ApprovalDateTime.Valid {
			panelEnt.ApprovalDateTime = panel.ApprovalDateTime.String
		} else {
			panelEnt.ApprovalDateTime = ""
		}

		panelSlice = append(panelSlice, panelEnt)
	}

	if panels == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "no records found!"})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"data": panelSlice})
	}

}
