# PerfectStay Booking API

API de réservation simple en Go pour gérer les réservations.

## Overview

Cette API expose 3 endpoints principaux pour les réservations :
- Ajouter une réservation
- Rechercher des réservations par statut
- Récupérer une réservation par identifiant

## Requirements

- Go 1.20+ installé
- Git
- Un terminal / invite de commandes

## Setup

1. Clonez le dépôt :

   ```bash
   git clone https://github.com/MarouaniHani/golang_mongodb.git
   cd golang_mongodb
   ```

2. Installez les dépendances :

   ```bash
   go mod tidy
   ```

3. Compilez l'application :

   ```bash
   go build -o perfectstay-api ./...
   ```

## Run

Démarrez le serveur API :

```bash
go run ./...
```

Par défaut, le serveur écoute sur le port `8080`.

## API Documentation

### Base URL

`http://localhost:8080`

### Endpoints

#### POST /bookings

Creates a new hotel booking in the system.

Request Body:

The request body must be sent as **JSON** (`application/json`).

| Field           | Type   | Required | Description                        |
|-----------------|--------|----------|------------------------------------|
| `customer_name` | string | Yes      | The full name of the customer.     |
| `hotel_name`    | string | Yes      | The name of the hotel to book.     |

Example:

```json
{
  "customer_name": "rayen",
  "hotel_name": "Hilton"
}
```

Responses:

| Status | Description                              |
|--------|------------------------------------------|
| `201`  | Booking created successfully. Returns the new booking's `id`. |
| `400`  | Bad request — missing or invalid fields. |

Example success response (`201`):

```json
{
  "id": "6a282985641e28a0c7aa3106"
}
```

#### GET /bookings

Retrieves a paginated list of bookings filtered by their status.

### Query Parameters

| Parameter | Type   | Required | Description                                                  |
|-----------|--------|----------|--------------------------------------------------------------|
| `status`  | string | Yes      | Filter bookings by status. Accepted values: `CONFIRMED`, `PENDING`, `CANCELLED`. |
| `page`    | number | Yes      | The page number to retrieve (1-based index).                 |
| `limit`   | number | Yes      | The maximum number of bookings to return per page.           |

### Example Request

```http
GET http://localhost:8080/bookings?status=CONFIRMED&page=1&limit=20
```

### Response

Returns a JSON array of booking objects matching the specified status. Returns an empty array `[]` if no bookings are found for the given criteria.

### Notes

- Use `status=CONFIRMED` to retrieve confirmed bookings and `status=PENDING` for pending ones.
- Combine `page` and `limit` to paginate through large result sets.

#### GET /bookings/{id}

Retrieves a single booking by its unique ID.

**Method:** GET  
**URL:** `http://localhost:8080/bookings/{id}`

### Path Parameters

| Parameter | Type   | Required | Description                        |
|-----------|--------|----------|------------------------------------|
| id        | string | Yes      | The unique identifier of the booking |

### Response

**200 OK**

Returns the booking object matching the provided ID.

```json
{
  "id": "6a282985641e28a0c7aa3106",
  "customer_name": "rayen",
  "hotel_name": "Hilton",
  "status": "PENDING",
  "created_at": "2026-06-09T14:56:05.635Z"
}
```
Updates an existing booking identified by its ID.

#### PUT /bookings/{id}

## Path Parameters

| Parameter | Type   | Required | Description                          |
|-----------|--------|----------|--------------------------------------|
| `id`      | string | Yes      | The unique identifier of the booking |

## Request Body

Content-Type: `application/json`

| Field        | Type   | Required | Description                                      |
|--------------|--------|----------|--------------------------------------------------|
| `status`     | string | No       | The new status of the booking (e.g. `CONFIRMED`) |
| `hotel_name` | string | No       | The name of the hotel for the booking            |

### Example Request Body

```json
{
  "status": "CONFIRMED",
  "hotel_name": "Hilton"
}
```

## Responses

| Status Code | Description                                      |
|-------------|--------------------------------------------------|
| `204`       | Booking updated successfully. No content returned. |
| `400`       | Bad request — invalid input data                 |
| `404`       | Booking not found for the given ID               |

#### DELETE /bookings/:id

Deletes an existing booking identified by its unique ID.

## Path Parameters

| Parameter | Type   | Required | Description                          |
|-----------|--------|----------|--------------------------------------|
| `id`      | string | Yes      | The unique identifier of the booking |

## Request Body

None.

## Responses

| Status Code | Description                                        |
|-------------|----------------------------------------------------|
| `204`       | Booking deleted successfully. No content returned. |
| `404`       | Booking not found for the given ID                 |

## Notes

- Ajustez la configuration et les variables d'environnement selon l'environnement de production.
- Ajoutez l'authentification et la validation si nécessaire.
- Étendez l'API avec d'autres fonctionnalités si besoin.
