# Rewind Topic Spike

This small project evaluates how easy it is to rewind a kafka topic partition by offset or timestamp.

## Description

This project uses the `github.com/segmentio/kafka-go` library internally. That library exposes a function on the `Reader` struct, [SetOffsetAt](https://pkg.go.dev/github.com/segmentio/kafka-go#Reader.SetOffsetAt), that accomplishes what the related spike was looking for.

The project creates four kafka readers, one for each partition, each one has a different behavior as explained here:

- `rewind-reader-0` waits for 60s and then rewinds 10s, starting at the instant 20
- `rewind-reader-1` doesn't wait and rewinds 30s, should start at the instant 0;
- `rewind-reader-2` waits for 10s and then rewinds 10s, starting at the instant 0;
- `rewind-reader-3` doesn't wait and doesn't rewind.

Where executing the project they all start fresh at the same time as the kafka writer.

Each container behavior can be seen in their logs with `docker compose logs <container_name> -f`.

## Running

```sh
docker compose up --build -d
```

## Outcome

The function `SetOffsetAt` works as expected, only when no Consumer Group is defined.
