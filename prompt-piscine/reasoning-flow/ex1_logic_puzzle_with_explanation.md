# Exercise 1: Logic Puzzle with Explanation
### Objective: Guide the model to solve a problem with explicit reasoning.

#### Model used:
 [openAI](https://chatgpt.com)

#### Logic used:
 - "Five people were eating apples, A finished before B, but behind C. D finished before E, but behind B. What was the finishing order?""

#### Prompt used:
- - Explain your reasoning before giving the final answer.
"Five people were eating apples, A finished before B, but behind C. D finished before E, but behind B. What was the finishing order?"" 

#### Model output: 
 - Let’s break it down step by step.

We are told:

1. **A finished before B, but behind C**

   * “Before B” → A came ahead of B
   * “Behind C” → C came ahead of A

   So from this we get:
   **C → A → B**

2. **D finished before E, but behind B**

   * “Before E” → D came ahead of E
   * “Behind B” → B came ahead of D

   So from this we get:
   **B → D → E**

Now combine both sequences:

From the first clue:
C → A → B

From the second clue (and since B connects both clues):
B → D → E

Putting everything together in order:

**C → A → B → D → E**

Final Answer:

**1st: C**
**2nd: A**
**3rd: B**
**4th: D**
**5th: E**

#### Verification.

From the output given so far by the model, there is a confirmation that there is a step by step answer, followed by a reasoning given.
