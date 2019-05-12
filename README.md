# Speedtest with Docker

```docker run -v $PWD/config.json:/home/config.json edasque/speedtest```

## Building the main app & docker container

If you want to build this yourself, this should work: 

### Building the main executable:

```GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build  -o speedtest.linux.amd64 speedtest.go```

### Building the docker image:
```docker build -t speedtest .```

