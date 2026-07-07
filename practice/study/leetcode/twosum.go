package main

import "fmt"

func twoSum(nums []int, target int) []int {
	result := []int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			// if nums[i] == target {
			// 	return []int{i}
			// }
			if nums[i]+nums[j] == target {
				result = append(result, i)
				result = append(result, j)

			}
		}
	}
	return result
}




func main() {
	a := []int{3, 2, 4}
	fmt.Println(twoSum(a, 6))
}



// func twoSum(nums []int, target int) []int {
//     seen := make(map[int]int)

//     for i, num := range nums {
//         complement := target - num

//         if idx, ok := seen[complement]; ok {
//             return []int{idx, i}
//         }

//         seen[num] = i
//     }

//     return nil
// }



/*what was two sum meant to teach me

The Two Sum problem is much less about adding numbers and much more about how you think about solving problems efficiently. It's often the first problem on coding interview sites because it introduces several core concepts.

Here are the main lessons.

1. Brute force vs. optimization

The most obvious solution is to compare every pair:

for each i
    for each j > i
        if nums[i] + nums[j] == target
            return

This works, but it takes O(n²) time.

The interviewer's real question is often:

"Can you do better?"

That leads you to think about what information you need to remember to avoid repeating work.

2. Using a hash map for fast lookups

Instead of asking:

"Does any future number complete this one?"

You ask:

"Have I already seen the number I need?"

For each number:

current = 7
target = 10
needed = 3

Check if 3 is already in a map.

If yes, you're done.

If no, store 7 and continue.

This teaches one of the most important programming patterns:

Trade memory for speed.

A hash map gives you approximately O(1) lookups, reducing the overall time to O(n).

3. Thinking in complements

Many beginners think:

"Which two numbers add to 10?"

Experienced programmers often think:

"I'm looking at 7, so I need 3."

That small shift—from searching for pairs to searching for the complement—appears in many algorithms involving sums, strings, and sets.

4. Choosing the right data structure

The problem demonstrates that selecting the right data structure can dramatically improve performance.

Data structure	Lookup time
Slice/array	O(n)
Hash map	O(1) average

Often, the biggest performance gains come from changing the data structure rather than writing more complex code.

5. Avoiding duplicate work

Your first version checked many pairs more than once.

For example:

(0,1)
(1,0)
(0,2)
(2,0)

Starting the inner loop at i + 1 eliminates duplicate comparisons. More generally, it teaches you to think about whether your algorithm is repeating work unnecessarily.

6. One-pass algorithms

The optimal solution processes the array only once:

2 → store it
7 → need 2 → found!

This introduces the idea of streaming through data while maintaining just enough state to solve the problem.

The broader lesson

Two Sum is really an introduction to a common interview pattern:

While iterating through data, maintain a lookup structure so future decisions become constant-time.

You'll see this pattern in many other problems, including:

Finding duplicates
Detecting cycles
Counting frequencies
Longest substring without repeating characters
Grouping anagrams
Checking set membership

Once you recognize this pattern, many seemingly different problems become much easier to approach.
*/