package controllers

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
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/JAFAX/giron-service/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CreateBuilding Add a building
//
//	@Summary		Create a new building
//	@Description	Create a new building
//	@Tags			buildings
//	@Accept			json
//	@Produce		json
//	@Param			building	body	model.ProposedBuilding	true	"Building data"
//	@Security		BasicAuth
//	@Success		200	{object}	model.SuccessMsg
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/building [post]
func (g *GironService) CreateBuilding(c *gin.Context) {
	var json model.ProposedBuilding
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

	s, err := model.CreateBuilding(json, userObject.Id)
	if s {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Building has been added to system"})
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// GetBuildings Retrieve list of all buildings
//
//	@Summary		Retrieve list of all panels
//	@Description	Retrieve list of all panels
//	@Tags			buildings
//	@Produce		json
//	@Success		200	{object}	model.BuildingList
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/buildings [get]
func (g *GironService) GetBuildings(c *gin.Context) {
	buildings, err := model.GetBuildings()
	if err != nil {
		log.Println("ERROR: Cannot retrieve list of buildings: " + string(err.Error()))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if buildings == nil {
		log.Println("WARN: No buildings returned")
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "no records found!"})
	} else {
		log.Println("INFO: Returned list of buildings")
		c.IndentedJSON(http.StatusOK, gin.H{"data": buildings})
	}
}

// GetBuildingById Retrieve building by Id
//
//	@Summary		Retrieve building by Id
//	@Description	Retrieve building by Id
//	@Tags			buildings
//	@Produce		json
//	@Param			id	path	string	true	"Building Id"
//	@Success		200	{object}	model.Building
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/building/{id} [get]
func (g *GironService) GetBuildingById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	ent, err := model.GetBuildingById(id)
	if err != nil {
		log.Println("ERROR: Cannot retrieve building by Id '" + strconv.Itoa(id) + "': " + string(err.Error()))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if ent.Name == "" {
		log.Println("WARN: No building returned")
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "no records found!"})
	} else {
		log.Println("INFO: Returned approved list of panels")
		c.IndentedJSON(http.StatusOK, ent)
	}
}

// UpdateBuildingById Update building by Id
//
//	@Summary		Update building information
//	@Description	Update building information
//	@Tags			buildings
//	@Produce		json
//	@Param			id	path	string	true	"Building Id"
//	@Param			building	body	model.BuildingUpdate	true	"Building data"
//	@Security		BasicAuth
//	@Success		200	{object}	model.SuccessMsg
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/building/{id} [patch]
func (g *GironService) UpdateBuildingById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": string(err.Error())})
		return
	}
	var json model.BuildingUpdate
	if err := c.ShouldBindJSON(&json); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// we don't need the status, since the error speaks for itself
	_, err = model.UpdateBuildingById(id, json)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": string(err.Error())})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Panel location updated"})
}

// DeleteBuildingById Delete a building by its Id
//
//	@Summary		Delete a building by Id
//	@Description	Delete a building by Id
//	@Tags			buildings
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"Building Id"
//	@Security		BasicAuth
//	@Success		200	{object}	model.SuccessMsg
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/building/{id} [delete]
func (g *GironService) DeleteBuildingById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	status, err := model.DeleteBuildingById(id)
	if err != nil {
		log.Println("ERROR: Cannot delete building: " + string(err.Error()))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Unable to remove building! " + string(err.Error())})
		return
	}

	if status {
		idString := strconv.Itoa(id)
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Building Id '" + idString + "' has been removed from system"})
	} else {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Unable to remove building!"})
	}
}
