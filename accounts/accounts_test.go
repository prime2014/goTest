package accounts

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/carlmjohnson/requests"
)

func TestSignup(t *testing.T) {

	credentials := map[string]string{
		"firstname": "Samra",
		"lastname":  "Dita",
		"email":     "samra.dita@gmail.com",
		"password":  "belindat2014",
	}

	bodyValue, _ := json.Marshal(credentials)

	t.Log(string(bodyValue))

	err := requests.
		URL("http://127.0.0.1:8080/api/v1/users/signup").
		BodyBytes([]byte(bodyValue)).
		Fetch(context.Background())

	if err != nil {
		t.Error(err.Error())
	}

}
