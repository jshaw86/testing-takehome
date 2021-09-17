# Summary
This is a simple Golang Rest API wrapper for Redis get and set functions. The purpose of this is to provide a SWET candidates the to show their ability in regard to contract testing an api. The repo ships with a functional docker-compose setup and once running candidates can make HTTP requests to their localhost and actually persist and fetch data through the rest API on port 8080.

## Task
Build a "framework" to contract test or integration test the Golang API to ensure the application is behaving as expected. As a Software Developer/Engineer in Test your primary focus is building amazing tooling/frameworks so that Software Engineers/Developers can build tests and retest as fast and reliably as possible. So keep that in mind when developing this tooling.

The expectation of this project isn't that you know Golang, docker or any other langauge, infrastructure or tool used, but rather do you understand the concepts of interfaces and how to effectively test those interfaces. If you have any questions about any of the tools used feel free to reach out to the hiring coordinator and we'll do our best to answer your questions.

### Deliverables 
1. Implement a test tool to correctly asserts the behavior of the api 
2. Provide a HOWTO explaining:
    - why you chose contract or integration testing for this exercise
    - how to run your test tooling
    - extra considerations you made or overlooked like race conditions etc..
3. Tar or Zip your project and HOWTO and send it back to the hiring coordinator


## Contract testing vs Integration testing
Sometimes these are conflated or have been misinterpreted over time, this website provides a good defintion of contract vs integration testing https://docs.pact.io/consumer/contract_tests_not_functional_tests/. TLDR; contract testing does not consider side effects rather just if the data was passed correctly between the provider and consumer. A side effect in this case is the storage of data to redis. 

# Setup
This is the recommended setup procedure as it reduces the varibility of your local environment. If you're comfortable with golang it is possible to run the application natively.

1. Ensure docker is installed for macOS see here: https://docs.docker.com/desktop/mac/install/ for ubuntu or other linux distributions see here: https://docs.docker.com/engine/install/ubuntu/
2. Run `docker-compose up --build` from the root directory of the project
3. The docker containers are ready to be used when you see output similar to this:
```
docker-compose up --build
Docker Compose is now in the Docker CLI, try `docker compose up`

Building api
[+] Building 7.9s (16/16) FINISHED
 => [internal] load build definition from Dockerfile                                                                                   0.0s
 => => transferring dockerfile: 44B                                                                                                    0.0s
 => [internal] load .dockerignore                                                                                                      0.0s
 => => transferring context: 2B                                                                                                        0.0s
 => [internal] load metadata for docker.io/library/alpine:3.11.3                                                                       0.9s
 => [internal] load metadata for docker.io/library/golang:1.13-alpine3.11                                                              0.9s
 => [auth] library/alpine:pull token for registry-1.docker.io                                                                          0.0s
 => [auth] library/golang:pull token for registry-1.docker.io                                                                          0.0s
 => [build 1/5] FROM docker.io/library/golang:1.13-alpine3.11@sha256:ec6dcf15073c307fbcfc3149efe8835f3ec2bd0a0cb49aaaee4949cfc4c86b65  0.0s
 => [internal] load build context                                                                                                      0.0s
 => => transferring context: 1.24kB                                                                                                    0.0s
 => [stage-1 1/3] FROM docker.io/library/alpine:3.11.3@sha256:ab00606a42621fb68f2ed6ad3c88be54397f981a7b70a79db3d1172b11c4367d         0.0s
 => CACHED [build 2/5] WORKDIR /app                                                                                                    0.0s
 => [build 3/5] COPY . /app                                                                                                            0.0s
 => [build 4/5] RUN apk add make                                                                                                       0.8s
 => [build 5/5] RUN make build                                                                                                         6.0s
 => CACHED [stage-1 2/3] WORKDIR /app                                                                                                  0.0s
 => CACHED [stage-1 3/3] COPY --from=build  /app/api .                                                                                 0.0s
 => exporting to image                                                                                                                 0.0s
 => => exporting layers                                                                                                                0.0s
 => => writing image sha256:b14fbf5df04ed0432f376078d49502ada4db20879ca8334f866d020e0be0165b                                           0.0s
 => => naming to docker.io/library/testing-takehome_api                                                                                0.0s

Use 'docker scan' to run Snyk tests against images to find vulnerabilities and learn how to fix them
Starting testing-takehome_redis_1 ... done
Starting testing-takehome_api_1   ... done
Attaching to testing-takehome_redis_1, testing-takehome_api_1
redis_1  | 1:C 17 Sep 2021 04:17:04.462 # oO0OoO0OoO0Oo Redis is starting oO0OoO0OoO0Oo
redis_1  | 1:C 17 Sep 2021 04:17:04.462 # Redis version=6.2.5, bits=64, commit=00000000, modified=0, pid=1, just started
redis_1  | 1:C 17 Sep 2021 04:17:04.462 # Warning: no config file specified, using the default config. In order to specify a config file use redis-server /path/to/redis.conf
api_1    | 2021/09/17 04:17:05 Listening https://0.0.0.0:8080
redis_1  | 1:M 17 Sep 2021 04:17:04.463 * monotonic clock: POSIX clock_gettime
redis_1  | 1:M 17 Sep 2021 04:17:04.464 * Running mode=standalone, port=6379.
redis_1  | 1:M 17 Sep 2021 04:17:04.464 # Server initialized
redis_1  | 1:M 17 Sep 2021 04:17:04.465 * Loading RDB produced by version 6.2.5
redis_1  | 1:M 17 Sep 2021 04:17:04.466 * RDB age 985 seconds
redis_1  | 1:M 17 Sep 2021 04:17:04.466 * RDB memory usage when created 0.77 Mb
redis_1  | 1:M 17 Sep 2021 04:17:04.466 * DB loaded from disk: 0.001 seconds
redis_1  | 1:M 17 Sep 2021 04:17:04.466 * Ready to accept connections
```

# Using the API
To interact with the api you can simply run the following two curl requests. First, set a key with the `PUT` verb to `/api/:key` with a payload which will be set to the value of the `:key` in redis. Second, fetch the key with the implicit `GET` to `/api/:key` and the api should return the prior payload out of redis and return as a simple string. 

```
> curl -XPUT -d"something" http://localhost:8080/api/test
OK%
> curl http://localhost:8080/api/test
something%
```

