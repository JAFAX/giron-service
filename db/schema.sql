--
-- File generated with SQLiteStudio v3.4.4 on Sat Apr 6 12:04:50 2024
--
-- Text encoding used: UTF-8
--
PRAGMA foreign_keys = off;
BEGIN TRANSACTION;

-- Table: BuildingFloors
DROP TABLE IF EXISTS BuildingFloors;

CREATE TABLE IF NOT EXISTS BuildingFloors (
    Id           INTEGER  PRIMARY KEY AUTOINCREMENT
                          UNIQUE
                          NOT NULL,
    FloorName    STRING   NOT NULL
                          UNIQUE,
    BuildingId   INTEGER  NOT NULL
                          REFERENCES Buildings (Id),
    CreatorId    INTEGER  NOT NULL
                          REFERENCES Users (Id),
    CreationDate DATETIME NOT NULL
                          DEFAULT (CURRENT_TIMESTAMP) 
);


-- Table: Buildings
DROP TABLE IF EXISTS Buildings;

CREATE TABLE IF NOT EXISTS Buildings (
    Id           INTEGER  PRIMARY KEY AUTOINCREMENT
                          UNIQUE
                          NOT NULL,
    Name         STRING   UNIQUE
                          NOT NULL,
    City         STRING   NOT NULL,
    Region       STRING   NOT NULL,
    CreatorId    INTEGER  REFERENCES Users (Id) 
                          NOT NULL,
    CreationDate DATETIME NOT NULL
                          DEFAULT (CURRENT_TIMESTAMP) 
);


-- Table: Locations
DROP TABLE IF EXISTS Locations;

CREATE TABLE IF NOT EXISTS Locations (
    Id           INTEGER  PRIMARY KEY AUTOINCREMENT
                          NOT NULL
                          UNIQUE,
    RoomName     STRING   NOT NULL
                          UNIQUE,
    FloorId      INTEGER  REFERENCES BuildingFloors (Id) 
                          NOT NULL,
    BuildingId   INTEGER  REFERENCES Buildings (Id) 
                          NOT NULL,
    CreatorId    INTEGER  REFERENCES Users (Id) 
                          NOT NULL,
    CreationDate DATETIME NOT NULL
                          DEFAULT (CURRENT_TIMESTAMP) 
);


-- Table: Panels
DROP TABLE IF EXISTS Panels;

CREATE TABLE IF NOT EXISTS Panels (
    Id                  INTEGER  PRIMARY KEY AUTOINCREMENT
                                 NOT NULL
                                 UNIQUE,
    Topic               STRING   NOT NULL
                                 UNIQUE,
    Description         TEXT     NOT NULL,
    PanelRequestorEmail STRING   NOT NULL,
    Location            STRING   DEFAULT "",
    ScheduledTime       DATETIME DEFAULT "",
    DurationInMinutes   INTEGER  NOT NULL
                                 DEFAULT (30),
    CreatorId           INTEGER  NOT NULL
                                 REFERENCES Users (Id),
    CreationDateTime    DATETIME NOT NULL
                                 DEFAULT (CURRENT_TIMESTAMP),
    ApprovalStatus      BOOL     NOT NULL
                                 DEFAULT (FALSE),
    ApprovedById        INTEGER  REFERENCES Users (Id) 
                                 DEFAULT (0),
    ApprovalDateTime    DATETIME DEFAULT ""
);


-- Table: ScheduledEvents
DROP TABLE IF EXISTS ScheduledEvents;

CREATE TABLE IF NOT EXISTS ScheduledEvents (
    Id                INTEGER  PRIMARY KEY AUTOINCREMENT
                               UNIQUE
                               NOT NULL,
    RoomId            INTEGER  REFERENCES Locations (Id) 
                               NOT NULL,
    EventId           INTEGER  REFERENCES Panels (Id) 
                               NOT NULL,
    ScheduledTime     DATETIME NOT NULL,
    DurationInMinutes INTEGER  NOT NULL
);


-- Table: Users
DROP TABLE IF EXISTS Users;

CREATE TABLE IF NOT EXISTS Users (
    Id              INTEGER  PRIMARY KEY AUTOINCREMENT
                             NOT NULL
                             UNIQUE,
    UserName        STRING   UNIQUE
                             NOT NULL,
    Status          STRING   DEFAULT enabled
                             NOT NULL,
    PasswordHash    STRING   NOT NULL,
    CreationDate    DATETIME NOT NULL
                             DEFAULT (CURRENT_TIMESTAMP),
    LastChangedDate DATETIME NOT NULL
                             DEFAULT (CURRENT_TIMESTAMP) 
);


COMMIT TRANSACTION;
PRAGMA foreign_keys = on;
