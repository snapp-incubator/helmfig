# Helmfig

![Release workflow](https://github.com/snapp-incubator/helmfig/actions/workflows/release.yaml/badge.svg)

Are you tired of writing `values.yaml` for `configmap` of your project when you are helmifying them? Helmfig is a handy 
tool that can generate the content of your `configmap` object and its parameters for `values.yaml` based on a config
example file.

Currently, we just support YAML config structure, but we will support JSON and ENV in the future.

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
