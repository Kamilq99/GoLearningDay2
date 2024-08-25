### Goroutines

1. **Basic Goroutines**: Write a program that runs two goroutines. Each goroutine should print a number from 1 to 5 with a delay. Use time.Sleep to simulate the delay.

2. **Sync Goroutines**: Write a program that runs two goroutines, each adding numbers to a common counter. Use sync.Mutex to ensure that only one goroutine at a time modifies the counter.

3. **Channel and Goroutines**: Write a function that runs a goroutine that calculates the sum of numbers from 1 to 10 and sends the result through a channel. The main function should receive the result through the channel and print it.

4. **Goroutine and Select**: Implement a function that runs two goroutines, each sending data through different channels. Use select in the main function to receive data from both channels.

5. **Goroutine and Buffers**: Write a program that uses a buffered channel to communicate between goroutines. One goroutine should send data to the channel, and the other goroutine should receive and print it.

6. **Goroutine and sync.WaitGroup**: Implement a program that starts multiple goroutines that perform different tasks. Use sync.WaitGroup to wait for all goroutines to finish before terminating the program.

7. **Goroutine and Context**: Write a function that starts a goroutine that runs for a specified amount of time (e.g. 5 seconds). Use context.Context to terminate the goroutine if the context is canceled.

8. **Goroutine with Timeout**: Implement a function that starts a goroutine that performs a long-running operation (simulated by time.Sleep). The function should have a timeout after which the goroutine is terminated.

9. **Gathering results from goroutines**: Write a program that runs several goroutines, each calculating different values ​​and returning them through channels. The main function should collect the results and sum them.

10. **Deadlock and goroutines**: Create an example program that has a deadlock due to incorrect use of channels and goroutines. Then fix the problem so that the program works correctly.