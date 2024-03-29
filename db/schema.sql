--
-- File generated with SQLiteStudio v3.4.4 on Fri Mar 29 16:42:21 2024
--
-- Text encoding used: UTF-8
--
PRAGMA foreign_keys = off;
BEGIN TRANSACTION;

-- Table: Panels
DROP TABLE IF EXISTS Panels;

CREATE TABLE IF NOT EXISTS Panels (
    Id                  INTEGER  PRIMARY KEY AUTOINCREMENT
                                 NOT NULL
                                 UNIQUE,
    Topic               STRING   NOT NULL
                                 UNIQUE,
    Description         TEXT     NOT NULL,
    PanelRequestorEmail STRING   NOT NULL
                                 DEFAULT "",
    Location            STRING   NOT NULL
                                 DEFAULT "",
    ScheduledTime       DATETIME NOT NULL,
    CreatorId           INTEGER  NOT NULL
                                 REFERENCES Users (Id),
    CreationDateTime    DATETIME NOT NULL
                                 DEFAULT (CURRENT_TIMESTAMP),
    ApprovalStatus      BOOL     NOT NULL
                                 DEFAULT (FALSE),
    ApprovedById        INTEGER  REFERENCES Users (Id),
    ApprovalDateTime    DATETIME
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
