# JUNO DB Client Library


JUNO DB is PayPal's home-grown secure, consistent, and highly available key-value store providing low, single digit millisecond latency at any scale. This document shows you how to integrate the JUNO DB client library into your Go project for simple CRUD operations.

---

## Table of Contents

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage Examples](#usage-examples)
  - [Creating a Client](#creating-a-client)
  - [CRUD Operations](#crud-operations)
- [Contributing](#contributing)
- [License](#license)

---

## Overview

The JUNO DB client library offers a simple API to interact with JUNO DB. It mainly uses the following packages:

```go
import (
    "github.com/paypal/junodb/pkg/client"
    "github.com/paypal/junodb/pkg/io"
)
```

You can create a configuration, initialize a client, and perform CRUD operations (e.g., GET and SET) using these packages.

---

## Prerequisites

- [Golang](https://golang.org/dl/) (version 1.11+ recommended)
- A running instance of JUNO DB  
  (Refer to the [JUNO DB repository](https://github.com/paypal/junodb) for server setup instructions)

---

## Installation

1. **Download the JUNO DB Client Library:**

   Use `go get` to fetch the modules:
   ```bash
   go get github.com/paypal/junodb/pkg/client
   go get github.com/paypal/junodb/pkg/io
   ```

2. **Initialize Your Module (if not already done):**

   ```bash
   mkdir my-junodb-app
   cd my-junodb-app
   go mod init github.com/yourusername/my-junodb-app
   ```

---

## Configuration

Create a configuration for your JUNO DB client. For example, create a configuration struct that defines your app name, namespace, and server endpoint:

```go
package main

import (
    "github.com/paypal/junodb/pkg/client"
    "github.com/paypal/junodb/pkg/io"
)

func main() {
    // Define your configuration
    ClientConfig := client.Config{
        Appname:   "ashok",
        Namespace: "ashok-junodb",
        Server: io.ServiceEndpoint{
            Addr: "192.168.1.6:8087",
            // Optionally, set Network and SSLEnabled if required:
            // Network: "tcp",
            // SSLEnabled: false,
        },
    }

    // Create a new JUNO DB client instance
    cl, err := client.New(ClientConfig)
    if err != nil {
        panic("Failed to create JUNO DB client: " + err.Error())
    }

    // Example CRUD operations
    runCRUD(cl)
}
```

---

## Usage Examples

### Creating a Client

Instantiate the client using the configuration:

```go
cl, err := client.New(ClientConfig)
if err != nil {
    // Handle error appropriately
    panic(err)
}
```

### CRUD Operations

#### SET Operation

Store a key-value pair in JUNO DB:
```go
err = cl.Set("key-name", "value")
if err != nil {
    // Handle error (e.g., log it)
    panic("Error setting key: " + err.Error())
}
```

#### GET Operation

Retrieve a value by key:
```go
value, err := cl.Get("key_name")
if err != nil {
    // Handle error
    panic("Error getting key: " + err.Error())
}
fmt.Println("Retrieved value:", value)
```

*Note: Extend this example with UPDATE and DELETE operations as supported by the client API.*

---
## Works only on linux and MacOS


*Happy Coding with JUNO DB!*

