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

import (
	"database/sql"
)

// primary object structs

type Building struct {
	Id           int    `json:"Id"`
	Name         string `json:"name"`
	City         string `json:"city"`
	Region       string `json:"region"`
	CreatorId    int    `json:"creatorId"`
	CreationDate string `json:"creationDateTime"`
}

type BuildingFloor struct {
	Id           int    `json:"Id"`
	FloorName    string `json:"floorName"`
	BuildingId   int    `json:"buildingId"`
	CreatorId    int    `json:"creatorId"`
	CreationDate string `json:"creationDateTime"`
}

type HealthCheck struct {
	Db           string `json:"db"`
	DiskSpace    string `json:"diskSpace"`
	DiskWritable string `json:"diskWritable"`
	Health       string `json:"health"`
	Status       int    `json:"status"`
}

type BuildingUpdate struct {
	Name   string `json:"name"`
	City   string `json:"city"`
	Region string `json:"region"`
}

type FloorUpdate struct {
	FloorName  string `json:"name"`
	BuildingId int    `json:"buildingId"`
}

type LocationUpdate struct {
	FloorId    int `json:"floorId"`
	BuildingId int `json:"buildingId"`
}

type Location struct {
	Id           int    `json:"Id"`
	Location     string `json:"location"`
	FloorId      int    `json:"floorId"`
	BuildingId   int    `json:"buildingId"`
	CreatorId    int    `json:"creatorId"`
	CreationDate string `json:"creationDateTime"`
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

type PanelApproval struct {
	State bool `json:"state"`
}

type PasswordChange struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type ScheduledEvent struct {
	Id                int    `json:"Id"`
	LocationId        int    `json:"locationId"`
	EventId           int    `json:"eventId"`
	ScheduledTime     string `json:"scheduledTime"`
	DurationInMinutes int    `json:"durationInMinutes"`
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

type UserStatus struct {
	Status string `json:"status" enum:"enabled,disabled"`
}

type UserStatusMsg struct {
	Message    string `json:"message"`
	UserStatus string `json:"userStatus" enum:"enabled,disabled"`
}

// proposed object structs. Normally used when creating new DB entries

type ProposedBuilding struct {
	Name   string `json:"name"`
	City   string `json:"city"`
	Region string `json:"region"`
}

type ProposedFloor struct {
	Name         string `json:"name"`
	BuildingName string `json:"buildingName"`
}

type ProposedLocation struct {
	RoomName   string `json:"name"`
	FloorId    int    `json:"floorId"`
	BuildingId int    `json:"buildingId"`
}

type ProposedPanel struct {
	Topic               string `json:"topic"`
	Description         string `json:"description"`
	PanelRequestorEmail string `json:"panelRequestorEmail"`
}

type ProposedUser struct {
	Id       int    `json:"Id"`
	UserName string `json:"userName"`
	Status   string `json:"status" enum:"enabled,disabled"`
	Password string `json:"password"`
}

// list object structs

type BuildingList struct {
	Data []Building `json:"data"`
}

type FloorList struct {
	Data []BuildingFloor `json:"data"`
}

type LocationList struct {
	Data []Location `json:"data"`
}

type PanelList struct {
	Data []Panel `json:"data"`
}

type UsersList struct {
	Data []User `json:"data"`
}

// generic message structs

type FailureMsg struct {
	Error string `json:"error"`
}

type SuccessMsg struct {
	Message string `json:"message"`
}
