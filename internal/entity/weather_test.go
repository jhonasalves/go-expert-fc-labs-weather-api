package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenCelsius_WhenConvertToKelvinAndFahrenheit_ThenShouldReturnCorrectValues(t *testing.T) {
	weather := &Weather{TempC: 25.0}
	weather.Convert()

	assert.Equal(t, 77.0, weather.TempF)
	assert.Equal(t, 298.15, weather.TempK)
}

func TestGivenFreezingPoint_WhenConvertToKelvinAndFahrenheit_ThenShouldReturnCorrectValues(t *testing.T) {
	weather := &Weather{TempC: 0.0}
	weather.Convert()

	assert.Equal(t, 32.0, weather.TempF)
	assert.Equal(t, 273.15, weather.TempK)
}

func TestGivenBoilingPoint_WhenConvertToKelvinAndFahrenheit_ThenShouldReturnCorrectValues(t *testing.T) {
	weather := &Weather{TempC: 100.0}
	weather.Convert()

	assert.Equal(t, 212.0, weather.TempF)
	assert.Equal(t, 373.15, weather.TempK)
}
