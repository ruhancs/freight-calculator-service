package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"freigth_service/internal/application/dto"
	"freigth_service/internal/application/usecase"
	"freigth_service/internal/domain/entity"
	"freigth_service/internal/infra/repository"
	"freigth_service/pkg/kafka"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db,err := sql.Open("mysql", os.Getenv("DB_SOURCE_LOCAL"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	msgChan := make(chan *ckafka.Message)
	topics := []string{"routes"}
	servers := os.Getenv("KAFKA_SERVER")
	go kafka.Consume(topics,servers,msgChan)

	routeRepo := repository.NewRouteRepository(db)
	freight := entity.NewFreight(10)

	createRouteUseCacse := usecase.NewCreateRouteUseCase(routeRepo,freight)
	changeRouteStatusUseCase := usecase.NewChangeStatusRouteUseaCase(routeRepo)

	for msg := range msgChan {
		createRouteInput := dto.InputCreateRouteDto{}
		json.Unmarshal(msg.Value,&createRouteInput)

		switch createRouteInput.Event {
		case "RouteCreated":
			output,err := createRouteUseCacse.Execute(createRouteInput)
			if err != nil{
				fmt.Println(err)
			}
			fmt.Println(output)
		case "RouteStarted", "RouteFinished":
			changeStatusInput := dto.InputChangeStatusRouteDto{}
			json.Unmarshal(msg.Value,&changeStatusInput)
			output,err := changeRouteStatusUseCase.Execute(changeStatusInput)
			if err != nil{
				fmt.Println(err)
			}
			fmt.Println(output)
		}
	}
}