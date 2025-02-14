# API Server in Go

This project is a simple REST API server built using the [Gin](https://github.com/gin-gonic/gin) framework in Go. The API provides basic CRUD operations for managing users, along with a dynamically generated random favicon.

## Features
- 🟢 **User Management**: Create, retrieve, and delete users.
- 🟢 **RESTful Endpoints**: Standard API conventions.
- 🟢 **Random Favicon**: Generates a new 16x16 favicon dynamically.
- 🟢 **Liveness Probe**: Kubernetes health check on `/` route.
- 🟢 **Environment Variables**: Configurable via Kubernetes `ConfigMap`.

---

## **1️⃣ Installation**
### **Prerequisites**
- Install [Go](https://golang.org/doc/install) (version 1.22 or later)
- Install [Docker](https://docs.docker.com/get-docker/)
- Install [Kubernetes (`kubectl`)](https://kubernetes.io/docs/setup/)

### **Clone the Repository**
```sh
git clone https://github.com/your-username/go-api-server.git
cd go-api-server
```

### **Install Dependencies**
```sh
go mod tidy
```

---

## **2️⃣ Running the Server**
```sh
go run main.go
```
The server starts at **`http://localhost:8080`**.

---

## **3️⃣ API Endpoints**

### 🔹 **Get All Users**
```sh
GET /users
```
**Response:**
```json
[
  {"id": 1, "name": "Alice", "email": "alice@example.com"},
  {"id": 2, "name": "Bob", "email": "bob@example.com"}
]
```

### 🔹 **Get a Specific User by ID**
```sh
GET /users/2
```
**Response:**
```json
{"id": 2, "name": "Bob", "email": "bob@example.com"}
```

### 🔹 **Create a New User**
```sh
POST /users
Content-Type: application/json
{
  "name": "Charlie",
  "email": "charlie@example.com"
}
```
**Response:**
```json
{"id": 3, "name": "Charlie", "email": "charlie@example.com"}
```

### 🔹 **Delete a User**
```sh
DELETE /users/2
```
**Response:**
```json
{"message": "User deleted"}
```

### 🔹 **Browser Hello Endpoint**
```sh
GET /
```
**Response:**
```text
Hello, world!
```

### 🔹 **Get Random Favicon**
```sh
GET /favicon.ico
```
This generates a **16x16 random favicon** each time.

---

## **4️⃣ Dockerizing the API**
### **Build and Run the Docker Image**
```sh
docker build -t your-dockerhub-username/go-api .
docker run -p 8080:8080 --env PORT=8080 --env APP_ENV=production --env LOG_LEVEL=info your-dockerhub-username/go-api
```

### **Push to Docker Hub**
```sh
docker tag your-dockerhub-username/go-api your-dockerhub-username/go-api:latest
docker push your-dockerhub-username/go-api:latest
```

---

## **5️⃣ Kubernetes Deployment**
### **Namespace Configuration**
To deploy in a separate namespace, first apply:
```sh
kubectl apply -f namespace.yaml
```

### **Create ConfigMap**
```sh
kubectl apply -f configmap.yaml --namespace=go-api-namespace
```

### **Deploy API & Service**
```sh
kubectl apply -f deployment.yaml --namespace=go-api-namespace
kubectl apply -f service.yaml --namespace=go-api-namespace
```

### **Verify Deployment**
```sh
kubectl get pods -n go-api-namespace
kubectl get services -n go-api-namespace
kubectl logs -f deployment/go-api -n go-api-namespace
```

---

## **6️⃣ Environment Variables**
The API reads configurations via **ConfigMap**.

| Variable  | Default | Description |
|-----------|---------|-------------|
| `PORT`    | `8080`  | API listening port |
| `APP_ENV` | `production` | Application environment |
| `LOG_LEVEL` | `info` | Log level (`debug`, `info`, `warn`, `error`) |

### **Example `configmap.yaml`**
```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: go-api-config
data:
  PORT: "8080"
  APP_ENV: "production"
  LOG_LEVEL: "debug"
```

---

## **7️⃣ Kubernetes Health Probes**
Kubernetes automatically restarts the container if it's unhealthy.

### **Liveness Probe**
```yaml
livenessProbe:
  httpGet:
    path: /
    port: 8080
  initialDelaySeconds: 5
  periodSeconds: 10
```

### **Readiness Probe**
```yaml
readinessProbe:
  httpGet:
    path: /
    port: 8080
  initialDelaySeconds: 3
  periodSeconds: 5
```

---

## **8️⃣ Exposing API via Kubernetes Service**
By default, the API runs inside a **ClusterIP** service.

To access it externally:
1. Change `type: LoadBalancer` in **`service.yaml`**
2. Use `kubectl port-forward`:
   ```sh
   kubectl port-forward svc/go-api-service 8080:80 -n go-api-namespace
   ```

---

## **9️⃣ Cleanup Kubernetes Resources**
To remove everything:
```sh
kubectl delete namespace go-api-namespace
```

---

## **🔟 Next Steps**
🔹 **Enhancements**: Add persistent storage (PostgreSQL, MySQL).  
🔹 **Authentication**: Implement JWT-based authentication.  
🔹 **Autoscaling**: Enable Kubernetes **Horizontal Pod Autoscaler (HPA)**.  
🔹 **Monitoring**: Integrate with **Prometheus & Grafana**.  
🔹 **Logging**: Centralized logging using **Loki, Elasticsearch, or Fluentd**.  
