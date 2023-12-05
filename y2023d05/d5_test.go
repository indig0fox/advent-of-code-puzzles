package y2023d05

import (
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func TestSeed(t *testing.T) {

	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	})
	log.Logger = log.Level(zerolog.TraceLevel)
	getInput("input_test.txt")

	type seedResult struct {
		soil        int
		fertilizer  int
		water       int
		light       int
		temperature int
		humidity    int
		location    int
	}
	testSeeds := []seedType{79, 14, 55, 13}
	want := map[seedType]seedResult{
		79: {
			soil:        81,
			fertilizer:  81,
			water:       81,
			light:       74,
			temperature: 78,
			humidity:    78,
			location:    82,
		},
		14: {
			soil:        14,
			fertilizer:  53,
			water:       49,
			light:       42,
			temperature: 42,
			humidity:    43,
			location:    43,
		},
		55: {
			soil:        57,
			fertilizer:  57,
			water:       53,
			light:       46,
			temperature: 82,
			humidity:    82,
			location:    86,
		},
		13: {
			soil:        13,
			fertilizer:  52,
			water:       41,
			light:       34,
			temperature: 34,
			humidity:    35,
			location:    35,
		},
	}
	results := make(map[seedType]seedResult)
	for _, seed := range testSeeds {
		results[seed] = seedResult{
			soil:        seed.getSoilType(),
			fertilizer:  seed.getFertilizerType(),
			water:       seed.getWaterType(),
			light:       seed.getLightType(),
			temperature: seed.getTemperatureType(),
			humidity:    seed.getHumidityType(),
			location:    seed.getLocationType(),
		}
	}

	for seed, result := range results {
		if result.soil != want[seed].soil {
			t.Errorf("seed %d: soil = %d, want %d", seed, result.soil, want[seed].soil)
		}
		if result.fertilizer != want[seed].fertilizer {
			t.Errorf("seed %d: fertilizer = %d, want %d", seed, result.fertilizer, want[seed].fertilizer)
		}
		if result.water != want[seed].water {
			t.Errorf("seed %d: water = %d, want %d", seed, result.water, want[seed].water)
		}
		if result.light != want[seed].light {
			t.Errorf("seed %d: light = %d, want %d", seed, result.light, want[seed].light)
		}
		if result.temperature != want[seed].temperature {
			t.Errorf("seed %d: temperature = %d, want %d", seed, result.temperature, want[seed].temperature)
		}
		if result.humidity != want[seed].humidity {
			t.Errorf("seed %d: humidity = %d, want %d", seed, result.humidity, want[seed].humidity)
		}
		if result.location != want[seed].location {
			t.Errorf("seed %d: location = %d, want %d", seed, result.location, want[seed].location)
		}
	}

}
