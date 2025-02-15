# API Server in Go - Blue-Green Deployment

This project is a simple REST API server built using the [Gin](https://github.com/gin-gonic/gin) framework in Go. The API provides basic CRUD operations for managing users, along with a dynamically generated random favicon.

## Features
- 🟢 **Blue-Green Deployment**: Seamless version switching with no downtime.
- 🟢 **User Management**: Create, retrieve, and delete users.
- 🟢 **RESTful Endpoints**: Standard API conventions.
- 🟢 **Liveness Probe**: Kubernetes health check on `/` route.
- 🟢 **Environment Variables**: Configurable via Kubernetes `ConfigMap`.

---

## **1️⃣ Blue-Green Deployment Setup**

| Component | Blue Version | Green Version |
|-----------|-------------|--------------|
| Deployment | `go-api-blue` | `go-api-green` |
| Service | `go-api-blue-service` | `go-api-green-service` |
| ConfigMap | `go-api-blue-config` | `go-api-green-config` |
| Traffic Controller | ✅ `go-api-primary-service` (switches between Blue/Green) |

---

## **2️⃣ Installation**
### **Prerequisites**
- Install [Go](https://golang.org/doc/install) (version 1.22 or later)
- Install [Docker](https://docs.docker.com/get-docker/)
- Install [Kubernetes (`kubectl`)](https://kubernetes.io/docs/setup/)

### **Clone the Repository**
```sh
git clone https://github.com/yBiodigitalJaz/bg-deploy-ha.git
cd bg-deploy-ha
```

### **Install Dependencies**
```sh
go mod tidy
```

---

## **3️⃣ Running the Server Locally**
```sh
go run main.go
```
The server starts at **`http://localhost:8080`**.

---

## **4️⃣ Deploying Blue-Green Setup in Kubernetes**

### **Apply Namespace**
```sh
kubectl apply -f namespace.yaml
```

### **Apply ConfigMaps**
```sh
kubectl apply -f configmap-blue.yaml --namespace=go-api-namespace
kubectl apply -f configmap-green.yaml --namespace=go-api-namespace
```

### **Apply Deployments & Services**
```sh
kubectl apply -f deployment-blue.yaml --namespace=go-api-namespace
kubectl apply -f deployment-green.yaml --namespace=go-api-namespace
kubectl apply -f service-blue.yaml --namespace=go-api-namespace
kubectl apply -f service-green.yaml --namespace=go-api-namespace
kubectl apply -f primary-service.yaml --namespace=go-api-namespace
```

---

## **5️⃣ Traffic Switching: Blue → Green**

To switch traffic from `blue` to `green`, update the `primary-service` selector:

```sh
kubectl patch service go-api-primary-service -n go-api-namespace -p '{"spec":{"selector":{"app":"go-api","version":"green"}}}'
```

✅ Now, `go-api-primary-service` routes traffic to **green deployment**.

### **Rollback If Needed (Green → Blue)**
If the new (green) version has issues, quickly roll back to **blue**:

```sh
kubectl patch service go-api-primary-service -n go-api-namespace -p '{"spec":{"selector":{"app":"go-api","version":"blue"}}}'
```

---

## **6️⃣ Verify Deployments**
```sh
kubectl get pods -n go-api-namespace
kubectl get services -n go-api-namespace
kubectl get deployment -n go-api-namespace
```

---

## **7️⃣ Cleanup Kubernetes Resources**
To remove everything:
```sh
kubectl delete namespace go-api-namespace
```

---

## **8️⃣ Next Steps**
🔹 **CI/CD Integration**: Automate blue-green deployments using **GitHub Actions**.  
🔹 **Health Checks**: Add **liveness/readiness probes** to ensure no downtime.  
🔹 **Traffic Splitting**: Implement **Canary Deployment** with Istio/NGINX.  

---
