# Polling
# Short Polling vs Long Polling

Polling is a technique used in web applications to fetch data from the server at regular intervals. There are two main types of polling: **Short Polling** and **Long Polling**. Below, we will discuss each method, provide examples, and illustrate them with diagrams.

## Short Polling

### Description

Short polling involves the client making requests to the server at regular intervals (e.g., every few seconds) to check for updates. This method is straightforward but can be inefficient because it generates a lot of HTTP requests, even if there are no updates.

### Example

**Frontend (JavaScript):**

```javascript
function fetchFileContent() {
    fetch('/setData')
        .then(response => response.text())
        .then(data => {
            console.log(data);
            document.getElementById('file-content').textContent = data;
        })
        .catch(error => console.error('Error fetching file:', error));
}

// Fetch file content every 2 seconds
setInterval(fetchFileContent, 2000);
```

**Backend (Go):**

```go
package main

import (
    "os"
    "time"
)

var DataToWrite []byte = []byte("default value")

func ShortPolling() {
    for {
        DataToWrite, _ = os.ReadFile("data/file.txt")
        time.Sleep(time.Second)
    }
}
```

### Diagram

```plaintext
Client                       Server
  |                            |
  |---- Request (every 2s) --->|
  |<---- Response (data) ------|
  |                            |
  |---- Request (every 2s) --->|
  |<---- Response (data) ------|
  |                            |
```

## Long Polling

### Description

Long polling is a technique where the client makes a request to the server, and the server holds the request open until there is new data to send. Once the data is sent, the client immediately makes a new request. This method is more efficient than short polling because it reduces the number of HTTP requests when there are no updates.

### Example

**Frontend (JavaScript):**

```javascript
function fetchFileContent() {
    fetch('/setData')
        .then(response => response.text())
        .then(data => {
            console.log(data);
            document.getElementById('file-content').textContent = data;
            // Immediately make another request
            fetchFileContent();
        })
        .catch(error => console.error('Error fetching file:', error));
}

// Initial fetch
fetchFileContent();
```

**Backend (Go):**

```go
package main

import (
    "os"
    "time"
)

var prevFileData []byte

var DataToWrite []byte = []byte("default value")

func LongPolling() {
    for {
        latestFileData, _ := os.ReadFile("data/file.txt")
        if string(prevFileData) != string(latestFileData) {
            DataToWrite = latestFileData
            prevFileData = latestFileData
        }
        time.Sleep(time.Second)
    }
}
```

### Diagram

```plaintext
Client                       Server
  |                            |
  |---- Request -------------->|
  |                            |
  |                            |
  |<---- Response (data) ------|
  |                            |
  |---- New Request ---------->|
  |                            |
  |                            |
  |<---- Response (data) ------|
  |                            |
```

## Comparison

| Feature              | Short Polling                          | Long Polling                          |
|----------------------|----------------------------------------|---------------------------------------|
| Efficiency           | Less efficient, more HTTP requests     | More efficient, fewer HTTP requests   |
| Complexity           | Simpler to implement                   | Slightly more complex                 |
| Latency              | Higher latency for updates             | Lower latency for updates             |
| Server Load          | Higher server load                     | Lower server load                     |
| Use Cases            | Suitable for less frequent updates     | Suitable for real-time updates        |

## Conclusion

Both short polling and long polling have their use cases. Short polling is simpler but can be inefficient for frequent updates. Long polling is more efficient for real-time updates but requires a slightly more complex implementation. Choose the method that best fits your application's requirements.
