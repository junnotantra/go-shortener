# Go-shortener

Go-shortener is an URL shortener service created using golang and [bbolt](https://github.com/etcd-io/bbolt#comparison-with-other-databases)

## How to run

1. Download this repo

    Using git

    ```sh
    $ git clone https://github.com/junnotantra/go-shortener
    ```

    or using go tool

    ```sh
    $ go get https://github.com/junnotantra/go-shortener

2. Run for development environment using makefile

    ```sh
    make run-shortener
    ```

3. Or build shortener binary

    ```sh
    $ cd $GOPATH/src/github.com/junnotantra/go-shortener/cmd/shortener

    $ go build && ./shortener
    ```

4. Access the API on port `:3002` (default value)

## API

### Creating new short URL `/shortener/v1/new`

You can create new short URL by passing two parameters:

- full_url: the original URL where requet will be redirected
- custom_unique_str (optional): Unique string for the short URL. If not provided, the system will generate random string for you.

cURL example:

```sh
curl --request POST \
  --url http://<YOUR-DOMAIN>/shortener/v1/new \
  --header 'Content-Type: multipart/form-data' \
  --form full_url=<FULL-URL> \
  --form custom_unique_str=<OPTIONAL-CUSTOM-UNIQUE-STRING>
```

### Updating existing short URL `/shortener/v1/update`

You can update existing short URL by using it's existing unique string. This will replace the unique string, but retain the statistics.

cURL example:

```sh
curl --request POST \
  --url http://<YOUR-DOMAIN>/shortener/v1/update \
  --header 'Content-Type: multipart/form-data' \
  --form full_url=<FULL-URL> \
  --form custom_unique_str=<OPTIONAL-CUSTOM-UNIQUE-STRING>
```

### Get short URL info `/shortener/v1/info/<UNIQUE-STRING>`

Get basic info about created short URL.

```sh
curl --request GET \
  --url http://<YOUR-DOMAIN>/shortener/v1/info/<UNIQUE-STRING>
```

### Get short URL statistic `/statistic/v1/info/<UNIQUE-STRING>`

Get statistic of created short URL. Currently available data is only for number clicks. There are two types of staistics: overall and daily. Only last 30 days of daily data will be retained (old data will be deleted when more recent daily stats is added).

```sh
curl --request GET \
  --url http://<YOUR-DOMAIN>/statistic/v1/info/<UNIQUE-STRING>
```