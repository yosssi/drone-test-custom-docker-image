package mongo

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type User struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

func Find(name string) (*User, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("droneTest").C("users")
	var user User
	err = c.Find(bson.M{"name": name}).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
