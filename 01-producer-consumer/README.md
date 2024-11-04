# Producer Consumer Problem - Pizza Parlor Simulation

This project demonstrates the classic Producer-Consumer problem using Go's concurrency features (goroutines and channels) through a pizza parlor simulation.

## Overview

The Producer-Consumer problem is a classic example of multi-process synchronization. It describes two processes:

- A **Producer** that generates data
- A **Consumer** that processes that data

The two processes share a fixed-size buffer and need to be synchronized so that:

- The producer won't try to add data when the buffer is full
- The consumer won't try to remove data when the buffer is empty

## Implementation Details

This implementation simulates a pizza parlor where:

- **Producer (Pizzeria)**

  - Represents the kitchen making pizzas
  - Runs as a separate goroutine
  - Generates pizza orders with random success/failure scenarios
  - Can fail due to:
    - Running out of ingredients
    - Pizza getting burnt

- **Consumer (Main routine)**
  - Processes the completed pizza orders
  - Handles successful and failed orders
  - Manages customer notifications
  - Tracks statistics about the day's operations

### Key Components

1. **Channels**

   - `data channel`: Transfers pizza orders from producer to consumer
   - `quit channel`: Handles graceful shutdown of the producer

2. **Structures**

   ```go
   type Producer struct {
       data chan PizzaOrder
       quit chan chan error
   }

   type PizzaOrder struct {
       pizzaNumber int
       message     string
       success     bool
   }
   ```

3. **Statistics Tracked**
   - Total pizzas attempted
   - Successfully made pizzas
   - Failed pizzas

## Running the Program

To run the program:

```bash
go run main.go
```

The program will simulate making 10 pizzas with random success/failure scenarios and delays. Each pizza's status is color-coded:

- 🟢 Green: Successful orders
- 🔴 Red: Failed orders
- 🟡 Yellow: Processing status
- 🔵 Cyan: General information

## Output

The program provides real-time feedback about:

- Order reception
- Making process with timing
- Delivery status
- Final statistics for the day
- Overall day rating based on performance

## Learning Outcomes

This implementation demonstrates several Go programming concepts:

- Goroutines for concurrent execution
- Channels for communication between concurrent processes
- Select statements for channel operations
- Structured error handling
- Use of custom types and methods
