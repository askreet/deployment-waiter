// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cloudsearchdomainiface provides an interface for the Amazon CloudSearch Domain.
package cloudsearchdomainiface

import (
	"github.com/awslabs/aws-sdk-go/service/cloudsearchdomain"
)

// CloudSearchDomainAPI is the interface type for cloudsearchdomain.CloudSearchDomain.
type CloudSearchDomainAPI interface {
	Search(*cloudsearchdomain.SearchInput) (*cloudsearchdomain.SearchOutput, error)

	Suggest(*cloudsearchdomain.SuggestInput) (*cloudsearchdomain.SuggestOutput, error)

	UploadDocuments(*cloudsearchdomain.UploadDocumentsInput) (*cloudsearchdomain.UploadDocumentsOutput, error)
}
