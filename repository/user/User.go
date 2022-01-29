package user

import (
	"database/sql"
	"fmt"
	"log"
	"sirclo/gql/entities"
	"sirclo/gql/util"

	"golang.org/x/crypto/bcrypt"
)

type Repository_User struct {
	db *sql.DB
}

func NewRepositoryUser(db *sql.DB) *Repository_User {
	return &Repository_User{db}
}

//GetUserIdByUsername check if a user exists in database by given username
func GetUserIdByUsername(name string) (int, error) {
	statement, err := util.Db.Prepare("select ID from Users WHERE name = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := statement.QueryRow(name)

	var Id int
	err = row.Scan(&Id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return 0, err
	}

	return Id, nil
}

func (r *Repository_User) Authenticate(name, password string) bool {
	statement, err := r.db.Prepare("select password from users WHERE name = ?")
	if err != nil {
		log.Fatal(err)
	}
	var user entities.User
	row := statement.QueryRow(name)
	fmt.Println(row)

	// // var hashedPassword string
	err = row.Scan(&user.Password)
	fmt.Println(err)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println(user.Password)
	return password == user.Password
	// return CheckPasswordHash(user.Password, hashedPassword)
}

//get users
func (r *Repository_User) GetUsers() ([]entities.User, error) {
	var users []entities.User
	results, err := r.db.Query("select id, name, email from users order by id asc")
	if err != nil {
		log.Fatalf("Error")
	}

	// defer results.Close()

	for results.Next() {
		var user entities.User

		err := results.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			fmt.Println("ini jalan")
			log.Fatalf("Error")
		}

		users = append(users, user)
	}
	return users, nil
}

//get user
func (r *Repository_User) GetUser(id int) (entities.User, error) {
	fmt.Println("ini id", id)
	var user entities.User
	results := r.db.QueryRow("select id, name, email from users WHERE id = ?", id)
	// if err != nil {
	// 	log.Fatalf("Error")
	// }

	err := results.Scan(&user.Id, &user.Name, &user.Email)
	fmt.Print(user)
	if err != nil {
		return user, err
	}
	return user, nil
}

//create users
func (r *Repository_User) CreateUser(user entities.User) (entities.User, error) {
	query := `INSERT INTO USERS (name, email, password) VALUES (?, ?, ?)`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return user, err
	}

	// hashedPassword, err := HashPassword(user.Password)
	_, err = statement.Exec(user.Name, user.Email, user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}

//update users
func (r *Repository_User) UpdateUser(user entities.User) (entities.User, error) {
	query := `UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return user, err
	}
	defer statement.Close()

	_, err = statement.Exec(user.Name, user.Email, user.Password, user.Id)
	if err != nil {
		return user, err
	}

	return user, nil

}

//delete user
func (r *Repository_User) DeleteUser(user entities.User) (entities.User, error) {
	query := `DELETE FROM users WHERE id = ?`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return user, err
	}

	defer statement.Close()

	_, err = statement.Exec(user.Id)
	if err != nil {
		return user, err
	}

	return user, nil
}

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
