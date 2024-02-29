# PaymentPlatform

This project is a payment processing platform that utilizes Docker to simplify the setup of development and production environments.

## Prerequisites

Before starting, ensure Docker and Docker Compose are installed on your system. Installation instructions can be found on the [official Docker website](https://docs.docker.com/get-docker/).

## Setting Up the Environment

Follow these steps to set up the development environment:

1. **Clone the Repository**

   Clone this repository to your local machine:

   ```bash
   git clone https://github.com/andrew221293/paymentplatform.git
   cd paymentplatform

# Configure Environment Variables

Create a `.env` file at the project's root and adjust the environment variables as needed. Example `.env` file content:

```env
POSTGRES_USER=postgres_user
POSTGRES_PASSWORD=secure_password
POSTGRES_DB=your_database_name
```

# Launch Containers with Docker Compose
docker-compose up --build

# Accessing the Application
With the containers up and running, access the application via a web browser or HTTP client at http://localhost:8080 (adjust if your Docker Compose configuration specifies a different port).

## API Endpoints
### Process Payment
**URL**: `/payments`
- **Method**: `POST`
- **Description**: Process a new payment.
- **Body**:
  ```json
  {
    "amount": 100.00,
    "currency": "USD",
    "cardNumber": "1234567890123456",
    "cvv": "123",
    "expiryDate": "12/23",
    "merchantID": 12345,
    "customerID": 67890
  }
- **Success Response:**: `HTTP 200 OK`

### Get Payment Details
**URL**: `/payments/{payment_id}`
- **Method**: `GET`
- **Description**: Retrieve details of a specific payment by its ID.
- **URL Parameters**: `payment_id=[integer]`
- **Success Response:**: `HTTP 200 OK`

### Process Refund
**URL**: `/payments/{payment_id}/refund`
- **Method**: `PATCH`
- **Description**: Process a refund for a specific payment.
- **Body**:
  ```json
  {
    "amount": 100.00,
  }
- **Success Response:**: `HTTP 200 OK`

For more detailed information about the API endpoints, please consult the [openapi.yaml](./openapi.yaml) file located at the root of the project.

# Stopping the Environment
```bash 
   docker-compose down 
   ``` 