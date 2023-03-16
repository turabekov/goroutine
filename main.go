package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v4"
	"github.com/google/uuid"
)

var users = []User{}
var products = []Product{}
var shopCarts = []ShopCart{}

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
	total := 0

	for _, pr := range products {
		if pr.Id == shopcart.Product_id {
			total += pr.Price * shopcart.Amount
		}
	}

	fmt.Printf("%d. Total price: %d \n", idx, total)
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)+500))
	empty <- struct{}{}
}
func Kassir2(idx int, shopcart *ShopCart, empty chan struct{}) {
	shopcart.IsFinished = true
	total := 0

	for _, pr := range products {
		if pr.Id == shopcart.Product_id {
			total += pr.Price * shopcart.Amount
		}
	}

	fmt.Printf("%d. Total price: %d \n", idx, total)

	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)+500))
	empty <- struct{}{}
}
func Kassir3(idx int, shopcart *ShopCart, empty chan struct{}) {
	shopcart.IsFinished = true

	total := 0

	for _, pr := range products {
		if pr.Id == shopcart.Product_id {
			total += pr.Price * shopcart.Amount
		}
	}

	fmt.Printf("%d. Total price: %d \n", idx, total)

	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)+500))
	empty <- struct{}{}
}

func main() {

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
	idx := 0
	for idx < len(shopCarts) {
		go Kassir1(idx, &shopCarts[idx], empty)
		if idx+1 < len(shopCarts) {
			go Kassir2(idx+1, &shopCarts[idx+1], empty)
		}
		if idx+2 < len(shopCarts) {
			go Kassir3(idx+2, &shopCarts[idx+2], empty)
		}
		idx += 3
	}
	<-empty
	<-empty
	<-empty
	// end
	end := time.Now().UnixNano() / int64(time.Millisecond)
	diff := end - start
	log.Printf("Duration(ms): %d", diff)

	// fmt.Scanln()

	// fmt.Println(shopCarts)

}
