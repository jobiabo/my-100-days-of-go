# STEP 1 THE MODIFIED PALINDROME FUNCTION I CREATED TO IGNORE SPACES AND PUNCTUATIONS, BE CASE SENSITIVE, AND RETURN THE POSITION WHERE THE STRING STOPS BEING A PALINDROME (IF NOT ONE)


````py
term = input("enter a term: ").lower().Upper(). ""()
 if term == term[::-1]:
   print("palindrome")
 else:
   print("not a palindrome")





 # STEP 2 AI REVEIEW ON MY FUNCTION ABOVE

 ## PROPER CODE REDEFINED BY THE AI

```py


import re
def Check_palindrome(str):
  text = re.sub(r'[^a-zA-Z0-9]', "", str)
  left = 0
  right = len(text) -1
  while left < right:
    if text[right] != text[left]:
      return False, left
      left += 1
      right -= 1
      return True, -1

def main():
  word = input("enter a word: ")
  is_palindrome, position = Check_palindrome(word)
  if is_palindrome:
    print("Palindrome")
  else:
    print(f" Not palindrome. Fails at position {position}")

main()


````py
## Did I Miss Anything?

Ai explained that my fuction misses the aspect of ignoring spaces, mixed letters and numbers(alphanum), and punctuations.

## Can It Be More Efficient?

Yes, i am made to undeerstand that due to the case of misses i have, my code will not be efficient in some levels. And a better effecient code was assigned to me for study and understanding to practice.






# STEP 3 Reflect on what AI added that you didn't consider initially.

    The ai made me understand that I cannot chain .upper() and .lower() together and that will bring indentation errors cause, calling both does not make sense. I did try to run the code so as it runs through the term entered by the user ignoring the cases and space in between. The ai also examine the strength of my codes and also fixed my syntax errors. it simply says my code doesn't ignore spaces or punctuation despite I wrote the code my self. I still have a lot to learn on this by taking more time to breakdown and absorb the correction AI made for me.


    
