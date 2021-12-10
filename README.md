# Helmfig

Are you tired of writing `values.yaml` for `configmap` of your project when you are helmifying them? Helmfig is a handy 
tool that can generate the content of your `configmap` object and its parameters for `values.yaml` based on a config
example file.

Currently, we just support YAML config structure, but we will support JSON and ENV in the future.

## How to use it?

### Build from source

1. Install a golang compiler (at least version 1.16).
2. Clone the project and compile it:
~~~bash
git clone 
cd 
go build .
~~~
3. Put your ```config.example.yml``` near the compiled binary and run it via:
~~~bash
./helmfig
~~~
3. If everything is OK, two files will be generated: ```configmap.yaml``` and ```values.yaml```. You can use them in
helm chart of your desired application
