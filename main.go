package main

import "fmt"

// Strategy interface
type NumberStorer interface {
	GetAll() ([]int, error)
	Put(number int) error
}

// Concrete strategies
type MongoDBNumberStorer struct {
}

func (m MongoDBNumberStorer) GetAll() ([]int, error) {
	return []int{1, 2, 3}, nil
}

func (m MongoDBNumberStorer) Put(number int) error {
	fmt.Println("store the number into the mongodb storage")
	return nil
}

type PostgresNumberStore struct {
}

func (m PostgresNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3}, nil
}

func (m PostgresNumberStore) Put(number int) error {
	fmt.Println("store the number into the postgres storage")
	return nil
}

// Context
type ApiServer struct {
	numberStore NumberStorer
}

func main() {
	apiServer := ApiServer{
		numberStore: MongoDBNumberStorer{},
	}

	err := apiServer.numberStore.Put(1)
	if err != nil {
		panic(err)
	}

	numbers, err := apiServer.numberStore.GetAll()
	if err != nil {
		panic(err)
	}

	fmt.Println("Numbers:", numbers)
}
