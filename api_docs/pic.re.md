# Pic.re API Documentation

## Base URL

`https://pic.re`

## Endpoints

### Get Random Image

Returns a random anime image.

*   **Method:** `GET`
*   **Endpoint:** `/image`
*   **Parameters:** See [Parameters](#parameters) section below.

**Example Request:**

```bash
curl -I -XGET "https://pic.re/image"
```

**Example Response Headers:**

The image is returned directly. Additional information is available in the response headers.

```
HTTP/1.1 200 OK
Date: Sat, 15 May 2021 15:01:25 GMT
Content-Type: image/jpeg
Content-Length: 689008
Connection: keep-alive
image_id: 242637
image_source: https://www.pixiv.net/member_illust.php?mode=medium&illust_id=63026388
image_tags: aqua_eyes,armor,fate/extra,fate/extra_ccc,fate/grand_order,fate_(series),long_hair,meltryllis,polychromatic,purple_hair,tagme_(artist)
```

### Get Random Image (Redirect)

Returns a 301 redirect to a random anime image on a CDN. This is the recommended method for fetching images.

*   **Method:** `GET`
*   **Endpoint:** `/images`
*   **Parameters:** See [Parameters](#parameters) section below.

**Example Request:**

```bash
curl "https://pic.re/images"
```

### Get Random Image Metadata

Returns a JSON object containing metadata for a random image.

*   **Method:** `GET` or `POST`
*   **Endpoint:** `/image.json` (also available via `POST` to `/image`)
*   **Parameters:** See [Parameters](#parameters) section below.

**Example Request:**

```bash
curl "https://pic.re/image.json"
```

**Example Response Body:**

```json
{
    "file_url":"https://konachan.com/image/c2013204d3c186f0b95e433eea9bce15/Konachan.com%20-%20208569%20animal%20aqua_hair%20bird%20boots%20bow%20building%20city%20clouds%20gloves%20hat%20hinanawi_tenshi%20night%20red_eyes%20short_hair%20skirt_lift%20sky%20touhou%20tree%20waira%20water.jpg",
    "md5":"c2013204d3c186f0b95e433eea9bce15",
    "tags":[
        "animal",
        "aqua_hair",
        "bird",
        "boots",
        "bow",
        "building",
        "city",
        "clouds",
        "gloves",
        "hat",
        "hinanawi_tenshi",
        "night",
        "red_eyes",
        "short_hair",
        "skirt_lift",
        "sky",
        "touhou",
        "tree",
        "waira",
        "water"
    ],
    "width":2047,
    "height":1447,
    "source":"http://i4.pixiv.net/img-original/img/2015/10/24/00/31/25/53177335_p0.jpg",
    "author":"Flandre93",
    "has_children":false,
    "_id":208569
}
```

### Get Tags

Returns a list of available tags.

*   **Method:** `GET` or `POST`
*   **Endpoint:** `/tags`

**Example Request:**

```bash
curl "https://pic.re/tags"
```

**Example Response Body:**

```json
[
    {
        "name": "long_hair",
        "count": 44045
    },
    {
        "name": "original",
        "count": 24985
    },
    {
        "name": "short_hair",
        "count": 19695
    },
    {
        "name": "blush",
        "count": 17494
    }
]
```

## Parameters

These parameters can be used with the `/image`, `/images`, and `/image.json` endpoints.

| Parameter  | Type    | Default | Description                                                                                             |
| :--------- | :------ | :------ | :------------------------------------------------------------------------------------------------------ |
| `in`       | string  |         | Comma-separated list of tags to **include**.                                                              |
| `nin`      | string  |         | Comma-separated list of tags to **exclude**.                                                              |
| `id`       | integer |         | A specific image ID to retrieve. If used, all other filter parameters are ignored.                      |
| `compress` | boolean | `true`  | If `true`, loads the image in WebP format for faster loading.                                           |
| `min_size` | integer |         | Sets a minimum size for the image dimensions.                                                           |
| `max_size` | integer | `6144`  | Sets a maximum size for the image dimensions. If you need an image larger than 6144x6144, set this value. |
