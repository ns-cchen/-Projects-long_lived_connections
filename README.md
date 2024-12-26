# Running Server
```
kubectl apply -f server.yaml
```

# Running Client
```
kubectl apply -f client.yaml
```

# Verifying traffic
```
kubectl logs {client pod} | grep "Response from Pod" | awk -F"Pod " '{print $2}' | cut -d' ' -f1 | sort | uniq -c
```
