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
	"net/http"

	"github.com/JAFAX/giron-service/globals"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (g *GironService) AdminUI(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.UserKey)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"content": "This is a dashboard",
		"user":    user,
	})
}
