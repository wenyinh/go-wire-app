package app

import (
	"context"
	"fmt"
	"github.com/wenyinh/go-wire-app/pkg/utils"
	"os"
)

func Run(rootCtx context.Context) {

	config, err := utils.LoadConfiguration()
	if err != nil {
		fmt.Printf("load configuration error: %s\n", err.Error())
		os.Exit(1)
	}

	app, cleanup, err := InitializeApp(config)
	if err != nil {
		fmt.Printf("failed to initialize: %s\n", err.Error())
		os.Exit(1)
	}
	defer cleanup()
	fmt.Println("[INFO] Entrypoint: starting app...")
	app.Serve(rootCtx)
	fmt.Println("[INFO] Entrypoint: exiting app...")
}
