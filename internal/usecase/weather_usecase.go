package usecase

import (
	"fmt"

	"github.com/jhonasalves/go-expert-fc-labs-weather-api/internal/entity"
	"github.com/jhonasalves/go-expert-fc-labs-weather-api/internal/infra"
)

type WeatherUseCase struct {
	WeatherRepository *infra.WeatherRepository
}

func NewWeatherUseCase(weatherRepository *infra.WeatherRepository) *WeatherUseCase {
	return &WeatherUseCase{
		WeatherRepository: weatherRepository,
	}
}

func (uc *WeatherUseCase) GetWeather(city string) (*entity.Weather, error) {
	weatherData, err := uc.WeatherRepository.GetWeather(city)
	if err != nil {
		return nil, fmt.Errorf("could not fetch weather data: %v", err)
	}

	weatherData.Convert()

	return weatherData, nil
}
