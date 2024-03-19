package main

import (
	"github.com/IlhamSetiaji/go-lms/internal/config"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)
	validate := config.NewValidator(viperConfig)
	app := config.NewGin(viperConfig)
	producer, _ := config.NewRabbitMQProducer(viperConfig, log)

	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		App:      app,
		Log:      log,
		Validate: validate,
		Config:   viperConfig,
		Producer: producer,
	})

	// webPort := strconv.Itoa(viperConfig.GetInt("web.port"))
	err := app.Run()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
