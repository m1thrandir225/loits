package main

import (
	"context"
	"m1thrandir225/loits/api"
	db "m1thrandir225/loits/db/sqlc"
	"m1thrandir225/loits/util"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
)

func main() {

	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal().Err(err).Msg("Cannot load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	connPool.Config().AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxUUID.Register(conn.TypeMap())
		return nil
	}

	store := db.NewStore(connPool)

	runGinServer(config, store)
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	err = server.Start(config.HTTPServerAddress)

	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}
