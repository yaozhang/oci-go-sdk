// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

type DatabaseClient struct {
	common.BaseClient
}

//Create a new default Database client for a given region
func NewDatabaseClientForRegion(region common.Region) (client DatabaseClient) {
	client = DatabaseClient{BaseClient: common.NewClientForRegion(region)}

	client.Host = fmt.Sprintf(common.DefaultHostUrlTemplate, "database", string(region))
	client.BasePath = "20160918"
	return
}

// Creates a new backup in the specified database based on the request parameters you provide. If you previously used RMAN or dbcli to configure backups and then you switch to using the Console or the API for backups, a new backup configuration is created and associated with your database. This means that you can no longer rely on your previously configured unmanaged backups to work.
func (client DatabaseClient) CreateBackup(ctx context.Context, request CreateBackupRequest) (response CreateBackupResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/backups", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Creates a new Data Guard association.  A Data Guard association represents the replication relationship between the
// specified database and a peer database. For more information, see [Using Oracle Data Guard]({{DOC_SERVER_URL}}/Content/Database/Tasks/usingdataguard.htm).
// All Oracle Cloud Infrastructure resources, including Data Guard associations, get an Oracle-assigned, unique ID
// called an Oracle Cloud Identifier (OCID). When you create a resource, you can find its OCID in the response.
// You can also retrieve a resource's OCID by using a List API operation on that resource type, or by viewing the
// resource in the Console. Fore more information, see
// [Resource Identifiers](http://localhost:8000/Content/General/Concepts/identifiers.htm).
func (client DatabaseClient) CreateDataGuardAssociation(ctx context.Context, request CreateDataGuardAssociationRequest) (response CreateDataGuardAssociationResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/databases/{databaseId}/dataGuardAssociations", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Creates a new DB Home in the specified DB System based on the request parameters you provide.
func (client DatabaseClient) CreateDbHome(ctx context.Context, request CreateDbHomeRequest) (response CreateDbHomeResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/dbHomes", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Performs an action, such as one of the power actions (start, stop, softreset, or reset), on the specified DB Node.
// **start** - power on
// **stop** - power off
// **softreset** - ACPI shutdown and power on
// **reset** - power off and power on
// Note that the **stop** state has no effect on the resources you consume.
// Billing continues for DB Nodes that you stop, and related resources continue
// to apply against any relevant quotas. You must terminate the DB System
// (TerminateDbSystem)
// to remove its resources from billing and quotas.
func (client DatabaseClient) DbNodeAction(ctx context.Context, request DbNodeActionRequest) (response DbNodeActionResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/dbNodes/{dbNodeId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Deletes a full backup. You cannot delete automatic backups using this API.
func (client DatabaseClient) DeleteBackup(ctx context.Context, request DeleteBackupRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/backups/{backupId}", request)
	if err != nil {
		return
	}

	_, err = client.Call(ctx, &httpRequest)
	return
}

// Deletes a DB Home. The DB Home and its database data are local to the DB System and will be lost when it is deleted. Oracle recommends that you back up any data in the DB System prior to deleting it.
func (client DatabaseClient) DeleteDbHome(ctx context.Context, request DeleteDbHomeRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/dbHomes/{dbHomeId}", request)
	if err != nil {
		return
	}

	_, err = client.Call(ctx, &httpRequest)
	return
}

// Performs a failover to transition the standby database identified by the `databaseId` parameter into the
// specified Data Guard association's primary role after the existing primary database fails or becomes unreachable.
// A failover might result in data loss depending on the protection mode in effect at the time of the primary
// database failure.
func (client DatabaseClient) FailoverDataGuardAssociation(ctx context.Context, request FailoverDataGuardAssociationRequest) (response FailoverDataGuardAssociationResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/databases/{databaseId}/dataGuardAssociations/{dataGuardAssociationId}/actions/failover", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Gets information about the specified backup.
func (client DatabaseClient) GetBackup(ctx context.Context, request GetBackupRequest) (response GetBackupResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/backups/{backupId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Gets the specified Data Guard association's configuration information.
func (client DatabaseClient) GetDataGuardAssociation(ctx context.Context, request GetDataGuardAssociationRequest) (response GetDataGuardAssociationResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/databases/{databaseId}/dataGuardAssociations/{dataGuardAssociationId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Gets information about a specific database.
func (client DatabaseClient) GetDatabase(ctx context.Context, request GetDatabaseRequest) (response GetDatabaseResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/databases/{databaseId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Gets information about the specified database home.
func (client DatabaseClient) GetDbHome(ctx context.Context, request GetDbHomeRequest) (response GetDbHomeResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/dbHomes/{dbHomeId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Gets information about a specified patch package.
func (client DatabaseClient) GetDbHomePatch(ctx context.Context, request GetDbHomePatchRequest) (response GetDbHomePatchResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/dbHomes/{dbHomeId}/patches/{patchId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Gets the patch history details for the specified patchHistoryEntryId
func (client DatabaseClient) GetDbHomePatchHistoryEntry(ctx context.Context, request GetDbHomePatchHistoryEntryRequest) (response GetDbHomePatchHistoryEntryResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/dbHomes/{dbHomeId}/patchHistoryEntries/{patchHistoryEntryId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Gets information about the specified database node.
func (client DatabaseClient) GetDbNode(ctx context.Context, request GetDbNodeRequest) (response GetDbNodeResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/dbNodes/{dbNodeId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Gets information about the specified DB System.
func (client DatabaseClient) GetDbSystem(ctx context.Context, request GetDbSystemRequest) (response GetDbSystemResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/dbSystems/{dbSystemId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Gets information about a specified patch package.
func (client DatabaseClient) GetDbSystemPatch(ctx context.Context, request GetDbSystemPatchRequest) (response GetDbSystemPatchResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/dbSystems/{dbSystemId}/patches/{patchId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Gets the patch history details for the specified patchHistoryEntryId.
func (client DatabaseClient) GetDbSystemPatchHistoryEntry(ctx context.Context, request GetDbSystemPatchHistoryEntryRequest) (response GetDbSystemPatchHistoryEntryResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/dbSystems/{dbSystemId}/patchHistoryEntries/{patchHistoryEntryId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Launches a new DB System in the specified compartment and Availability Domain. You'll specify a single Oracle
// Database Edition that applies to all the databases on that DB System. The selected edition cannot be changed.
// An initial database is created on the DB System based on the request parameters you provide and some default
// options. For more information,
// see [Default Options for the Initial Database]({{DOC_SERVER_URL}}/Content/Database/Tasks/launchingDB.htm#Default_Options_for_the_Initial_Database).
// The DB System will include a command line interface (CLI) that you can use to create additional databases and
// manage existing databases. For more information, see the
// [Oracle Database CLI Reference]({{DOC_SERVER_URL}}/Content/Database/References/odacli.htm#Oracle_Database_CLI_Reference).
func (client DatabaseClient) LaunchDbSystem(ctx context.Context, request LaunchDbSystemRequest) (response LaunchDbSystemResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/dbSystems", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Gets a list of backups based on the databaseId or compartmentId specified. Either one of the query parameters must be provided.
func (client DatabaseClient) ListBackups(ctx context.Context, request ListBackupsRequest) (response ListBackupsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/backups", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Lists all Data Guard associations for the specified database.
func (client DatabaseClient) ListDataGuardAssociations(ctx context.Context, request ListDataGuardAssociationsRequest) (response ListDataGuardAssociationsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/databases/{databaseId}/dataGuardAssociations", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Gets a list of the databases in the specified database home.
func (client DatabaseClient) ListDatabases(ctx context.Context, request ListDatabasesRequest) (response ListDatabasesResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/databases", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Gets history of the actions taken for patches for the specified database home.
func (client DatabaseClient) ListDbHomePatchHistoryEntries(ctx context.Context, request ListDbHomePatchHistoryEntriesRequest) (response ListDbHomePatchHistoryEntriesResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/dbHomes/{dbHomeId}/patchHistoryEntries", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Lists patches applicable to the requested database home.
func (client DatabaseClient) ListDbHomePatches(ctx context.Context, request ListDbHomePatchesRequest) (response ListDbHomePatchesResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/dbHomes/{dbHomeId}/patches", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Gets a list of database homes in the specified DB System and compartment. A database home is a directory where Oracle database software is installed.
func (client DatabaseClient) ListDbHomes(ctx context.Context, request ListDbHomesRequest) (response ListDbHomesResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/dbHomes", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Gets a list of database nodes in the specified DB System and compartment. A database node is a server running database software.
func (client DatabaseClient) ListDbNodes(ctx context.Context, request ListDbNodesRequest) (response ListDbNodesResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/dbNodes", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Gets the history of the patch actions performed on the specified DB System.
func (client DatabaseClient) ListDbSystemPatchHistoryEntries(ctx context.Context, request ListDbSystemPatchHistoryEntriesRequest) (response ListDbSystemPatchHistoryEntriesResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/dbSystems/{dbSystemId}/patchHistoryEntries", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Lists the patches applicable to the requested DB System.
func (client DatabaseClient) ListDbSystemPatches(ctx context.Context, request ListDbSystemPatchesRequest) (response ListDbSystemPatchesResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/dbSystems/{dbSystemId}/patches", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Gets a list of the shapes that can be used to launch a new DB System. The shape determines resources to allocate to the DB system - CPU cores and memory for VM shapes; CPU cores, memory and storage for non-VM (or bare metal) shapes.
func (client DatabaseClient) ListDbSystemShapes(ctx context.Context, request ListDbSystemShapesRequest) (response ListDbSystemShapesResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/dbSystemShapes", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Gets a list of the DB Systems in the specified compartment.
//
func (client DatabaseClient) ListDbSystems(ctx context.Context, request ListDbSystemsRequest) (response ListDbSystemsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/dbSystems", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Gets a list of supported Oracle database versions.
func (client DatabaseClient) ListDbVersions(ctx context.Context, request ListDbVersionsRequest) (response ListDbVersionsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/dbVersions", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Reinstates the database identified by the `databaseId` parameter into the standby role in a Data Guard association.
func (client DatabaseClient) ReinstateDataGuardAssociation(ctx context.Context, request ReinstateDataGuardAssociationRequest) (response ReinstateDataGuardAssociationResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/databases/{databaseId}/dataGuardAssociations/{dataGuardAssociationId}/actions/reinstate", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Restore a Database based on the request parameters you provide.
func (client DatabaseClient) RestoreDatabase(ctx context.Context, request RestoreDatabaseRequest) (response RestoreDatabaseResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/databases/{databaseId}/actions/restore", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Performs a switchover to transition the primary database of a Data Guard association into a standby role. The
// standby database associated with the `dataGuardAssociationId` assumes the primary database role.
// A switchover guarantees no data loss.
func (client DatabaseClient) SwitchoverDataGuardAssociation(ctx context.Context, request SwitchoverDataGuardAssociationRequest) (response SwitchoverDataGuardAssociationResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/databases/{databaseId}/dataGuardAssociations/{dataGuardAssociationId}/actions/switchover", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Terminates a DB System and permanently deletes it and any databases running on it, and any storage volumes attached to it. The database data is local to the DB System and will be lost when the system is terminated. Oracle recommends that you back up any data in the DB System prior to terminating it.
func (client DatabaseClient) TerminateDbSystem(ctx context.Context, request TerminateDbSystemRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/dbSystems/{dbSystemId}", request)
	if err != nil {
		return
	}

	_, err = client.Call(ctx, &httpRequest)
	return
}

// Update a Database based on the request parameters you provide.
func (client DatabaseClient) UpdateDatabase(ctx context.Context, request UpdateDatabaseRequest) (response UpdateDatabaseResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/databases/{databaseId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Patches the specified dbHome.
func (client DatabaseClient) UpdateDbHome(ctx context.Context, request UpdateDbHomeRequest) (response UpdateDbHomeResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/dbHomes/{dbHomeId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Updates the properties of a DB System, such as the CPU core count.
func (client DatabaseClient) UpdateDbSystem(ctx context.Context, request UpdateDbSystemRequest) (response UpdateDbSystemResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/dbSystems/{dbSystemId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}