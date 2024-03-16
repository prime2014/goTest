package accounts

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Users struct {
	ID                  uint      `json:"id"`
	FirstName           string    `json:"firstname" gorm:"not null"`
	LastName            string    `json:"lastname"`
	Email               string    `json:"email" gorm:"unique;not null"`
	Password            string    `json:"password" gorm:"not null;check:char_length(password) >= 8"`
	IsActive            bool      `json:"is_active" gorm:"default:false"`
	DateJoined          time.Time `json:"date_joined" gorm:"default:current_timestamp"`
	PasswordResetToken  string    `json:"-"`
	PasswordResetExpiry time.Time `json:"-"`
}

type AuthenticationStruct struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func checkPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}

// This function is executed before any create operation
func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {

	// first check the plain text password isn't less than 8 characters
	if len(u.Password) < 8 {
		return fmt.Errorf("password too short")
	}

	// create a Hash of the plain text password
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)

	if err != nil {
		panic("there was a problem hashing password")
	}
	u.Password = string(bytes)
	return nil
}

// saves the user credentials to the database
func (u *Users) Save(tx *gorm.DB) (Users, error) {
	var user Users

	// first check the plain text password isn't less than 8 characters
	if len(u.Password) < 8 {
		return Users{}, fmt.Errorf("password too short")
	}

	// create a Hash of the plain text password
	bytes, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 14)

	u.Password = string(bytes)

	result := tx.Raw("INSERT INTO users(first_name, last_name, email, password) VALUES($1, $2, $3, $4) RETURNING id, first_name, last_name, email, password, is_active, date_joined",
		u.FirstName,
		u.LastName,
		u.Email,
		u.Password,
	).Scan(&user)

	fmt.Printf("RETURNED USER: %v", user)

	if result.Error != nil {
		fmt.Println("Navigant: " + result.Error.Error())
		return Users{}, result.Error
	}

	return user, nil
}

func GetUsers(tx *gorm.DB, offset int, limit int) ([]Users, error) {
	var users []Users
	// Find all users in the db
	result := tx.Find(&users).Offset(offset).Limit(limit).Order("date_joined desc")

	if result.Error != nil {
		return []Users{}, fmt.Errorf("there was a problem getting users")
	}
	return users, nil
}

func (u *AuthenticationStruct) Login(tx *gorm.DB) (map[string]string, error) {

	var result Users
	var tokenString string
	var err error

	myresult := tx.Raw("SELECT id, first_name, last_name, email, password FROM users WHERE email = $1", u.Email).Scan(&result)

	// Immediately return an error if the provided email does not exist in the db
	if myresult.Error != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	//confirm that the provided password matches the hashed one in the database
	passwordMatches := checkPassword(result.Password, u.Password)

	if passwordMatches {
		tokenString, err = GenerateJWT(result.Email, result.ID)

		if err != nil {
			return nil, fmt.Errorf("ERROR: Generating Token")
		}
	} else {
		return nil, fmt.Errorf("invalid credentials")
	}

	return map[string]string{
		"token": tokenString,
	}, nil
}

func GenerateJWT(email string, id uint) (string, error) {
	var mySigningKey = []byte("secrecyAWeapon")
	token := jwt.New(jwt.SigningMethodHS256)

	// create the claims to append data/payload to the token. This is also the way data shall
	// be extracted from the token when authenticating
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["sub"] = id
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		err := fmt.Errorf("ERROR: Token generation failure: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}
