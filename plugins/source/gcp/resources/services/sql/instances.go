// Code generated by codegen; DO NOT EDIT.

package sql

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Instances() *schema.Table {
	return &schema.Table{
		Name:      "gcp_sql_instances",
		Resolver:  fetchInstances,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "self_link",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SelfLink"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "available_maintenance_versions",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("AvailableMaintenanceVersions"),
			},
			{
				Name:     "backend_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BackendType"),
			},
			{
				Name:     "connection_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ConnectionName"),
			},
			{
				Name:     "create_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreateTime"),
			},
			{
				Name:     "current_disk_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("CurrentDiskSize"),
			},
			{
				Name:     "database_installed_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DatabaseInstalledVersion"),
			},
			{
				Name:     "database_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DatabaseVersion"),
			},
			{
				Name:     "disk_encryption_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DiskEncryptionConfiguration"),
			},
			{
				Name:     "disk_encryption_status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DiskEncryptionStatus"),
			},
			{
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Etag"),
			},
			{
				Name:     "failover_replica",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("FailoverReplica"),
			},
			{
				Name:     "gce_zone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GceZone"),
			},
			{
				Name:     "instance_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("InstanceType"),
			},
			{
				Name:     "ip_addresses",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IpAddresses"),
			},
			{
				Name:     "ipv_6_address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Ipv6Address"),
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "maintenance_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MaintenanceVersion"),
			},
			{
				Name:     "master_instance_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MasterInstanceName"),
			},
			{
				Name:     "max_disk_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxDiskSize"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "on_premises_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("OnPremisesConfiguration"),
			},
			{
				Name:     "out_of_disk_report",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("OutOfDiskReport"),
			},
			{
				Name:     "project",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Project"),
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Region"),
			},
			{
				Name:     "replica_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ReplicaConfiguration"),
			},
			{
				Name:     "replica_names",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ReplicaNames"),
			},
			{
				Name:     "root_password",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RootPassword"),
			},
			{
				Name:     "satisfies_pzs",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("SatisfiesPzs"),
			},
			{
				Name:     "scheduled_maintenance",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ScheduledMaintenance"),
			},
			{
				Name:     "secondary_gce_zone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SecondaryGceZone"),
			},
			{
				Name:     "server_ca_cert",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ServerCaCert"),
			},
			{
				Name:     "service_account_email_address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceAccountEmailAddress"),
			},
			{
				Name:     "settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Settings"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "suspension_reason",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SuspensionReason"),
			},
			{
				Name:     "server_response",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ServerResponse"),
			},
			{
				Name:     "force_send_fields",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ForceSendFields"),
			},
			{
				Name:     "null_fields",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("NullFields"),
			},
		},
	}
}
