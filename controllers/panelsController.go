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
	"log"
	"net/http"
	"strconv"

	"github.com/JAFAX/giron-service/model"
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

	userObject, _ := g.GetUserId(c)
	s, err := model.CreatePanel(json, userObject.Id)
	if s {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Panel has been added to system"})
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// DeletePanelById Delete a panel by its Id
//
//	@Summary		Delete a panel by Id
//	@Description	Delete a panel by Id
//	@Tags			panels
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"Panel Id"
//	@Security		BasicAuth
//	@Success		200	{object}	model.SuccessMsg
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/panel/{id} [delete]
func (g *GironService) DeletePanelById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	_, authed := g.GetUserId(c)
	if authed {
		status, err := model.DeletePanelById(id)
		if err != nil {
			log.Println("ERROR: Cannot delete panel: " + string(err.Error()))
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Unable to remove Panel! " + string(err.Error())})
			return
		}

		if status {
			idString := strconv.Itoa(id)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Panel Id '" + idString + "' has been removed from system"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Unable to remove Panel!"})
		}
	} else {
		c.IndentedJSON(http.StatusForbidden, gin.H{"error": "Insufficient access. Access denied!"})
		return
	}
}

// GetPanels Retrieve list of all panels
//
//	@Summary		Retrieve list of all panels
//	@Description	Retrieve list of all panels
//	@Tags			panels
//	@Produce		json
//	@Security		BasicAuth
//	@Success		200	{object}	model.PanelList
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/panels/all [get]
func (g *GironService) GetPanels(c *gin.Context) {
	_, authed := g.GetUserId(c)
	if authed {
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
			panelEnt.DurationInMinutes = panel.DurationInMinutes
			panelEnt.AgeRestricted = panel.AgeRestricted
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
	} else {
		c.IndentedJSON(http.StatusForbidden, gin.H{"error": "Insufficient access. Access denied!"})
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
			panelEnt.DurationInMinutes = panel.DurationInMinutes
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

// GetPanelsByLocationId Retrieve list of all approved panels by location Id
//
//	@Summary		Retrieve list of all approved panels by location Id
//	@Description	Retrieve list of all approved panels by location Id
//	@Tags			panels
//	@Produce		json
//	@Success		200	{object}	model.PanelList
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/panels/ByLocationId/{id} [get]
func (g *GironService) GetPanelsByLocationId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": string(err.Error())})
		return
	}
	panels, err := model.GetPanelsByLocationId(id)
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
			panelEnt.DurationInMinutes = panel.DurationInMinutes
			panelEnt.AgeRestricted = panel.AgeRestricted
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
//	@Param			id	path	string	true	"Panel Id"
//	@Success		200	{object}	model.Panel
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/panel/{id} [get]
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
//	@Param			id	path	string	true	"Panel Id"
//	@Success		200	{object}	model.Location
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/panel/{id}/location [get]
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
//	@Param			id	path	string	true	"Panel Id"
//	@Success		200	{object}	model.Schedule
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/panel/{id}/schedule [get]
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
//	@Param			id	path	string	true	"Building Id"
//	@Security		BasicAuth
//	@Success		200	{object}	model.SuccessMsg
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/panel/{id}/location [post]
func (g *GironService) SetPanelLocation(c *gin.Context) {
	_, authed := g.GetUserId(c)
	if authed {
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
	} else {
		c.IndentedJSON(http.StatusForbidden, gin.H{"error": "Insufficient access. Access denied!"})
	}
}

// SetPanelScheduledTimeById Set the panel's scheduled time
//
//	@Summary		Set the scheduled time for a panel
//	@Description	Set the scheduled time for a panel
//	@Tags			panels
//	@Produce		json
//	@Param			id	path	string	true	"Panel Id"
//	@Param			json	body	model.PanelScheduledTime	true	"Scheduled Time"
//	@Security		BasicAuth
//	@Success		200	{object}	model.SuccessMsg
//	@Failure		400 {object}	model.FailureMsg
//	@Router			/panel/{id}/schedule [post]
func (g *GironService) SetPanelScheduledTimeById(c *gin.Context) {
	_, authed := g.GetUserId(c)
	if authed {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": string(err.Error())})
			return
		}

		var json model.PanelScheduledTime
		if err := c.ShouldBindJSON(&json); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": string(err.Error())})
			return
		}

		status, msg, err := model.SetPanelScheduledTimeById(id, json)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": string(err.Error())})
			return
		}

		if status {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Panel scheduled for " + msg})
		} else {
			c.IndentedJSON(http.StatusConflict, gin.H{"message": "Panel cannot be scheduled. Reason " + msg})
		}
	} else {
		c.IndentedJSON(http.StatusForbidden, gin.H{"error": "Insufficient access. Access denied!"})
	}
}

// SetPanelAgeRestrictionById Set the age restriction status of a panel
//
//	 @Summary		Set panel age restriction
//		@Description	Set panel age restriction
//		@Tags			panels
//		@Produce		json
//		@Param			id	path	string	true	"Panel Id"
//		@Param			json body	model.PanelAgeRestrictionState true	"Age restriction state"
//		@Security		BasicAuth
//		@Success		200	{object}	model.SuccessMsg
//		@Failure		400	{object}	model.FailureMsg
//		@Router			/panel/{id}/restricted [post]
func (g *GironService) SetPanelAgeRestrictionById(c *gin.Context) {
	_, authed := g.GetUserId(c)
	if authed {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": string(err.Error())})
			return
		}

		var json model.PanelAgeRestrictionState
		if err := c.ShouldBindJSON(&json); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": string(err.Error())})
			return
		}

		status, err := model.SetPanelAgeRestrictionById(id, json)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": string(err.Error())})
			return
		}

		if status {
			if json.RestrictionState {
				c.IndentedJSON(http.StatusOK, gin.H{"message": "Panel is age restricted"})
			} else {
				c.IndentedJSON(http.StatusOK, gin.H{"message": "Panel is not age restricted"})
			}
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Panel is not age restricted"})
		}
	} else {
		c.IndentedJSON(http.StatusForbidden, gin.H{"error": "Insufficient access. Access denied!"})
	}
}

// SetApprovalStatusPanelById Set the approval status of a panel
//
//	@Summary		Set panel location
//	@Description	Set panel location
//	@Tags			panels
//	@Produce		json
//	@Param			id	path	string	true	"Panel Id"
//	@Param			json body	model.PanelApproval true	"Approval data"
//	@Security		BasicAuth
//	@Success		200	{object}	model.SuccessMsg
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/panel/{id}/approve [post]
func (g *GironService) SetApprovalStatusPanelById(c *gin.Context) {
	userObject, authed := g.GetUserId(c)
	if authed {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": string(err.Error())})
			return
		}

		var json model.PanelApproval
		if err := c.ShouldBindJSON(&json); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": string(err.Error())})
			return
		}

		status, err := model.SetApprovalStatusPanelById(id, json, userObject.Id)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": string(err.Error())})
			return
		}

		if status {
			if json.State {
				c.IndentedJSON(http.StatusOK, gin.H{"message": "Panel approved"})
			} else {
				c.IndentedJSON(http.StatusOK, gin.H{"message": "Panel unapproved"})
			}
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Panel not approved"})
		}
	} else {
		c.IndentedJSON(http.StatusForbidden, gin.H{"error": "Insufficient access. Access denied!"})
	}
}
