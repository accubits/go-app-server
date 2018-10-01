package user

import (
	"db"
	"logger"

	mgo "gopkg.in/mgo.v2"
)

var log = logger.GetLogger()

type User struct {
	ID       string `json:"id"`
	FName    string `json:"fName"`
	LName    string `json:"lName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) Save() {
	mdb := db.GetDB()
	c := mdb.C("User")
	err := c.Insert(u)
	if err != nil {
		log.Error(err)
	}
}

func (u *User) Update(newUser User) {
	mdb := db.GetDB()
	c := mdb.C("User")
	err := c.Update(u, newUser)
	if err != nil {
		log.Error(err)
	}
}
func GetByID(u User) User {
	mdb := db.GetDB()
	c := mdb.C("User")
	var user User
	err := c.Find(u).One(&user)
	if err != nil {
		log.Error(err)
	}
	return user
}

func ensureIndex() {
	mdb := db.GetDB()
	c := mdb.C("User")

	index := mgo.Index{
		Key:        []string{"id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}
