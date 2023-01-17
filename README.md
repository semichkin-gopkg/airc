## Air toml config builder
A tool for generating .toml config for [air](https://github.com/cosmtrek/air) with using ENV

### Installing
`go install github.com/semichkin-gopkg/airc/cmd/airc@v0.0.2`

### Usage
`AIRC_ROOT=/app AIRC_BUILD_INCLUDE_EXT=go,html airc build -o path/to/output/.air.toml`
