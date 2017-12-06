// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// DatabaseSummary. An Oracle database on a DB System. For more information, see [Managing Oracle Databases]({{DOC_SERVER_URL}}/Content/Database/Concepts/overview.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type DatabaseSummary struct {

	// The OCID of the compartment.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The database name.
	DbName *string `mandatory:"true" json:"dbName,omitempty"`

	// A system-generated name for the database to ensure uniqueness within an Oracle Data Guard group (a primary database and its standby databases). The unique name cannot be changed.
	DbUniqueName *string `mandatory:"true" json:"dbUniqueName,omitempty"`

	// The OCID of the database.
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The current state of the database.
	LifecycleState DatabaseSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// The character set for the database.
	CharacterSet *string `mandatory:"false" json:"characterSet,omitempty"`

	DbBackupConfig *DbBackupConfig `mandatory:"false" json:"dbBackupConfig,omitempty"`

	// The OCID of the database home.
	DbHomeID *string `mandatory:"false" json:"dbHomeId,omitempty"`

	// Database workload type.
	DbWorkload *string `mandatory:"false" json:"dbWorkload,omitempty"`

	// Additional information about the current lifecycleState.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails,omitempty"`

	// The national character set for the database.
	NcharacterSet *string `mandatory:"false" json:"ncharacterSet,omitempty"`

	// Pluggable database name. It must begin with an alphabetic character and can contain a maximum of eight alphanumeric characters. Special characters are not permitted. Pluggable database should not be same as database name.
	PdbName *string `mandatory:"false" json:"pdbName,omitempty"`

	// The date and time the database was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated,omitempty"`
}

func (model DatabaseSummary) String() string {
	return common.PointerString(model)
}

type DatabaseSummaryLifecycleStateEnum string

const (
	DATABASE_SUMMARY_LIFECYCLE_STATE_PROVISIONING       DatabaseSummaryLifecycleStateEnum = "PROVISIONING"
	DATABASE_SUMMARY_LIFECYCLE_STATE_AVAILABLE          DatabaseSummaryLifecycleStateEnum = "AVAILABLE"
	DATABASE_SUMMARY_LIFECYCLE_STATE_UPDATING           DatabaseSummaryLifecycleStateEnum = "UPDATING"
	DATABASE_SUMMARY_LIFECYCLE_STATE_BACKUP_IN_PROGRESS DatabaseSummaryLifecycleStateEnum = "BACKUP_IN_PROGRESS"
	DATABASE_SUMMARY_LIFECYCLE_STATE_TERMINATING        DatabaseSummaryLifecycleStateEnum = "TERMINATING"
	DATABASE_SUMMARY_LIFECYCLE_STATE_TERMINATED         DatabaseSummaryLifecycleStateEnum = "TERMINATED"
	DATABASE_SUMMARY_LIFECYCLE_STATE_RESTORE_FAILED     DatabaseSummaryLifecycleStateEnum = "RESTORE_FAILED"
	DATABASE_SUMMARY_LIFECYCLE_STATE_FAILED             DatabaseSummaryLifecycleStateEnum = "FAILED"
	DATABASE_SUMMARY_LIFECYCLE_STATE_UNKNOWN            DatabaseSummaryLifecycleStateEnum = "UNKNOWN"
)

var mapping_databasesummary_lifecycleState = map[string]DatabaseSummaryLifecycleStateEnum{
	"PROVISIONING":       DATABASE_SUMMARY_LIFECYCLE_STATE_PROVISIONING,
	"AVAILABLE":          DATABASE_SUMMARY_LIFECYCLE_STATE_AVAILABLE,
	"UPDATING":           DATABASE_SUMMARY_LIFECYCLE_STATE_UPDATING,
	"BACKUP_IN_PROGRESS": DATABASE_SUMMARY_LIFECYCLE_STATE_BACKUP_IN_PROGRESS,
	"TERMINATING":        DATABASE_SUMMARY_LIFECYCLE_STATE_TERMINATING,
	"TERMINATED":         DATABASE_SUMMARY_LIFECYCLE_STATE_TERMINATED,
	"RESTORE_FAILED":     DATABASE_SUMMARY_LIFECYCLE_STATE_RESTORE_FAILED,
	"FAILED":             DATABASE_SUMMARY_LIFECYCLE_STATE_FAILED,
	"UNKNOWN":            DATABASE_SUMMARY_LIFECYCLE_STATE_UNKNOWN,
}

func GetDatabaseSummaryLifecycleStateEnumValues() []DatabaseSummaryLifecycleStateEnum {
	values := make([]DatabaseSummaryLifecycleStateEnum, 0)
	for _, v := range mapping_databasesummary_lifecycleState {
		if v != DATABASE_SUMMARY_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}