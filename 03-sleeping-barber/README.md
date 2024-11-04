# The Sleeping Barber Problem

## Problem Description

The Sleeping Barber problem is a classic computer science synchronization problem that illustrates the complexities of process scheduling and synchronization. It was originally proposed by Edsger Dijkstra in 1965.

### Scenario

- A barbershop has one barber, one barber chair, and a waiting room with n chairs
- If there are no customers, the barber goes to sleep in the barber chair
- When a customer arrives:
  - If the barber is sleeping, they wake the barber and get their haircut
  - If the barber is busy but chairs are available, they sit in the waiting room
  - If all chairs are occupied, they leave the shop

## Implementation Details

This implementation uses Go's concurrency features to demonstrate the problem:

- Channels for communication between barber and customers
- Mutexes for protecting shared resources
- WaitGroups for synchronization

### Components

- `Barber`: Represents the barber who serves customers
- `Customer`: Represents customers arriving at random intervals
- `BarberShop`: Manages the shop's state and coordinates interactions

## Running the Program

```bash
go run main.go
```

## Solution Overview

The solution demonstrates:

- Proper synchronization between multiple goroutines
- Prevention of race conditions
- Efficient handling of shared resources
- Implementation of producer-consumer pattern

## Expected Output

The program will show:

- Customers arriving at random intervals
- Barber serving customers
- Customers waiting or leaving when shop is full
- Barber sleeping when no customers are present

## Key Concepts Demonstrated

- Mutex locks
- Channel communication
- Goroutines
- Wait groups
- Race condition prevention
- Deadlock avoidance

## References

- Dijkstra, E. W. (1965). "Cooperating Sequential Processes"
- [Go Concurrency Patterns](https://golang.org/doc/effective_go#concurrency)
