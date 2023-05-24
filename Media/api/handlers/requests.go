package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"media/middlewares"
	"media/models"
	"net/http"
)

// This function will send the request
func SendRequest(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("id")
	var request models.FollowerDetail
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}
	db := r.Context().Value("database").(*sql.DB)
	if _, err := db.Exec(
		"INSERT INTO request (following_id,follower_id) values($1,$2)", id, request.FollowerId); err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("Request sent successfully"))
	log.Print("successfully got followings details\n")
}

// This function will give us the list of people who has requested us
func PendingFollowers(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("id")
	db := r.Context().Value("database").(*sql.DB)
	rows, err := db.Query("select * from request where follower_id =$1", id)
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}
	defer rows.Close()

	requests := []models.Requests{}
	for rows.Next() {
		var request models.Requests
		if err := rows.Scan(&request.FollowingId, &request.FollowerId); err != nil {
			log.Println(err)
			w.Write([]byte(err.Error()))
			return
		}
		requests = append(requests, request)

	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}
	middlewares.ResponseWithJsonPayload(w, requests)
	log.Print("successfully got pending followers requests\n")
}

// This function will give us the list of people whom we have sent the request
func PendingFollowing(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("id")
	db := r.Context().Value("database").(*sql.DB)
	rows, err := db.Query("select * from request where following_id =$1", id)
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}
	defer rows.Close()

	requests := []models.Requests{}
	for rows.Next() {
		var request models.Requests
		if err := rows.Scan(&request.FollowingId, &request.FollowerId); err != nil {
			log.Println(err)
			w.Write([]byte(err.Error()))
			return
		}
		requests = append(requests, request)

	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}
	middlewares.ResponseWithJsonPayload(w, requests)
	log.Print("successfully got pending following requests\n")
}

// This function will accept or reject the pending request
func AcceptRejectRequest(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("id")
	following_id := r.URL.Query().Get("following_id")
	var action models.Actions
	err := json.NewDecoder(r.Body).Decode(&action)
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}
	db := r.Context().Value("database").(*sql.DB)

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

	query := `delete from request where following_id=$1 and follower_id=$2`
	_, err = tx.Exec(query, following_id, id)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

	if action.Action == "accept" {
		query1 := `insert into following(follower_id,following_id) values($1,$2)`
		_, err = tx.Exec(query1, id, following_id)
		if err != nil {
			tx.Rollback()
			log.Println(err)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write([]byte("Request accepted successfully"))
		log.Println("Successfully accepted the request")

	} else if action.Action == "reject" {
		w.Write([]byte("Request rejected successfully"))
		log.Println("You rejected the request")

	} else {
		w.Write([]byte("Please enter valid action"))
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}
}

// This function will remove people from list whom we have requested and which is stil pending
func RemoveFollowingRequest(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("id")
	var request models.FollowerDetail
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}
	db := r.Context().Value("database").(*sql.DB)
	if _, err := db.Exec(
		"delete from request where following_id=$1 and follower_id=$2", id, request.FollowerId); err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("successfully removed pending following request"))
	log.Print("successfully removed pending following request\n")
}
