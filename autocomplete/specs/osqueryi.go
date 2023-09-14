// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["osqueryi"] = model.Subcommand{
		Name:        []string{"osqueryi"},
		Description: `Your OS as a high-performance relational database`,
		Options: []model.Option{{
			Name:        []string{"--flagfile"},
			Description: `Line-delimited file of additional flags`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "path",
			}},
		}, {
			Name:        []string{"--D"},
			Description: `Run as a daemon process`,
		}, {
			Name:        []string{"--S"},
			Description: `Run as a shell process`,
		}, {
			Name:        []string{"--alarm_timeout"},
			Description: `Seconds to allow for shutdown. Minimum is 10`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--carver_block_size"},
			Description: `Size of blocks used for POSTing data back to remote endpoints`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--carver_compression"},
			Description: `Compress archives using zstd prior to upload (default false)`,
		}, {
			Name:        []string{"--carver_continue_endpoint"},
			Description: `TLS/HTTPS endpoint that receives carved content after session creation`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--carver_disable_function"},
			Description: `Disable the osquery file carver function (default true)`,
		}, {
			Name:        []string{"--carver_expiry"},
			Description: `Seconds to store successful carve result metadata (in carves table)`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--carver_start_endpoint"},
			Description: `TLS/HTTPS init endpoint for forensic carver`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--config_accelerated_refresh"},
			Description: `Interval to wait if reading a configuration fails`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--config_check"},
			Description: `Check the format of an osquery config and exit`,
		}, {
			Name:        []string{"--config_dump"},
			Description: `Dump the contents of the configuration, then exit`,
		}, {
			Name:        []string{"--config_enable_backup"},
			Description: `Backup config and use it when refresh fails`,
		}, {
			Name:        []string{"--config_path"},
			Description: `Path to JSON config file`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "value",
			}},
		}, {
			Name:        []string{"--config_plugin"},
			Description: `Config plugin name`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--config_refresh"},
			Description: `Optional interval in seconds to re-read configuration`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--config_tls_endpoint"},
			Description: `TLS/HTTPS endpoint for config retrieval`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--config_tls_max_attempts"},
			Description: `Number of attempts to retry a TLS config request`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--daemonize"},
			Description: `Attempt to daemonize (POSIX only)`,
		}, {
			Name:        []string{"--database_dump"},
			Description: `Dump the contents of the backing store`,
		}, {
			Name:        []string{"--database_path"},
			Description: `If using a disk-based backing store, specify a path`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "value",
			}},
		}, {
			Name:        []string{"--disable_carver"},
			Description: `Disable the osquery file carver (default true)`,
		}, {
			Name:        []string{"--disable_enrollment"},
			Description: `Disable enrollment functions on related config/logger plugins`,
		}, {
			Name:        []string{"--disable_extensions"},
			Description: `Disable extension API`,
		}, {
			Name:        []string{"--disable_reenrollment"},
			Description: `Disable re-enrollment attempts if related plugins return invalid`,
		}, {
			Name:        []string{"--disable_tables"},
			Description: `Comma-delimited list of table names to be disabled`,
			Args: []model.Arg{{
				Name:      "value",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--disable_watchdog"},
			Description: `Disable userland watchdog process`,
		}, {
			Name:        []string{"--enable_extensions_watchdog"},
			Description: `Enable userland watchdog for extensions processes`,
		}, {
			Name:        []string{"--enable_tables"},
			Description: `Comma-delimited list of table names to be enabled`,
			Args: []model.Arg{{
				Name:      "value",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--enroll_always"},
			Description: `On startup, send a new enrollment request`,
		}, {
			Name:        []string{"--enroll_secret_env"},
			Description: `Name of environment variable holding enrollment-auth secret`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--enroll_secret_path"},
			Description: `Path to an optional client enrollment-auth secret`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "value",
			}},
		}, {
			Name:        []string{"--enroll_tls_endpoint"},
			Description: `TLS/HTTPS endpoint for client enrollment`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--extensions_autoload"},
			Description: `Optional path to a list of autoloaded & managed extensions`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "value",
			}},
		}, {
			Name:        []string{"--extensions_interval"},
			Description: `Seconds delay between connectivity checks`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--extensions_require"},
			Description: `Comma-separated list of required extensions`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--extensions_socket"},
			Description: `Path to the extensions UNIX domain socket`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "value",
			}},
		}, {
			Name:        []string{"--extensions_timeout"},
			Description: `Seconds to wait for autoloaded extensions`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--force"},
			Description: `Force osqueryd to kill previously-running daemons`,
		}, {
			Name:        []string{"--install"},
			Description: `Install osqueryd as a service`,
		}, {
			Name:        []string{"--logger_mode"},
			Description: `Octal mode for log files (default '0640')`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--logger_plugin"},
			Description: `Logger plugin name`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--logger_stderr"},
			Description: `Write status logs to stderr`,
		}, {
			Name:        []string{"--logtostderr"},
			Description: `Log messages to stderr in addition to the logger plugin(s)`,
		}, {
			Name:        []string{"--pidfile"},
			Description: `Path to the daemon pidfile mutex`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "value",
			}},
		}, {
			Name:        []string{"--proxy_hostname"},
			Description: `Optional HTTP proxy hostname`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--stderrthreshold"},
			Description: `Stderr log level threshold`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--tls_client_cert"},
			Description: `Optional path to a TLS client-auth PEM certificate`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "value",
			}},
		}, {
			Name:        []string{"--tls_client_key"},
			Description: `Optional path to a TLS client-auth PEM private key`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "value",
			}},
		}, {
			Name:        []string{"--tls_enroll_max_attempts"},
			Description: `The total number of attempts that will be made to the enroll endpoint if a request fails, 0 for infinite`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--tls_enroll_max_interval"},
			Description: `Maximum wait time in seconds between enroll retry attempts`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--tls_hostname"},
			Description: `TLS/HTTPS hostname for Config, Logger, and Enroll plugins`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--tls_server_certs"},
			Description: `Optional path to a TLS server PEM certificate(s) bundle`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "value",
			}},
		}, {
			Name:        []string{"--tls_session_reuse"},
			Description: `Reuse TLS session sockets`,
		}, {
			Name:        []string{"--tls_session_timeout"},
			Description: `TLS session keep alive timeout in seconds`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--uninstall"},
			Description: `Uninstall osqueryd as a service`,
		}, {
			Name:        []string{"--watchdog_delay"},
			Description: `Initial delay in seconds before watchdog starts`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--watchdog_forced_shutdown_delay"},
			Description: `Seconds that the watchdog will wait to do a forced shutdown after a graceful shutdown request, when a resource limit is hit`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--watchdog_latency_limit"},
			Description: `Override watchdog profile CPU utilization latency limit`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--watchdog_level"},
			Description: `Performance limit level`,
			Args: []model.Arg{{
				Name: "value",
				Suggestions: []model.Suggestion{{
					Name:        []string{`0`},
					Description: `Normal`,
				}, {
					Name:        []string{`1`},
					Description: `Restrictive`,
				}, {
					Name:        []string{`-1`},
					Description: `Off`,
				}},
			}},
		}, {
			Name:        []string{"--watchdog_memory_limit"},
			Description: `Override watchdog profile memory limit (e.g., 300, for 300MB)`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--watchdog_utilization_limit"},
			Description: `Override watchdog profile CPU utilization limit`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--audit_allow_config"},
			Description: `Allow the audit publisher to change auditing configuration`,
		}, {
			Name:        []string{"--audit_allow_fim_events"},
			Description: `Allow the audit publisher to install filesystem-related rules`,
		}, {
			Name:        []string{"--audit_allow_process_events"},
			Description: `Allow the audit publisher to install process-related rules`,
		}, {
			Name:        []string{"--audit_allow_sockets"},
			Description: `Allow the audit publisher to install socket-related rules`,
		}, {
			Name:        []string{"--audit_allow_user_events"},
			Description: `Allow the audit publisher to install user-related rules`,
		}, {
			Name:        []string{"--augeas_lenses"},
			Description: `Directory that contains augeas lenses files`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--aws_access_key_id"},
			Description: `AWS access key ID`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--aws_debug"},
			Description: `Enable AWS SDK debug logging`,
		}, {
			Name:        []string{"--aws_enable_proxy"},
			Description: `Enable proxying of HTTP/HTTPS requests in AWS client config`,
		}, {
			Name:        []string{"--aws_firehose_endpoint"},
			Description: `Custom Firehose endpoint`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--aws_firehose_period"},
			Description: `Seconds between flushing logs to Firehose (default 10)`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--aws_firehose_stream"},
			Description: `Name of Firehose stream for logging`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--aws_kinesis_disable_log_status"},
			Description: `Disable status logs processing`,
		}, {
			Name:        []string{"--aws_kinesis_endpoint"},
			Description: `Custom Kinesis endpoint`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--aws_kinesis_period"},
			Description: `Seconds between flushing logs to Kinesis (default 10)`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--aws_kinesis_random_partition_key"},
			Description: `Enable random kinesis partition keys`,
		}, {
			Name:        []string{"--aws_kinesis_stream"},
			Description: `Name of Kinesis stream for logging`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--aws_profile_name"},
			Description: `AWS profile for authentication and region configuration`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--aws_proxy_host"},
			Description: `Proxy host for use in AWS client config`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--aws_proxy_password"},
			Description: `Proxy password for use in AWS client config`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--aws_proxy_port"},
			Description: `Proxy port for use in AWS client config`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--aws_proxy_scheme"},
			Description: `Proxy HTTP scheme for use in AWS client config (http or https, default https)`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--aws_proxy_username"},
			Description: `Proxy username for use in AWS client config`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--aws_region"},
			Description: `AWS region`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--aws_secret_access_key"},
			Description: `AWS secret access key`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--aws_session_token"},
			Description: `AWS STS session token`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--aws_sts_arn_role"},
			Description: `AWS STS ARN role`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--aws_sts_region"},
			Description: `AWS STS region`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--aws_sts_session_name"},
			Description: `AWS STS session name`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--aws_sts_timeout"},
			Description: `AWS STS assume role credential validity in seconds (default 3600)`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--buffered_log_max"},
			Description: `Maximum number of logs in buffered output plugins (0 = unlimited)`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--decorations_top_level"},
			Description: `Add decorators as top level JSON objects`,
		}, {
			Name:        []string{"--disable_audit"},
			Description: `Disable receiving events from the audit subsystem`,
		}, {
			Name:        []string{"--disable_caching"},
			Description: `Disable scheduled query caching`,
		}, {
			Name:        []string{"--disable_database"},
			Description: `Disable the persistent RocksDB storage`,
		}, {
			Name:        []string{"--disable_decorators"},
			Description: `Disable log result decoration`,
		}, {
			Name:        []string{"--disable_distributed"},
			Description: `Disable distributed queries (default true)`,
		}, {
			Name:        []string{"--disable_endpointsecurity"},
			Description: `Disable receiving events from the EndpointSecurity subsystem`,
		}, {
			Name:        []string{"--disable_endpointsecurity_fim"},
			Description: `Disable file events from the EndpointSecurity subsystem`,
		}, {
			Name:        []string{"--disable_events"},
			Description: `Disable osquery publish/subscribe system`,
		}, {
			Name:        []string{"--disable_hash_cache"},
			Description: `Cache calculated file hashes, re-calculate only if inode times change`,
		}, {
			Name:        []string{"--disable_logging"},
			Description: `Disable ERROR/INFO logging`,
		}, {
			Name:        []string{"--distributed_interval"},
			Description: `Seconds between polling for new queries (default 60)`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--distributed_loginfo"},
			Description: `Log the running distributed queries name at INFO level`,
		}, {
			Name:        []string{"--distributed_plugin"},
			Description: `Distributed plugin name`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--distributed_tls_max_attempts"},
			Description: `Number of times to attempt a request`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--distributed_tls_read_endpoint"},
			Description: `TLS/HTTPS endpoint for distributed query retrieval`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--distributed_tls_write_endpoint"},
			Description: `TLS/HTTPS endpoint for distributed query results`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--docker_socket"},
			Description: `Docker UNIX domain socket path`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "value",
			}},
		}, {
			Name:        []string{"--enable_file_events"},
			Description: `Enables the file_events publisher`,
		}, {
			Name:        []string{"--enable_foreign"},
			Description: `Enable no-op foreign virtual tables`,
		}, {
			Name:        []string{"--enable_keyboard_events"},
			Description: `Enable listening for keyboard events`,
		}, {
			Name:        []string{"--enable_mouse_events"},
			Description: `Enable listening for mouse events`,
		}, {
			Name:        []string{"--enable_numeric_monitoring"},
			Description: `Enable numeric monitoring system`,
		}, {
			Name:        []string{"--ephemeral"},
			Description: `Skip pidfile and database state checks`,
		}, {
			Name:        []string{"--es_fim_mute_path_literal"},
			Description: `Comma delimited list of path literals to be muted for FIM`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--es_fim_mute_path_prefix"},
			Description: `Comma delimited list of path prefxes to be muted for FIM`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--events_expiry"},
			Description: `Timeout to expire event subscriber results`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--events_max"},
			Description: `Maximum number of event batches per type to buffer`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--events_optimize"},
			Description: `Optimize subscriber select queries (scheduler only)`,
		}, {
			Name:        []string{"--extensions_default_index"},
			Description: `Enable INDEX on all extension table columns (default true)`,
		}, {
			Name:        []string{"--hash_cache_max"},
			Description: `Size of LRU file hash cache`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--host_identifier"},
			Description: `Field used to identify the host running osquery (hostname, uuid, instance, ephemeral, specified)`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--logger_event_type"},
			Description: `Log scheduled results as events`,
		}, {
			Name:        []string{"--logger_kafka_acks"},
			Description: `The number of acknowledgments the leader has to receive (0, 1, 'all')`,
			Args: []model.Arg{{
				Name:        "value",
				Suggestions: []model.Suggestion{{Name: []string{`0`}}, {Name: []string{`1`}}, {Name: []string{`all`}}},
			}},
		}, {
			Name:        []string{"--logger_kafka_brokers"},
			Description: `Bootstrap broker(s) as a comma-separated list of host or host:port (default port 9092)`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--logger_kafka_compression"},
			Description: `Compression codec to use for compressing message sets ('none' or 'gzip')`,
			Args: []model.Arg{{
				Name:        "value",
				Suggestions: []model.Suggestion{{Name: []string{`none`}}, {Name: []string{`gzip`}}},
			}},
		}, {
			Name:        []string{"--logger_kafka_topic"},
			Description: `Kafka topic to publish logs under`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--logger_min_status"},
			Description: `Minimum level for status log recording`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--logger_min_stderr"},
			Description: `Minimum level for statuses written to stderr`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--logger_numerics"},
			Description: `Use numeric JSON syntax for numeric values`,
		}, {
			Name:        []string{"--logger_path"},
			Description: `Directory path for ERROR/WARN/INFO and results logging`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "value",
			}},
		}, {
			Name:        []string{"--logger_rotate"},
			Description: `Use filesystem log rotation`,
		}, {
			Name:        []string{"--logger_rotate_max_files"},
			Description: `Max number of files to keep in rotation`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--logger_rotate_size"},
			Description: `Size for each filesystem log in bytes`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--logger_snapshot_event_type"},
			Description: `Log scheduled snapshot results as events`,
		}, {
			Name:        []string{"--logger_syslog_facility"},
			Description: `Syslog facility for status and results logs (0-23, default 19)`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--logger_syslog_prepend_cee"},
			Description: `Prepend @cee: tag to logged JSON messages`,
		}, {
			Name:        []string{"--logger_tls_compress"},
			Description: `GZip compress TLS/HTTPS request body`,
		}, {
			Name:        []string{"--logger_tls_endpoint"},
			Description: `TLS/HTTPS endpoint for results logging`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--logger_tls_max_lines"},
			Description: `Max number of logs to send per period`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--logger_tls_max_linesize"},
			Description: `Max size in bytes allowed per log line`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--logger_tls_period"},
			Description: `Seconds between flushing logs over TLS/HTTPS`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--nullvalue"},
			Description: `Set string for NULL values, default ''`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--numeric_monitoring_filesystem_path"},
			Description: `File to dump numeric monitoring records one per line. The format of the line is <PATH><TAB><VALUE><TAB><TIMESTAMP>`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--numeric_monitoring_plugins"},
			Description: `Comma separated numeric monitoring plugins names`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--numeric_monitoring_pre_aggregation_time"},
			Description: `Time period in seconds for numeric monitoring pre-aggregation buffer`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--pack_delimiter"},
			Description: `Delimiter for pack and query names`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--pack_refresh_interval"},
			Description: `Cache expiration for a packs discovery queries`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--read_max"},
			Description: `Maximum file read size`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--schedule_default_interval"},
			Description: `Query interval to use if none is provided`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--schedule_epoch"},
			Description: `Epoch for scheduled queries`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--schedule_lognames"},
			Description: `Log the running scheduled query name at INFO level`,
		}, {
			Name:        []string{"--schedule_max_drift"},
			Description: `Max time drift in seconds`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--schedule_reload"},
			Description: `Interval in seconds to reload database arenas`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--schedule_splay_percent"},
			Description: `Percent to splay config times`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--schedule_timeout"},
			Description: `Limit the schedule to a duration in seconds, 0 for no limit`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--specified_identifier"},
			Description: `Field used to specify the host_identifier when set to 'specified'`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--table_delay"},
			Description: `Add an optional microsecond delay between table scans`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--table_exceptions"},
			Description: `Allow tables to throw exceptions`,
		}, {
			Name:        []string{"--thrift_string_size_limit"},
			Description: `Sets the maximum string size allowed in a thrift message, use 0 for unlimited`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--thrift_timeout"},
			Description: `Timeout for thrift socket operations`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--thrift_verbose"},
			Description: `Enable the thrift log handler`,
		}, {
			Name:        []string{"--tls_disable_status_log"},
			Description: `Disable sending status logs`,
		}, {
			Name:        []string{"--verbose"},
			Description: `Enable verbose informational messages`,
		}, {
			Name:        []string{"--worker_threads"},
			Description: `Number of work dispatch threads`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--yara_delay"},
			Description: `Time in ms to sleep after scan of each file (default 50) to reduce memory spikes`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--A"},
			Description: `Select all from a table`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--L"},
			Description: `List all table names`,
		}, {
			Name:        []string{"--connect"},
			Description: `Connect to an extension socket`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--csv"},
			Description: `Set output mode to 'csv'`,
		}, {
			Name:        []string{"--extension"},
			Description: `Path to a single extension to autoload`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths, model.TemplateFolders},
				Name:      "value",
			}},
		}, {
			Name:        []string{"--header"},
			Description: `Toggle column headers true/false`,
		}, {
			Name:        []string{"--json"},
			Description: `Set output mode to 'json'`,
		}, {
			Name:        []string{"--json_pretty"},
			Description: `Set output mode to 'json_pretty'`,
		}, {
			Name:        []string{"--line"},
			Description: `Set output mode to 'line'`,
		}, {
			Name:        []string{"--list"},
			Description: `Set output mode to 'list'`,
		}, {
			Name:        []string{"--pack"},
			Description: `Run all queries in a pack`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--planner"},
			Description: `Enable osquery runtime planner output`,
		}, {
			Name:        []string{"--profile"},
			Description: `Enable profile mode when non-0, set number of iterations`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}, {
			Name:        []string{"--separator"},
			Description: `Set output field separator, default '|'`,
			Args: []model.Arg{{
				Name: "value",
			}},
		}},
	}
}
