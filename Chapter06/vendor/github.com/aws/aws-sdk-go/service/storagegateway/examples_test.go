// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/storagegateway"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleStorageGateway_ActivateGateway() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.ActivateGatewayInput{
		ActivationKey:     aws.String("ActivationKey"),   // Required
		GatewayName:       aws.String("GatewayName"),     // Required
		GatewayRegion:     aws.String("RegionId"),        // Required
		GatewayTimezone:   aws.String("GatewayTimezone"), // Required
		GatewayType:       aws.String("GatewayType"),
		MediumChangerType: aws.String("MediumChangerType"),
		TapeDriveType:     aws.String("TapeDriveType"),
	}
	resp, err := svc.ActivateGateway(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_AddCache() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.AddCacheInput{
		DiskIds: []*string{ // Required
			aws.String("DiskId"), // Required
			// More values...
		},
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.AddCache(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_AddTagsToResource() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.AddTagsToResourceInput{
		ResourceARN: aws.String("ResourceARN"), // Required
		Tags: []*storagegateway.Tag{ // Required
			{ // Required
				Key:   aws.String("TagKey"),   // Required
				Value: aws.String("TagValue"), // Required
			},
			// More values...
		},
	}
	resp, err := svc.AddTagsToResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_AddUploadBuffer() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.AddUploadBufferInput{
		DiskIds: []*string{ // Required
			aws.String("DiskId"), // Required
			// More values...
		},
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.AddUploadBuffer(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_AddWorkingStorage() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.AddWorkingStorageInput{
		DiskIds: []*string{ // Required
			aws.String("DiskId"), // Required
			// More values...
		},
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.AddWorkingStorage(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_CancelArchival() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.CancelArchivalInput{
		GatewayARN: aws.String("GatewayARN"), // Required
		TapeARN:    aws.String("TapeARN"),    // Required
	}
	resp, err := svc.CancelArchival(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_CancelRetrieval() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.CancelRetrievalInput{
		GatewayARN: aws.String("GatewayARN"), // Required
		TapeARN:    aws.String("TapeARN"),    // Required
	}
	resp, err := svc.CancelRetrieval(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_CreateCachediSCSIVolume() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.CreateCachediSCSIVolumeInput{
		ClientToken:        aws.String("ClientToken"),        // Required
		GatewayARN:         aws.String("GatewayARN"),         // Required
		NetworkInterfaceId: aws.String("NetworkInterfaceId"), // Required
		TargetName:         aws.String("TargetName"),         // Required
		VolumeSizeInBytes:  aws.Int64(1),                     // Required
		SnapshotId:         aws.String("SnapshotId"),
		SourceVolumeARN:    aws.String("VolumeARN"),
	}
	resp, err := svc.CreateCachediSCSIVolume(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_CreateNFSFileShare() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.CreateNFSFileShareInput{
		ClientToken: aws.String("ClientToken"), // Required
		GatewayARN:  aws.String("GatewayARN"),  // Required
		LocationARN: aws.String("LocationARN"), // Required
		Role:        aws.String("Role"),        // Required
		ClientList: []*string{
			aws.String("IPV4AddressCIDR"), // Required
			// More values...
		},
		DefaultStorageClass: aws.String("StorageClass"),
		KMSEncrypted:        aws.Bool(true),
		KMSKey:              aws.String("KMSKey"),
		NFSFileShareDefaults: &storagegateway.NFSFileShareDefaults{
			DirectoryMode: aws.String("PermissionMode"),
			FileMode:      aws.String("PermissionMode"),
			GroupId:       aws.Int64(1),
			OwnerId:       aws.Int64(1),
		},
		ReadOnly: aws.Bool(true),
		Squash:   aws.String("Squash"),
	}
	resp, err := svc.CreateNFSFileShare(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_CreateSnapshot() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.CreateSnapshotInput{
		SnapshotDescription: aws.String("SnapshotDescription"), // Required
		VolumeARN:           aws.String("VolumeARN"),           // Required
	}
	resp, err := svc.CreateSnapshot(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_CreateSnapshotFromVolumeRecoveryPoint() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.CreateSnapshotFromVolumeRecoveryPointInput{
		SnapshotDescription: aws.String("SnapshotDescription"), // Required
		VolumeARN:           aws.String("VolumeARN"),           // Required
	}
	resp, err := svc.CreateSnapshotFromVolumeRecoveryPoint(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_CreateStorediSCSIVolume() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.CreateStorediSCSIVolumeInput{
		DiskId:               aws.String("DiskId"),             // Required
		GatewayARN:           aws.String("GatewayARN"),         // Required
		NetworkInterfaceId:   aws.String("NetworkInterfaceId"), // Required
		PreserveExistingData: aws.Bool(true),                   // Required
		TargetName:           aws.String("TargetName"),         // Required
		SnapshotId:           aws.String("SnapshotId"),
	}
	resp, err := svc.CreateStorediSCSIVolume(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_CreateTapeWithBarcode() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.CreateTapeWithBarcodeInput{
		GatewayARN:      aws.String("GatewayARN"),  // Required
		TapeBarcode:     aws.String("TapeBarcode"), // Required
		TapeSizeInBytes: aws.Int64(1),              // Required
	}
	resp, err := svc.CreateTapeWithBarcode(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_CreateTapes() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.CreateTapesInput{
		ClientToken:       aws.String("ClientToken"),       // Required
		GatewayARN:        aws.String("GatewayARN"),        // Required
		NumTapesToCreate:  aws.Int64(1),                    // Required
		TapeBarcodePrefix: aws.String("TapeBarcodePrefix"), // Required
		TapeSizeInBytes:   aws.Int64(1),                    // Required
	}
	resp, err := svc.CreateTapes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DeleteBandwidthRateLimit() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.DeleteBandwidthRateLimitInput{
		BandwidthType: aws.String("BandwidthType"), // Required
		GatewayARN:    aws.String("GatewayARN"),    // Required
	}
	resp, err := svc.DeleteBandwidthRateLimit(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DeleteChapCredentials() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.DeleteChapCredentialsInput{
		InitiatorName: aws.String("IqnName"),   // Required
		TargetARN:     aws.String("TargetARN"), // Required
	}
	resp, err := svc.DeleteChapCredentials(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DeleteFileShare() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.DeleteFileShareInput{
		FileShareARN: aws.String("FileShareARN"), // Required
	}
	resp, err := svc.DeleteFileShare(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DeleteGateway() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.DeleteGatewayInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.DeleteGateway(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DeleteSnapshotSchedule() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.DeleteSnapshotScheduleInput{
		VolumeARN: aws.String("VolumeARN"), // Required
	}
	resp, err := svc.DeleteSnapshotSchedule(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DeleteTape() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.DeleteTapeInput{
		GatewayARN: aws.String("GatewayARN"), // Required
		TapeARN:    aws.String("TapeARN"),    // Required
	}
	resp, err := svc.DeleteTape(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DeleteTapeArchive() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.DeleteTapeArchiveInput{
		TapeARN: aws.String("TapeARN"), // Required
	}
	resp, err := svc.DeleteTapeArchive(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DeleteVolume() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.DeleteVolumeInput{
		VolumeARN: aws.String("VolumeARN"), // Required
	}
	resp, err := svc.DeleteVolume(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DescribeBandwidthRateLimit() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.DescribeBandwidthRateLimitInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.DescribeBandwidthRateLimit(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DescribeCache() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.DescribeCacheInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.DescribeCache(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DescribeCachediSCSIVolumes() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.DescribeCachediSCSIVolumesInput{
		VolumeARNs: []*string{ // Required
			aws.String("VolumeARN"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeCachediSCSIVolumes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DescribeChapCredentials() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.DescribeChapCredentialsInput{
		TargetARN: aws.String("TargetARN"), // Required
	}
	resp, err := svc.DescribeChapCredentials(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DescribeGatewayInformation() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.DescribeGatewayInformationInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.DescribeGatewayInformation(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DescribeMaintenanceStartTime() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.DescribeMaintenanceStartTimeInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.DescribeMaintenanceStartTime(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DescribeNFSFileShares() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.DescribeNFSFileSharesInput{
		FileShareARNList: []*string{ // Required
			aws.String("FileShareARN"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeNFSFileShares(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DescribeSnapshotSchedule() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.DescribeSnapshotScheduleInput{
		VolumeARN: aws.String("VolumeARN"), // Required
	}
	resp, err := svc.DescribeSnapshotSchedule(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DescribeStorediSCSIVolumes() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.DescribeStorediSCSIVolumesInput{
		VolumeARNs: []*string{ // Required
			aws.String("VolumeARN"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeStorediSCSIVolumes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DescribeTapeArchives() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.DescribeTapeArchivesInput{
		Limit:  aws.Int64(1),
		Marker: aws.String("Marker"),
		TapeARNs: []*string{
			aws.String("TapeARN"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeTapeArchives(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DescribeTapeRecoveryPoints() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.DescribeTapeRecoveryPointsInput{
		GatewayARN: aws.String("GatewayARN"), // Required
		Limit:      aws.Int64(1),
		Marker:     aws.String("Marker"),
	}
	resp, err := svc.DescribeTapeRecoveryPoints(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DescribeTapes() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.DescribeTapesInput{
		GatewayARN: aws.String("GatewayARN"), // Required
		Limit:      aws.Int64(1),
		Marker:     aws.String("Marker"),
		TapeARNs: []*string{
			aws.String("TapeARN"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeTapes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DescribeUploadBuffer() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.DescribeUploadBufferInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.DescribeUploadBuffer(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DescribeVTLDevices() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.DescribeVTLDevicesInput{
		GatewayARN: aws.String("GatewayARN"), // Required
		Limit:      aws.Int64(1),
		Marker:     aws.String("Marker"),
		VTLDeviceARNs: []*string{
			aws.String("VTLDeviceARN"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeVTLDevices(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DescribeWorkingStorage() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.DescribeWorkingStorageInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.DescribeWorkingStorage(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DisableGateway() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.DisableGatewayInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.DisableGateway(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_ListFileShares() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.ListFileSharesInput{
		GatewayARN: aws.String("GatewayARN"),
		Limit:      aws.Int64(1),
		Marker:     aws.String("Marker"),
	}
	resp, err := svc.ListFileShares(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_ListGateways() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.ListGatewaysInput{
		Limit:  aws.Int64(1),
		Marker: aws.String("Marker"),
	}
	resp, err := svc.ListGateways(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_ListLocalDisks() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.ListLocalDisksInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.ListLocalDisks(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_ListTagsForResource() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.ListTagsForResourceInput{
		ResourceARN: aws.String("ResourceARN"), // Required
		Limit:       aws.Int64(1),
		Marker:      aws.String("Marker"),
	}
	resp, err := svc.ListTagsForResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_ListTapes() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.ListTapesInput{
		Limit:  aws.Int64(1),
		Marker: aws.String("Marker"),
		TapeARNs: []*string{
			aws.String("TapeARN"), // Required
			// More values...
		},
	}
	resp, err := svc.ListTapes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_ListVolumeInitiators() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.ListVolumeInitiatorsInput{
		VolumeARN: aws.String("VolumeARN"), // Required
	}
	resp, err := svc.ListVolumeInitiators(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_ListVolumeRecoveryPoints() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.ListVolumeRecoveryPointsInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.ListVolumeRecoveryPoints(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_ListVolumes() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.ListVolumesInput{
		GatewayARN: aws.String("GatewayARN"),
		Limit:      aws.Int64(1),
		Marker:     aws.String("Marker"),
	}
	resp, err := svc.ListVolumes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_RefreshCache() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.RefreshCacheInput{
		FileShareARN: aws.String("FileShareARN"), // Required
	}
	resp, err := svc.RefreshCache(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_RemoveTagsFromResource() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.RemoveTagsFromResourceInput{
		ResourceARN: aws.String("ResourceARN"), // Required
		TagKeys: []*string{ // Required
			aws.String("TagKey"), // Required
			// More values...
		},
	}
	resp, err := svc.RemoveTagsFromResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_ResetCache() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.ResetCacheInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.ResetCache(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_RetrieveTapeArchive() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.RetrieveTapeArchiveInput{
		GatewayARN: aws.String("GatewayARN"), // Required
		TapeARN:    aws.String("TapeARN"),    // Required
	}
	resp, err := svc.RetrieveTapeArchive(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_RetrieveTapeRecoveryPoint() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.RetrieveTapeRecoveryPointInput{
		GatewayARN: aws.String("GatewayARN"), // Required
		TapeARN:    aws.String("TapeARN"),    // Required
	}
	resp, err := svc.RetrieveTapeRecoveryPoint(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_SetLocalConsolePassword() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.SetLocalConsolePasswordInput{
		GatewayARN:           aws.String("GatewayARN"),           // Required
		LocalConsolePassword: aws.String("LocalConsolePassword"), // Required
	}
	resp, err := svc.SetLocalConsolePassword(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_ShutdownGateway() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.ShutdownGatewayInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.ShutdownGateway(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_StartGateway() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.StartGatewayInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.StartGateway(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_UpdateBandwidthRateLimit() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.UpdateBandwidthRateLimitInput{
		GatewayARN:                           aws.String("GatewayARN"), // Required
		AverageDownloadRateLimitInBitsPerSec: aws.Int64(1),
		AverageUploadRateLimitInBitsPerSec:   aws.Int64(1),
	}
	resp, err := svc.UpdateBandwidthRateLimit(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_UpdateChapCredentials() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.UpdateChapCredentialsInput{
		InitiatorName:                 aws.String("IqnName"),    // Required
		SecretToAuthenticateInitiator: aws.String("ChapSecret"), // Required
		TargetARN:                     aws.String("TargetARN"),  // Required
		SecretToAuthenticateTarget:    aws.String("ChapSecret"),
	}
	resp, err := svc.UpdateChapCredentials(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_UpdateGatewayInformation() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.UpdateGatewayInformationInput{
		GatewayARN:      aws.String("GatewayARN"), // Required
		GatewayName:     aws.String("GatewayName"),
		GatewayTimezone: aws.String("GatewayTimezone"),
	}
	resp, err := svc.UpdateGatewayInformation(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_UpdateGatewaySoftwareNow() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.UpdateGatewaySoftwareNowInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.UpdateGatewaySoftwareNow(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_UpdateMaintenanceStartTime() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.UpdateMaintenanceStartTimeInput{
		DayOfWeek:    aws.Int64(1),             // Required
		GatewayARN:   aws.String("GatewayARN"), // Required
		HourOfDay:    aws.Int64(1),             // Required
		MinuteOfHour: aws.Int64(1),             // Required
	}
	resp, err := svc.UpdateMaintenanceStartTime(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_UpdateNFSFileShare() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.UpdateNFSFileShareInput{
		FileShareARN: aws.String("FileShareARN"), // Required
		ClientList: []*string{
			aws.String("IPV4AddressCIDR"), // Required
			// More values...
		},
		DefaultStorageClass: aws.String("StorageClass"),
		KMSEncrypted:        aws.Bool(true),
		KMSKey:              aws.String("KMSKey"),
		NFSFileShareDefaults: &storagegateway.NFSFileShareDefaults{
			DirectoryMode: aws.String("PermissionMode"),
			FileMode:      aws.String("PermissionMode"),
			GroupId:       aws.Int64(1),
			OwnerId:       aws.Int64(1),
		},
		ReadOnly: aws.Bool(true),
		Squash:   aws.String("Squash"),
	}
	resp, err := svc.UpdateNFSFileShare(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_UpdateSnapshotSchedule() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.UpdateSnapshotScheduleInput{
		RecurrenceInHours: aws.Int64(1),            // Required
		StartAt:           aws.Int64(1),            // Required
		VolumeARN:         aws.String("VolumeARN"), // Required
		Description:       aws.String("Description"),
	}
	resp, err := svc.UpdateSnapshotSchedule(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_UpdateVTLDeviceType() {
	sess := session.Must(session.NewSession())

	svc := storagegateway.New(sess)

	params := &storagegateway.UpdateVTLDeviceTypeInput{
		DeviceType:   aws.String("DeviceType"),   // Required
		VTLDeviceARN: aws.String("VTLDeviceARN"), // Required
	}
	resp, err := svc.UpdateVTLDeviceType(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
