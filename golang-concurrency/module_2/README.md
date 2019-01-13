## Race condition

Race condition is defined as a problem where the outcome of a program depends on the interleaving of machine instructions. This may occur due to the communication between tasks without proper synchronization between them.

The program demonstrates the work of two goroutines - producer and consumer. Producer generates a random int and assigns it to the shared variable "sharedInt". Consumer reads the current "sharedInt" value and prints it to the console. These two goroutines communicate with each other through the "sharedInt" variable.

So, these two goroutines use a shared resource, they are not synchronized, which leads to a race condition.