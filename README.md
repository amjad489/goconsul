# goconsul
goconsul a cli which makes easy to list all the keys and its values available on `consul-kv` store. It also supports exporting all the kv store in flat files. Various file formats are supported. 
It also support `prefix` to list keys and its values in a specific folder/path.

## Installation
### Clone repo
```shell script
$ git clone https://github.com/amjad489/goconsul.git
```
### Build
```shell script
$ cd goconsul
$ go build
```
### Run
```shell script
$ ./goconsul
```
### Help
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

### List
```shell script
$ ./goconsul -a <consul_host>:8500 -t <token>
+----+---------------------------------------------------------+----------------------------+
| ID |                           KEY                           |           VALUE            |
+----+---------------------------------------------------------+----------------------------+
| 1  | elasticsearch/action.disable_close_all_indices          | true                       |
| 2  | elasticsearch/action.disable_delete_all_indices         | true                       |
| 3  | elasticsearch/action.disable_shutdown                   | true                       |
| 4  | elasticsearch/bootstrap.mlockall                        | true                       |
| 5  | elasticsearch/cluster.name                              | od-fts1                    |
| 6  | elasticsearch/gateway.expected_nodes                    | 2                          |
| 7  | elasticsearch/gateway.recover_after_nodes               | 1                          |
| 8  | elasticsearch/gateway.recover_after_time                | 10m                        |
| 9  | elasticsearch/index.number_of_replicas                  | 1                          |
| 10 | elasticsearch/index.number_of_shards                    | 2                          |
| 11 | elasticsearch/indices.recovery.max_bytes_per_sec        | 100mb                      |
| 12 | elasticsearch/node.data                                 | true                       |
| 13 | elasticsearch/node.master                               | true                       |
| 14 | elasticsearch/node.name                                 | od-fts1a                   |
| 15 | logstash/config.test_and_exit                           | false                      |
| 16 | logstash/dead_letter_queue.enable                       | false                      |
| 17 | logstash/dead_letter_queue.max_bytes                    | 1024mb                     |
| 18 | logstash/http.host                                      | 127.0.0.1                  |
| 19 | logstash/http.port                                      | 9600-9700                  |
| 20 | logstash/log.level                                      | info                       |
| 21 | logstash/monitoring.cluster_uuid                        | elasticsearch_cluster_uuid |
| 22 | logstash/monitoring.collection.interval                 | 10s                        |
| 23 | logstash/monitoring.collection.pipeline.details.enabled | true                       |
| 24 | logstash/monitoring.elasticsearch.cloud_id              | monitoring_cluster_id      |
| 25 | logstash/monitoring.elasticsearch.password              | password                   |
| 26 | logstash/monitoring.elasticsearch.sniffing              | false                      |
| 27 | logstash/monitoring.elasticsearch.ssl.keystore.password | password                   |
| 28 | logstash/monitoring.elasticsearch.ssl.keystore.path     | /path/to/file              |
| 29 | logstash/monitoring.elasticsearch.ssl.truststore.passwo | password                   |
| 30 | logstash/monitoring.elasticsearch.ssl.truststore.path   | path/to/file               |
| 31 | logstash/monitoring.elasticsearch.ssl.verification_mode | certificate                |
| 32 | logstash/monitoring.elasticsearch.username              | logstash_system            |
| 33 | logstash/monitoring.enabled                             | false                      |
| 34 | logstash/path.dead_letter_queue                         | asdad                      |
| 35 | logstash/pipeline.batch.delay                           | 50                         |
| 36 | logstash/pipeline.batch.size                            | 125                        |
| 37 | logstash/pipeline.id                                    | main                       |
| 38 | logstash/pipeline.ordered                               | auto                       |
| 39 | logstash/pipeline.unsafe_shutdown                       | false                      |
| 40 | logstash/pipeline.workers                               | 2                          |
| 41 | logstash/queue.checkpoint.acks                          | 1024                       |
| 42 | logstash/queue.checkpoint.interval                      | 1000                       |
| 43 | logstash/queue.checkpoint.writes                        | 1024                       |
| 44 | logstash/queue.max_bytes                                | 1024mb                     |
| 45 | logstash/queue.type                                     | memory                     |
| 46 | logstash/xpack.management.elasticsearch.password        | password                   |
| 47 | logstash/xpack.management.elasticsearch.sniffing        | false                      |
| 48 | logstash/xpack.management.elasticsearch.ssl.keystore.pa | password                   |
| 49 | logstash/xpack.management.elasticsearch.ssl.keystore.pa | /path/to/file              |
| 50 | logstash/xpack.management.elasticsearch.ssl.truststore. | password                   |
| 51 | logstash/xpack.management.elasticsearch.ssl.truststore. | /path/to/file              |
| 52 | logstash/xpack.management.elasticsearch.ssl.verificatio | certificate                |
| 53 | logstash/xpack.management.elasticsearch.username        | logstash_admin_user        |
| 54 | logstash/xpack.management.enabled                       | false                      |
| 55 | logstash/xpack.management.logstash.poll_interval        | 5s                         |
+----+---------------------------------------------------------+----------------------------+
Total keys: 55
```

### Export
```shell script
$ ./goconsul -a <consul_host>:8500 -t <token> -o json
data exported to file: data.json
```
### list with prefix
```shell script
$ ./goconsul -a <consul_host>:8500 -t <token> -p config/
```
### export with prefix
```shell script
$ ./goconsul -a <consul_host>:8500 -t <token> -p config/ -o json
```
