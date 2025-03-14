#Clone the repository: git clone https://github.com/sanidhyanagpal/task.git
Set up environment variables:Create a .env file in the root directory and add necessary environment variables for your services (DB credentials, API keys, etc.).
    DB_HOST=localhost
    DB_USER=admin
    DB_PASSWORD=your_password
    API_KEY=your_api_key
Start services with Docker Compose:docker-compose up --build
      The up command starts the services defined in your docker-compose.yml file, while the --build flag rebuilds images if there are any changes in the Dockerfile or       dependencies.
Run Kubernetes locally with Minikube: minikube start
                                      kubectl apply -f k8s/
    Minikube is a tool that runs a single-node Kubernetes cluster locally.
Run tests:
          Backend:go test ./...                  //Runs all Go test files in the directory
          Frontend:npm run test                  //Runs the frontend unit or integration tests using the configured test runner
