# Exercise 2: Data Extraction Template
### Objective: Extract structured data from unstructured input.

#### 1 - 2. Choose an unstructured text (e.g., Extract the following fields from the text: Name, Age, Location, Occupation. Return the output in JSON format.).

- **Input:** Extract the following fields from the text: Name, Age, Location, Occupation. Return the output in JSON format "John Doe, age 29, lives in Paris and works as a software engineer."
- **Model Otput:** 

```json

{
  "Name": "John Doe",
  "Age": 29,
  "Location": "Paris",
  "Occupation": "Software Engineer"
}



```

# 3. Validate the output and ensure consistency across multiple inputs.

- **Input:** Extract the following fields from the text: Name, Age, Location, Occupation, Marita status, No. of children. Return the output in JSON format "Jack Fiender, age 29, lives in Paris and works as a software engineer with two kids."

- **Model Output:** 

```json

{
  "Name": "Jack Fiender",
  "Age": 29,
  "Location": "Paris",
  "Occupation": "Software Engineer",
  "Marital Status": null,
  "No. of Children": 2
}
```

- The output was validated with multiple inputs and consistency was ensured.