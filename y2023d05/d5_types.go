package y2023d05

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

type resultStruct struct {
	Part1_LowestValueOfIndividualSeeds int
	Part2_LowestValueOfSeedsInPairs    int
}

type almanacEntry struct {
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLength           int
}

func (a almanacEntry) getMappedValue(value int) int {
	if value < a.SourceRangeStart || value > a.SourceRangeStart+a.RangeLength-1 {
		return -1
	}

	return value - a.SourceRangeStart + a.DestinationRangeStart
}

type seedType int

func (s seedType) getSoilType() int {
	for i, mapEntry := range *conversionMaps["seed-to-soil map:"] {
		if mapEntry.getMappedValue(int(s)) >= 0 {
			if debug {
				log.Trace().
					Str("seed", fmt.Sprintf("%d", s)).Str("func", "getSoilType").
					Msgf("Found soil type %d for seed %d in soil map %d", mapEntry.getMappedValue(int(s)), s, i)
			}
			return mapEntry.getMappedValue(int(s))
		}
	}
	if debug {
		log.Trace().
			Str("seed", fmt.Sprintf("%d", s)).Str("func", "getSoilType").
			Msgf("No soil type found for seed %d, using same value", s)
	}
	return int(s)
}

func (s seedType) getFertilizerType() int {
	soilType := s.getSoilType()
	for i, mapEntry := range *conversionMaps["soil-to-fertilizer map:"] {
		if mapEntry.getMappedValue(int(soilType)) >= 0 {
			if debug {
				log.Trace().
					Str("seed", fmt.Sprintf("%d", s)).Str("func", "getFertilizerType").
					Msgf("Found fertilizer type %d for soil type %d in fertilizer map %d", mapEntry.getMappedValue(int(soilType)), soilType, i)
			}
			return mapEntry.getMappedValue(int(soilType))
		}
	}
	if debug {
		log.Trace().
			Str("seed", fmt.Sprintf("%d", s)).Str("func", "getFertilizerType").
			Msgf("No fertilizer type found for soil type %d, using same value", soilType)
	}
	return soilType
}

func (s seedType) getWaterType() int {
	fertilizerType := s.getFertilizerType()
	for i, mapEntry := range *conversionMaps["fertilizer-to-water map:"] {
		if mapEntry.getMappedValue(int(fertilizerType)) >= 0 {
			if debug {
				log.Trace().
					Str("seed", fmt.Sprintf("%d", s)).Str("func", "getWaterType").
					Msgf("Found water type %d for fertilizer type %d in water map %d", mapEntry.getMappedValue(int(fertilizerType)), fertilizerType, i)
			}
			return mapEntry.getMappedValue(int(fertilizerType))
		}
	}
	if debug {
		log.Trace().
			Str("seed", fmt.Sprintf("%d", s)).Str("func", "getWaterType").
			Msgf("No water type found for fertilizer type %d, using same value", fertilizerType)
	}
	return fertilizerType
}

func (s seedType) getLightType() int {
	waterType := s.getWaterType()
	for i, mapEntry := range *conversionMaps["water-to-light map:"] {
		if mapEntry.getMappedValue(int(waterType)) >= 0 {
			if debug {
				log.Trace().
					Str("seed", fmt.Sprintf("%d", s)).Str("func", "getLightType").
					Msgf("Found light type %d for water type %d in light map %d", mapEntry.getMappedValue(int(waterType)), waterType, i)
			}
			return mapEntry.getMappedValue(int(waterType))
		}
	}
	if debug {
		log.Trace().
			Str("seed", fmt.Sprintf("%d", s)).Str("func", "getLightType").
			Msgf("No light type found for water type %d, using same value", waterType)
	}
	return waterType
}

func (s seedType) getTemperatureType() int {
	lightType := s.getLightType()
	for i, mapEntry := range *conversionMaps["light-to-temperature map:"] {
		if mapEntry.getMappedValue(int(lightType)) >= 0 {
			if debug {
				log.Trace().
					Str("seed", fmt.Sprintf("%d", s)).Str("func", "getTemperatureType").
					Msgf("Found temperature type %d for light type %d in temperature map %d", mapEntry.getMappedValue(int(lightType)), lightType, i)
			}
			return mapEntry.getMappedValue(int(lightType))
		}
	}
	if debug {
		log.Trace().
			Str("seed", fmt.Sprintf("%d", s)).Str("func", "getTemperatureType").
			Msgf("No temperature type found for light type %d, using same value", lightType)
	}
	return lightType
}

func (s seedType) getHumidityType() int {
	temperatureType := s.getTemperatureType()
	for i, mapEntry := range *conversionMaps["temperature-to-humidity map:"] {
		if mapEntry.getMappedValue(int(temperatureType)) >= 0 {
			if debug {
				log.Trace().
					Str("seed", fmt.Sprintf("%d", s)).Str("func", "getHumidityType").
					Msgf("Found humidity type %d for temperature type %d in humidity map %d", mapEntry.getMappedValue(int(temperatureType)), temperatureType, i)
			}
			return mapEntry.getMappedValue(int(temperatureType))
		}
	}
	if debug {
		log.Trace().
			Str("seed", fmt.Sprintf("%d", s)).Str("func", "getHumidityType").
			Msgf("No humidity type found for temperature type %d, using same value", temperatureType)
	}
	return temperatureType
}

func (s seedType) getLocationType() int {
	humidityType := s.getHumidityType()
	for i, mapEntry := range *conversionMaps["humidity-to-location map:"] {
		if mapEntry.getMappedValue(int(humidityType)) >= 0 {
			if debug {
				log.Trace().
					Str("seed", fmt.Sprintf("%d", s)).Str("func", "getLocationType").
					Msgf("Found location type %d for humidity type %d in location map %d", mapEntry.getMappedValue(int(humidityType)), humidityType, i)
			}
			return mapEntry.getMappedValue(int(humidityType))
		}
	}
	if debug {
		log.Trace().
			Str("seed", fmt.Sprintf("%d", s)).Str("func", "getLocationType").
			Msgf("No location type found for humidity type %d, using same value", humidityType)
	}
	return humidityType
}

type seedPairRange struct {
	RangeStart    int
	NumberOfSeeds int
}
