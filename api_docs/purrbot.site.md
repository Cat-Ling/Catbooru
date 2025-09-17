# PurrBot API Documentation

## Base URL

`https://api.purrbot.site/v2`

## Endpoints

The PurrBot API provides a large number of endpoints for SFW and NSFW images, as well as some miscellaneous endpoints.

### Image Endpoints

The image endpoints are divided into SFW (Safe for Work) and NSFW (Not Safe for Work) categories.

#### SFW Images

*   **Endpoint Structure:** `GET /img/sfw/<category>/<type>`
*   **Description:** Returns a random SFW image or GIF from the specified category.
*   **Example Request:**

    ```bash
    curl "https://api.purrbot.site/v2/img/sfw/pat/gif"
    ```

*   **Example Response:**

    ```json
    {
      "link": "https://cdn.purrbot.site/sfw/pat/gif/pat_001.gif",
      "error": false,
      "time": 0
    }
    ```

A full list of SFW categories can be found in the [official documentation](https://docs.purrbot.site/api/#sfw).

#### NSFW Images

*   **Endpoint Structure:** `GET /img/nsfw/<category>/<type>`
*   **Description:** Returns a random NSFW image or GIF from the specified category.
*   **Example Request:**

    ```bash
    curl "https://api.purrbot.site/v2/img/nsfw/neko/gif"
    ```

*   **Example Response:**

    ```json
    {
      "link": "https://cdn.purrbot.site/nsfw/neko/gif/neko_001.gif",
      "error": false,
      "time": 0
    }
    ```

A full list of NSFW categories can be found in the [official documentation](https://docs.purrbot.site/api/#nsfw).

### List Endpoints

These endpoints return a list of all available images for a given path.

*   **SFW List:** `GET /list/sfw/<path>`
*   **NSFW List:** `GET /list/nsfw/<path>`

### Miscellaneous Endpoints

#### OWOify Text

This endpoint "OWOifies" text.

*   **Method:** `POST` or `GET`
*   **Endpoint:** `/owoify`
*   **GET Parameters:** `text`, `replace-words`, `stutter`, `emoticons`
*   **POST Body:**

    ```json
    {
      "text": "Hello World!",
      "replace-words": true,
      "stutter": true,
      "emoticons": true
    }
    ```

#### API Info

*   **Method:** `GET`
*   **Endpoint:** `/info`
*   **Description:** Provides info about the API.
