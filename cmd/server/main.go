package main

import (
	"context"
	noteV1 "github.com/AndrewEminov/auth/internal/api/note_v1"
	noteRepository "github.com/AndrewEminov/auth/internal/repository/note"
	noteService "github.com/AndrewEminov/auth/internal/service/note"
	desc "github.com/AndrewEminov/auth/pkg/note_v1"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const grpcPort = ":50051"

func main() {
	ctx := context.Background()

	list, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to get lisner: %s", err.Error())
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pgCfg, err := pgxpool.ParseConfig("host=localhost port=54321 dbname=note user=note-user password=note-password sslmode=disable")
	if err != nil {
		log.Fatalf("failed to get db config: %s", err.Error())
	}

	dbc, err := pgxpool.ConnectConfig(ctx, pgCfg)
	if err != nil {
		log.Fatalf("failed to get db connection: %s", err.Error())
	}
	defer dbc.Close()

	err = dbc.Ping(ctx)
	if err != nil {
		log.Fatalf("ping error: %s", err.Error())
	}

	noteRepo := noteRepository.NewRepository(dbc)
	noteSrv := noteService.NewService(noteRepo)
	desc.RegisterNoteV1Server(s, noteV1.NewImplementation(noteSrv))

	err = s.Serve(list)
	if err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}
}
