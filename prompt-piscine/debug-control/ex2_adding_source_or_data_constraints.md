# Exercise 2: Adding Source or Data Constraints
##### Objective: Force responses to be tied to given material.

#### Model Used:
[OpenAI](https://chatgpt.com)

### 1. Write a factual question prompt.

- **factual question prompt:**

" Who is the president of Nigeria? "

- **Model Output:**

As of 2026, the President of Nigeria is Bola Ahmed Tinubu. He has been serving as the 16th President of the Federal Republic of Nigeria since 29 May 2023 after winning the 2023 presidential election. ([Wikipedia][1])

Here’s a quick summary:

* Name: Bola Ahmed Adekunle Tinubu
* Office: President of Nigeria (16th)
* Assumed Office: 29 May 2023
* Predecessor: Muhammadu Buhari ([Wikipedia][1])

Tinubu previously served as Governor of Lagos State and has been a major figure in Nigerian politics for decades. ([Wikipedia][1])

If you want more details about his background or policies, just let me know!

[1]: https://en.wikipedia.org/wiki/Bola_Tinubu?utm_source=chatgpt.com "Bola Tinubu"

### 2. Add instructions such as:

- **Prompt with Instruction:**

Cite the source in your answer, Use only the text provided below.
" Bola Ahmed Tinubu is the President of Nigeria, sworn into office on May 29, 2023. A seasoned politician, he previously served as Governor of Lagos State from 1999 to 2007, where he was widely credited with improving the state’s revenue generation and infrastructure.

A leading figure in the ruling All Progressives Congress (APC), Tinubu played a key role in shaping modern Nigerian politics. His presidency has focused on economic reforms, including fuel subsidy removal and efforts to attract foreign investment. Supporters praise his bold policy decisions, while critics express concerns about inflation and cost of living challenges.

Tinubu remains an influential and sometimes controversial leader, steering Africa’s most populous nation through a period of economic transition and reform." 
-  who is the president of nigeria? 


* MODEL OUTPUT:

The president of Nigeria is **Bola Ahmed Tinubu**.

According to the provided text:

 “Bola Ahmed Tinubu is the President of Nigeria, sworn into office on May 29, 2023.”


### 3. Provide a short passage and check that the answer stays within the provided text.

The answer of the vague and the prompt with instruction stays within the provided text. The prompt with instruction took it source from the article i added to the prompt as the text which provided the information it must use to answer the question(who is the president of nigeria). Using the source, the model has be constrained.


