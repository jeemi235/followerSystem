package models

type FollowingDetails struct {
	FollowerId  int `json:"followerId"`
	FollowingId int `json:"followingId"`
}

type Requests struct {
	FollowingId int `json:"followingId"`
	FollowerId  int `json:"followerId"`
}

type Actions struct {
	Action      string `json:"action"`
}

type FollowingDetail struct {
	FollowingId int `json:"followingId"`
}

type FollowerDetail struct{
	FollowerId  int `json:"followerId"`
}