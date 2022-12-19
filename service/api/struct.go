package api

import (
	"github.com/InfernalPyro/WASA-Homework/service/database"
)

type Session struct {
	Username string `json:"username"`
}

type Comment struct {
	CommentId uint64 `json:"commentId"`
	PhotoID   uint64 `json:"photoId"`
	UserId    uint64 `json:"userId"`
	Comment   string `json:"comment"`
	Time      string `json:"time"`
}

type Photo struct {
	PhotoId   uint64    `json:"photoId"`
	ProfileId uint64    `json:"profileId"`
	Image     string    `json:"image"`
	Time      string    `json:"time"`
	Likes     []uint64  `json:"likes"`
	Comments  []Comment `json:"comments"`
}

type User struct {
	UserId    uint64   `json:"userId"`
	Username  string   `json:"username"`
	Follows   []uint64 `json:"follows"`
	Followers []uint64 `json:"followers"`
	Banned    []uint64 `json:"banned"`
	Photos    []Photo  `json:"photos"`
}

// This function convert all the data of a single user taken from the db into a single user in api form
func (u *User) UserFromDatabase(user database.User, follow []database.Follow, followed []database.Follow, bans []database.Ban, comments []database.Comment, photos []Photo, likes []database.Like) {
	u.UserId = user.UserId
	u.Username = user.Username
	for _, f := range follow {
		u.Follows = append(u.Follows, f.Follows)
	}
	for _, fd := range followed {
		u.Followers = append(u.Followers, fd.UserId)
	}
	for _, b := range bans {
		u.Banned = append(u.Banned, b.Banned)
	}
	for _, p := range photos {
		u.Photos = append(u.Photos, p)
	}

}

// This function convert all the data of a single photo taken from the db into a single photo in api form
func (p *Photo) PhotoFromDatabase(photo database.Photo, comments []database.Comment, likes []database.Like) {
	p.PhotoId = photo.PhotoId
	p.ProfileId = photo.UserId
	p.Image = photo.Image
	p.Time = photo.Time
	for _, l := range likes {
		p.Likes = append(p.Likes, l.UserId)
	}
	for _, c := range comments {
		var comment Comment
		comment.CommentId = c.CommentId
		comment.PhotoID = c.PhotoId
		comment.UserId = c.UserId
		comment.Comment = c.Content
		comment.Time = c.Time
		p.Comments = append(p.Comments, comment)
	}
	return
}