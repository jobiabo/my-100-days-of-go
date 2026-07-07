# Exercise 3: Chaining Prompts for Agent Flow
#### Objective: Combine multiple prompts to simulate a tool-using agent.
## Chained process and outputs. 

- The Agent project a prompt to the user:
  "What are we doing today"
- Simulated User input:
  "What is the temperature in london"
- Agent Tool detection.
   
   Then Agent tool will have to format user input in this way
  * Request type: Weather information 
  * Location Identified: London
  * Request Data: Current temperature and weather condition.
     The agent understand the users intents. It does not have the required data so it has to make a call to an external tool.

  * The Agent Tool simulates a virtual request to a weather API using the extracted location with a API request
   **Simulated API Request**

```json

{
  "endpoint": "GET /weather/current",
  "parameters": {
    "city": "London",
    "units": "metric"
  }
}

```
 * The API responds with the json format given to it

  **API response**

```json 

{
  "city": "London",
  "temperature": 18,
  "unit": "C",
  "condition": "Partly Cloudy",
  "humidity": 72
}
```
 * The Agent tool formats the API final result in json format to a natural language.

    **Agent Response**
    
    - from the API json response, the output was not user friendly, so the agent will have to reformat the API response to a user friendly language.

    From the JSON response:

Temperature = 18°C

Weather condition = Partly Cloudy

Humidity = 72%


## Chanined Process documentation.

1. The agent prompt for a user input.
2. The agent detects user input.
3. The agent simulates a call to a weather API.
4. The API responded with a json formatted data containing, temperature, weather condition, and humidity.
5. The agent tool converts raw json format to a natural language.
6. The final output will be the result formatted by agent and display to the user in natural language.
