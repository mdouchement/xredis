# xredis

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/mdouchement/xredis)

Handy wrapper around `github.com/redis/go-redis/v9` package. All the Redis client methods are usable on the xredis client.

## Usage

- Script
    ```go
    const lua = `local age=redis.call("get", KEYS[1])
    local result="Hello "..ARGV[1].." "..KEYS[1]..", you are "..age.." years old!"
    return result`

    func main() {
        ctx := context.Background()
        xclient := xredis.New(redisClient)

        script := xredis.NewScript(lua)

        err := xclient.Set(ctx, "Abitbol", 42, time.Hour).Err()
        if err != nil {
            panic(err)
        }

        value, err := xclient.Run(ctx, script, []string{"Abitbol"}, "George")
        if err != nil {
            panic(err)
        }

        fmt.Println(value.String())
        // => Hello George Abitbol, you are 42 years old!
    }
    ```

- Value object is a JSON value?
    ```go
    json.Unmarshal(value.Bytes(), &object)
    ```

## License

**MIT**


## Contributing

All PRs are welcome.

1. Fork it
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
5. Push to the branch (git push origin my-new-feature)
6. Create new Pull Request