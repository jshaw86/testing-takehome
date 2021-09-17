package main

import (
    "fmt"
    "net/http"
    "errors"
    "strings"
    "net/url"
    "os"
    "io/ioutil"
    "github.com/go-redis/redis/v8"
    "context"
    "log"
)

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}

func handleReq(req *http.Request) ([]byte, error) {
    defer req.Body.Close()

    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        return nil, err
    }

    return body, nil
}

func getKeyFromURL(u *url.URL) (string, error) {
    uriFragments := strings.Split(u.Path, "/")
    if len(uriFragments) == 3 {
        return uriFragments[2], nil

    }

    return "", errors.New(fmt.Sprintf("malformed get request should be /api/:key given %s\n", u.Path))
    
}

func upsertKey(ctx context.Context, key string, r *http.Request, rdb *redis.Client) (interface{}, error) {
    body, err := handleReq(r)

    if err != nil {
        return nil, err
    }

    val, err := rdb.Set(ctx, key, body, 0).Result()

    return val, err
}

func main() {
    redisAddr := getEnv("REDIS_ADDR", "localhost:6379") 
    rdb := redis.NewClient(&redis.Options{
        Addr:     redisAddr,
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    ctx := context.Background()

    http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
        key, keyErr := getKeyFromURL(r.URL)

        if keyErr != nil {
            http.Error(w, keyErr.Error(), 404)
            return
        }

        var val interface{} 
        var err error 
        if r.Method == "PUT" {
            val, err = upsertKey(ctx, key, r, rdb)
        } else if r.Method == "GET" {
            val, err = rdb.Get(ctx, key).Result()
        } else {
            err = errors.New("Method Not Supported")
        }

        if (err == nil) {
            w.Write([]byte(val.(string)))
            return
        } 
        
        http.Error(w, err.Error(), 400)

    })

    srv := &http.Server{Addr: ":8080"}
    log.Printf("Listening https://0.0.0.0:8080")
    log.Fatal(srv.ListenAndServe())

}
