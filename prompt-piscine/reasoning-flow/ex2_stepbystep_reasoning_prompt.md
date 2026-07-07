# Exercise 2: Step-by-Step Reasoning Prompt
#### Objective: Use structured prompting to encourage reasoning.

### Model Used:
   [openAI](https://www.chatgpt.com)

- 1. Choose a math or coding problem (e.g., "Calculate the factorial of 5").
###### Math problem:
 "Calculate the factorial of 7"

- 2. Prompt the model: "Let's think step by step before answering."

###### Prompt 
  Let's think step by step before answering.
  "Calculate the factorial of 7"


###### Model Output without reasoning instruction:
The factorial of 7 is:

[
7! = 7 \times 6 \times 5 \times 4 \times 3 \times 2 \times 1
]

[
7! = 5040
]

So, **7! = 5040**.


 
###### Model Output with reasoning instruction:

   To calculate **7! (7 factorial)**:

[
7! = 7 \times 6 \times 5 \times 4 \times 3 \times 2 \times 1
]

Factorial means multiplying the number by every positive integer below it down to 1.

### ✅ Final Answer:

**7! = 5040**

## Comparison and confirmation:
The prompt without the reasoning instruction output was been abstract. The model gave a little visual without explaining in details WHILE, The prompt with reasonning instruction was detailed in every step starting detailing how the multiplication was processed before the final answer.
