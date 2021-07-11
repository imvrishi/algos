##Advent of Code
https://adventofcode.com/

##Top Interview Questions
https://leetcode.com/problems/reverse-string/
https://leetcode.com/problems/fizz-buzz/
https://leetcode.com/problems/single-number/
https://leetcode.com/problems/maximum-depth-of-binary-tree/
https://leetcode.com/problems/move-zeroes/
https://leetcode.com/problems/sum-of-two-integers/
https://leetcode.com/problems/reverse-linked-list/
https://leetcode.com/problems/excel-sheet-column-number/
https://leetcode.com/problems/majority-element/
https://leetcode.com/problems/roman-to-integer/
https://leetcode.com/problems/delete-node-in-a-linked-list/
https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
https://leetcode.com/problems/valid-anagram/
https://leetcode.com/problems/contains-duplicate/
https://leetcode.com/problems/first-unique-character-in-a-string/
https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
https://leetcode.com/problems/missing-number/
https://leetcode.com/problems/intersection-of-two-arrays-ii/
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
https://leetcode.com/problems/merge-two-sorted-lists/
https://leetcode.com/problems/happy-number/
https://leetcode.com/problems/pascals-triangle/
https://leetcode.com/problems/climbing-stairs
https://leetcode.com/problems/symmetric-tree
https://leetcode.com/problems/maximum-subarray
https://leetcode.com/problems/power-of-three
https://leetcode.com/problems/number-of-1-bits
https://leetcode.com/problems/house-robber/
https://leetcode.com/problems/plus-one/
https://leetcode.com/problems/two-sum/
https://leetcode.com/problems/count-and-say/
https://leetcode.com/problems/remove-duplicates-from-sorted-array/
https://leetcode.com/problems/factorial-trailing-zeroes/
https://leetcode.com/problems/valid-parentheses/
https://leetcode.com/problems/linked-list-cycle/
https://leetcode.com/problems/palindrome-linked-list/
https://leetcode.com/problems/merge-sorted-array/
https://leetcode.com/problems/min-stack/
https://leetcode.com/problems/longest-common-prefix
https://leetcode.com/problems/intersection-of-two-linked-lists
https://leetcode.com/problems/implement-strstr/
https://leetcode.com/problems/sqrtx/
https://leetcode.com/problems/reverse-bits/
https://leetcode.com/problems/valid-palindrome/
https://leetcode.com/problems/rotate-array/
https://leetcode.com/problems/count-primes/
https://leetcode.com/problems/reverse-integer/
https://leetcode.com/problems/binary-tree-inorder-traversal/

##Amazon Interview Questions
https://www.glassdoor.ca/Interview/Amazon-Software-Development-Engineer-Interview-Questions-EI_IE6036.0,6_KO7,36.htm
https://leetcode.com/problems/two-sum/
https://leetcode.com/problems/add-two-numbers/
https://leetcode.com/problems/longest-substring-without-repeating-characters/
https://leetcode.com/problems/number-of-islands/
https://leetcode.com/problems/valid-parentheses/
https://leetcode.com/problems/longest-palindromic-substring/
https://leetcode.com/problems/copy-list-with-random-pointer/
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
https://leetcode.com/problems/merge-two-sorted-lists/

####Bonus Question asked by Amazon:

####From: https://www.dailycodingproblem.com/

There's a staircase with N steps, and you can climb 1 or 2 steps at a time. Given N, write a function that returns the number of unique ways you can climb the staircase. The order of the steps matters.

For example, if N is 4, then there are 5 unique ways:

1, 1, 1, 1

2, 1, 1

1, 2, 1

1, 1, 2

2, 2

What if, instead of being able to climb 1 or 2 steps at a time, you could climb any number from a set of positive integers X? For example, if X = {1, 3, 5}, you could climb 1, 3, or 5 steps at a time. Generalize your function to take in X.


Solution

It's always good to start off with some test cases. Let's start with small cases and see if we can find some sort of pattern.

N = 1: [1]

N = 2: [1, 1], [2]

N = 3: [1, 2], [1, 1, 1], [2, 1]

N = 4: [1, 1, 2], [2, 2], [1, 2, 1], [1, 1, 1, 1], [2, 1, 1]

What's the relationship?

The only ways to get to N = 3, is to first get to N = 1, and then go up by 2 steps, or get to N = 2 and go up by 1 step. So f(3) = f(2) + f(1).

Does this hold for N = 4? Yes, it does. Since we can only get to the 4th step by getting to the 3rd step and going up by one, or by getting to the 2nd step and going up by two. So f(4) = f(3) + f(2).

To generalize, f(n) = f(n - 1) + f(n - 2). That's just the Fibonacci sequence!

def staircase(n):
    if n <= 1:
        return 1
    return staircase(n - 1) + staircase(n - 2)
Of course, this is really slow (O(2N)) — we are doing a lot of repeated computations! We can do it a lot faster by just computing iteratively:

def staircase(n):
    a, b = 1, 2
    for _ in range(n - 1):
        a, b = b, a + b
    return a
Now, let's try to generalize what we've learned so that it works if you can take a number of steps from the set X. Similar reasoning tells us that if X = {1, 3, 5}, then our algorithm should be f(n) = f(n - 1) + f(n - 3) + f(n - 5). If n < 0, then we should return 0 since we can't start from a negative number of steps.

def staircase(n, X):
    if n < 0:
        return 0
    elif n == 0:
        return 1
    else:
        return sum(staircase(n - x, X) for x in X)
This is again, very slow (O(|X|N)) since we are repeating computations again. We can use dynamic programming to speed it up.

Each entry cache[i] will contain the number of ways we can get to step i with the set X. Then, we'll build up the array from zero using the same recurrence as before:

def staircase(n, X):
    cache = [0 for _ in range(n + 1)]
    cache[0] = 1
    for i in range(1, n + 1):
        cache[i] += sum(cache[i - x] for x in X if i - x >= 0)
    return cache[n]
This now takes O(N * |X|) time and O(N) space.

##Facebook Interview Questions
https://www.glassdoor.ca/Interview/Facebook-Software-Engineer-Interview-Questions-EI_IE40772.0,8_KO9,26.htm
https://leetcode.com/problems/two-sum/
https://leetcode.com/problems/number-of-islands/
https://leetcode.com/problems/valid-parentheses/
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
https://leetcode.com/problems/merge-intervals/
https://leetcode.com/problems/reverse-linked-list/

####Bonus Interview Question asked by Facebook:

Determine the 10 most frequent words given a terabyte of strings.
Solution: https://stackoverflow.com/questions/12525455/most-frequent-words-in-a-terabyte-of-data

##Google Interview Questions
https://www.careercup.com/page?pid=google-interview-questions
https://leetcode.com/problems/min-stack/
https://leetcode.com/problems/number-of-islands/
https://leetcode.com/problems/valid-parentheses/
https://leetcode.com/problems/trapping-rain-water/
https://leetcode.com/problems/merge-intervals/
https://leetcode.com/problems/next-closest-time/
https://leetcode.com/problems/word-break/
https://leetcode.com/problems/next-permutation/

####Bonus Interview Question asked by Google:

Write a function which will remove any given character from a String
Solution: http://javarevisited.blogspot.sg/2015/04/how-to-remove-given-character-from.html

##Domain Specific Questions
https://github.com/imvrishi/awesome-interview-questions
https://www.codementor.io/@nihantanu/21-essential-javascript-tech-interview-practice-questions-answers-du107p62z
https://www.toptal.com/javascript/interview-questions
http://www.thatjsdude.com/interview/js2.html
