# Waifu.pics API Documentation

## Base URL

`https://api.waifu.pics`

## Endpoints

The API has two types of endpoints: SFW (Safe for Work) and NSFW (Not Safe for Work).

### SFW Endpoints

These endpoints return SFW images.

#### Get Single SFW Image

*   **Method:** `GET`
*   **Endpoint:** `/sfw/<category>`

**Example Request:**

```bash
curl "https://api.waifu.pics/sfw/waifu"
```

**Example Response:**

```json
{
  "url": "https://i.waifu.pics/some-image.jpg"
}
```

#### Get Multiple SFW Images

*   **Method:** `POST`
*   **Endpoint:** `/many/sfw/<category>`
*   **Body:**

    ```json
    {
      "exclude": []
    }
    ```

**Example Request:**

```bash
curl -X POST -H "Content-Type: application/json" -d '{"exclude": []}' "https://api.waifu.pics/many/sfw/waifu"
```

**Example Response:**

```json
{
  "files": [
    "https://i.waifu.pics/another-image.jpg",
    "https://i.waifu.pics/a-third-image.png"
  ]
}
```

**SFW Categories:** `waifu`, `neko`, `shinobu`, `bully`, `cry`, `hug`, `kiss`, `lick`, `pat`, `smug`, `highfive`, `nom`, `bite`, `slap`, `wink`, `poke`, `dance`, `cringe`, `blush`

### NSFW Endpoints

These endpoints return NSFW images.

#### Get Single NSFW Image

*   **Method:** `GET`
*   **Endpoint:** `/nsfw/<category>`

**Example Request:**

```bash
curl "https://api.waifu.pics/nsfw/waifu"
```

**Example Response:**

```json
{
  "url": "https://i.waifu.pics/some-nsfw-image.jpg"
}
```

#### Get Multiple NSFW Images

*   **Method:** `POST`
*   **Endpoint:** `/many/nsfw/<category>`
*   **Body:**

    ```json
    {
      "exclude": []
    }
    ```

**Example Request:**

```bash
curl -X POST -H "Content-Type: application/json" -d '{"exclude": []}' "https://api.waifu.pics/many/nsfw/waifu"
```

**Example Response:**

```json
{
  "files": [
    "https://i.waifu.pics/another-nsfw-image.jpg",
    "https://i.waifu.pics/a-third-nsfw-image.png"
  ]
}
```

**NSFW Categories:** `waifu`, `neko`, `trap`, `blowjob`
