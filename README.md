## Air toml config builder
A tool for generating .toml config for [air](https://github.com/cosmtrek/air) with using ENV.  
You can see the full list of supported env variables [here](https://github.com/semichkin-gopkg/airc/blob/main/internal/cmd/build.go#L14)

### Installing
`go install github.com/semichkin-gopkg/airc/cmd/airc@v0.0.6`

### Usage
`AIRC_ROOT=/app AIRC_BUILD_INCLUDE_EXT=go,html airc build -o path/to/output/.air.toml`

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
2. If your main.go file located in another place, you can specify it by AIRC variables
   ```yaml
    version: "3.9"
    services:
      some_go_application:
        image: semichkin/airc:latest
        environment:
          - AIRC_ROOT=/some_go_application
          - AIRC_BUILD_CMD_PATH_TO_SOURCE=$$AIRC_ROOT/cmd/api/main.go
          - AIRC_BUILD_KILL_DELAY=2s
          - ETC=...
        volumes:
          - ./src/some_go_application:/some_go_application
    ```