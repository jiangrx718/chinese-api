package server

import (
	"crm/gopkg/gins"
	"crm/gopkg/graceful"
	"crm/handler/api"
	"net/http"

	"github.com/urfave/cli/v2"
)

func Run(*cli.Context) error {
	go func() {
		_ = http.ListenAndServe(":8999", nil)
	}()

	server := gins.NewHttpServer(":8080")
	server.RegisterHandler(
		api.NewHandler,
	)
	graceful.Start(server)
	graceful.Wait()
	return nil
}
