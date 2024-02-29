# PaymentPlatform

This project is a payment processing platform that utilizes Docker to simplify the setup of development and production environments.

# Execution of the Solution

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
BASIC_AUTH_PASSWORD=test;
BASIC_AUTH_USERNAME=test
```

# Launch Containers with Docker Compose
docker-compose up --build

# Accessing the Application
With the containers up and running, access the application via a web browser or HTTP client at http://localhost:8080 (adjust if your Docker Compose configuration specifies a different port).

## API Endpoints
####Basic Auth required
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

## Stopping the Environment
```bash 
   docker-compose down 
   ``` 

# Assumptions:
The decision was made to use specific technologies, such as Docker and Docker Compose, based on the familiarity of the development team and their suitability to simplify the development and deployment process.

End-user requirements and expectations were considered to align with the platform's objectives and were reflected in the use cases and solution architecture.

# Areas for Improvement:
In the current implementation of the online payment platform application, I identify a main area of improvement: the migration from the monolithic architecture to a microservices-based architecture. This transition offers numerous benefits, such as increased system scalability, flexibility, and maintainability. Microservices allow you to decompose your application into smaller, independent components, making it easier to deploy, evolve, and manage each service individually.

Additionally, I recognize that the current architecture was designed to meet a specific goal: to develop a working prototype in a limited period of time, approximately 5 hours. This approach demonstrated my versatility and ability to meet established deadlines, highlighting my effectiveness and speed in delivering solutions. However, it is important to note that while I achieved the goal effectively, the monolithic approach may have limitations in terms of scalability and long-term maintainability.

In this sense, migrating to a microservices architecture is presented as an option to improve my efficiency and scalability as I grow and evolve. While monolithic development allowed me a quick and efficient deployment in the short term, adopting microservices offers a more flexible and sustainable solution in the long term, allowing me to better adapt to the changing demands of my business and the market.

# Cloud Technologies:
No cloud technologies will be used in this project. To be honest, I tried to deploy the application on AWS as it is the technology I have the most experience with. However, I found that the high costs associated with AWS were not feasible for me at this time. Therefore, I decided to keep the application in a local environment.

However, I can describe the cloud technologies I had planned to use. These include:

`Amazon Aurora DB`: Aurora DB provides a highly scalable and secure relational database instance on AWS. Aurora was considered to host the application's PostgreSQL database due to its performance, scalability, and advanced security capabilities.

`Amazon Cognito`: Cognito is an AWS service that provides authentication and authorization capabilities for web and mobile applications. I was planning to use Cognito to manage the security and authentication of online payment platform users. With Cognito, you would have been able to easily integrate JWT token-based authentication and manage users and groups securely.

`Amazon CloudWatch`: CloudWatch is an AWS monitoring and observability service that provides visibility into the operational performance of AWS applications and resources. I would have used CloudWatch for application auditing and logging, allowing me to monitor and analyze performance, as well as detect and resolve potential issues.

`AWS Elastic Beanstalk`: Elastic Beanstalk is an AWS service that simplifies the deployment and management of web applications and cloud services. I would have opted for Elastic Beanstalk due to its ease of use and ability to automate deployment, scaling, and management tasks, which would have accelerated the online payment platform API development and deployment process.

These technologies were selected for their ability to provide a robust and scalable infrastructure, as well as their tight integration with other AWS services, which would have facilitated the development and deployment of the application in the cloud. However, due to cost considerations, I chose to keep the application in a local environment at this time.
