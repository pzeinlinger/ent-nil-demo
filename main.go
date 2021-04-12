package main

import (
	"context"
	"demo/ent"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	zlog, _ := zap.NewDevelopment()
	log := zlog.Sugar()
	defer log.Sync()

	ctx, cancel := context.WithCancel(context.Background())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		log.Infof("received signal: %s", sig)
		log.Info("initiating shutdown...")
		cancel()
	}()

	client, err := ent.Open("sqlite3", "file:ent.sqlite?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	note := client.Note.Create().SetText("Test").SaveX(ctx)
	if note.ExternalReference != nil {
		log.Fatalf("not nil on create")
	}

	notes := client.Note.Query().AllX(ctx)
	for _, note := range notes {
		if note.ExternalReference != nil {
			log.Fatalf("not nil on query")
		}
	}
}
