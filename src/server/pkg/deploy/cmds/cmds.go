package cmds

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/pachyderm/pachyderm/src/client/version"
	"github.com/pachyderm/pachyderm/src/server/pkg/deploy"
	"github.com/pachyderm/pachyderm/src/server/pkg/deploy/assets"
	_metrics "github.com/pachyderm/pachyderm/src/server/pkg/metrics"
	"github.com/spf13/cobra"
	"go.pedge.io/pkg/cobra"
	"go.pedge.io/pkg/exec"
)

func maybeKcCreate(dryRun bool, manifest *bytes.Buffer) error {
	if dryRun {
		_, err := os.Stdout.Write(manifest.Bytes())
		return err
	}
	return pkgexec.RunIO(
		pkgexec.IO{
			Stdin:  manifest,
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		}, "kubectl", "create", "-f", "-")
}

// DeployCmd returns a cobra command for deploying a pachyderm cluster.
func DeployCmd(noMetrics *bool) *cobra.Command {
	metrics := !*noMetrics
	var pachdShards int
	var rethinkShards int
	var hostPath string
	var dev bool
	var dryRun bool
	var deployRethinkAsRc bool
	var deployRethinkAsStatefulSet bool
	var rethinkdbCacheSize string
	var logLevel string
	var opts *assets.AssetOpts

	deployLocal := &cobra.Command{
		Use:   "local",
		Short: "Deploy a single-node Pachyderm cluster with local metadata storage.",
		Long:  "Deploy a single-node Pachyderm cluster with local metadata storage.",
		Run: pkgcobra.RunBoundedArgs(pkgcobra.Bounds{Min: 0, Max: 0}, func(args []string) (retErr error) {
			if metrics && !dev {
				metricsFn := _metrics.ReportAndFlushUserAction("Deploy")
				defer func(start time.Time) { metricsFn(start, retErr) }(time.Now())
			}
			manifest := &bytes.Buffer{}
			if dev {
				opts.Version = deploy.DevVersionTag
			}
			if err := assets.WriteLocalAssets(manifest, opts, hostPath); err != nil {
				return err
			}
			return maybeKcCreate(dryRun, manifest)
		}),
	}
	deployLocal.Flags().StringVar(&hostPath, "host-path", "/var/pachyderm", "Location on the host machine where PFS metadata will be stored.")
	deployLocal.Flags().BoolVarP(&dev, "dev", "d", false, "Don't use a specific version of pachyderm/pachd.")

	deployGoogle := &cobra.Command{
		Use:   "google <GCS bucket> <GCE persistent disks> <size of disks (in GB)>",
		Short: "Deploy a Pachyderm cluster running on GCP.",
		Long: "Deploy a Pachyderm cluster running on GCP.\n" +
			"Arguments are:\n" +
			"  <GCS bucket>: A GCS bucket where Pachyderm will store PFS data.\n" +
			"  <GCE persistent disks>: A comma-separated list of GCE persistent disks, one per rethink shard (see --rethink-shards).\n" +
			"  <size of disks>: Size of GCE persistent disks in GB (assumed to all be the same).\n",
		Run: pkgcobra.RunBoundedArgs(pkgcobra.Bounds{Min: 3, Max: 3}, func(args []string) (retErr error) {
			if metrics && !dev {
				metricsFn := _metrics.ReportAndFlushUserAction("Deploy")
				defer func(start time.Time) { metricsFn(start, retErr) }(time.Now())
			}
			volumeNames := strings.Split(args[1], ",")
			volumeSize, err := strconv.Atoi(args[2])
			if err != nil {
				return fmt.Errorf("volume size needs to be an integer; instead got %v", args[2])
			}
			manifest := &bytes.Buffer{}
			if err = assets.WriteGoogleAssets(manifest, opts, args[0], volumeNames, volumeSize); err != nil {
				return err
			}
			return maybeKcCreate(dryRun, manifest)
		}),
	}

	deployAmazon := &cobra.Command{
		Use:   "amazon <S3 bucket> <id> <secret> <token> <region> <EBS volume names> <size of volumes (in GB)>",
		Short: "Deploy a Pachyderm cluster running on AWS.",
		Long: "Deploy a Pachyderm cluster running on AWS. Arguments are:\n" +
			"  <S3 bucket>: An S3 bucket where Pachyderm will store PFS data.\n" +
			"  <id>, <secret>, <token>: Session token details, used for authorization. You can get these by running 'aws sts get-session-token'\n" +
			"  <region>: The aws region where pachyderm is being deployed (e.g. us-west-1)\n" +
			"  <EBS volume names>: A comma-separated list of EBS volumes, one per rethink shard (see --rethink-shards).\n" +
			"  <size of volumes>: Size of EBS volumes, in GB (assumed to all be the same).\n",
		Run: pkgcobra.RunBoundedArgs(pkgcobra.Bounds{Min: 7, Max: 7}, func(args []string) (retErr error) {
			if metrics && !dev {
				metricsFn := _metrics.ReportAndFlushUserAction("Deploy")
				defer func(start time.Time) { metricsFn(start, retErr) }(time.Now())
			}
			volumeNames := strings.Split(args[5], ",")
			volumeSize, err := strconv.Atoi(args[6])
			if err != nil {
				return fmt.Errorf("volume size needs to be an integer; instead got %v", args[6])
			}
			manifest := &bytes.Buffer{}
			if err = assets.WriteAmazonAssets(manifest, opts, args[0], args[1], args[2], args[3], args[4], volumeNames, volumeSize); err != nil {
				return err
			}
			return maybeKcCreate(dryRun, manifest)
		}),
	}

	deployMicrosoft := &cobra.Command{
		Use:   "microsoft <container> <storage account name> <storage account key> <volume URIs> <size of volumes (in GB)>",
		Short: "Deploy a Pachyderm cluster running on Microsoft Azure.",
		Long: "Deploy a Pachyderm cluster running on Microsoft Azure. Arguments are:\n" +
			"  <container>: An Azure container where Pachyderm will store PFS data.\n" +
			"  <volume URIs>: A comma-separated list of persistent volumes, one per rethink shard (see --rethink-shards).\n" +
			"  <size of volumes>: Size of persistent volumes, in GB (assumed to all be the same).\n",
		Run: pkgcobra.RunBoundedArgs(pkgcobra.Bounds{Min: 5, Max: 5}, func(args []string) (retErr error) {
			if metrics && !dev {
				metricsFn := _metrics.ReportAndFlushUserAction("Deploy")
				defer func(start time.Time) { metricsFn(start, retErr) }(time.Now())
			}
			if _, err := base64.StdEncoding.DecodeString(args[2]); err != nil {
				return fmt.Errorf("storage-account-key needs to be base64 encoded; instead got '%v'", args[2])
			}
			volumeURIs := strings.Split(args[3], ",")
			for i, uri := range volumeURIs {
				tempURI, err := url.ParseRequestURI(uri)
				if err != nil {
					return fmt.Errorf("All volume-uris needs to be a well-formed URI; instead got '%v'", uri)
				}
				volumeURIs[i] = tempURI.String()
			}
			volumeSize, err := strconv.Atoi(args[4])
			if err != nil {
				return fmt.Errorf("volume size needs to be an integer; instead got %v", args[4])
			}
			manifest := &bytes.Buffer{}
			if err = assets.WriteMicrosoftAssets(manifest, opts, args[0], args[1], args[2], volumeURIs, volumeSize); err != nil {
				return err
			}
			return maybeKcCreate(dryRun, manifest)
		}),
	}

	cmd := &cobra.Command{
		Use:   "deploy amazon|google|microsoft|basic",
		Short: "Deploy a Pachyderm cluster.",
		Long:  "Deploy a Pachyderm cluster.",
		PersistentPreRun: pkgcobra.Run(func([]string) error {
			if deployRethinkAsRc {
				if deployRethinkAsStatefulSet {
					return fmt.Errorf("Error: pachctl deploy received contradictory flags: " +
						"--deploy-rethink-as-rc and --deploy-rethink-as-stateful set")
				}
				fmt.Fprintf(os.Stderr, "Warning: --deploy-rethink-as-rc is no longer "+
					"necessary (and is ignored). The default behavior since Pachyderm "+
					"1.3.2 is to manage RethinkDB with a Kubernetes Replication Controller. "+
					"This flag will be removed by Pachyderm's 1.4 release, so please remove "+
					"it from your scripts. Also see --deploy-rethink-as-stateful-set.\n")
			}
			opts = &assets.AssetOpts{
				PachdShards:                uint64(pachdShards),
				RethinkShards:              uint64(rethinkShards),
				RethinkdbCacheSize:         rethinkdbCacheSize,
				DeployRethinkAsStatefulSet: deployRethinkAsStatefulSet,
				Version:                    version.PrettyPrintVersion(version.Version),
				LogLevel:                   logLevel,
				Metrics:                    metrics,
			}
			return nil
		}),
	}
	cmd.PersistentFlags().IntVar(&pachdShards, "shards", 1, "Number of Pachd nodes (stateless Pachyderm API servers).")
	cmd.PersistentFlags().IntVar(&rethinkShards, "rethink-shards", 1, "Number of RethinkDB shards (for pfs metadata storage) if "+
		"--deploy-rethink-as-stateful-set is used.")
	cmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "Don't actually deploy pachyderm to Kubernetes, instead just print the manifest.")
	cmd.PersistentFlags().StringVar(&rethinkdbCacheSize, "rethinkdb-cache-size", "768M", "Size of in-memory cache to use for Pachyderm's RethinkDB instance, "+
		"e.g. \"2G\". Size is specified in bytes, with allowed SI suffixes (M, K, G, Mi, Ki, Gi, etc).")
	cmd.PersistentFlags().StringVar(&logLevel, "log-level", "info", "The level of log messages to print options are, from least to most verbose: \"error\", \"info\", \"debug\".")
	cmd.PersistentFlags().BoolVar(&deployRethinkAsRc, "deploy-rethink-as-rc", false, "Defunct flag (does nothing). The default behavior since "+
		"Pachyderm 1.3.2 is to manage RethinkDB with a Kubernetes Replication Controller.")
	cmd.PersistentFlags().BoolVar(&deployRethinkAsStatefulSet, "deploy-rethink-as-stateful-set", false, "Deploy RethinkDB as a multi-node cluster "+
		"controlled by kubernetes StatefulSet, instead of a single-node instance controlled by a Kubernetes Replication Controller. Note that both "+
		"your local kubectl binary and the kubernetes server must be at least version 1.5.")
	cmd.AddCommand(deployLocal)
	cmd.AddCommand(deployAmazon)
	cmd.AddCommand(deployGoogle)
	cmd.AddCommand(deployMicrosoft)
	return cmd
}
