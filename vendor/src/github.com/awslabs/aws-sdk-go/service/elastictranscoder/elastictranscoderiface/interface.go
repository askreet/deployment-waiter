// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package elastictranscoderiface provides an interface for the Amazon Elastic Transcoder.
package elastictranscoderiface

import (
	"github.com/awslabs/aws-sdk-go/service/elastictranscoder"
)

// ElasticTranscoderAPI is the interface type for elastictranscoder.ElasticTranscoder.
type ElasticTranscoderAPI interface {
	CancelJob(*elastictranscoder.CancelJobInput) (*elastictranscoder.CancelJobOutput, error)

	CreateJob(*elastictranscoder.CreateJobInput) (*elastictranscoder.CreateJobResponse, error)

	CreatePipeline(*elastictranscoder.CreatePipelineInput) (*elastictranscoder.CreatePipelineOutput, error)

	CreatePreset(*elastictranscoder.CreatePresetInput) (*elastictranscoder.CreatePresetOutput, error)

	DeletePipeline(*elastictranscoder.DeletePipelineInput) (*elastictranscoder.DeletePipelineOutput, error)

	DeletePreset(*elastictranscoder.DeletePresetInput) (*elastictranscoder.DeletePresetOutput, error)

	ListJobsByPipeline(*elastictranscoder.ListJobsByPipelineInput) (*elastictranscoder.ListJobsByPipelineOutput, error)

	ListJobsByStatus(*elastictranscoder.ListJobsByStatusInput) (*elastictranscoder.ListJobsByStatusOutput, error)

	ListPipelines(*elastictranscoder.ListPipelinesInput) (*elastictranscoder.ListPipelinesOutput, error)

	ListPresets(*elastictranscoder.ListPresetsInput) (*elastictranscoder.ListPresetsOutput, error)

	ReadJob(*elastictranscoder.ReadJobInput) (*elastictranscoder.ReadJobOutput, error)

	ReadPipeline(*elastictranscoder.ReadPipelineInput) (*elastictranscoder.ReadPipelineOutput, error)

	ReadPreset(*elastictranscoder.ReadPresetInput) (*elastictranscoder.ReadPresetOutput, error)

	TestRole(*elastictranscoder.TestRoleInput) (*elastictranscoder.TestRoleOutput, error)

	UpdatePipeline(*elastictranscoder.UpdatePipelineInput) (*elastictranscoder.UpdatePipelineOutput, error)

	UpdatePipelineNotifications(*elastictranscoder.UpdatePipelineNotificationsInput) (*elastictranscoder.UpdatePipelineNotificationsOutput, error)

	UpdatePipelineStatus(*elastictranscoder.UpdatePipelineStatusInput) (*elastictranscoder.UpdatePipelineStatusOutput, error)
}
