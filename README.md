## goconsul

cli to list all the keys

### Synopsis

goconsul  is a library for Go to list all the keys and their values in consul-kv.

```
goconsul [flags]
```

### Options

```
  -a, --consulAddr string          Consul http address. (default "127.0.0.1:8500")
  -t, --consulToken string         Consul http token.
  -h, --help                       help for goconsul
  -o, --output json|yaml|xml|csv   Export all the content to json|yaml|xml|csv format.
  -p, --prefix string              Consul prefix. ex:- config/ (default "all")
```

