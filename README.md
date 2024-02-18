# API for スーパー支払い君.com
Go API for creating and showing invoice data.

## Endpoints
Endpoints for this API
```
POST /api/invoice/create
GET /api/invoice/list
```

## How to run this API locally
This project uses go, PostgreSQL, and make. Please make sure these tools are installed before setting up.

You can use the environment variables defined in `.envrc.dev` for testing locally.

To run locally, follow the below steps:


Set up DB:
```
docker compose up
```
Seed db with test data
```
make seed-db
```
Build and run server
```
make run
```
Run unit tests
```
make test
```

# Notes
This project's core packages are as follows:
- `cmd/server/main.go` is the entry point of this API.
- `auth` contains some very basic auth functions
- `http` contains the transport layer related logic, such as http client, handlers and handler functions.
- `service` is the usecase layer, where the business logic is implemented.
- `db` is the repository layer, containing the data entities and models. 

This directory structure is organized around packages rather than layers and is simplified due to the scope of this exercise. Each package represents a unit of functionality and DI is used to define separation and hierarchy between packages.

### Improvements for the future
- Implement a more secure Auth
  - Current auth is basic auth, but it would be better to use some kind of token based auth
  - With more time, an auth package with refreshable token could be implemented
- Setup healthcheck
- Setup monitoring/tracers
- Setup CI pipeline
- Increase test coverage, like integration tests, e2e tests, and more unit tests, for auth and http layer
  - service/use-case layer logic has been tested thoroughly, and main validation logic in http layer
- For ListInvoices endpoint:
  - Add sort (by issue_date, due_date, etc)
  - Add filtering (by status, service_provider_id, etc)
  - Support pagination
- Use swagger or some kind of tool for generating and keeping up-to-date documents for the API
- Use ADR for documenting technical decisions within this repository

# Example requests
The below curl requests require you set up the server and seed the database with test data using the `make seed-db` command.
If you use Postman, you can import the `api.postman_collection.json` for testing the endpoint locally with Postman.

## `GET /api/invoice`
Sample request
```bash
curl -i -X GET \
     -H "Authorization: Basic $(echo -n 'arthur@example.com:password123' | base64)" \
     -H 'accept: application/json' \
     http://localhost:8080/api/invoices?start_date=2024-02-01&end_date=2024-07-01
```
Sample response:
```json
{
  "result": "OK",
  "data": [
    {
      "id": "b796e13d-bd76-46a6-af4e-159f8e19587f",
      "issue_date": "2024-02-15T11:12:57.06461+09:00",
      "payment_amount": 100,
      "commission": 4,
      "commission_rate": 0.04,
      "tax": 1.1,
      "tax_rate": 0.1,
      "amount": 104.4,
      "due_date": "2024-03-16T11:12:57.06461+09:00",
      "company_id": 1,
      "service_provider_id": 1,
      "status": "OUTSTANDING"
    }
  ]
}
```
Sample invalid request
```bash
curl -i -X GET \
     -H "Authorization: Basic $(echo -n 'arthur@example.com:password123' | base64)" \
     -H 'accept: application/json' \
     http://localhost:8080/api/invoices?start_date=2024-06-01&end_date=2024-04-01
```
Sample response:
```json
{
  "result": "error",
  "code": 400,
  "message": "http.validateListInvoicesRequest: start_date 2024-06-01 is after end_date 2024-04-01"
}
```

## `POST /api/invoice`
Sample request
```bash
curl -i -X POST \
  -H "Authorization: Basic $(echo -n 'arthur@example.com:password123' | base64)" \
  -H 'Content-Type: application/json' \
  -d '{"payment_amount":10000,"service_provider_id":3,"due_date":"2024-05-12T17:11:15.504318+09:00" }' \
  http://localhost:8080/api/invoices
```
Sample response:
```json
{
  "result": "OK",
  "data": {
    "id": "1291401c-6766-470c-bb1e-298a6b6f0928",
    "issue_date": "2024-02-16T10:10:17.893791+09:00",
    "payment_amount": 10000,
    "commission": 400,
    "commission_rate": 0.04,
    "tax": 40,
    "tax_rate": 0.1,
    "amount": 10440,
    "due_date": "2024-06-12T17:11:15.504318+09:00",
    "company_id": 2,
    "service_provider_id": 3,
    "status": "OUTSTANDING"
  }
}
```
Sample invalid request
```bash
curl -i -X POST \
  -H "Authorization: Basic $(echo -n 'arthur@example.com:wrongpass' | base64)" \
  -H 'Content-Type: application/json' \
  -d '{"payment_amount":10000,"service_provider_id":3,"due_date":"2024-06-12T17:11:15.504318+09:00" }' \
  http://localhost:8080/api/invoices
```
Sample response
```json
{
  "result": "error",
  "code": 401,
  "message": "service.(*Impl).validateUser: failed to authenticate user Arthur (arthur@example.com), wrong password"
}
```
