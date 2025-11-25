(# Library Management — Documentation)

## Concurrency: Book Reservation Strategy

This project implements a simple concurrent reservation system for books. The concurrency design focuses on offloading reservation requests to a pool of worker goroutines and automatically cancelling reservations after a timeout.

- **Worker pool**: `concurrency.Init(workerCount, lib)` starts `workerCount` goroutines, each running `reservation_worker()` which reads reservation requests from a shared `reserve_chan` channel and processes them by calling `Library.ReserveBook`.

- **Request queue**: Reservation requests are submitted to the system by sending a `reservation` value into the channel via `concurrency.ReserveBook(bookID, memberID)`. The channel acts as a FIFO queue and decouples the caller from the actual reservation processing.

- **Auto-cancel timer**: After a reservation is processed, the worker sends the reservation into `reserved_chan`. A separate goroutine listens on `reserved_chan` and for each reserved item launches `auto_cancel_reservation`, which waits for a fixed timeout (currently 5s) and then calls `Library.UnReserveBook(bookID, memberID)` to automatically unreserve the book if applicable.

- **Library sharing**: The concurrency package currently receives a `services.Library` value in `Init` and stores it in a package-level variable `lib`. Because the `Library` type contains maps (which are reference types), mutations to maps are visible across copies. However, passing the `Library` by value can be confusing and brittle. Passing a pointer (i.e. `*services.Library`) is recommended for clarity and to signal shared mutable state explicitly.

- **Synchronization**: Individual `Book` values have a `Mu *sync.Mutex` used by `BorrowBook` to protect book-level operations. The reservation workers rely on `Library` methods for correctness; those methods should ensure proper synchronization when mutating shared state.