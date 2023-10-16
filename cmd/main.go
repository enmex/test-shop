package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"test-shop/internal/config"
	"test-shop/internal/dto"
	"test-shop/internal/repositories"
	"test-shop/internal/usecases"
	"test-shop/pkg/db"
	"test-shop/pkg/logger"
	"time"
)

func main() {
	//init application
	logger.Init()

	log := &logger.Logger
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	if err := db.Migrations(cfg.DB, log); err != nil {
		log.Warnln("failed to load migrations, retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
		if err := db.Migrations(cfg.DB, log); err != nil {
			log.Fatalln(err)
		}
	}

	conn, err := db.GetConnect(cfg.DB, log)
	if err != nil {
		log.Warnln("failed to establish db connection, retrying in 10 seconds...")
		time.Sleep(5 * time.Second)
		conn, err = db.GetConnect(cfg.DB, log)
		if err != nil {
			log.Fatalln(err)
		}	
	}
	log.Infoln("database connection established")

	orderRepo := repositories.NewOrderRepository(conn)
	productRepo := repositories.NewProductRepository(conn)

	orderAssemblyUsecase := usecases.NewOrderAssemblyUsecase(orderRepo, productRepo)

	//client code
	arg := os.Args[1]
	ordersStr := strings.Split(arg, ",")
	orderIds := make([]int, 0)
	for _, orderStr := range ordersStr {
		orderId, err := strconv.Atoi(orderStr)
		if err != nil {
			log.Fatalln("wrong input: must be numbers divided by comma like '1,2,3'")
		}
		orderIds = append(orderIds, orderId)
	}
	res, err := orderAssemblyUsecase.GetOrderAssembly(&dto.GetOrderAssemblyRequest{
		OrderIds: orderIds,
	})
	if err != nil {
		log.Errorln(err)
		return
	}
	racksToOrders := res.RacksToOrders
	output := "=+=+=+=\n"
	output += fmt.Sprintf("Страница сборки заказов %s\n\n", arg)

	for rack, orders := range racksToOrders {
		output += fmt.Sprintf("===Стеллаж %s\n", rack)
		for _, order := range orders {
			for _, product := range order.Products {
				if !strings.EqualFold(product.BasicRack, rack) {
					continue
				}
				output += fmt.Sprintf("%s (id=%d)\nзаказ %d, %d шт\n", product.Name, product.ID, order.ID, product.Quantity)
				if len(product.AdditionalRacks) > 0 {
					output += fmt.Sprintf("доп стеллаж: %s\n", product.AdditionalRacks)
				}
				output += "\n"
			}
		}
	}

	fmt.Println(output)
}