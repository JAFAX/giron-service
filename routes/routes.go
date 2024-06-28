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
	g.GET("/health", i.GetHealth) // service health
}

func PrivateRoutes(g *gin.RouterGroup, i *controllers.GironService) {
	// building related routes
	g.GET("/buildings", i.GetBuildings)             // get all buildings
	g.GET("/building/:id", i.GetBuildingById)       // get building by Id
	g.POST("/building", i.CreateBuilding)           // create a new building
	g.PATCH("/building/:id", i.UpdateBuildingById)  // update a building
	g.DELETE("/building/:id", i.DeleteBuildingById) // delete a building
	// floor related routes
	g.GET("/floors", i.GetAllFloors)                         // get all floor records
	g.GET("/floors/buildingId/:id", i.GetFloorsByBuildingId) // get all floors in a building
	g.GET("/floor/:id", i.GetFloorById)                      // get floor by Id
	g.POST("/floor", i.CreateFloor)                          // create a new floor in a building
	g.PATCH("/floor/:id", i.UpdateFloorById)                 // update a floor by its Id
	g.DELETE("/floor/:id", i.DeleteFloorById)                // delete a floor by its Id
	// location related routes
	g.GET("/locations", i.GetAllLocations)                           // get all locations at the event
	g.GET("/locations/byFloorId/:id", i.GetLocationsByFloorId)       // get locations by the floor id
	g.GET("/locations/byBuildingId/:id", i.GetLocationsByBuildingId) // get locations by building id
	g.GET("/location/:id", i.GetLocationById)                        // get location by id
	g.POST("/location", i.CreateLocation)                            // create locations in the building
	g.PATCH("/location/:id", i.UpdateLocationById)                   // update locations in the building by id
	g.DELETE("/location/:id", i.DeleteLocationById)                  // delete a location by id
	// panel related routes
	g.GET("/panels", i.GetApprovedPanels)                         // get all approved panels
	g.GET("/panels/ByLocationId/:id", i.GetPanelsByLocationId)    // get all approved panels by location ID
	g.GET("/panel/:id", i.GetPanelById)                           // get panel details
	g.GET("/panel/:id/location", i.GetPanelLocationByPanelId)     // get the location of a panel
	g.GET("/panel/:id/schedule", i.GetPanelScheduleByPanelId)     // get the time and date of a panel
	g.GET("/panel/:id/tags")                                      // get a list of tags associated with a panel
	g.GET("/panels/all", i.GetPanels)                             // get all panels
	g.POST("/panel", i.CreatePanel)                               // create a new panel event
	g.POST("/panel/:id/location", i.SetPanelLocation)             // set/update the location of a panel
	g.POST("/panel/:id/schedule", i.SetPanelScheduledTimeById)    // set/update the time and date of a panel
	g.POST("/panel/:id/approve", i.SetApprovalStatusPanelById)    // approve a panel
	g.POST("/panel/:id/restricted", i.SetPanelAgeRestrictionById) // set whether the panel is age restricted
	g.POST("/panel/:id/assignTag")                                // assign a tag to a panel
	g.PATCH("/panel/:id/unassignTag")                             // unassign a tag to a panel
	g.DELETE("/panel/:id", i.DeletePanelById)                     // delete a panel
	// screening related routes
	g.GET("/screenings")                  // get all screenings
	g.GET("/screenings/ByLocationId/:id") // get all screenings by location ID
	g.GET("/screening/:id")               // get screening details
	g.GET("/screening/:id/location")      // get the location of a screening
	g.GET("/screening/:id/schedule")      // get the time and date of a screening
	g.GET("/screening/:id/tags")          // get a list of tags associated with a screening
	g.POST("/screening")                  // create a new screening event
	g.POST("/screening/:id/location")     // set/update the location of a screening
	g.POST("/screening/:id/schedule")     // set/update the time and date of a screening
	g.POST("/screening/:id/restricted")   // set/update whether a screening is age restricted
	g.POST("/screening/:id/assignTag")    // assign a tag to a screening
	g.PATCH("/screening/:id/unassignTag") // unassign a tag to a screening
	g.DELETE("/screening/:id")            // delete a screening
	// tag related routes
	g.GET("/tags")       // get all tags
	g.GET("/tag/:id")    // get tag details
	g.POST("/tag")       // create a new tag
	g.PATCH("/tag/:id")  // update a new tag
	g.DELETE("/tag/:id") // delete a tag
	// user related routes
	g.GET("/user/id/:id", i.GetUserById)            // get user by id
	g.GET("/user/name/:name", i.GetUserByUserName)  // get user by username
	g.GET("/user/:name/status", i.GetUserStatus)    // get whether a user is locked or not
	g.GET("/users", i.GetUsers)                     // get users
	g.POST("/user", i.CreateUser)                   // create new user
	g.PATCH("/user/:name", i.ChangeAccountPassword) // update a user password
	g.PATCH("/user/:name/status", i.SetUserStatus)  // lock a user
	g.DELETE("/user/:name", i.DeleteUser)           // trash a user
}
