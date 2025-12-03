package main

import (
	"fmt"
	"log"
)

type MockDataStore struct {
	users map[int]User
}

type User struct {
	id    int
	first string
}

type DataStore interface {
	GetUser(id int) (User, error)
	SaveUser(user User) error
}

func (m MockDataStore) GetUser(id int) (User, error) {
	user, ok := m.users[id]
	if !ok {
		return User{}, fmt.Errorf("ID: %v, Not Found", id)
	}

	return user, nil
}

func (m MockDataStore) SaveUser(u User) error {
	_, ok := m.users[u.id]
	if ok {
		return fmt.Errorf("ID: %v, already exists", u.id)
	}

	m.users[u.id] = u
	return nil
}

type Service struct {
	ds DataStore
}

func (s Service) GetUser(id int) (User, error) {
	return s.ds.GetUser(id)
}

func (s Service) SaveUser(u User) error {
	return s.ds.SaveUser(u)
}

func main() {
	md := MockDataStore{
		users: make(map[int]User),
	}

	s := Service{
		ds: md,
	}

	u := User{
		id:    1,
		first: "Amir",
	}

	err := s.SaveUser(u)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	//u2 := User{
	//	id:    1,
	//	first: "koala",
	//}
	//
	//err2 := s.SaveUser(u2)
	//if err2 != nil {
	//	log.Fatalf("Error: %s", err2)
	//}

	myUser, _ := s.GetUser(1)
	fmt.Println(myUser)
}
