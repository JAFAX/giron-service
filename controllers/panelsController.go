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

// CreateDomain Add a panel event
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
func (i *GironService) CreatePanel(c *gin.Context) {
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
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Domain has been added to system"})
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
