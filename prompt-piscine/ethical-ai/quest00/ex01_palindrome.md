# STEP 1

## MY PSEUDOCDE

BEGIN


  PYTHON INBUILT FUNCTION


  TAKE "WORD INPUT"


  CONVERTS WORD INTO LOWER CASE


  WORD = WORD REVERSED


  IF "WORD INPUT" = "WORD REVERSED OUTPUT"


  PRINT "PALINDROME"


  ELSE


  PRINT "NOT A PALINDROME"


END 












# PSEUDOCODE IMPLEMENTED TO PYTHON



term = input("Enter a term: ").lower()


//built in function to take a user input


  if term == term[::-1]:


//runs the conditions and slices the user input being a string and print true if conditions are met.


   print("Palindrome")

//prints if conditions are met


else:
   print("not a palindrome")

//vice versa on the prints



   # THE PYTHON SCRIPTS WITHOUT A COMMENT

   term = input("enter a term: ").lower()
  
if term == term[::-1]:
  print("palindrome")
else:
  print("not a palindrome")  


# STEP 2
## USING AI TO LEARN 
 
  I gave Ai this
  
   "I wrote a palindrome function. Here's my code:
term = input("enter a term: ").lower()
  
if term == term[::-1]:
  print("palindrome")
else:
  print("not a palindrome")

  ## Questions and answers structured for Ai  
*What's the time complexity?*

**answer** 

Term[::-1] creates a reversed copy of the string and the time complexity is linear i.e O(n)

Comparing two strings of length n is O(n)

Space complexity is O(n) because reversed string is new copy

The overall time complexity is O(n)


### What edge cases might I miss?

##### answer

Ai confirmed that my function works perfectly for a simple single word in palindrome but here are my edge cases

1. Empty strings (" ") is identified and executed successfully.

2. simple character ("a") is executed successfully 

3. Space and punctuation with examples such as "nurses run", " a man, a plan, a canal: panama" which are palindrome but with my function, it will not be acceptable as a palindrome.


Are there better approaches?


#### answer 


Yes the two pointer approaches was suggested by the ai in which the reversed string that i know looks more efficient to me.


**And it's response was that my function is linear, which is actually pretty efficient for small to medium strings. that my function will fails when it comes to be in case sensitive, ignoring spaces and punctuations and also number symbols.**




# Step 3 - REFLECTION
1. What did you learn from solving it before asking AI?

   I learnt the fundamentals of python on a different web gradually before asking the Ai

2. How is your understanding different now?


   my understanding is different now on the level that there are various ways in which the palindrome functions can be determined in terms of (let me use the word "typescripts)

3. Could you now write similar functions (e.g., reverse a string) without help?


   Yes, using the simple and basic python way
