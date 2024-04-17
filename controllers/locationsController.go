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

// CreateLocation Add a location to a building
//
//	@Summary		Create a new location
//	@Description	Create a new location
//	@Tags			locations
//	@Accept			json
//	@Produce		json
//	@Param			location	body	model.ProposedLocation	true	"Location data"
//	@Security		BasicAuth
//	@Success		200	{object}	model.SuccessMsg
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/location [post]
func (g *GironService) CreateLocation(c *gin.Context) {
	var json model.ProposedLocation
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

	s, err := model.CreateLocation(json, userObject.Id)
	if s {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Location has been added to system"})
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// DeleteLocationById Delete a location by its Id
//
//	@Summary		Delete a location by Id
//	@Description	Delete a location by Id
//	@Tags			locations
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"Location Id"
//	@Security		BasicAuth
//	@Success		200	{object}	model.SuccessMsg
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/location/{id} [delete]
func (g *GironService) DeleteLocationById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	status, err := model.DeleteLocationById(id)
	if err != nil {
		log.Println("ERROR: Cannot delete location: " + string(err.Error()))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Unable to remove Location! " + string(err.Error())})
		return
	}

	if status {
		idString := strconv.Itoa(id)
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Location Id '" + idString + "' has been removed from system"})
	} else {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Unable to remove Location!"})
	}
}

// GetAllLocations Retrieve list of all location objects
//
//	@Summary		Retrieve list of all location objects
//	@Description	Retrieve list of all location objects
//	@Tags			locations
//	@Produce		json
//	@Success		200	{object}	model.LocationList
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/locations [get]
func (g *GironService) GetAllLocations(c *gin.Context) {
	locations, err := model.GetAllLocations()
	if err != nil {
		log.Println("ERROR: Cannot retrieve list of locations: " + string(err.Error()))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if locations == nil {
		log.Println("WARN: No locations returned")
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "no records found!"})
	} else {
		log.Println("INFO: Returned list of locations")
		c.IndentedJSON(http.StatusOK, gin.H{"data": locations})
	}
}

// GetLocationById Retrieve location by Id
//
//	@Summary		Retrieve location by Id
//	@Description	Retrieve location by Id
//	@Tags			locations
//	@Produce		json
//	@Param			id	path	string	true	"Location Id"
//	@Success		200	{object}	model.Location
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/location/{id} [get]
func (g *GironService) GetLocationById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	ent, err := model.GetLocationById(id)
	if err != nil {
		log.Println("ERROR: Cannot retrieve location by Id '" + strconv.Itoa(id) + "': " + string(err.Error()))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if ent.Location == "" {
		log.Println("WARN: No Location returned")
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "no records found!"})
	} else {
		log.Println("INFO: Returned location object for Id '" + strconv.Itoa(id) + "'")
		c.IndentedJSON(http.StatusOK, ent)
	}
}

// GetLocationsByFloorId Retrieve list of locations by floor Id
//
//	@Summary		Retrieve list of locations by floor Id
//	@Description	Retrieve list of locations by floor Id
//	@Tags			locations
//	@Produce		json
//	@Param			id	path	string	true	"Floor Id"
//	@Success		200	{object}	model.LocationList
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/location/byFloorId/{id} [get]
func (g *GironService) GetLocationsByFloorId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	ent, err := model.GetLocationsByFloorId(id)
	if err != nil {
		log.Println("ERROR: Cannot retrieve locations by floor Id '" + strconv.Itoa(id) + "': " + string(err.Error()))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	log.Println("INFO: Returned location object for floor Id '" + strconv.Itoa(id) + "'")
	c.IndentedJSON(http.StatusOK, gin.H{"data": ent})
}

// GetLocationsByBuildingId Retrieve list of locations by building Id
//
//	@Summary		Retrieve list of locations by building Id
//	@Description	Retrieve list of locations by building Id
//	@Tags			locations
//	@Produce		json
//	@Param			id	path	string	true	"Building Id"
//	@Success		200	{object}	model.LocationList
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/location/byBuildingId/{id} [get]
func (g *GironService) GetLocationsByBuildingId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	ent, err := model.GetLocationsByBuildingId(id)
	if err != nil {
		log.Println("ERROR: Cannot retrieve locations by building Id '" + strconv.Itoa(id) + "': " + string(err.Error()))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	log.Println("INFO: Returned location object for building Id '" + strconv.Itoa(id) + "'")
	c.IndentedJSON(http.StatusOK, gin.H{"data": ent})
}
