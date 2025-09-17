# Nekos API Documentation

## Base URL

`https://api.nekosapi.com/v4`

## Endpoints

### Search for an Image

This endpoint allows you to search for an image, filtering by tags, characters, artists, etc.

*   **Method:** `GET`
*   **Endpoint:** `/images`
*   **Parameters:**
    *   `rating` (Array of strings): The (age) rating of the image. (e.g. `safe`, `suggestive`, `borderline`, `explicit`)
    *   `artist` (Array of integers): The artist's ID.
    *   `tags` (Array of strings): The tags names, comma-delimited.
    *   `without_tags` (Array of strings): The tags to exclude's names, comma-delimited.
    *   `limit` (Integer): The amount of images to return. [1-100]
    *   `offset` (Integer): The amount of images to skip. [0-...]

**Example Request:**

```bash
curl "https://api.nekosapi.com/v4/images"
```

### Get an Image by ID

This endpoint allows you to get an image by its ID.

*   **Method:** `GET`
*   **Endpoint:** `/images/{id}`
*   **Parameters:**
    *   `id` (string, path): The image's ID.

**Example Request:**

```bash
curl "https://api.nekosapi.com/v4/images/1"
```

### Get Random Images

This endpoint allows you to get x random images, filtering by tags, characters, artists, etc.

*   **Method:** `GET`
*   **Endpoint:** `/images/random`
*   **Parameters:**
    *   `rating` (Array of strings): The (age) rating of the image. (e.g. `safe`, `suggestive`, `borderline`, `explicit`)
    *   `artist` (Array of integers): The artist's ID.
    *   `tags` (Array of strings): The tags names, comma-delimited.
    *   `without_tags` (Array of strings): The tags to exclude's names, comma-delimited.
    *   `limit` (Integer): The amount of images to return. [1-100]

**Example Request:**

```bash
curl "https://api.nekosapi.com/v4/images/random"
```

### Get a Random Image File

This endpoint allows you to get a redirect to a random image's file URL, filtering by tags, characters, artists, etc.

*   **Method:** `GET`
*   **Endpoint:** `/images/random/file`
*   **Parameters:**
    *   `rating` (Array of strings): The (age) rating of the image. (e.g. `safe`, `suggestive`, `borderline`, `explicit`)
    *   `artist` (Array of integers): The artist's ID.
    *   `tags` (Array of strings): The tags names, comma-delimited.
    *   `without_tags` (Array of strings): The tags to exclude's names, comma-delimited.

**Example Request:**

```bash
curl "https://api.nekosapi.com/v4/images/random/file"
```
