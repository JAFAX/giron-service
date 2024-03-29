package routes

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
	"github.com/gin-gonic/gin"

	"github.com/JAFAX/giron-service/controllers"
)

func PublicRoutes(g *gin.RouterGroup, i *controllers.GironService) {
	// panel related routes
	g.GET("/panels")    // get all approved panels
	g.GET("/panel/:id") // get approved panel details
	// user related routes
	g.GET("/user/id/:id", i.GetUserById)
	g.GET("/user/name/:name", i.GetUserByUserName)
	g.GET("/users", i.GetUsers)
	// service related routes
	g.OPTIONS("/")   // API options
	g.GET("/health") // service health
}

func PrivateRoutes(g *gin.RouterGroup, i *controllers.GironService) {
	// panel related routes
	g.GET("/panels/all")            // get all panels
	g.POST("/panel", i.CreatePanel) // create a new panel event
	g.POST("/panel/location")       // set the location of a panel
	g.POST("/panel/schedule")       // set the time and date of a panel
	g.POST("/panel/approve")        // approve a panel
	g.DELETE("/panel/:id")          // delete a panel
	// user related routes
	g.POST("/user", i.CreateUser)                   // create new user
	g.PATCH("/user/:name", i.ChangeAccountPassword) // update a user password
	g.PATCH("/user/:name/status", i.SetUserStatus)  // lock a user
	g.GET("/user/:name/status", i.GetUserStatus)    // get whether a user is locked or not
	g.DELETE("/user/:name", i.DeleteUser)           // trash a user
}
