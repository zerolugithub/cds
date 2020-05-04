package workflow_test

import (
	"context"
	"testing"

	"github.com/ovh/cds/engine/api/test"
	"github.com/ovh/cds/engine/api/workflow"
	"github.com/stretchr/testify/require"
)

func TestLoadAllWorkflows(t *testing.T) {
	db, _, end := test.SetupPG(t)
	defer end()

	var opts = []workflow.LoadAllWorkflowsOptions{
		{},
		{
			Filters: workflow.LoadAllWorkflowsOptionsFilters{
				ProjectKey: "test",
			},
		},
		{
			Filters: workflow.LoadAllWorkflowsOptionsFilters{
				WorkflowName: "test",
			},
		},
		{
			Filters: workflow.LoadAllWorkflowsOptionsFilters{
				Repository: "test",
			},
		},
		{
			Filters: workflow.LoadAllWorkflowsOptionsFilters{
				VCSServer: "test",
			},
		},
		{
			Filters: workflow.LoadAllWorkflowsOptionsFilters{
				GroupIDs: []int64{1, 2, 3, 4},
			},
		},
		{
			Filters: workflow.LoadAllWorkflowsOptionsFilters{
				ProjectKey:   "test",
				WorkflowName: "test",
				Repository:   "test",
				VCSServer:    "test",
				GroupIDs:     []int64{1, 2, 3, 4},
			},
		},
		{
			Filters: workflow.LoadAllWorkflowsOptionsFilters{
				ProjectKey: "test",
				Repository: "test",
				GroupIDs:   []int64{1, 2, 3, 4},
			},
			Loaders: workflow.LoadAllWorkflowsOptionsLoaders{
				WithAsCodeUpdateEvents: true,
				WithEnvironments:       true,
				WithApplications:       true,
				WithIcon:               true,
				WithIntegrations:       true,
				WithPipelines:          true,
				WithTemplate:           true,
			},
		},
		{
			Filters: workflow.LoadAllWorkflowsOptionsFilters{},
			Loaders: workflow.LoadAllWorkflowsOptionsLoaders{
				WithAsCodeUpdateEvents: true,
				WithEnvironments:       true,
				WithApplications:       true,
				WithIcon:               true,
				WithIntegrations:       true,
				WithPipelines:          true,
				WithTemplate:           true,
			},
		},
	}

	for _, opt := range opts {
		_, err := workflow.LoadAllWorkflows(context.TODO(), db, opt)
		require.NoError(t, err)
	}
}
