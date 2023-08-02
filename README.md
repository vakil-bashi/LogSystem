# REST API LogSystem with Golang and Go-Gin

This is a simple log system built using Golang and the Go-Gin web framework. The system allows users to submit logs through HTTP requests, and the logs are stored in Elasticsearch for easy retrieval and analysis.

## Features

- Submit logs via HTTP requests.
- Store logs in Elasticsearch for efficient indexing and searching.
- Retrieve logs based on various search criteria.
- Simple and intuitive RESTful API.

## Requirements

- Golang: Make sure you have Golang installed on your system.
- Elasticsearch: Install and set up Elasticsearch for log storage and retrieval.
- Go-Gin: Install the Go-Gin framework to handle HTTP requests.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/vakil-bashi/logSystem.git
   cd log-system
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Set up Elasticsearch:

   Ensure that Elasticsearch is installed and running on your system or accessible via a remote endpoint. Update the Elasticsearch connection details in the `main.go` file to point to your Elasticsearch instance.

4. Build and run the application:

   ```bash
   docker compose build && docker-compose up -d
   ```

## API Endpoints

### Submit a log

- Endpoint: `/logs/insert`
- Method: `POST`
- Request body: JSON object containing the log data.
- Response: JSON object with the status of the log submission.

### Search logs

- Endpoint: `/api/logs`
- Method: `GET`
- Query parameters: You can use various query parameters to filter and search for logs based on different criteria such as timestamp, severity level, source, etc.
- Response: JSON array containing the matched logs.

## Example Usage

1. Submit a log:

   ```bash
   curl -X POST -H "Content-Type: application/json" -d '{"message": "Something went wrong!", "severity": "error", "source": "app_server"}' http://localhost:8080/api/logs
   ```

2. Retrieve logs:

   ```bash
   curl -X GET "http://localhost:8080/api/logs?severity=error&source=app_server"
   ```

## License

This project is licensed under the [MIT License](LICENSE).

## Contribution

Contributions are welcome! If you find any issues or have suggestions for improvements, feel free to open an issue or create a pull request.

## Authors

- [Your Name](https://github.com/your_username)

## Acknowledgments

- Thanks to the creators and maintainers of Golang, Go-Gin, and Elasticsearch for their amazing tools and frameworks.

## Disclaimer

This log system is meant for educational and demonstration purposes. Always use secure and production-ready solutions for real-world applications.

---

Note: Make sure to replace "your_username" with your actual GitHub username and update other relevant details according to your project. This README.md file provides a general template and can be expanded further based on the complexity and features of your log system.