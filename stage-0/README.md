# Name Classification API (Genderize Wrapper)

A lightweight service that classifies a given name using the [Genderize API](https://api.genderize.io), applies additional business rules, and returns a structured and consistent response.

## Overview

This API accepts a name as input, queries the external Genderize service, processes the response, and returns a normalized classification result with confidence scoring and metadata.

## Endpoint

### Classify Name

```bash
GET /api/classify?name={name}
```

## Query Parameters

| Parameter | Type   | Required | Description      |
| --------- | ------ | -------- | ---------------- |
| name      | string | Yes      | Name to classify |

## Success Response (200 OK)

```json
{
  "status": "success",
  "data": {
    "name": "john",
    "gender": "male",
    "probability": 0.99,
    "sample_size": 1234,
    "is_confident": true,
    "processed_at": "2026-04-01T12:00:00Z"
  }
}
```

## Processing Rules

#### Field Mapping

- Extract: `gender`, `probability`
- Rename: `count → sample_size`

#### Confidence Rule

`is_confident = true` if:

- probability >= 0.7
- sample_size >= 100

> Both conditions must be met.

#### Timestamp Rule

`processed_at` = current UTC time in ISO 8601 format

**Example**:
2026-04-01T12:00:00Z

> Must be generated dynamically per request.

## Error Handling

All errors follow:

```json
{ "status": "error", "message": "<error message>" }
```

#### 400 Bad Request

Missing or empty name parameter

```json
{
  "status": "error",
  "message": "name query parameter is required"
}
```

#### 422 Unprocessable Entity

Invalid type for name

```json
{
  "status": "error",
  "message": "name must be a string"
}
```

#### 500 / 502

Upstream or server failure

```json
{
  "status": "error",
  "message": "upstream service error"
}
```

## Edge Case Handling

If Genderize returns:

- gender: null OR
- count: 0

**Return**:

```json
{
  "status": "error",
  "message": "No prediction available for the provided name"
}
```

## Example Request

```bash
curl "https://go-tutorial-eight.vercel.app/api/classify?name=john"
```

## Summary Logic

1. Validate input
2. Call Genderize API
3. Transform response
4. Apply confidence rules
5. Attach processed_at timestamp
6. Return structured output
