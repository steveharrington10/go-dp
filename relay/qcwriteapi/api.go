// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * QC Write API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 */

package qcwriteapi

import (
	"context"
	"net/http"
)



// WriteApiAPIRouter defines the required methods for binding the api requests to a responses for the WriteApiAPI
// The WriteApiAPIRouter implementation should parse necessary information from the http request,
// pass the data to a WriteApiAPIServicer to perform the required actions, then write the service results to the http response.
type WriteApiAPIRouter interface { 
	SendTransaction(http.ResponseWriter, *http.Request)
}


// WriteApiAPIServicer defines the api actions for the WriteApiAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type WriteApiAPIServicer interface { 
	SendTransaction(context.Context, SendTransactionRequest) (ImplResponse, error)
}
