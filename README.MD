# SecureMailHub: Feedback and Complaint Management System

This project is a microservice-based application developed in Go, designed to function as a robust **feedback and complaint management system**. It utilizes RabbitMQ for asynchronous messaging and is orchestrated with Kubernetes for scalable deployment. The frontend is built using Go HTML templates.

## 💡 Introduction

**SecureMailHub** is a comprehensive microservice-based application built in Go, dedicated to streamlining the process of **collecting and managing user feedback and complaints**. It provides a secure, scalable, and modular platform where users can submit their concerns, and administrators can track, respond to, and resolve them efficiently. The system leverages asynchronous communication patterns with RabbitMQ and is designed for cloud-native environments, supporting both local development with Docker Compose and production deployment on Kubernetes.

---

## ✨ Features

- **Modular and Scalable Architecture:** Each service is self-contained, allowing for independent development, deployment, and scaling.
- **Asynchronous Event-Driven Communication:** Services communicate via RabbitMQ message queues, enhancing system resilience and responsiveness.
- **Centralized Logging:** A dedicated `logger-service` aggregates and stores all application logs, powered by **MongoDB** for efficient log management.
- **Secure User Authentication:** The `authentication-service` provides robust user registration, login, and token-based authentication, backed by **PostgreSQL**.
- **Automated Email Notifications:** The `mailer-service` handles all outbound email communications, such as feedback confirmations and complaint status updates, integrated with **MailHog** for development.
- **Intuitive Web Frontend:** A user-friendly web interface, crafted with Go HTML templates, allows seamless interaction for submitting and viewing feedback/complaints.
- **Robust Deployment Options:** Supports local setup with Docker Compose for quick development cycles and production-grade deployment with Kubernetes for high availability and scalability.
- **RPC Communication:** Utilizes RPC for direct, synchronous inter-service communication where immediate responses are required.
- **Health Checks:** Includes readiness and liveness probes for Kubernetes deployments to ensure service reliability.

---

## ⚙️ Technologies

The project utilizes a robust stack of modern technologies:

- **Backend Language:** Go (Golang)
- **Containerization:** Docker
- **Orchestration:** Kubernetes (k8s)
- **Messaging Queue:** RabbitMQ
- **Databases:**
    - **PostgreSQL:** Relational database for `authentication-service` (user data).
    - **MongoDB:** NoSQL database for `logger-service` (log data).
- **Inter-service Communication:** HTTP/REST, RPC (Remote Procedure Call)
- **Email Testing:** MailHog
- **Web Frontend:** Standard Go HTML templates

---

## 🛠️ Services

This monorepo comprises the following Go microservices and supporting infrastructure:

| Service | Purpose | Database/Communication | Container Image |
| :--- | :--- | :--- | :--- |
| `authentication-service` | Manages user registration, login, and session authentication. | **PostgreSQL**, RPC/HTTP | `eyupcanbay/auth-service:v1.0.3` |
| `broker-service` | Acts as the API Gateway, handling incoming HTTP requests, routing them to the appropriate backend services, and producing events to RabbitMQ. | HTTP API, RabbitMQ, RPC | `eyupcanbay/broker-service:v1.0.2` |
| `frontend` | The user-facing web application, providing a UI for users to submit feedback/complaints and view their status. | Go HTML templates | `eyupcanbay/frontend:v1.0.6` |
| `listener-service` | Consumes events from RabbitMQ and orchestrates further actions, often calling other services via RPC or HTTP based on the event type. | RabbitMQ Consumer, RPC/HTTP Client | `eyupcanbay/listener-service:v1.0.2` |
| `logger-service` | Centralized logging service that receives log entries (via RPC or events) from other services and stores them. | **MongoDB**, RPC Server, RabbitMQ Consumer | `eyupcanbay/logger-service:latest` |
| `mailer-service` | Responsible for sending transactional emails (e.g., confirmations, status updates) based on events received from RabbitMQ. | MailHog (for dev), External SMTP | `eyupcanbay/mail-service:v1.0.2` |

---


## 🧑‍💻 Usage

Once SecureMailHub is deployed and accessible, users can:

- **Register and Log In:** Create and manage their accounts through the web frontend.
- **Submit Feedback:** Provide general feedback or suggestions using dedicated forms.
- **File Complaints:** Submit detailed complaints, which can be tracked and managed.
- **Receive Email Notifications:** Get automated email updates regarding their submissions (e.g., status changes, responses).
- **View Logs (Admin):** Monitor system activities and troubleshoot issues (potentially via an admin interface or direct access to `logger-service`).



---

## 🚀 Setup

### Prerequisites

Ensure you have the following installed:

- Git
- Docker
- Docker Compose
- Kubernetes CLI (`kubectl`)
- A Kubernetes cluster (e.g., Minikube, Kind) for local Kubernetes deployment testing.

### Local Development (Docker Compose)

The `docker-compose.yml` file in the project root defines the entire application stack for easy local development.

1.  **Clone the repository:**

    ```bash
    git clone [https://github.com/EyupCanbay/SecureMailHub.git](https://github.com/EyupCanbay/SecureMailHub.git)
    cd SecureMailHub
    ```

2.  **Build and run all services:**
    Navigate to the root directory containing `docker-compose.yml` and execute:

    ```bash
    docker-compose up --build
    ```
    This command will build Docker images for all Go services (broker, auth, logger, mailer, listener, frontend) and start them along with PostgreSQL, MongoDB, RabbitMQ, and MailHog.

3.  **Verify services are running:**

    ```bash
    docker-compose ps
    ```

4.  **Access the application:**

    - **Frontend:** `http://localhost:80`
    - **MailHog UI:** `http://localhost:8025` (to view emails sent by `mailer-service`)
    - **RabbitMQ Management:** `http://localhost:15672` (Default credentials: `guest`/`guest`)

### Deployment on Kubernetes

The `k8s/` directory contains all the Kubernetes manifests required to deploy SecureMailHub to a Kubernetes cluster.

1.  **Connect to your Cluster:**
    Ensure `kubectl` is configured to connect to your target Kubernetes cluster.

2.  **Review Kubernetes Manifests:**
    The `k8s/` directory includes: `auth-db-secret.yaml`, `authentication-service.yaml`, `broker-service.yaml`, `frontend.yaml`, `ingress.yaml`, `listener-service.yaml`, `logger-service.yaml`, `mailer-service.yaml`, `mailhog.yaml`, `mongo.yaml`, `postgres.yaml`, `rabbitmq.yaml`.

3.  **Apply all manifests to your cluster:**

    ```bash
    kubectl apply -f k8s/
    ```

4.  **Check the status of your deployments:**

    ```bash
    kubectl get deployments
    kubectl get services
    kubectl get pods
    kubectl get ingress
    ```
    Wait for all pods to be in a `Running` state.

5.  **Access the application via Ingress:**
    If you have an Ingress controller set up (like NGINX Ingress), you might need to configure your `/etc/hosts` file (or equivalent) to map the `frontend.info` and `broker-service.info` domains to your Ingress controller's IP address.

    **Example `/etc/hosts` entry** (replace `<INGRESS_IP>` with your Ingress controller's external IP):

    ```code
    <INGRESS_IP> frontend.info
    <INGRESS_IP> broker-service.info
    ```
    Then, access the frontend at `http://frontend.info` and the broker service at `http://broker-service.info:8080`.

---

## 🏛️ Architecture

The project employs an architecture using Go microservices and data layer components like RabbitMQ, MongoDB, and PostgreSQL, as illustrated below.

```mermaid
graph TD
    Frontend(Frontend) -- "Via Browser" --> Broker(Broker);
    
    Broker -- RPC --> Logger(Logger);
    Logger --> MongoDB(MongoDB);

    Broker -- "Publishes Message" --> RabbitMQ(RabbitMQ);
    RabbitMQ -- "Consumes Message" --> Consumer(Consumer);

    Consumer -- "RPC/Request" --> Authentication(Authentication);
    Authentication --> Postgres(Postgres);

    Consumer -- "Email Request" --> Mailer(Mailer);
    Mailer -- SMTP --> MailHog_SMTP[MailHog / SMTP Server];

    style Frontend fill:#fff,stroke:#333,stroke-width:2px,color:#333;
    style Broker fill:#fff,stroke:#333,stroke-width:2px,color:#333;
    style Logger fill:#fff,stroke:#333,stroke-width:2px,color:#333;
    style MongoDB fill:#fff,stroke:#333,stroke-width:2px,color:#333;
    style RabbitMQ fill:#fff,stroke:#333,stroke-width:2px,color:#333;
    style Consumer fill:#fff,stroke:#333,stroke-width:2px,color:#333;
    style Authentication fill:#fff,stroke:#333,stroke-width:2px,color:#333;
    style Postgres fill:#fff,stroke:#333,stroke-width:2px,color:#333;
    style Mailer fill:#fff,stroke:#333,stroke-width:2px,color:#333;
    style MailHog_SMTP fill:#fff,stroke:#333,stroke-width:2px,color:#333;