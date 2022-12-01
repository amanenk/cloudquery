// Code generated by codegen; DO NOT EDIT.

package compute

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"

	pb "google.golang.org/genproto/googleapis/cloud/compute/v1"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Disks() *schema.Table {
	return &schema.Table{
		Name:      "gcp_compute_disks",
		Resolver:  fetchDisks,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "architecture",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Architecture"),
			},
			{
				Name:     "creation_timestamp",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreationTimestamp"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "disk_encryption_key",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DiskEncryptionKey"),
			},
			{
				Name:     "guest_os_features",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("GuestOsFeatures"),
			},
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "label_fingerprint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LabelFingerprint"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "last_attach_timestamp",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastAttachTimestamp"),
			},
			{
				Name:     "last_detach_timestamp",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastDetachTimestamp"),
			},
			{
				Name:     "license_codes",
				Type:     schema.TypeIntArray,
				Resolver: schema.PathResolver("LicenseCodes"),
			},
			{
				Name:     "licenses",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Licenses"),
			},
			{
				Name:     "location_hint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LocationHint"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "options",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Options"),
			},
			{
				Name:     "params",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Params"),
			},
			{
				Name:     "physical_block_size_bytes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PhysicalBlockSizeBytes"),
			},
			{
				Name:     "provisioned_iops",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ProvisionedIops"),
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Region"),
			},
			{
				Name:     "replica_zones",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ReplicaZones"),
			},
			{
				Name:     "resource_policies",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ResourcePolicies"),
			},
			{
				Name:     "satisfies_pzs",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("SatisfiesPzs"),
			},
			{
				Name:     "size_gb",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("SizeGb"),
			},
			{
				Name:     "source_disk",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceDisk"),
			},
			{
				Name:     "source_disk_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceDiskId"),
			},
			{
				Name:     "source_image",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceImage"),
			},
			{
				Name:     "source_image_encryption_key",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SourceImageEncryptionKey"),
			},
			{
				Name:     "source_image_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceImageId"),
			},
			{
				Name:     "source_snapshot",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceSnapshot"),
			},
			{
				Name:     "source_snapshot_encryption_key",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SourceSnapshotEncryptionKey"),
			},
			{
				Name:     "source_snapshot_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceSnapshotId"),
			},
			{
				Name:     "source_storage_object",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceStorageObject"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "users",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Users"),
			},
			{
				Name:     "zone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Zone"),
			},
		},
	}
}

func fetchDisks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &pb.AggregatedListDisksRequest{
		Project: c.ProjectId,
	}
	it := c.Services.ComputeDisksClient.AggregatedList(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return errors.WithStack(err)
		}

		res <- resp.Value.Disks

	}
	return nil
}
