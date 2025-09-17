# Waifu.im API Documentation

## Base URL

`https://api.waifu.im`

## Endpoints

### Search Images

Retrieves images randomly or by tag based on the specified search criteria.

*   **Method:** `GET`
*   **Endpoint:** `/search`
*   **Parameters:** See [Parameters](#parameters) section below.

**Example Request:**

```bash
curl -X GET 'https://api.waifu.im/search?included_tags=maid&height=>=2000'
```

**Example Response Body:**

```json
{
  "images": [
    {
      "artist": {
        "artist_id": 1,
        "deviant_art": "https://www.deviantart.com/4thwallzart",
        "name": "fourthwallzart",
        "patreon": "string",
        "pixiv": "string",
        "twitter": "https://twitter.com/4thWallzArt"
      },
      "byte_size": 3299586,
      "dominant_color": "#bbb7b2",
      "extension": ".png",
      "favorites": 1,
      "height": 2304,
      "image_id": 8108,
      "is_nsfw": false,
      "liked_at": "string",
      "preview_url": "https://www.waifu.im/preview/8108/",
      "signature": "58e6f0372364abda",
      "source": "https://www.patreon.com/posts/persephone-78224476",
      "tags": [
        {
          "description": "A female anime/manga character.",
          "is_nsfw": false,
          "name": "waifu",
          "tag_id": 12
        }
      ],
      "uploaded_at": "2023-05-03T18:40:04.381354+02:00",
      "url": "https://cdn.waifu.im/8108.png",
      "width": 1536
    }
  ]
}
```

### Get Tags

Get all the tags available.

*   **Method:** `GET`
*   **Endpoint:** `/tags`
*   **Parameters:**
    *   `full` (boolean): Returns more information about the tags, such as a description.

**Example Request:**

```bash
curl -X GET 'https://api.waifu.im/tags'
```

**Example Response Body:**

```json
{
  "versatile": [
    "maid",
    "waifu",
    "marin-kitagawa",
    "mori-calliope",
    "raiden-shogun",
    "oppai",
    "selfies",
    "uniform",
    "kamisato-ayaka"
  ],
  "nsfw": [
    "ass",
    "hentai",
    "milf",
    "oral",
    "paizuri",
    "ecchi",
    "ero"
  ]
}
```

## Parameters

These parameters can be used with the `/search` endpoint.

| Parameter         | Type           | Description                                                                                             |
| :---------------- | :------------- | :------------------------------------------------------------------------------------------------------ |
| `included_tags`   | array[string]  | Force the API to return images with at least all the provided tags.                                     |
| `excluded_tags`   | array[string]  | Force the API to return images without any of the provided tags.                                        |
| `included_files`  | array[string]  | Force the API to provide only the specified file IDs or signatures.                                     |
| `excluded_files`  | array[string]  | Force the API to not list the specified file IDs or signatures.                                         |
| `is_nsfw`         | string         | Default to `false`. Force or exclude lewd files. You can provide `null` to make it random.                |
| `gif`             | boolean        | Force or prevent the API to return .gif files.                                                          |
| `order_by`        | string         | Ordering criteria for the images.                                                                       |
| `orientation`     | string         | Image orientation criteria.                                                                             |
| `limit`           | integer        | Return an array of the number provided. A value greater than 30 requires admin permissions. Default is 1. |
| `full`            | boolean        | Returns the full result without any limit (admins only).                                                |
| `width`           | string         | Filter images by width (in pixels). Accepted operators: `<=`, `>=`, `>`, `<`, `!=`, `=`.                   |
| `height`          | string         | Filter images by height (in pixels). Accepted operators: `<=`, `>=`, `>`, `<`, `!=`, `=`.                  |
| `byte_size`       | string         | Filter images by byte size. Accepted operators: `<=`, `>=`, `>`, `<`, `!=`, `=`.                          |
