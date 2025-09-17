# Nekos.moe API Documentation

## Base URL

`https://nekos.moe/api/v1`

## Endpoints

### Get Image

Returns an Image matching the given ID.

*   **Method:** `GET`
*   **Endpoint:** `/images/{id}`
*   **Parameters:**
    *   `id` (string, path): The image's ID.

**Example Request:**

```bash
curl "https://nekos.moe/api/v1/images/1"
```

### Get Random Images

Returns the requested number of random Images.

*   **Method:** `GET`
*   **Endpoint:** `/random/image`
*   **Parameters:**
    *   `nsfw` (boolean): Whether to include NSFW images.
    *   `count` (number): The number of images to return (1-100).

**Example Request:**

```bash
curl "https://nekos.moe/api/v1/random/image?count=5"
```

### Search Images

Search for Images.

*   **Method:** `POST`
*   **Endpoint:** `/images/search`
*   **Body:**
    *   `id` (string)
    *   `nsfw` (boolean)
    *   `uploader` (string | object)
    *   `artist` (string)
    *   `tags` (Array<string>)
    *   `sort` (string): `newest`, `likes`, `oldest`, `relevance`
    *   `posted_before` (number, milliseconds)
    *   `posted_after` (number, milliseconds)
    *   `skip` (number): 0-2500
    *   `limit` (number): 1-50

**Example Request:**

```bash
curl -X POST -H "Content-Type: application/json" -d '{"tags": ["catgirl"]}' "https://nekos.moe/api/v1/images/search"
```

### Upload Image

Creates a new Pending Image.

*   **Method:** `POST`
*   **Endpoint:** `/images`
*   **Authentication:** Required
*   **Form fields:**
    *   `image` (file): The image to upload.
    *   `nsfw` (boolean)
    *   `artist` (string)
    *   `tags` (Array<string>)

**Example Request:**

```bash
# This endpoint requires authentication and multipart/form-data,
# so a curl example is not provided here.
```
