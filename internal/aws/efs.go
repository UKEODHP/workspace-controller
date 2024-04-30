package aws

import (
	"context"

	"github.com/UKEODHP/workspace-controller/api/v1alpha1"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/efs"
	"github.com/google/uuid"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

func (c *AWSClient) ReconcileEFSAccessPoint(ctx context.Context, efsID string,
	awsEFS *v1alpha1.AWSEFSSpec) (*string, error) {

	log := log.FromContext(ctx)

	// Create a new EFS service client
	svc := efs.New(c.sess)

	// Get the access point
	describeAccessPointsParams := &efs.DescribeAccessPointsInput{
		FileSystemId: aws.String(efsID),
	}
	accessPoints, err := svc.DescribeAccessPoints(describeAccessPointsParams)
	if err != nil {
		log.Error(err, "Failed to describe EFS access points", "file system ID", efsID)
		return nil, err
	}

	// Find the access point with the desired root directory
	var accessPointID string
	for _, ap := range accessPoints.AccessPoints {
		if aws.StringValue(ap.RootDirectory.Path) == awsEFS.RootDirectory {
			accessPointID = aws.StringValue(ap.AccessPointId)
			break
		}
	}

	// If the access point is found, return its ID
	if accessPointID != "" {
		return &accessPointID, nil
	}

	// Access point not found, create a new one
	// Define the parameters for the access point
	var uid, gid int64
	if awsEFS.PosixUser.UID == 0 {
		uid = 1000
	} else {
		uid = awsEFS.PosixUser.UID
	}
	if awsEFS.PosixUser.GID == 0 {
		gid = 1000
	} else {
		gid = awsEFS.PosixUser.GID
	}
	accessPointParams := &efs.CreateAccessPointInput{
		ClientToken:  aws.String(uuid.New().String()),
		FileSystemId: aws.String(efsID),
		PosixUser: &efs.PosixUser{
			Uid: aws.Int64(uid),
			Gid: aws.Int64(gid),
		},
		RootDirectory: &efs.RootDirectory{
			Path: aws.String(awsEFS.RootDirectory),
		},
	}

	// Create the access point
	if ap, err := svc.CreateAccessPoint(accessPointParams); err == nil {
		return ap.AccessPointId, nil
	} else {
		log.Error(err, "Failed to create EFS access point", "access point", accessPointParams)
		return nil, err
	}
}

func (c *AWSClient) DeleteEFSAccessPoint(ctx context.Context,
	accessPointID string) error {

	// Create a new EFS service client
	svc := efs.New(c.sess)

	// Delete the access point
	deleteAccessPointParams := &efs.DeleteAccessPointInput{
		AccessPointId: aws.String(accessPointID),
	}
	_, err := svc.DeleteAccessPoint(deleteAccessPointParams)
	if err != nil {
		return err
	}

	return nil
}