# Final Project

The final project, you will create `Rest API with Gin Gonic`. The requirements are below.

- You have to create three tables. The tables are `artists`, `albums` and `songs`. `An artist` can have `many albums` and `an album` can have `many songs`.
- You have to use `Postgres` for database.
- You can use `Redis` for caching (optional).
- You have upload the rest api to `Railway` or `another Cloud` that support `Go`.

### Response Format

```
{
  "status": bool,
  "message": string,
  "errors": []string,
  "data": interface{}
}
```

- `status` must be a boolean, if success it will `true` and error it will `false`
- `message` must be a string
- `errors` must be array of strings
- `data` can be an `object` or `array object` or `can be anything`

### API Lists

| API Address  | Method | Description                                                                                                                                                                                                                     | Cache (Optional) |
| ------------ | ------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------- |
| /artists     | POST   | Create an artist                                                                                                                                                                                                                | Yes              |
| /artists     | GET    | Get all artists or (optional) you can add filter and pagination                                                                                                                                                                 | No               |
| /artists/:id | GET    | Get details information of an artist                                                                                                                                                                                            | Yes              |
| /artists/:id | PUT    | Update an artist                                                                                                                                                                                                                | Yes              |
| /artists/:id | DELETE | Delete an artist or (optional) you can use `soft delete` if you want                                                                                                                                                            | Yes              |
| /albums      | POST   | Create an album                                                                                                                                                                                                                 | Yes              |
| /albums      | GET    | Get all albums, can filter albums by id of artists and include who the artist is or (optional) you can add pagination, ex: `/albums?artists_id=1`, it will show the albums of an specific artists                               | No               |
| /albums/:id  | GET    | Get details information of an album include who is the artist                                                                                                                                                                   | Yes              |
| /albums/:id  | PUT    | Update an album                                                                                                                                                                                                                 | Yes              |
| /albums/:id  | DELETE | Delete an album or (optional) you can use `soft delete` if you want                                                                                                                                                             | Yes              |
| /songs       | POST   | Create an song                                                                                                                                                                                                                  | Yes              |
| /songs       | GET    | Get songs of an artists or of an album and include who is the artist and what the album or (optional) you can add pagination, ex: `/songs?artists_id=1` or `/songs?albums_id=1`, it will show songs by specific artist or album | No               |
| /songs/:id   | GET    | Get details information of an song include who is the artist and what the album                                                                                                                                                 | Yes              |
| /songs/:id   | PUT    | Update an song                                                                                                                                                                                                                  | Yes              |
| /songs/:id   | DELETE | Delete an song or (optional) you can use `soft delete` if you want                                                                                                                                                              | Yes              |

**Notes**:

- Cache (Optional) `Yes` means that you can add set/get/del data to/from Cache (Redis).
- Cache (Optional) `No` means that you do not need to add cache for he specific API.
- Cache is not required, it is optional but if you want to go with it, just do it.

### Tables

#### Artists

| Column | Type      | Description    | Required |
| ------ | --------- | -------------- | -------- |
| id     | BIGSERIAL | Primary key    | Yes      |
| name   | VARCHAR   | Name of artist | Yes      |

#### Albums

| Column    | Type    | Description           | Required |
| --------- | ------- | --------------------- | -------- |
| id        | SERIAL  | Primary key           | Yes      |
| artist_id | BIGINT  | Foreign key of artist | Yes      |
| title     | VARCHAR | Name of album         | Yes      |
| price     | NUMERIC | Price of album        | Yes      |

#### Songs

| Column   | Type    | Description          | Required |
| -------- | ------- | -------------------- | -------- |
| id       | SERIAL  | Primary key          | Yes      |
| album_id | BIGINT  | Foreign key of album | Yes      |
| title    | VARCHAR | Name of album        | Yes      |
| lyrics   | VARCHAR | Lyrics               | No       |
