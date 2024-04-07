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

// CreateFloor Add a floor to a building
//
//	@Summary		Create a new floor
//	@Description	Create a new floor
//	@Tags			floors
//	@Accept			json
//	@Produce		json
//	@Param			floor	body	model.ProposedFloor	true	"Floor data"
//	@Security		BasicAuth
//	@Success		200	{object}	model.SuccessMsg
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/floor [post]
func (g *GironService) CreateFloor(c *gin.Context) {
	var json model.ProposedFloor
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

	s, err := model.CreateFloor(json, userObject.Id)
	if s {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Floor has been added to system"})
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// GetAllFloors Retrieve list of all floor records
//
//	@Summary		Retrieve list of all floor records
//	@Description	Retrieve list of all floor records
//	@Tags			floors
//	@Produce		json
//	@Success		200	{object}	model.FloorList
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/floors [get]
func (g *GironService) GetAllFloors(c *gin.Context) {
	floors, err := model.GetAllFloors()
	if err != nil {
		log.Println("ERROR: Cannot retrieve list of floor records: " + string(err.Error()))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if floors == nil {
		log.Println("WARN: No floors returned")
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "no records found!"})
	} else {
		log.Println("INFO: Returned floor list")
		c.IndentedJSON(http.StatusOK, gin.H{"data": floors})
	}
}

// GetFloorsByBuildingId Retrieve list of all floors based on building Id
//
//	@Summary		Retrieve list of all floors based on building Id
//	@Description	Retrieve list of all floors based on building Id
//	@Tags			floors
//	@Produce		json
//	@Param			id	path	string	true	"Building Id"
//	@Success		200	{object}	model.FloorList
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/floors/buildingId/{id} [get]
func (g *GironService) GetFloorsByBuildingId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	floors, err := model.GetFloorsByBuildingId(id)
	if err != nil {
		log.Println("ERROR: Cannot retrieve list of floor records: " + string(err.Error()))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if floors == nil {
		log.Println("WARN: No floors returned")
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "no records found!"})
	} else {
		log.Println("INFO: Returned floor list")
		c.IndentedJSON(http.StatusOK, gin.H{"data": floors})
	}
}

// GetFloorById Retrieve a floor based on Id
//
//	@Summary		Retrieve floor based on Id
//	@Description	Retrieve floor based on Id
//	@Tags			floors
//	@Produce		json
//	@Param			id	path	string	true	"Floor Id"
//	@Success		200	{object}	model.BuildingFloor
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/floor/{id} [get]
func (g *GironService) GetFloorById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	floor, err := model.GetFloorById(id)
	if err != nil {
		log.Println("ERROR: Cannot retrieve list of floor records: " + string(err.Error()))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	log.Println("INFO: Returned floor list")
	c.IndentedJSON(http.StatusOK, gin.H{"data": floor})
}
