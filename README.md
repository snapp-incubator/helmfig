# Helmfig

![Release workflow](https://github.com/snapp-incubator/helmfig/actions/workflows/release.yaml/badge.svg)

Are you tired of writing `values.yaml` for `configmap` of your project when you are helmifying them? Helmfig is a handy 
tool that can generate the content of your `configmap` object and its parameters for `values.yaml` based on a config
example file.

Currently, we just support YAML config structure, but we will support JSON and ENV in the future.

## What actually it does

Assume that you have a file named `config.example.yaml` like the one below. You want to helmify your project, so you need
a `configmap` object file and an appropriate values for that `configmap`. Helmfig generate them for you.

`config.example.yaml`
```yaml
logger:
  level: info
  syslog:
    enabled: false
    server_address: localhost:514
    network: udp
    priority: info
nats_streaming:
  address: "127.0.0.1:4222"
  connect_wait: 1s
  pub_ack_wait: 2s
  max_pub_acks_in_flight: 5
  ping_interval: 2
  ping_max_out: 150
  cluster_id: test-cluster
  client_id: app
```

`Configmap` output:
```yaml
logger:
  level: '{{ logger.level }}'
  syslog:
    enabled: '{{ logger.syslog.enabled }}'
    network: '{{ logger.syslog.network }}'
    priority: '{{ logger.syslog.priority }}'
    server_address: '{{ logger.syslog.serverAddress }}'
nats_streaming:
  address: '{{ natsStreaming.address }}'
  client_id: '{{ natsStreaming.clientId }}'
  cluster_id: '{{ natsStreaming.clusterId }}'
  connect_wait: '{{ natsStreaming.connectWait }}'
  max_pub_acks_in_flight: '{{ natsStreaming.maxPubAcksInFlight }}'
  ping_interval: '{{ natsStreaming.pingInterval }}'
  ping_max_out: '{{ natsStreaming.pingMaxOut }}'
  pub_ack_wait: '{{ natsStreaming.pubAckWait }}'
```

`Values` output:
```yaml
logger:
  level: info
  syslog:
    enabled: false
    network: udp
    priority: info
    serverAddress: localhost:514
natsStreaming:
  address: 127.0.0.1:4222
  clientId: app
  clusterId: test-cluster
  connectWait: 1s
  maxPubAcksInFlight: 5
  pingInterval: 2
  pingMaxOut: 150
  pubAckWait: 2s
```

## How to use it?

### Download released binary

1. Go to release page of the repo and download the appropriate released binary with regard to your OS and arch.

2. Put it in one of PATH directories

3. Run it by simply typing `helmfig` in your desired terminal.

### Build from source

1. Install a golang compiler (at least version 1.16).
2. Clone the project and compile it:
~~~bash
git clone https://github.com/snapp-incubator/helmfig.git
cd helmfig
go build .
~~~
3. Put your ```config.example.yml``` near the compiled binary and run it via:
~~~bash
./helmfig yaml
~~~
4. If everything is OK, two files will be generated: ```configmap.yaml``` and ```values.yaml```. You can use them in
helm chart of your desired application

## License

Apache-2.0 License, see [LICENSE](LICENSE).
