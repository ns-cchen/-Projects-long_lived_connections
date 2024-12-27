# Long-Lived Connections Simulation

This project simulates a scenario where long-lived connections are handled in a Kubernetes environment. It includes:
1. A server application that delays responses randomly.
2. A client application that sends concurrent requests to the server to test load balancing and traffic distribution across Kubernetes pods.

## Components

### Server
The server is a Flask application that:
- Simulates a delay in response (20-60 seconds).
- Identifies the pod responding to the request using an environment variable `POD_ID`.

### Client
The client application:
- Sends up to **10,000 concurrent requests** to the server using Python's `ThreadPoolExecutor`.
- Uses a persistent session to simulate real-world usage of RESTful APIs.

## Setup Instructions

### Step 1: Build Docker Images
#### Server
Build the server Docker image:
```bash
docker build -t your-docker-hub-username/long_lived_connections_server -f Dockerfile .
docker push your-docker-hub-username/long_lived_connections_server
```

#### Client
Build the client Docker image:
```bash
docker build -t your-docker-hub-username/long_lived_connections_client -f Dockerfile_client .
docker push your-docker-hub-username/long_lived_connections_client
```

### Step 2: Deploy Server
```
kubectl apply -f server.yaml
```

### Step 3: Deploy Client
```
kubectl apply -f client.yaml
```

### Step 4: Verifying traffic
```
kubectl logs {client pod} | grep "Response from Pod" | awk -F"Pod " '{print $2}' | cut -d' ' -f1 | sort | uniq -c
```

## Result

The server will delay responses randomly between 20-60 seconds.
The client will send 10,000 requests concurrently with a delay of 1-3 seconds between each request.
The traffic is distributed evenly across the server pods.

```
counts server-name
3377   server-5f6cf9cbc4-mg7js
3288   server-5f6cf9cbc4-w6kwx
3241   server-5f6cf9cbc4-wdf4v
```

## Traffic Distribution Differences Between Go and Flask in Kubernetes
### Python Flask Service

**Setup**:

Flask service with three pods (a, b, c) managed by a Kubernetes deployment.

A Kubernetes service load balances traffic across the pods.

Flask runs directly (not with uWSGI).

**Behavior**:

When one or more pods restart, traffic continues to be distributed evenly across all pods.

### Go Service

**Setup**:

Replaced the Flask service with a Go-based HTTP server.

Deployment and Kubernetes service configuration remain the same.

**Behavior**:

When pods b and c restart, the traffic sticks to pod a instead of being distributed evenly.
