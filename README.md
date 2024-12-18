# CargoPlot Backend Challenge

A logistics platform backend for calculating expected rates based on freight shipping prices.

## Features

- Handle shipping prices from multiple companies and origins.

- Calculate the expected rate for each origin based on the top 10 cheapest prices.

## Getting Started

### Prerequisites

- Go 1.18 or newer installed on your system.

### Installation

Clone the repository:

```bash
git clone https://github.com/your-repo/cargoplot-backend.git
cd cargoplot-backend
```

### Build and Run

To build the application:

```bash
make build
```

To run the application:

```bash
make run
```

## API Endpoints

### POST /

- Description: Accepts a shipping price entry.

- Request Body:

```json
{
  "Company": 123,
  "Price": 500,
  "Origin": "CNSGH",
  "Date": "2023-01-01"
}
```

- Response: HTTP 200 OK if successful.

### GET /

- Description: Retrieves the expected rates for all origins.

- Response Body:

```json
{
  "CNSGH": 1050,
  "CNGGZ": 1100
}
```

## Running Tests

The current test coverage is 95%.
To run all tests:

```bash
make test
```

To check test coverage:

```bash
make coverage
```

## Project Structure

cargoplot/
├── cmd/
│   └── main.go              # Application entry point
├── internal/
│   ├── handlers/            # HTTP request handlers
│   ├── models/              # Data models
│   ├── services/            # Business logic
│   └── storage/             # Data storage
├── tests/                   # Unit and integration tests
├── Makefile                 # Automation commands
└── README.md                # Project documentation

## License

This project is licensed under the MIT License.

