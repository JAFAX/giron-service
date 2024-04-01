package model

import "database/sql"

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

type Location struct {
	Location string `json:"location"`
	// Floor    string `json:"locationFloor"`
}

type ProposedPanel struct {
	Topic               string `json:"topic"`
	Description         string `json:"description"`
	PanelRequestorEmail string `json:"panelRequestorEmail"`
}

type PanelList struct {
	Data []Panel `json:"data"`
}

type PanelSQL struct {
	Id                  int            `json:"Id"`
	Topic               string         `json:"topic"`
	Description         string         `json:"description"`
	PanelRequestorEmail string         `json:"panelRequestorEmail"`
	Location            string         `json:"location"`
	ScheduledTime       sql.NullString `json:"scheduledTime"`
	DurationInMinutes   int            `json:"durationInMinutes"`
	CreatorId           int            `json:"creatorId"`
	CreationDateTime    string         `json:"creationDateTime"`
	ApprovalStatus      bool           `json:"approvalStatus"`
	ApprovedById        sql.NullInt64  `json:"approvedById"`
	ApprovalDateTime    sql.NullString `json:"approvalDateTime"`
}

type Panel struct {
	Id                  int    `json:"Id"`
	Topic               string `json:"topic"`
	Description         string `json:"description"`
	PanelRequestorEmail string `json:"panelRequestorEmail"`
	Location            string `json:"location"`
	ScheduledTime       string `json:"scheduledTime"`
	DurationInMinutes   int    `json:"durationInMinutes"`
	CreatorId           int    `json:"creatorId"`
	CreationDateTime    string `json:"creationDateTime"`
	ApprovalStatus      bool   `json:"approvalStatus"`
	ApprovedById        int    `json:"approvedById"`
	ApprovalDateTime    string `json:"approvalDateTime"`
}

type Schedule struct {
	StartTime         string `json:"startTime"`
	DurationInMinutes int    `json:"durationInMinutes"`
}

type User struct {
	Id              int    `json:"Id"`
	UserName        string `json:"userName"`
	Status          string `json:"status"`
	PasswordHash    string `json:"passwordHash"`
	CreationDate    string `json:"creationDate"`
	LastChangedDate string `json:"lastChangedDate"`
}

type ProposedUser struct {
	Id           int    `json:"Id"`
	UserName     string `json:"userName"`
	Status       string `json:"status"`
	Password     string `json:"password"`
	CreationDate string `json:"creationDate"`
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
