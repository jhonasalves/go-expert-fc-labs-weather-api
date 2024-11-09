package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jhonasalves/go-expert-fc-labs-weather-api/configs"
	"github.com/jhonasalves/go-expert-fc-labs-weather-api/internal/handler"
	"github.com/jhonasalves/go-expert-fc-labs-weather-api/internal/infra"
	"github.com/jhonasalves/go-expert-fc-labs-weather-api/internal/usecase"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	weatherAPIClient := infra.WeatherAPIClient{APIKey: configs.WeatherAPIKey}
	viaCepClient := infra.ViaCepClient{}

	weatherRepository := infra.NewWeatherRepository(weatherAPIClient, viaCepClient)

	weatherUseCase := usecase.NewWeatherUseCase(weatherRepository)
	locationUseCase := usecase.NewLocationUseCase(weatherRepository)

	weatherHandler := handler.NewWeatherHandler(weatherUseCase, locationUseCase)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/weather/{zipCode}", weatherHandler.GetWeatherByZip)

	http.ListenAndServe(":8080", r)
}
