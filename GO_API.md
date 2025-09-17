# Go Booru Server API

This document describes the API for the Go Booru Server.

## Authentication

If authentication is enabled in the server configuration, requests to `/api/search` must include a valid bearer token in the `Authorization` header.

`Authorization: Bearer <your-token>`

If the token is missing or invalid, the server will respond with a `401 Unauthorized` error.

## Rate Limiting

The server may enforce rate limiting. If you exceed the configured rate limit, the server will respond with a `429 Too Many Requests` error.

## Endpoints

### `GET /api/search`

Searches for images across multiple booru sources.

#### Query Parameters

| Parameter | Type    | Description                                                                 | Example                               |
| :-------- | :------ | :-------------------------------------------------------------------------- | :------------------------------------ |
| `tags`    | string  | A comma-separated list of tags to search for.                               | `tags=waifu,neko`                     |
| `nsfw`    | boolean | Whether to include NSFW (Not Safe for Work) content. `true` or `false`.     | `nsfw=false`                          |
| `limit`   | integer | The maximum number of results to return from each provider.                 | `limit=10`                            |
| `width`   | integer | A minimum width for the images. (Only supported by some providers)          | `width=1920`                          |
| `height`  | integer | A minimum height for the images. (Only supported by some providers)         | `height=1080`                         |
| `orderBy` | string  | The order in which to sort the results. (Only supported by some providers)  | `orderBy=likes`                       |

#### Responses

- **`200 OK`**: A successful response will contain a JSON array of `Image` objects.
- **`401 Unauthorized`**: Returned if authentication is enabled and a valid `Authorization` header is not provided.
- **`429 Too Many Requests`**: Returned if the client has exceeded the rate limit.

**Image Object Structure:**

| Field       | Type      | Description                                                              |
| :---------- | :-------- | :----------------------------------------------------------------------- |
| `id`        | string    | The unique ID of the image from the provider.                            |
| `url`       | string    | The direct URL to the image file.                                        |
| `source`    | string    | The original source URL of the image (e.g., Pixiv, Twitter).             |
| `tags`      | []string  | A list of tags associated with the image.                                |
| `width`     | integer   | The width of the image in pixels.                                        |
| `height`    | integer   | The height of the image in pixels.                                       |
| `score`     | integer   | The score or number of favorites the image has.                          |
| `nsfw`      | boolean   | Whether the image is marked as NSFW.                                     |
| `createdAt` | time.Time | The timestamp when the image was created or uploaded.                    |
| `provider`  | string    | The name of the provider that supplied the image (e.g., "waifu.im").     |

**Example Success Response (`200 OK`):**

```json
[
  {
    "id": "7369",
    "url": "https://cdn.waifu.im/7369.jpg",
    "source": "https://www.pixiv.net/en/artworks/95252525",
    "tags": ["waifu"],
    "width": 1054,
    "height": 1500,
    "score": 2,
    "nsfw": false,
    "created_at": "2022-01-02T21:46:58.959235+01:00",
    "provider": "waifu.im"
  }
]
```
