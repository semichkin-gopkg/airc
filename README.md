## [Air](https://github.com/cosmtrek/air)-based utility for live reloading with config building by throwing env variables

### Installing
`go install github.com/semichkin-gopkg/airc/cmd/airc@v0.0.9`

### Usage
#### Build configuration
`AIRC_ROOT=/app AIRC_INCLUDE_EXT=go,html airc build -c path/to/output/.air.toml`
#### Run
`airc run -c path/to/output/.air.toml`
#### Build and Run
`AIRC_ROOT=/app AIRC_INCLUDE_EXT=go,html airc build-run -c path/to/output/.air.toml`

### Integration with docker
You can use [this](https://hub.docker.com/r/semichkin/airc/tags) docker image for building your go application with live reloading.  
Configuration for air stored inside a container, and you can modify it with AIRC_... env variables.

#### Examples
1. You can omit AIRC variables if your main.go file located in /app/cmd/app/main.go 
    ```yaml
    version: "3.9"
    services:
      some_go_application:
        image: semichkin/airc:latest
        volumes:
          - ./src/some_go_application:/app
    ```
2. If your main.go file located in another place, or you want to change other configuration parameters, you can specify it by AIRC variables
   ```yaml
    version: "3.9"
    services:
      some_go_application:
        image: semichkin/airc:latest
        environment:
          - AIRC_ROOT=/some_go_application
          - AIRC_SRC=$$AIRC_ROOT/cmd/api/main.go
          - AIRC_KILL_DELAY=2s
          - ETC=...
        volumes:
          - ./src/some_go_application:/some_go_application
    ```