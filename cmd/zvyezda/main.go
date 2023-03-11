package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	socialapp_v1 "github.com/mezmerizxd/zvyezda/api/social_app/v1"
	"github.com/mezmerizxd/zvyezda/data"
	socialAppData "github.com/mezmerizxd/zvyezda/data/social_app"
	"github.com/mezmerizxd/zvyezda/features"
	socialAppFeatures "github.com/mezmerizxd/zvyezda/features/social_app"
	"github.com/mezmerizxd/zvyezda/pkg/server"
)

func main() {
	port := ":3000"
	
	// Data
	sad := socialAppData.New(&socialAppData.Config{})
	data := data.New(&data.Config{
		SocialApp: sad,
	})

	// Features
	saf := socialAppFeatures.New(&socialAppFeatures.Config{
		Data: data,
	})
	features := features.New(&features.Config{
		SocialApp: saf,
	})

	srv := server.New(port, &socialapp_v1.Config{
		Features: features,
	})

	quitChannel := make(chan bool, 1)

	go func() {
		log.Println("Starting at http://localhost" + port)
		if err := srv.Start(); err != nil {
			select {
			case <-quitChannel:
				return
			default:
				log.Println("Failed to start server: " + err.Error())
			}
		}
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc
	quitChannel <- true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := srv.Stop(ctx); err != nil {
		log.Println("Failed to stop server: " + err.Error())
	}

	log.Println("exiting")
}