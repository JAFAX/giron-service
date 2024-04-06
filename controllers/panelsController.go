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
//	@Router			/panels/all [get]
func (g *GironService) GetPanels(c *gin.Context) {
	panels, err := model.GetPanels()
	if err != nil {
		log.Println("ERROR: Cannot retrieve list of panels: " + string(err.Error()))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

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
		log.Println("WARN: No panels returned")
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "no records found!"})
	} else {
		log.Println("INFO: Returned approved list of panels")
		c.IndentedJSON(http.StatusOK, gin.H{"data": panelSlice})
	}
}

// GetApprovedPanels Retrieve list of all approved panels
//
//	@Summary		Retrieve list of all approved panels
//	@Description	Retrieve list of all approved panels
//	@Tags			panels
//	@Produce		json
//	@Success		200	{object}	model.PanelList
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/panels [get]
func (g *GironService) GetApprovedPanels(c *gin.Context) {
	panels, err := model.GetPanels()
	if err != nil {
		log.Println("ERROR: Cannot retrieve list of panels: " + string(err.Error()))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	panelSlice := make([]model.Panel, 0)
	for _, panel := range panels {
		if panel.ApprovalStatus {
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
	}

	if panels == nil {
		log.Println("WARN: No panels returned")
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "no records found!"})
	} else {
		log.Println("INFO: Returned approved list of panels")
		c.IndentedJSON(http.StatusOK, gin.H{"data": panelSlice})
	}
}

// GetPanelById Retrieve panel by Id
//
//	@Summary		Retrieve panel by Id
//	@Description	Retrieve panel by Id
//	@Tags			panels
//	@Produce		json
//	@Success		200	{object}	model.Panel
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/panel/{Id} [get]
func (g *GironService) GetPanelById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	ent, err := model.GetPanelById(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": string(err.Error())})
		return
	}

	if ent.Topic == "" {
		strId := strconv.Itoa(id)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "no records found with panel id " + strId})
	} else {
		c.IndentedJSON(http.StatusOK, ent)
	}
}

// GetPanelLocationById Retrieve panel location by the panel Id
//
//	@Summary		Retrieve panel location by the panel Id
//	@Description	Retrieve panel location by the panel Id
//	@Tags			panels
//	@Produce		json
//	@Success		200	{object}	model.Location
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/panel/{Id}/location [get]
func (g *GironService) GetPanelLocationByPanelId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	ent, err := model.GetPanelLocationByPanelId(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": string(err.Error())})
		return
	}

	c.IndentedJSON(http.StatusOK, ent)
}

// GetPanelScheduleById Retrieve panel schedule by the panel Id
//
//	@Summary		Retrieve panel schedule by the panel Id
//	@Description	Retrieve panel schedule by the panel Id
//	@Tags			panels
//	@Produce		json
//	@Success		200	{object}	model.Schedule
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/panel/{Id}/schedule [get]
func (g *GironService) GetPanelScheduleByPanelId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	ent, err := model.GetPanelScheduleByPanelId(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": string(err.Error())})
		return
	}

	c.IndentedJSON(http.StatusOK, ent)
}

// SetPanelLocation Set panel location
//
//	@Summary		Set panel location
//	@Description	Set panel location
//	@Tags			panels
//	@Produce		json
//	@Success		200	{object}	model.SuccessMsg
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/panel/{Id}/location [post]
func (g *GironService) SetPanelLocation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": string(err.Error())})
		return
	}
	var json model.Location
	if err := c.ShouldBindJSON(&json); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// we don't need the status, since the error speaks for itself
	_, err = model.SetPanelLocation(id, json)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": string(err.Error())})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Panel location updated"})
}
