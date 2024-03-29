package model

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

type User struct {
	Id              int    `json:"Id"`
	UserName        string `json:"UserName"`
	Status          string `json:"Status"`
	PasswordHash    string `json:"PasswordHash"`
	CreationDate    string `json:"CreationDate"`
	LastChangedDate string `json:"LastChangedDate"`
}

type ProposedUser struct {
	Id           int    `json:"Id"`
	UserName     string `json:"UserName"`
	Status       string `json:"Status"`
	Password     string `json:"Password"`
	CreationDate string `json:"CreationDate"`
}

type UserStatus struct {
	Status string `json:"status"`
}

type FailureMsg struct {
	Error string `json:"error"`
}

type SuccessMsg struct {
	Message string `json:"message"`
}

type UserStatusMsg struct {
	Message    string `json:"message"`
	UserStatus string `json:"userStatus"`
}

type PasswordChange struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type UsersList struct {
	Data []User `json:"data"`
}
