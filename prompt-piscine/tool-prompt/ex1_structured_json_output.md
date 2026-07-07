# Exercise 1: Structured JSON Output
## Objective: Create prompts that produce structured machine-readable output.



- **Prompt 1:**
List three cities with their populations.

- **Model Output:**
Here are three major cities and their approximate populations:

1. * *Tokyo* – ~37 million (metropolitan area)
2. * *New York City* – ~8.5 million (city proper)
3. * *Paris* – ~2.1 million (city proper)
Note: Populations are approximate and based on recent estimates.

- **Improved prompts:**
* List three cities with their populations.
* Return the following in JSON format: [{city: string, population: number}].


- **Model Output:**

```json
[
  { "city": "Tokyo", "population": 37000000 },
  { "city": "New York City", "population": 8500000 },
  { "city": "Paris", "population": 2100000 }
]
```

- **Prompt 2:**

List three books and their authors.

- **Model Output:**
Here are three well-known books and their authors:

1. **1984* – by George Orwell
2. **Pride and Prejudice* – by Jane Austen
3. **To Kill a Mockingbird* – by Harper Lee


- **Improved Prompt:**

* List three books and their authors.
* Return the following in JSON format: [{author: string, year of publish: number}].

- **Model Output:**


```json
[
  { "author": "George Orwell", "year of publish": 1949 },
  { "author": "Jane Austen", "year of publish": 1813 },
  { "author": "Harper Lee", "year of publish": 1960 }
]
```


- **Prompt 3:**
List three universities and the year they were founded.

- **Model Output:**
Here are three well-known universities and their founding years:

1. **Harvard University* – Founded in 1636
2. **University of Oxford* – Teaching began in 1096
3. **Stanford University* – Founded in 1885


- **Improved Prompt:**
* List three universities and the year they were founded.
* Return the following in JSON format: [{university: string, year: number}].

- **Model Output:**

```json

[
  { "university": "Harvard University", "year": 1636 },
  { "university": "University of Oxford", "year": 1096 },
  { "university": "Stanford University", "year": 1885 }
]
``` 


# Reflection
The prompt to the model returned expected output, all output from the multiple examples resulted in JSON format as prompted to the model.