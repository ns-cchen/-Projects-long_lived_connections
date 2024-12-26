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

#### Client
Build the client Docker image:
```bash
docker build -t your-docker-hub-username/long_lived_connections_client -f Dockerfile_client .
docker push your-docker-hub-username/long_lived_connections_client

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
