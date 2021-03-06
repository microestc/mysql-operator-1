package backup

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"github.com/grtl/mysql-operator/pkg/cmd/util/config"
	"github.com/grtl/mysql-operator/pkg/cmd/util/fail"
	"github.com/grtl/mysql-operator/pkg/cmd/util/options"
)

var (
	clusterName string
	backupName  string
	storage     string
	time        string
)

var backupCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new MySQL Backup Schedule",
	Long: `Creates a new MySQL Backup Schedule for specified cluster. The backups will
be created according to cron-style time provided.
msp backup create --name "my-backup" --cluster "my-cluster"`,
	Args: cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		options := options.ExtractOptions(cmd)

		if backupName == "" {
			backupName = clusterName
		}

		err := createMySQLBackup(options)
		if err != nil {
			fail.Error(err)
		}
	},
}

func init() {
	Cmd.AddCommand(backupCreateCmd)

	backupCreateCmd.Flags().StringVarP(&clusterName, "cluster", "c", "",
		"name of the cluster to be backed up")
	backupCreateCmd.MarkFlagRequired("cluster")
	backupCreateCmd.Flags().StringVarP(&time, "time", "t", "", "cron-style time")
	backupCreateCmd.MarkFlagRequired("time")
	backupCreateCmd.Flags().StringVarP(&storage, "storage", "s", "1Gi", "storage value")
	backupCreateCmd.Flags().StringVar(&backupName, "name", "", "backupschedule name")
}

func createMySQLBackup(options *options.Options) error {
	fmt.Printf("Creating: %s/%s for %s\n", options.Namespace, backupName, clusterName)

	var (
		storageQuantity resource.Quantity
		err             error
	)

	if storage != "" {
		storageQuantity, err = resource.ParseQuantity(storage)
		if err != nil {
			return err
		}
	}

	mySQLBackupInterface := config.GetConfig().Clientset().CrV1().MySQLBackupSchedules(options.Namespace)
	_, err = mySQLBackupInterface.Create(&crv1.MySQLBackupSchedule{
		ObjectMeta: metav1.ObjectMeta{
			Name: backupName,
		},
		Spec: crv1.MySQLBackupScheduleSpec{
			Cluster: clusterName,
			Time:    time,
			Storage: storageQuantity,
		},
	})

	return err
}
