package main

/*

  Giron-Service - Golang-based web service for managing panel events

  Author:  Gary L. Greene, Jr.
  License: Apache v2.0

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
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/JAFAX/giron-service/controllers"
	_ "github.com/JAFAX/giron-service/docs"
	"github.com/JAFAX/giron-service/globals"
	"github.com/JAFAX/giron-service/helpers"
	"github.com/JAFAX/giron-service/middleware"
	"github.com/JAFAX/giron-service/model"
	"github.com/JAFAX/giron-service/routes"
)

//	@title			Giron-Service
//	@version		0.0.11
//	@description	An API for managing panel events

//	@contact.name	Gary Greene
//	@contact.url	https://github.com/JAFAX/giron-service

//	@securityDefinitions.basic	BasicAuth

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:5000
//	@BasePath	/api/v1

// @schemas	http https
func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	// lets get our working directory
	appdir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	helpers.FatalCheckError(err)

	// config path is derived from app working directory
	configDir := filepath.Join(appdir, "config")

	// now that we have our appdir and configDir, lets read in our app config
	// and marshall it to the Config struct
	config := globals.Config{}
	jsonContent, err := os.ReadFile(filepath.Join(configDir, "config.json"))
	helpers.FatalCheckError(err)
	err = json.Unmarshal(jsonContent, &config)
	helpers.FatalCheckError(err)

	// create an app object that contains our routes and the configuration
	GironService := new(controllers.GironService)
	GironService.AppPath = appdir
	GironService.ConfigPath = configDir
	GironService.ConfStruct = config

	err = model.ConnectDatabase(GironService.ConfStruct.DbPath)
	helpers.FatalCheckError(err)

	// set up our static assets
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*.html")

	// some defaults for using session support
	r.Use(sessions.Sessions("session", cookie.NewStore(globals.Secret)))
	// frontend
	fePublic := r.Group("/")
	routes.FePublicRoutes(fePublic, GironService)

	fePrivate := r.Group("/")
	fePrivate.Use(middleware.AuthCheck)
	routes.FePrivateRoutes(fePrivate, GironService)

	// API
	public := r.Group("/api/v1")
	routes.PublicRoutes(public, GironService)

	private := r.Group("/api/v1")
	private.Use(middleware.AuthCheck)
	routes.PrivateRoutes(private, GironService)

	// swagger doc
	r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	tcpPort := strconv.Itoa(GironService.ConfStruct.TcpPort)
	tlsTcpPort := strconv.Itoa(GironService.ConfStruct.TLSTcpPort)
	tlsPemFile := GironService.ConfStruct.TLSPemFile
	tlsKeyFile := GironService.ConfStruct.TLSKeyFile
	if GironService.ConfStruct.UseTLS {
		r.RunTLS(":"+tlsTcpPort, tlsPemFile, tlsKeyFile)
	} else {
		r.Run(":" + tcpPort)
	}
}
