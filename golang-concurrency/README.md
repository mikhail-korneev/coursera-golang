# Concurrency in Go

Coursera "Concurrency in Go" assignments.

## Module 1: Moore's law explanation

Define Moore’s law and explain why it has now stopped being true. Be sure to describe all of the physical limitations that have prevented Moore’s law from continuing to be true.

## Module 2: Race condition

Write two goroutines which have a race condition when executed concurrently. Explain what the race condition is and how it can occur.

## Module 3: Sort an array

Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.

## Module 4: Dining Philosophers

Implement the dining philosopher’s problem with the following constraints:
1.  There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
2.  Each philosopher should eat only 3 times.
3.  The philosophers pick up the chopsticks in any order, not lowest-numbered first (which is Dijkstra's solution).
4.  No more than 2 philosophers are allowed to eat concurrently.
5.  Each philosopher is numbered, 1 through 5.
6.  When a philosopher starts eating (after it has obtained necessary locks) it prints "starting to eat [number]" on a line by itself, where [number] is the number of the philosopher.
7.  When a philosopher finishes eating (before it has released its locks) it prints "finishing eating [number]" on a line by itself, where [number] is the number of the philosopher.