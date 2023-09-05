//go:build unit || !integration

package scheduler

import (
	"context"
	"testing"

	"github.com/bacalhau-project/bacalhau/pkg/jobstore"
	"github.com/bacalhau-project/bacalhau/pkg/models"
	"github.com/bacalhau-project/bacalhau/pkg/orchestrator"
	"github.com/bacalhau-project/bacalhau/pkg/test/mock"
	"github.com/google/uuid"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type OpsJobSchedulerTestSuite struct {
	suite.Suite
	jobStore       *jobstore.MockStore
	planner        *orchestrator.MockPlanner
	nodeDiscoverer *orchestrator.MockNodeDiscoverer
	nodeRanker     *orchestrator.MockNodeRanker
	scheduler      *OpsJobScheduler
}

func (s *OpsJobSchedulerTestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.jobStore = jobstore.NewMockStore(ctrl)
	s.planner = orchestrator.NewMockPlanner(ctrl)
	s.nodeDiscoverer = orchestrator.NewMockNodeDiscoverer(ctrl)
	s.nodeRanker = orchestrator.NewMockNodeRanker(ctrl)

	s.scheduler = NewOpsJobScheduler(OpsJobSchedulerParams{
		JobStore:       s.jobStore,
		Planner:        s.planner,
		NodeDiscoverer: s.nodeDiscoverer,
		NodeRanker:     s.nodeRanker,
	})
}

func TestOpsJobSchedulerTestSuite(t *testing.T) {
	suite.Run(t, new(OpsJobSchedulerTestSuite))
}

func (s *OpsJobSchedulerTestSuite) TestProcess_ShouldCreateNewExecutions() {
	ctx := context.Background()
	job, executions, evaluation := mockOpsJob()
	executions = []models.Execution{}
	s.jobStore.EXPECT().GetJob(gomock.Any(), job.ID).Return(*job, nil)
	s.jobStore.EXPECT().GetExecutions(gomock.Any(), job.ID).Return(executions, nil)

	nodeInfos := []models.NodeInfo{
		*mockNodeInfo(s.T(), nodeIDs[0]),
		*mockNodeInfo(s.T(), nodeIDs[1]),
		*mockNodeInfo(s.T(), nodeIDs[2]),
	}
	s.mockNodeSelection(job, nodeInfos)

	matcher := NewPlanMatcher(s.T(), PlanMatcherParams{
		Evaluation:               evaluation,
		NewExecutionDesiredState: models.ExecutionDesiredStateRunning,
		NewExecutionsNodes: []peer.ID{
			nodeInfos[0].PeerInfo.ID,
			nodeInfos[1].PeerInfo.ID,
			nodeInfos[2].PeerInfo.ID,
		},
	})
	s.planner.EXPECT().Process(gomock.Any(), matcher).Times(1)
	s.Require().NoError(s.scheduler.Process(ctx, evaluation))
}

func (s *OpsJobSchedulerTestSuite) TestProcess_ShouldMarkJobAsCompleted() {
	ctx := context.Background()
	job, executions, evaluation := mockOpsJob()
	executions[0].ComputeState = models.NewExecutionState(models.ExecutionStateCompleted) // Simulate a completed execution
	s.jobStore.EXPECT().GetJob(gomock.Any(), job.ID).Return(*job, nil)
	s.jobStore.EXPECT().GetExecutions(gomock.Any(), job.ID).Return(executions, nil)

	matcher := NewPlanMatcher(s.T(), PlanMatcherParams{
		Evaluation: evaluation,
		JobState:   models.JobStateTypeCompleted,
	})
	s.planner.EXPECT().Process(gomock.Any(), matcher).Times(1)
	s.Require().NoError(s.scheduler.Process(ctx, evaluation))
}

func (s *OpsJobSchedulerTestSuite) TestProcess_ShouldMarkLostExecutionsOnUnhealthyNodes() {
	ctx := context.Background()
	job, executions, evaluation := mockOpsJob()
	executions[0].ComputeState = models.NewExecutionState(models.ExecutionStateBidAccepted)
	executions[1].ComputeState = models.NewExecutionState(models.ExecutionStateBidAccepted)
	s.jobStore.EXPECT().GetJob(gomock.Any(), job.ID).Return(*job, nil)
	s.jobStore.EXPECT().GetExecutions(gomock.Any(), job.ID).Return(executions, nil)

	// mock node discoverer to exclude the first node
	nodeInfos := []models.NodeInfo{
		*mockNodeInfo(s.T(), executions[1].NodeID),
	}
	s.nodeDiscoverer.EXPECT().ListNodes(gomock.Any()).Return(nodeInfos, nil)

	matcher := NewPlanMatcher(s.T(), PlanMatcherParams{
		Evaluation:         evaluation,
		NewExecutionsNodes: []peer.ID{},
		StoppedExecutions: []string{
			executions[0].ID,
		},
	})
	s.planner.EXPECT().Process(gomock.Any(), matcher).Times(1)
	s.Require().NoError(s.scheduler.Process(ctx, evaluation))
}

func (s *OpsJobSchedulerTestSuite) TestProcess_ShouldMarkJobAsFailed() {
	ctx := context.Background()
	job, executions, evaluation := mockOpsJob()
	executions[0].ComputeState = models.NewExecutionState(models.ExecutionStateBidAccepted) // running, but lost node
	executions[1].ComputeState = models.NewExecutionState(models.ExecutionStateFailed)      // failed execution
	s.jobStore.EXPECT().GetJob(gomock.Any(), job.ID).Return(*job, nil)
	s.jobStore.EXPECT().GetExecutions(gomock.Any(), job.ID).Return(executions, nil)

	// mock node discoverer to exclude the first node
	nodeInfos := []models.NodeInfo{
		*mockNodeInfo(s.T(), executions[1].NodeID),
	}
	s.nodeDiscoverer.EXPECT().ListNodes(gomock.Any()).Return(nodeInfos, nil)

	matcher := NewPlanMatcher(s.T(), PlanMatcherParams{
		Evaluation: evaluation,
		JobState:   models.JobStateTypeFailed,
		StoppedExecutions: []string{
			executions[0].ID,
		},
	})
	s.planner.EXPECT().Process(gomock.Any(), matcher).Times(1)
	s.Require().NoError(s.scheduler.Process(ctx, evaluation))
}

func (s *OpsJobSchedulerTestSuite) TestProcess_WhenJobIsStopped_ShouldMarkNonTerminalExecutionsAsStopped() {
	ctx := context.Background()
	job, executions, evaluation := mockOpsJob()
	job.State = models.NewJobState(models.JobStateTypeStopped) // Simulate a canceled job
	s.jobStore.EXPECT().GetJob(gomock.Any(), job.ID).Return(*job, nil)
	s.jobStore.EXPECT().GetExecutions(gomock.Any(), job.ID).Return(executions, nil)

	matcher := NewPlanMatcher(s.T(), PlanMatcherParams{
		Evaluation: evaluation,
		StoppedExecutions: []string{
			executions[0].ID,
		},
	})
	s.planner.EXPECT().Process(gomock.Any(), matcher).Times(1)
	s.Require().NoError(s.scheduler.Process(ctx, evaluation))
}

func (s *OpsJobSchedulerTestSuite) TestProcessFail_NoMatchingNodes() {
	ctx := context.Background()
	job, executions, evaluation := mockOpsJob()
	executions = []models.Execution{} // no executions yet
	s.jobStore.EXPECT().GetJob(gomock.Any(), job.ID).Return(*job, nil)
	s.jobStore.EXPECT().GetExecutions(gomock.Any(), job.ID).Return(executions, nil)
	s.mockNodeSelection(job, []models.NodeInfo{})

	matcher := NewPlanMatcher(s.T(), PlanMatcherParams{
		Evaluation: evaluation,
		JobState:   models.JobStateTypeFailed,
	})
	s.planner.EXPECT().Process(gomock.Any(), matcher).Times(1)
	s.Require().NoError(s.scheduler.Process(ctx, evaluation))
}

func (s *OpsJobSchedulerTestSuite) mockNodeSelection(job *models.Job, nodeInfos []models.NodeInfo) {
	s.nodeDiscoverer.EXPECT().FindNodes(gomock.Any(), *job).Return(nodeInfos, nil)

	nodeRanks := make([]orchestrator.NodeRank, len(nodeInfos))
	for i, nodeInfo := range nodeInfos {
		nodeRanks[i] = orchestrator.NodeRank{
			NodeInfo: nodeInfo,
			Rank:     i,
		}
	}
	s.nodeRanker.EXPECT().RankNodes(gomock.Any(), *job, nodeInfos).Return(nodeRanks, nil)
}

func mockOpsJob() (*models.Job, []models.Execution, *models.Evaluation) {
	job := mock.Job()

	executionCount := 2
	executions := make([]models.Execution, executionCount)
	for i, e := range mock.Executions(job, executionCount) {
		e.NodeID = nodeIDs[i]
		executions[i] = *e
	}

	executions[0].ComputeState = models.NewExecutionState(models.ExecutionStateBidAccepted)
	executions[1].ComputeState = models.NewExecutionState(models.ExecutionStateCompleted)

	evaluation := &models.Evaluation{
		JobID: job.ID,
		Type:  models.JobTypeOps,
		ID:    uuid.NewString(),
	}

	return job, executions, evaluation
}