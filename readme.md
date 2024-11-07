# logsync

This project implements a basic yet powerful distributed log system, which serves as a foundation for reliable, high-performance data storage and retrieval. It manages data through structured segments, indexes, and a high-level log abstraction.

The goal is to design a log that supports a distributed architecture, enabling seamless scaling, fault tolerance, and high availability.

## Table of Contents
- [Project Overview](#project-overview)
- [Structure and Terminology](#structure-and-terminology)
- [Setup and Installation](#setup-and-installation)
- [Usage](#usage)
  - [Basic Operations](#basic-operations)
- [Testing](#testing)
- [Future Work](#future-work)

## Project Overview

This log system has been designed from the ground up to handle data persistence, reading, and management efficiently. It uses a layered approach with the following components:

- **Record**: The raw data entries stored in the log
- **Store**: Manages the storage file where records are written and read
- **Index**: Tracks offsets and positions for efficient record lookup
- **Segment**: Wraps the store and index, serving as a unit of storage
- **Log**: The primary interface that manages segments, appending, reading, and more

Each of these layers is built with the intent to simplify distributed log architecture by organizing data into manageable segments, facilitating data durability and quick retrieval.

## Structure and Terminology

Throughout this system, specific terms are used for clarity:

- **Record**: A data entry in the log
- **Store**: The file where records are saved
- **Index**: The file where index entries for records are saved
- **Segment**: The unit combining both a store and an index
- **Log**: The overarching system managing multiple segments

### Directories and Files

- `internal/log/`: Contains the primary log package
  - `store.go`: Manages the underlying storage
  - `index.go`: Provides index functionality
  - `segment.go`: Manages segments, combining stores and indexes
  - `log.go`: The main log that ties all segments together

## Setup and Installation

### Clone the Repository

```bash
git clone <repository-url>
cd <repository-directory>
```

### Install Dependencies

This project is written in Go. Ensure that you have Go installed and then fetch the dependencies:

```bash
go mod download
```

### Run Tests

Ensure the functionality of each component with:

```bash
go test ./internal/log
```

## Usage

### Basic Operations

This log system exposes methods for basic operations like appending and reading records, truncating old data, and handling log segments. Here's a quick guide:

#### 1. Append Records

To add data, use the Append method on the log. This method writes data to the current active segment and updates the index.

```go
log, err := NewLog("<directory-path>", config)
if err != nil {
    log.Fatalf("failed to create log: %v", err)
}

offset, err := log.Append(&api.Record{Value: []byte("your data")})
if err != nil {
    log.Fatalf("failed to append record: %v", err)
}
```

#### 2. Read Records

Retrieve records using the Read method, which accesses records by their offset.

```go
record, err := log.Read(offset)
if err != nil {
    log.Fatalf("failed to read record: %v", err)
}
fmt.Println("Record value:", string(record.Value))
```

#### 3. Truncate Logs

To remove old data and free up space, use the Truncate method to delete segments with offsets below a specified threshold.

```go
err := log.Truncate(lowestOffset)
if err != nil {
    log.Fatalf("failed to truncate log: %v", err)
}
```

## Testing

Comprehensive tests ensure each component (store, index, segment, and log) operates as expected. Run all tests to validate the setup:

```bash
go test ./internal/log
```

The tests cover:
- Appending and reading records
- Handling offsets and out-of-range errors
- Segment initialization and restoration
- Log truncation and cleanup

## Future Work

This log system is designed to grow into a distributed architecture. Future goals include:

- **Distributed Log Replication**: Implementing consensus algorithms for fault-tolerant data replication across nodes
- **Snapshot and Restore**: Adding capabilities for snapshotting the log state and restoring from snapshots
- **Optimized Indexing**: Enhancing indexing strategies to improve retrieval times for high-volume records
- **Graceful and Ungraceful Shutdown Handling**: Expanding functionality to handle recovery from crashes and data corruption
