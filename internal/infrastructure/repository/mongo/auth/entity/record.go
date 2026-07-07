package entity

type AuthUserRecord struct {
	Id       string `bson:"_id,omitempty"`
	Login    string `bson:"login"`
	Password string `bson:"password"`
}
