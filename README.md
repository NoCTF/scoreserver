# Score Server for NoCTF

## CommandLine Options

| Name | Description |
|:-----|:------------|
| config | config filename without extension. you can use JSON/TOML/YAML/HCL. |
| port | listen port of this application. default is ":8080". |

## Configfile Options

Configfile Options overrides the value of CommandLine Options.
"."(dot) syntax is representation of nest.

| Name | Description |
|:-----|:------------|
| port | listen port of this application. |
| database.DSN | data source name for mysql database driver. default value is "root@tcp(localhost:3306)/noctf?parseTime=true". |

## Generation

some model and db scripts are auto-generated.
to generate, run below

``` shell
$ go build internal/genmodel/*.go
$ go build internal/gendb/*.go
$ PATH=.:${PATH} go generate .
```
