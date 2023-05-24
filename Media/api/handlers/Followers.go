package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"media/models"
	"media/middlewares"
	"net/http"
)

//This function will give us the list of users whom we are following
func GetFollowers(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("id")
		db := r.Context().Value("database").(*sql.DB)

		rows, err := db.Query("select * from following where following_id =$1", id)
		if err != nil {
			log.Println(err)
			w.Write([]byte(err.Error()))
			return
		}
		defer rows.Close()

		followings := []models.FollowingDetails{}
		for rows.Next() {
			var following models.FollowingDetails
			if err := rows.Scan(&following.FollowerId, &following.FollowingId); err != nil {
				log.Println(err)
				w.Write([]byte(err.Error()))
				return
			}
			followings = append(followings, following)

		}
		if err := rows.Err(); err != nil {
			log.Println(err)
			w.Write([]byte(err.Error()))
			return
		}
		middlewares.ResponseWithJsonPayload(w, followings)
		log.Print("successfully got followers details\n")
	}

//This function will remove follower who are following us 
func RemoveFollower(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("id")
		var following models.FollowingDetail
		err := json.NewDecoder(r.Body).Decode(&following)
		if err != nil {
			log.Println(err)
			w.Write([]byte(err.Error()))
			return
		}
		db := r.Context().Value("database").(*sql.DB)
		if _, err := db.Exec(
			"delete from following where follower_id=$1 and following_id=$2", id, following.FollowingId); err != nil {
			log.Println(err)
			w.Write([]byte(err.Error()))
			return
		}
		
		w.Write([]byte("successfully removed from followers"))
		log.Print("successfully removed from followers\n")
	}

