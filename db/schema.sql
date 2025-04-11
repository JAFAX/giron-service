--
-- File generated with SQLiteStudio v3.4.4 on Thu Apr 10 20:01:48 2025
--
-- Text encoding used: UTF-8
--
PRAGMA foreign_keys = off;
BEGIN TRANSACTION;

-- Table: Audit
DROP TABLE IF EXISTS Audit;

CREATE TABLE IF NOT EXISTS Audit (
    Id           INTEGER  PRIMARY KEY AUTOINCREMENT
                          UNIQUE
                          NOT NULL,
    ChangedById  INTEGER  REFERENCES Users (Id) 
                          NOT NULL,
    TableChanged STRING   NOT NULL,
    ChangeClass  STRING   NOT NULL,
    ChangeDate   DATETIME NOT NULL
                          DEFAULT (CURRENT_TIMESTAMP) 
);


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


-- Table: LiveEventRatings
DROP TABLE IF EXISTS LiveEventRatings;

CREATE TABLE IF NOT EXISTS LiveEventRatings (
    Id          INTEGER PRIMARY KEY AUTOINCREMENT
                        UNIQUE
                        NOT NULL,
    UserId      INTEGER REFERENCES Users (Id) 
                        NOT NULL,
    LiveEventId INTEGER REFERENCES LiveEvents (Id) 
                        NOT NULL,
    Rating      INTEGER NOT NULL
                        DEFAULT (0) 
);


-- Table: LiveEvents
DROP TABLE IF EXISTS LiveEvents;

CREATE TABLE IF NOT EXISTS LiveEvents (
    Id                INTEGER  PRIMARY KEY AUTOINCREMENT
                               NOT NULL
                               UNIQUE,
    Topic             STRING   NOT NULL
                               UNIQUE,
    Description       TEXT     NOT NULL,
    Location          STRING   DEFAULT "",
    ScheduledTime     DATETIME DEFAULT "",
    DurationInMinutes INTEGER  NOT NULL
                               DEFAULT (30),
    Rating            INTEGER  DEFAULT (0),
    AgeRestricted     BOOL     NOT NULL
                               DEFAULT (FALSE),
    CreatorId         INTEGER  NOT NULL
                               REFERENCES Users (Id),
    CreationDateTime  DATETIME NOT NULL
                               DEFAULT (CURRENT_TIMESTAMP) 
);


-- Table: LiveEventTagAssignments
DROP TABLE IF EXISTS LiveEventTagAssignments;

CREATE TABLE IF NOT EXISTS LiveEventTagAssignments (
    Id          INTEGER PRIMARY KEY AUTOINCREMENT
                        UNIQUE
                        NOT NULL,
    TagId       INTEGER NOT NULL
                        REFERENCES Tags (Id),
    LiveEventId INTEGER REFERENCES LiveEvents (Id) 
                        NOT NULL
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


-- Table: PanelRatings
DROP TABLE IF EXISTS PanelRatings;

CREATE TABLE IF NOT EXISTS PanelRatings (
    Id           INTEGER  PRIMARY KEY AUTOINCREMENT
                          UNIQUE
                          NOT NULL,
    UserId       INTEGER  REFERENCES Users (Id) 
                          NOT NULL,
    PanelId      INTEGER  NOT NULL
                          REFERENCES Panels (Id),
    Rating       INTEGER  NOT NULL,
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
    Rating              INTEGER  NOT NULL
                                 DEFAULT (0),
    AgeRestricted       BOOL     NOT NULL
                                 DEFAULT (FALSE),
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


-- Table: PanelTagAssignments
DROP TABLE IF EXISTS PanelTagAssignments;

CREATE TABLE IF NOT EXISTS PanelTagAssignments (
    Id      INTEGER PRIMARY KEY AUTOINCREMENT
                    UNIQUE
                    NOT NULL,
    TagId   INTEGER NOT NULL
                    REFERENCES Tags (Id),
    PanelId INTEGER REFERENCES Panels (Id) 
                    NOT NULL
);


-- Table: PrivilegeAssignments
DROP TABLE IF EXISTS PrivilegeAssignments;

CREATE TABLE IF NOT EXISTS PrivilegeAssignments (
    Id     INTEGER PRIMARY KEY AUTOINCREMENT
                   UNIQUE
                   NOT NULL,
    RoleId INTEGER REFERENCES Roles (Id) 
                   NOT NULL,
    PrivId INTEGER REFERENCES Privileges (Id) 
                   NOT NULL
);


-- Table: Privileges
DROP TABLE IF EXISTS Privileges;

CREATE TABLE IF NOT EXISTS Privileges (
    Id              INTEGER PRIMARY KEY AUTOINCREMENT
                            UNIQUE
                            NOT NULL,
    PrivShortName   STRING  UNIQUE
                            NOT NULL,
    PrivDescription STRING  NOT NULL
);


-- Table: Roles
DROP TABLE IF EXISTS Roles;

CREATE TABLE IF NOT EXISTS Roles (
    Id           INTEGER  PRIMARY KEY AUTOINCREMENT
                          UNIQUE
                          NOT NULL,
    RoleName     STRING   NOT NULL,
    Description  STRING   NOT NULL,
    CreationDate DATETIME NOT NULL
                          DEFAULT (CURRENT_TIMESTAMP) 
);


-- Table: Tags
DROP TABLE IF EXISTS Tags;

CREATE TABLE IF NOT EXISTS Tags (
    Id      INTEGER PRIMARY KEY AUTOINCREMENT,
    TagName STRING  UNIQUE
                    NOT NULL
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


-- Table: VideoScreeningRating
DROP TABLE IF EXISTS VideoScreeningRating;

CREATE TABLE IF NOT EXISTS VideoScreeningRating (
    Id               INTEGER  PRIMARY KEY AUTOINCREMENT
                              NOT NULL
                              UNIQUE,
    UserId           INTEGER  REFERENCES Users (Id) 
                              NOT NULL,
    VideoScreeningId INTEGER  REFERENCES VideoScreenings (Id) 
                              NOT NULL,
    Rating           INTEGER  NOT NULL
                              DEFAULT (0),
    CreationDate     DATETIME NOT NULL
                              DEFAULT (CURRENT_TIMESTAMP) 
);


-- Table: VideoScreenings
DROP TABLE IF EXISTS VideoScreenings;

CREATE TABLE IF NOT EXISTS VideoScreenings (
    Id                INTEGER  PRIMARY KEY AUTOINCREMENT
                               NOT NULL,
    Title             STRING   NOT NULL,
    Synopsis          TEXT     NOT NULL,
    Location          STRING,
    ScheduledTime     DATETIME,
    DurationInMinutes INTEGER  NOT NULL,
    AgeRestricted     BOOL     NOT NULL
                               DEFAULT (FALSE),
    CreatorId         INTEGER  REFERENCES Users (Id) 
                               NOT NULL,
    CreationDateTime  DATETIME NOT NULL
                               DEFAULT (CURRENT_TIMESTAMP) 
);


-- Table: VideoScreeningTagAssignments
DROP TABLE IF EXISTS VideoScreeningTagAssignments;

CREATE TABLE IF NOT EXISTS VideoScreeningTagAssignments (
    Id      INTEGER PRIMARY KEY AUTOINCREMENT
                    UNIQUE
                    NOT NULL,
    TagId   INTEGER NOT NULL
                    REFERENCES Tags (Id),
    PanelId INTEGER REFERENCES VideoScreenings (Id) 
                    NOT NULL
);


COMMIT TRANSACTION;
PRAGMA foreign_keys = on;
