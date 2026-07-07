# Exercise 2: Role-Based Mentoring and Feedback
#### Objective: Use role-based prompts for personalized feedback.
## 1 - 2. Python function with bug.

import re
def CheckPalindrome(str):
  cleaned = re.sub(r'[^a-zA-z0-9]', "", str)
  if cleaned == cleaned[::-1]:
      stand += 1
      check if -= 56
    return True
  else:
    return False  

def main():
  word = input("enter a word: ")
  if CheckPalindrome(word):
    print("Palindrome")
  else:
    print("Not Palindrome")

main()          


## 3. Prompt the model: "Act as a mentor reviewing this code. Provide corrections and explain improvements."

- **Model used:** [openAI](https://chatgpt.com)

- **Prompt for model:**  Act as a mentor reviewing this code. Provide corrections and explain improvements.

import re
def CheckPalindrome(str):
  cleaned = re.sub(r'[^a-zA-z0-9]', "", str)
  if cleaned == cleaned[::-1]:
      stand += 1
      check if -= 56
    return True
  else:
    return False  

def main():
  word = input("enter a word: ")
  if CheckPalindrome(word):
    print("Palindrome")
  else:
    print("Not Palindrome")

main()          

- **Model output:** Got it! Let’s go through your code carefully. I’ll point out **errors**, **bad practices**, and **how to improve it**, line by line.

---

### **Your code**

```python
import re
def CheckPalindrome(str):
  cleaned = re.sub(r'[^a-zA-z0-9]', "", str)
  if cleaned == cleaned[::-1]:
      stand += 1
      check if -= 56
    return True
  else:
    return False  

def main():
  word = input("enter a word: ")
  if CheckPalindrome(word):
    print("Palindrome")
  else:
    print("Not Palindrome")

main()
```

---

### **Problems & Corrections**

1. **Function naming and parameter naming**

   ```python
   def CheckPalindrome(str):
   ```

   * Python convention: function names should be **snake_case** → `check_palindrome`.
   * Avoid using `str` as a variable name because it **overrides the built-in `str()` type**.

   ✅ Corrected:

   ```python
   def check_palindrome(word):
   ```

2. **Regex cleaning**

   ```python
   cleaned = re.sub(r'[^a-zA-z0-9]', "", str)
   ```

   * Typo in regex: `A-z` includes **extra characters**, use `A-Z` (uppercase) properly.
   * Also, make it **case-insensitive** by converting to lower-case.

   ✅ Corrected:

   ```python
   cleaned = re.sub(r'[^a-zA-Z0-9]', '', word).lower()
   ```

3. **Invalid lines inside `if`**

   ```python
       stand += 1
       check if -= 56
   ```

   * These lines don’t make sense in this context and will throw errors. Remove them.

4. **Indentation**

   ```python
    return True
   ```

   * Original code had inconsistent indentation. Python is strict about this.

5. **Optional improvement: two-pointer method**

   * Using `cleaned == cleaned[::-1]` works but **reversing a string creates a new string**, so O(n) memory.
   * For large strings, a **two-pointer approach** is more efficient.

---

### **Corrected & Improved Version**

```python
import re

def check_palindrome(word):
    # Remove non-alphanumeric characters and convert to lowercase
    cleaned = re.sub(r'[^a-zA-Z0-9]', '', word).lower()
    
    # Two-pointer check (more efficient for long strings)
    left, right = 0, len(cleaned) - 1
    while left < right:
        if cleaned[left] != cleaned[right]:
            return False
        left += 1
        right -= 1
    return True

def main():
    word = input("Enter a word: ")
    if check_palindrome(word):
        print("Palindrome")
    else:
        print("Not Palindrome")

main()
```

---

### **Why this is better**

1. **Readability:** Function names follow Python conventions. Variable names are clear.
2. **Bug-free:** Removed invalid lines that would crash the program.
3. **Robust cleaning:** Ignores spaces, punctuation, and case.
4. **Efficient:** Two-pointer approach avoids creating a reversed copy of the string.
5. **Scalable:** Works well even with very long input strings.

---

If you want, I can also **show an even cleaner one-liner version** using Python built-ins that’s compact but still readable.

Do you want me to do that?



## 4. Reflection/Comments

Prompting the model to act as coding mentor, Yes, it affects the style of the out which was structured, educational, and supportive, on a mentorship tone. It made emphasis on every aspect of the function i gave it with a bug.