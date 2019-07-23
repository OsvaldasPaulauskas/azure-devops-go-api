﻿// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package pyPiApi

import (
    "bytes"
    "context"
    "encoding/json"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
    "net/url"
    "strconv"
)

var ResourceAreaId, _ = uuid.Parse("92f0314b-06c5-46e0-abe7-15fd9d13276a")

type Client struct {
    Client azureDevops.Client
}

func NewClient(ctx context.Context, connection azureDevops.Connection) (*Client, error) {
    client, err := connection.GetClientByResourceAreaId(ctx, ResourceAreaId)
    if err != nil {
        return nil, err
    }
    return &Client {
        Client: *client,
    }, nil
}

// [Preview API] Download a python package file directly. This API is intended for manual UI download options, not for programmatic access and scripting.
// ctx
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
// fileName (required): Name of the file in the package
// project (optional): Project ID or project name
func (client Client) DownloadPackage(ctx context.Context, feedId *string, packageName *string, packageVersion *string, fileName *string, project *string) (interface{}, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageName == nil || *packageName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *packageName
    if packageVersion == nil || *packageVersion == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
    }
    routeValues["packageVersion"] = *packageVersion
    if fileName == nil || *fileName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "fileName"} 
    }
    routeValues["fileName"] = *fileName

    locationId, _ := uuid.Parse("97218bae-a64d-4381-9257-b5b7951f0b98")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, responseValue)
    return responseValue, err
}

// [Preview API] Delete a package version from the feed, moving it to the recycle bin.
// ctx
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
// project (optional): Project ID or project name
func (client Client) DeletePackageVersionFromRecycleBin(ctx context.Context, feedId *string, packageName *string, packageVersion *string, project *string) error {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageName == nil || *packageName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *packageName
    if packageVersion == nil || *packageVersion == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
    }
    routeValues["packageVersion"] = *packageVersion

    locationId, _ := uuid.Parse("07143752-3d94-45fd-86c2-0c77ed87847b")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get information about a package version in the recycle bin.
// ctx
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
// project (optional): Project ID or project name
func (client Client) GetPackageVersionMetadataFromRecycleBin(ctx context.Context, feedId *string, packageName *string, packageVersion *string, project *string) (*PyPiPackageVersionDeletionState, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageName == nil || *packageName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *packageName
    if packageVersion == nil || *packageVersion == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
    }
    routeValues["packageVersion"] = *packageVersion

    locationId, _ := uuid.Parse("07143752-3d94-45fd-86c2-0c77ed87847b")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue PyPiPackageVersionDeletionState
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Restore a package version from the recycle bin to its associated feed.
// ctx
// packageVersionDetails (required): Set the 'Deleted' state to 'false' to restore the package to its feed.
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
// project (optional): Project ID or project name
func (client Client) RestorePackageVersionFromRecycleBin(ctx context.Context, packageVersionDetails *PyPiRecycleBinPackageVersionDetails, feedId *string, packageName *string, packageVersion *string, project *string) error {
    if packageVersionDetails == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "packageVersionDetails"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageName == nil || *packageName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *packageName
    if packageVersion == nil || *packageVersion == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
    }
    routeValues["packageVersion"] = *packageVersion

    body, marshalErr := json.Marshal(*packageVersionDetails)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("07143752-3d94-45fd-86c2-0c77ed87847b")
    _, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Delete a package version, moving it to the recycle bin.
// ctx
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
// project (optional): Project ID or project name
func (client Client) DeletePackageVersion(ctx context.Context, feedId *string, packageName *string, packageVersion *string, project *string) (*Package, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageName == nil || *packageName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *packageName
    if packageVersion == nil || *packageVersion == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
    }
    routeValues["packageVersion"] = *packageVersion

    locationId, _ := uuid.Parse("d146ac7e-9e3f-4448-b956-f9bb3bdf9b2e")
    resp, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Package
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get information about a package version.
// ctx
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
// project (optional): Project ID or project name
// showDeleted (optional): True to show information for deleted package versions.
func (client Client) GetPackageVersion(ctx context.Context, feedId *string, packageName *string, packageVersion *string, project *string, showDeleted *bool) (*Package, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageName == nil || *packageName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *packageName
    if packageVersion == nil || *packageVersion == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
    }
    routeValues["packageVersion"] = *packageVersion

    queryParams := url.Values{}
    if showDeleted != nil {
        queryParams.Add("showDeleted", strconv.FormatBool(*showDeleted))
    }
    locationId, _ := uuid.Parse("d146ac7e-9e3f-4448-b956-f9bb3bdf9b2e")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Package
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update state for a package version.
// ctx
// packageVersionDetails (required): Details to be updated.
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
// project (optional): Project ID or project name
func (client Client) UpdatePackageVersion(ctx context.Context, packageVersionDetails *PackageVersionDetails, feedId *string, packageName *string, packageVersion *string, project *string) error {
    if packageVersionDetails == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "packageVersionDetails"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageName == nil || *packageName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *packageName
    if packageVersion == nil || *packageVersion == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
    }
    routeValues["packageVersion"] = *packageVersion

    body, marshalErr := json.Marshal(*packageVersionDetails)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("d146ac7e-9e3f-4448-b956-f9bb3bdf9b2e")
    _, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}
