### Model Used:

[OpenAI](https://www.chatgpt.com)

### Model with different cities and compare outputs.


#### 1st city

- **Prompt:** 
Simulate a weather API response for Paris today.


- **Model Output:**

```json
{
  "location": {
    "city": "Paris",
    "country": "France",
    "latitude": 48.8566,
    "longitude": 2.3522,
    "timezone": "Europe/Paris"
  },
  "date": "2026-02-21",
  "current": {
    "temperature_c": 8,
    "temperature_f": 46,
    "feels_like_c": 7,
    "feels_like_f": 45,
    "condition": "Mostly cloudy with light rain possible",
    "humidity_percent": 78,
    "wind_kph": 15,
    "wind_mph": 9,
    "uv_index": 1,
    "sunrise": "08:14",
    "sunset": "17:57"
  },
  "forecast": {
    "high_temp_c": 12,
    "low_temp_c": 6,
    "high_temp_f": 53,
    "low_temp_f": 43,
    "condition_summary": "Cloudy with light rain in the afternoon and evening",
    "rain_chance_percent": 30,
    "rain_amount_mm": 0.1
  },
  "hourly": [
    {"time": "06:00", "temp_c": 7, "condition": "Cloudy", "rain_chance": 20},
    {"time": "09:00", "temp_c": 9, "condition": "Cloudy with light rain", "rain_chance": 40},
    {"time": "12:00", "temp_c": 11, "condition": "Light rain", "rain_chance": 50},
    {"time": "15:00", "temp_c": 12, "condition": "Cloudy", "rain_chance": 30},
    {"time": "18:00", "temp_c": 10, "condition": "Cloudy", "rain_chance": 25},
    {"time": "21:00", "temp_c": 8, "condition": "Cloudy", "rain_chance": 10}
  ],
  "alerts": [
    {
      "type": "Flood Warning",
      "severity": "Yellow",
      "description": "Moderate flooding risk; keep updated with forecasts.",
      "effective_from": "2026-02-21T00:00:00+01:00",
      "effective_to": "2026-02-22T00:00:00+01:00"
    }
  ]
}

```


- **Prompt:** 
Simulate a weather API response for London today.


- **Model Output:**

```json

{
  "location": {
    "city": "London",
    "country": "United Kingdom",
    "latitude": 51.5074,
    "longitude": -0.1278,
    "timezone": "Europe/London"
  },
  "date": "2026-02-21",
  "current": {
    "temperature_c": 6,
    "temperature_f": 43,
    "feels_like_c": 4,
    "feels_like_f": 39,
    "condition": "Overcast with occasional drizzle",
    "humidity_percent": 82,
    "wind_kph": 18,
    "wind_mph": 11,
    "uv_index": 1,
    "sunrise": "07:52",
    "sunset": "17:15"
  },
  "forecast": {
    "high_temp_c": 9,
    "low_temp_c": 3,
    "high_temp_f": 48,
    "low_temp_f": 37,
    "condition_summary": "Cloudy with light drizzle through the day",
    "rain_chance_percent": 40,
    "rain_amount_mm": 0.3
  },
  "hourly": [
    {"time": "06:00", "temp_c": 4, "condition": "Cloudy", "rain_chance": 20},
    {"time": "09:00", "temp_c": 6, "condition": "Light drizzle", "rain_chance": 50},
    {"time": "12:00", "temp_c": 8, "condition": "Cloudy with drizzle", "rain_chance": 40},
    {"time": "15:00", "temp_c": 9, "condition": "Overcast", "rain_chance": 30},
    {"time": "18:00", "temp_c": 7, "condition": "Cloudy", "rain_chance": 20},
    {"time": "21:00", "temp_c": 5, "condition": "Cloudy", "rain_chance": 10}
  ],
  "alerts": [
    {
      "type": "Wind Advisory",
      "severity": "Yellow",
      "description": "Moderate winds expected; secure loose objects outdoors.",
      "effective_from": "2026-02-21T00:00:00+00:00",
      "effective_to": "2026-02-21T23:59:59+00:00"
    }
  ]
}
```



- **Prompt:** 
Simulate a weather API response for Lagos today.



- **Model Output:**

```json
{
  "location": {
    "city": "Lagos",
    "country": "Nigeria",
    "latitude": 6.5244,
    "longitude": 3.3792,
    "timezone": "Africa/Lagos"
  },
  "date": "2026-02-21",
  "current": {
    "temperature_c": 31,
    "temperature_f": 88,
    "feels_like_c": 35,
    "feels_like_f": 95,
    "condition": "Partly cloudy with high humidity",
    "humidity_percent": 78,
    "wind_kph": 12,
    "wind_mph": 7,
    "uv_index": 10,
    "sunrise": "06:39",
    "sunset": "18:30"
  },
  "forecast": {
    "high_temp_c": 33,
    "low_temp_c": 27,
    "high_temp_f": 91,
    "low_temp_f": 81,
    "condition_summary": "Hot, partly cloudy, chance of isolated showers",
    "rain_chance_percent": 20,
    "rain_amount_mm": 1.0
  },
  "hourly": [
    {"time": "06:00", "temp_c": 28, "condition": "Partly cloudy", "rain_chance": 10},
    {"time": "09:00", "temp_c": 30, "condition": "Partly cloudy", "rain_chance": 15},
    {"time": "12:00", "temp_c": 32, "condition": "Mostly sunny", "rain_chance": 5},
    {"time": "15:00", "temp_c": 33, "condition": "Partly cloudy", "rain_chance": 20},
    {"time": "18:00", "temp_c": 31, "condition": "Cloudy with isolated showers", "rain_chance": 25},
    {"time": "21:00", "temp_c": 29, "condition": "Cloudy", "rain_chance": 10}
  ],
  "alerts": [
    {
      "type": "Heat Advisory",
      "severity": "Yellow",
      "description": "High temperatures with strong UV index; stay hydrated and limit sun exposure.",
      "effective_from": "2026-02-21T00:00:00+01:00",
      "effective_to": "2026-02-21T18:30:00+01:00"
    }
  ]
}

```

# Comparison

The output of all the prompts maintained the syntatic characters of JSON, but the dfference is temprature results of the different cities.