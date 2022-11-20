---
title: 'Installation'
date: 2019-02-11T19:27:37+10:00
weight: 1
---

To use this utility you should have it installed on your local system.

### Download released binary

1. Go to release page of the repo and download the appropriate released binary with regard to your OS and arch.

2. Put it in one of PATH directories

3. Run it by simply typing `helmfig` in your desired terminal.

### Build from source

1. Install a golang compiler (at least version 1.16).

2. Clone the project and compile it:

```bash
git clone https://github.com/snapp-incubator/helmfig.git
cd helmfig
go build .
```

3. Put your ```config.example.yml``` near the compiled binary and run it via:

```bash
./helmfig yaml
```

```
hugo version
```
