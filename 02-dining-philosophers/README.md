# The Dining Philosophers Problem

## Problem Description

The Dining Philosophers problem is a classic computer science problem that illustrates synchronization issues and techniques for resolving them. It was originally formulated by Edsger Dijkstra in 1965 to illustrate the challenges of avoiding deadlock and resource starvation in concurrent systems.

### The Setup

- Five philosophers sit around a circular table
- Each philosopher has a plate of spaghetti in front of them
- Between each pair of philosophers is a single fork (5 forks total)
- A philosopher needs two forks to eat (both left and right fork)
- Each philosopher alternates between thinking and eating

### The Challenge

The challenge is to design a solution where:

1. No philosopher will starve (they all get to eat eventually)
2. No deadlocks occur (where all philosophers are stuck waiting forever)
3. Maximum concurrency is maintained (as many philosophers can eat simultaneously as possible)

## Common Issues

### Deadlock Scenario

If each philosopher picks up their left fork simultaneously and waits for their right fork, a deadlock occurs as no philosopher can obtain their second fork.

### Resource Starvation

Poor synchronization strategies might lead to some philosophers never getting a chance to eat while others monopolize the forks.

## Solution Strategies

1. **Resource Hierarchy Solution**

   - Assign numbers to forks (0 to 4)
   - Philosophers must pick up lower-numbered fork first
   - Prevents circular wait condition

2. **Arbitrator Solution**

   - Use a waiter (mutex) to control access
   - Philosophers must ask permission before picking up forks
   - Waiter ensures only N-1 philosophers can attempt to eat simultaneously

3. **Chandry/Misra Solution**
   - Forks are requested with messages
   - Clean/dirty fork states track usage
   - Philosophers pass forks to neighbors upon request

## Implementation

The implementation in this directory demonstrates the Resource Hierarchy solution using:

- Mutex locks for fork synchronization
- Goroutines for concurrent philosopher behavior
- Channels for communication and synchronization

## Running the Program

```bash
go run main.go
```

## Expected Output

The program will show philosophers thinking and eating in a concurrent manner, demonstrating proper resource sharing and deadlock prevention.

```bash
Philosopher 1 is thinking
Philosopher 2 is eating
Philosopher 4 is thinking
Philosopher 3 is eating
Philosopher 5 is thinking
```

## Testing

The implementation includes comprehensive tests to verify the correct behavior of the dining philosophers solution.

### Test Cases

1. **Basic Dining Test**

   ```go
   func Test_dine(t *testing.T) {
       // Tests multiple iterations with zero delays
       // Verifies that all 5 philosophers complete their meal
   }
   ```

2. **Variable Timing Test**
   ```go
   func Test_dineWithVaryingDelays(t *testing..T) {
       // Tests the solution with different time delays:
       // - Zero delay
       // - Quarter second delay (250ms)
       // - Half second delay (500ms)
   }
   ```

### What the Tests Verify

- All philosophers successfully complete their meals
- The solution works under different timing conditions
- No deadlocks occur during execution
- The program maintains proper synchronization
- All 5 philosophers are accounted for in the results

### Running the Tests

To run the tests, use the following command:

```bash
go test -v
```

Expected output:

```
=== RUN   Test_dine
--- PASS: Test_dine
=== RUN   Test_dineWithVaryingDelays
--- PASS: Test_dineWithVaryingDelays
PASS
ok      dining-philosophers    X.XXs
```

### Test Coverage

To run tests with coverage:

```bash
go test -cover
```

## Learning Objectives

- Understanding concurrent programming challenges
- Implementing deadlock prevention strategies
- Working with mutual exclusion and synchronization
- Managing shared resources in concurrent systems

## References

- Dijkstra, E. W. (1971). Hierarchical ordering of sequential processes
- Operating System Concepts (Silberschatz, Galvin, Gagne)
- The Little Book of Semaphores by Allen B. Downey
