package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v4"
	"github.com/google/uuid"
)

type User struct {
	Id    string
	Name  string
	Email string
}

type Product struct {
	Id    string
	Name  string
	Price int
}

type ShopCart struct {
	User_id    string
	Product_id string
	Amount     int
	IsFinished bool
}

func Kassir1(idx int, shopcart *ShopCart, empty chan struct{}) {
	shopcart.IsFinished = true

	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)+500))
	empty <- struct{}{}
}
func Kassir2(idx int, shopcart ShopCart, empty chan struct{}) {
	shopcart.IsFinished = true
}
func Kassir3(idx int, shopcart ShopCart, empty chan struct{}) {
	shopcart.IsFinished = true
}

func main() {
	users := []User{}
	products := []Product{}
	shopCarts := []ShopCart{}

	rand.Seed(time.Now().UnixNano())

	// generate users
	for i := 0; i < 100; i++ {
		users = append(users, User{
			Id:    uuid.NewString(),
			Name:  faker.Name(),
			Email: faker.Email(),
		})
	}

	// generate products
	for i := 0; i < 100; i++ {
		products = append(products, Product{
			Id:    uuid.NewString(),
			Name:  faker.Name(),
			Price: rand.Intn(100) * 100,
		})
	}

	// generate shopcarts
	for i := 0; i < 100; i++ {
		shopCarts = append(shopCarts, ShopCart{
			User_id:    users[rand.Intn(99-0)+0].Id,
			Product_id: products[rand.Intn(99-0)+0].Id,
			Amount:     rand.Intn(10-1) + 1,
			IsFinished: false,
		})
	}
	//-----------------------------------------------------------------------------
	empty := make(chan struct{})

	fmt.Println("Started")

	start := time.Now().UnixNano() / int64(time.Millisecond)
	// start
	for idx, _ := range shopCarts {
		go Kassir1(idx, &shopCarts[idx], empty)
	}
	<-empty
	// end
	end := time.Now().UnixNano() / int64(time.Millisecond)
	diff := end - start
	log.Printf("Duration(ms): %d", diff)

	// fmt.Scanln()

	fmt.Println(shopCarts)

}
