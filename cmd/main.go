package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v9"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/nurmuhammaddeveloper/blog_db/api"
	_ "github.com/nurmuhammaddeveloper/blog_db/api/docs"
	"github.com/nurmuhammaddeveloper/blog_db/config"
	"github.com/nurmuhammaddeveloper/blog_db/storage"
)

func main() {
	cfg := config.Load(".")

	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)

	psqlConn, err := sqlx.Connect("postgres", psqlUrl)

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	fmt.Println("Configuration: ", cfg)
	fmt.Println("Connected Succesfully!")

	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Addr,
	})

	strg := storage.NewStoragePg(psqlConn)
	inMemory := storage.NewInMemoryStorage(rdb)

	apiServer := api.New(&api.RoutetOptions{
		Cfg:      &cfg,
		Storage:  strg,
		InMemory: inMemory,
	})

	err = apiServer.Run(cfg.HttpPort)
	if err != nil {
		log.Fatalf("failed to run server: %s", err)
	}

	log.Print("Server Stopped!")
}
