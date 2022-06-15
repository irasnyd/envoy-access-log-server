# Envoy Access Log Server

Dummy implementation of an Envoy Access Log Server. Dumps all output received
to `stdout` for further collection and processing by a full featured log
aggregation system.

## Container Image

```bash
$ docker build --pull -t irasnyd/envoy-access-log-server:latest .
$ docker push irasnyd/envoy-access-log-server:latest
```

## Kudos

This project was heavily inspired by: <https://github.com/dio/metricsink>

Also: <https://github.com/istio/istio/tree/09c12541385aeeebb6e464b675d966a1aa96bc5e/samples/als-sink>

## License

This project is licensed under the [MIT License](https://choosealicense.com/licenses/mit/).
For further details please refer to the [LICENSE](./LICENSE) file in this repository.
