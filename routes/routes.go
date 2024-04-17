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

func FePublicRoutes(g *gin.RouterGroup, i *controllers.GironService) {
	// login page
	g.GET("/", i.LoginUI)           // login UI
	g.POST("/login", i.LoginUIPost) // the actual action of logging a person in
	g.GET("/logout")                // log out UI
}

func FePrivateRoutes(g *gin.RouterGroup, i *controllers.GironService) {
	// admin panel
	g.GET("/admin")
}

func PublicRoutes(g *gin.RouterGroup, i *controllers.GironService) {
	// building related routes
	g.GET("/buildings", i.GetBuildings)       // get all buildings
	g.GET("/building/:id", i.GetBuildingById) // get building by Id
	// floor related routes
	g.GET("/floors", i.GetAllFloors)                         // get all floor records
	g.GET("/floors/buildingId/:id", i.GetFloorsByBuildingId) // get all floors in a building
	g.GET("/floor/:id", i.GetFloorById)                      // get floor by Id
	// location related routes
	g.GET("/locations", i.GetAllLocations)                           // get all locations at the event
	g.GET("/locations/byFloorId/:id", i.GetLocationsByFloorId)       // get locations by the floor id
	g.GET("/locations/byBuildingId/:id", i.GetLocationsByBuildingId) // get locations by building id
	g.GET("/location/:id", i.GetLocationById)                        // get location by id
	// panel related routes
	g.GET("/panels", i.GetApprovedPanels)                     // get all approved panels
	g.GET("/panel/:id", i.GetPanelById)                       // get approved panel details
	g.GET("/panel/:id/location", i.GetPanelLocationByPanelId) // get the location of a panel
	g.GET("/panel/:id/schedule", i.GetPanelScheduleByPanelId) // get the time and date of a panel
	// user related routes
	g.GET("/user/id/:id", i.GetUserById)           // get user by id
	g.GET("/user/name/:name", i.GetUserByUserName) // get user by username
	g.GET("/users", i.GetUsers)                    // get users
	// service related routes
	g.OPTIONS("/")   // API options
	g.GET("/health") // service health
}

func PrivateRoutes(g *gin.RouterGroup, i *controllers.GironService) {
	// building related routes
	g.POST("/building", i.CreateBuilding)           // create a new building
	g.PATCH("/building/:id", i.UpdateBuildingById)  // update a building
	g.DELETE("/building/:id", i.DeleteBuildingById) // delete a building
	// floor related routes
	g.POST("/floor", i.CreateFloor)           // create a new floor in a building
	g.PATCH("/floor/:id", i.UpdateFloorById)  // update a floor by its Id
	g.DELETE("/floor/:id", i.DeleteFloorById) // delete a floor by its Id
	// location related routes
	g.POST("/location", i.CreateLocation)           // create locations in the building
	g.PATCH("/location/:id")                        // update locations in the building by id
	g.DELETE("/location/:id", i.DeleteLocationById) // delete a location by id
	// panel related routes
	g.GET("/panels/all", i.GetPanels)                 // get all panels
	g.POST("/panel", i.CreatePanel)                   // create a new panel event
	g.POST("/panel/:id/location", i.SetPanelLocation) // set the location of a panel
	g.POST("/panel/:id/schedule")                     // set the time and date of a panel
	g.POST("/panel/:id/approve")                      // approve a panel
	g.DELETE("/panel/:id", i.DeletePanelById)         // delete a panel
	// user related routes
	g.POST("/user", i.CreateUser)                   // create new user
	g.PATCH("/user/:name", i.ChangeAccountPassword) // update a user password
	g.PATCH("/user/:name/status", i.SetUserStatus)  // lock a user
	g.GET("/user/:name/status", i.GetUserStatus)    // get whether a user is locked or not
	g.DELETE("/user/:name", i.DeleteUser)           // trash a user
}
