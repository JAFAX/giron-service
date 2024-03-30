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

	"github.com/JAFAX/giron-service/globals"
	"github.com/JAFAX/giron-service/helpers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (g *GironService) LoginUI(c *gin.Context) {
	log.Println("INFO: Displaying the login UI")
	session := sessions.Default(c)
	user := session.Get(globals.UserKey)
	if user != nil {
		log.Println("WARN: User session is empty")
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"content": "Please login",
			"user":    user,
		})
		return
	}
	log.Println("INFO: User session is known")
	c.HTML(http.StatusOK, "login.html", gin.H{
		"content": "",
		"user":    user,
	})
}

func (g *GironService) LoginUIPost(c *gin.Context) {
	log.Println("INFO: Requesting Login action POST")
	session := sessions.Default(c)
	user := session.Get(globals.UserKey)
	if user != nil {
		log.Println("INFO: No user session. Might be first time logging in")
	}

	username := c.PostForm("username")
	password := c.PostForm("password")

	if helpers.EmptyUserPass(username, password) {
		log.Println("ERROR: username or password is empty! Please login")
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "Parameters can't be empty"})
		return
	}

	if !helpers.CheckUserPass(username, password) {
		log.Println("ERROR: Invalid username or password! Please login")
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"content": "Incorrect username or password"})
		return
	}

	session.Set(globals.UserKey, username)
	if err := session.Save(); err != nil {
		log.Println("ERROR: Cannot save session: error: " + string(err.Error()))
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{"content": "Failed to save session"})
		return
	}

	log.Println("INFO: Login succeeded. Redirecting to /admin")
	c.Redirect(http.StatusMovedPermanently, "/admin")
}
