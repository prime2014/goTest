package accounts

import (
	"db"
	"testing"
)

func TestSignup(t *testing.T) {
	db.ConnectTestDB()

	db.Db.AutoMigrate(&Users{})

	var user = Users{
		FirstName: "Sasha",
		LastName:  "Banks",
		Email:     "sasha.banks@mail.com",
		Password:  "belindat2014",
	}

	_, err := user.Save(db.Db)

	if err != nil {
		t.Errorf(err.Error())
	}

}

func TestLoginUser(t *testing.T) {
	var user = AuthenticationStruct{
		Email:    "sasha.banks@mail.com",
		Password: "belindat2014",
	}

	_, err := user.Login(db.Db)

	if err != nil {
		t.Errorf(err.Error())
	}

	TearDown()
}

func TearDown() {
	pg, _ := db.Db.DB()
	pg.Exec("DROP TABLE users")
	pg.Close()
}
