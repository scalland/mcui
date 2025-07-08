# mcui

A Memcached UI

Create a YAML config file and use the same to run this software:

```
app:
  host: "localhost"
  port: 8080
memcached:
  mc_host: "localhost"
  mc_port: 11211
```
Definitions
---

| Section     | Key       | Description                          |
|-------------|-----------|--------------------------------------|
| app         | host      | The host where the server will listen on |
|             | port      | The port on which the server will listen |
| memcached   | mc_host   | Memcached Hostname                   |
|             | mc_port   | Memcached Port number                |


Running the server

```
mcui serve -c /path/to/config.yaml
```

**Note:** If you have the config.yaml file in the current directory, it will be read on its own.
