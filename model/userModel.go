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
	"crypto/sha512"
	"database/sql"
	"encoding/hex"
	"errors"
	"log"
	"strconv"
	"time"
)

func getStoredPasswordHash(username string) (string, error) {
	stmt, err := DB.Prepare("SELECT PasswordHash FROM Users WHERE UserName = ?")
	if err != nil {
		log.Println("ERROR: Could not prepare the DB query!" + string(err.Error()))
		return "", err
	}
	defer stmt.Close()

	passwordHash := ""
	record, err := stmt.Query(username)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("ERROR: No such user found in DB: " + string(err.Error()))
			return "", nil
		}
		log.Println("ERROR: Cannot retrieve user from DB: " + string(err.Error()))
		return "", err
	}
	defer record.Close()
	err = record.Scan(
		&passwordHash,
	)
	if err != nil {
		log.Println("ERROR: Cannot retrieve password hash from DB: " + string(err.Error()))
		return "", err
	}

	log.Println("INFO: Retrieved password hash from DB")
	return passwordHash, nil
}

func storeNewPassword(hashedPassword string, username string) (bool, error) {
	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Could not start DB transaction!" + string(err.Error()))
		return false, err
	}
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: DB transaction failed: " + string(err.Error()))
			t.Rollback()
		}
		if err != nil {
			log.Println("ERROR: DB transaction failed: " + string(err.Error()))
			t.Rollback()
		}
	}()

	// now we need to create a new transaction to SET the password hash into the DB
	q, err := DB.Prepare("UPDATE Users SET PasswordHash = ?, LastChangedDate = ? WHERE UserName = ?")
	if err != nil {
		return false, err
	}

	// get time stamp
	tStamp := time.Now().Format("2006-01-02 15:04:05") // force into SQL DateTime format

	_, err = q.Exec(hashedPassword, tStamp, username)
	if err != nil {
		return false, err
	}

	err = t.Commit()
	if err != nil {
		log.Println("ERROR: Cannot commit DB transaction: " + string(err.Error()))
		return false, err
	}

	log.Println("INFO: Password hash stored in DB")
	return true, nil
}

func ChangeAccountPassword(username string, oldPassword string, newPassword string) (bool, error) {
	log.Println("INFO: Password change requested")
	hashedOldPassword := sha512.Sum512([]byte(oldPassword))
	encodedHashedOldPassword := hex.EncodeToString(hashedOldPassword[:])

	storedHash, err := getStoredPasswordHash(username)
	if err != nil {
		log.Println("ERROR: Cannot retrieve stored password hash from DB: " + string(err.Error()))
		return false, err
	}
	log.Println("INFO: Retrieved stored hash for comparison")

	// now get password hash from the db
	if storedHash != encodedHashedOldPassword {
		log.Println("ERROR: Hashed value of old password does not match stored hashed value")
		p := new(PasswordHashMismatch)
		return false, p
	}

	// matches, so hash new password
	hashedNewPassword := sha512.Sum512([]byte(newPassword))
	encodedHashedNewPassword := hex.EncodeToString(hashedNewPassword[:])
	_, err = storeNewPassword(encodedHashedNewPassword, username)
	if err != nil {
		log.Println("ERROR: Cannot store updated password hash in DB: " + string(err.Error()))
		return false, err
	}

	log.Println("INFO: Stored updated hash")
	return true, nil
}

func GetUserById(id int) (User, error) {
	log.Println("INFO: User by Id requested: " + strconv.Itoa(id))
	stmt, err := DB.Prepare("SELECT * FROM Users WHERE Id = ?")
	if err != nil {
		log.Println("ERROR: Could not prepare the DB query!" + string(err.Error()))
		return User{}, err
	}
	defer stmt.Close()

	user := User{}
	record, err := stmt.Query(id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("ERROR: No such user found in DB: " + string(err.Error()))
			return User{}, nil
		}
		log.Println("ERROR: Cannot retrieve user from DB: " + string(err.Error()))
		return User{}, err
	}
	defer record.Close()
	err = record.Scan(
		&user.Id,
		&user.UserName,
		&user.Status,
		&user.PasswordHash,
		&user.CreationDate,
		&user.LastChangedDate,
	)
	if err != nil {
		log.Println("ERROR: Cannot unmarshal the user object!" + string(err.Error()))
		return User{}, err
	}

	log.Println("INFO: User by Id '" + strconv.Itoa(id) + "' retrieved")
	return user, nil
}

func GetUserByUserName(username string) (User, error) {
	log.Println("INFO: User by username requested: " + username)
	stmt, err := DB.Prepare("SELECT * FROM Users WHERE UserName = ?")
	if err != nil {
		log.Println("ERROR: Could not prepare the DB query!" + string(err.Error()))
		return User{}, err
	}
	defer stmt.Close()

	user := User{}
	record, err := stmt.Query(username)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("ERROR: No such user found in DB: " + string(err.Error()))
			return User{}, nil
		}
		log.Println("ERROR: Cannot retrieve user from DB: " + string(err.Error()))
		return User{}, err
	}
	defer record.Close()
	err = record.Scan(
		&user.Id,
		&user.UserName,
		&user.Status,
		&user.PasswordHash,
		&user.CreationDate,
		&user.LastChangedDate,
	)
	if err != nil {
		log.Println("ERROR: Cannot unmarshal the user object!" + string(err.Error()))
		return User{}, err
	}

	log.Println("INFO: User by username '" + username + "' retrieved")
	return user, nil
}

func CreateUser(p ProposedUser) (bool, error) {
	log.Println("INFO: User creation requested: " + p.UserName)
	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Could not start DB transaction!" + string(err.Error()))
		return false, err
	}
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: DB transaction failed: " + string(err.Error()))
			t.Rollback()
		}
		if err != nil {
			log.Println("ERROR: DB transaction failed: " + string(err.Error()))
			t.Rollback()
		}
	}()

	q, err := t.Prepare("INSERT INTO Users (UserName, PasswordHash) VALUES (?, ?)")
	if err != nil {
		log.Println("ERROR: Could not prepare the DB query!" + string(err.Error()))
		return false, err
	}

	// take password and hash it
	hash := sha512.Sum512([]byte(p.Password))
	passwdHash := hex.EncodeToString(hash[:])

	_, err = q.Exec(p.UserName, passwdHash)
	if err != nil {
		log.Println("ERROR: Cannot create user '" + p.UserName + "': " + string(err.Error()))
		return false, err
	}

	err = t.Commit()
	if err != nil {
		log.Println("ERROR: Cannot commit DB transaction: " + string(err.Error()))
		return false, err
	}

	log.Println("INFO: User '" + p.UserName + "' created")
	return true, nil
}

func DeleteUser(username string) (bool, error) {
	log.Println("INFO: User deletion requested: " + username)
	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Could not start DB transaction!" + string(err.Error()))
		return false, err
	}
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: DB transaction failed: " + string(err.Error()))
			t.Rollback()
		}
		if err != nil {
			log.Println("ERROR: DB transaction failed: " + string(err.Error()))
			t.Rollback()
		}
	}()

	q, err := DB.Prepare("DELETE FROM Users WHERE UserName IS ?")
	if err != nil {
		log.Println("ERROR: Could not prepare the DB query!" + string(err.Error()))
		return false, err
	}

	_, err = q.Exec(username)
	if err != nil {
		log.Println("ERROR: Cannot delete user '" + username + "': " + string(err.Error()))
		return false, err
	}

	err = t.Commit()
	if err != nil {
		log.Println("ERROR: Cannot commit DB transaction: " + string(err.Error()))
		return false, err
	}

	log.Println("INFO: User '" + username + "' has been deleted")
	return true, nil
}

func GetUsers() ([]User, error) {
	log.Println("INFO: List of user object requested")
	rows, err := DB.Query("SELECT * FROM Users")
	if err != nil {
		log.Println("ERROR: Could not run the DB query!" + string(err.Error()))
		return nil, err
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		user := User{}
		err = rows.Scan(
			&user.Id,
			&user.UserName,
			&user.Status,
			&user.PasswordHash,
			&user.CreationDate,
			&user.LastChangedDate,
		)
		if err != nil {
			log.Println("ERROR: Cannot marshal the user objects!" + string(err.Error()))
			return nil, err
		}
		users = append(users, user)
	}

	log.Println("INFO: List of all users retrieved")
	return users, nil
}

func GetUserStatus(username string) (string, error) {
	log.Println("INFO: User status requested for user '" + username + "'")
	stmt, err := DB.Prepare("SELECT Status FROM Users WHERE UserName IS ?")
	if err != nil {
		log.Println("ERROR: Could not prepare DB query! " + string(err.Error()))
		return "", err
	}
	defer stmt.Close()

	status := ""
	record, err := stmt.Query(username)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("ERROR: No such user found in DB: " + string(err.Error()))
			return "", nil
		}
		log.Println("ERROR: Cannot retrieve user from DB: " + string(err.Error()))
		return "", err
	}
	defer record.Close()
	err = record.Scan(
		&status,
	)
	if err != nil {
		log.Println("ERROR: Cannot unmarshal the user object!" + string(err.Error()))
		return "", err
	}

	log.Println("INFO: User '" + username + "' status: " + status)
	return status, nil
}

func SetUserStatus(username string, j UserStatus) (bool, error) {
	log.Println("INFO: Set user status for user '" + username + "'")
	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Could not start DB transaction: " + string(err.Error()))
		return false, err
	}
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: DB transaction failed: " + string(err.Error()))
			t.Rollback()
		}
		if err != nil {
			log.Println("ERROR: DB transaction failed: " + string(err.Error()))
			t.Rollback()
		}
	}()

	q, err := DB.Prepare("UPDATE Users SET Status = ? WHERE UserName = ?")
	if err != nil {
		log.Println("ERROR: Could not prepare DB query! " + string(err.Error()))
		return false, err
	}
	// ensure the UserStatus.Status value is either 'enabled' or 'locked'
	log.Println("INFO: user to set status of: " + username)
	log.Println("INFO: requested state to set user to: " + j.Status)
	if j.Status != "enabled" && j.Status != "locked" {
		return false, &InvalidStatusValue{Err: errors.New("invalid value: " + j.Status)}
	}

	result, err := q.Exec(j.Status, username)
	if err != nil {
		log.Println("ERROR: Could not execute query for user '" + username + "': " + string(err.Error()))
		return false, err
	}
	numberOfRows, err := result.RowsAffected()
	if err != nil {
		log.Println("ERROR: Could not get number of rows affected: " + string(err.Error()))
		return false, err
	}

	err = t.Commit()
	if err != nil {
		log.Println("ERROR: Cannot commit DB transaction: " + string(err.Error()))
		return false, err
	}

	log.Println("INFO: SQL result: Rows: " + strconv.Itoa(int(numberOfRows)))
	return true, nil
}
