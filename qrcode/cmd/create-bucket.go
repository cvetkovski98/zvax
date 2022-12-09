package cmd

import (
	"github.com/cvetkovski98/zvax/zvax-qrcode/internal/config"
	"github.com/cvetkovski98/zvax/zvax-qrcode/pkg/minio"
	"github.com/spf13/cobra"
)

var createBucketCommand = &cobra.Command{
	Use:   "create-bucket",
	Short: "Create MinIO bucket",
	Long:  `Create MinIO bucket`,
	RunE:  createBucket,
}

func init() {
	createBucketCommand.Flags().StringP("bucket", "b", "", "Bucket name")
}

func createBucket(cmd *cobra.Command, args []string) error {
	cfg := config.GetConfig()
	minioClient, err := minio.NewMinioClient(&cfg.Minio)
	if err != nil {
		return err
	}
	bucket := cfg.Minio.BucketName
	if cmd.Flag("bucket").Value.String() != "" {
		bucket = cmd.Flag("bucket").Value.String()
	}
	if err = minio.CreateBucket(cmd.Context(), minioClient, bucket); err != nil {
		return err
	}
	if err = minio.MakeReadOnly(cmd.Context(), minioClient, bucket); err != nil {
		return err
	}
	return nil
}
