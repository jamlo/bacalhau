// This file is auto-generated by @hey-api/openapi-ts

import { createClient, createConfig, type Options } from '@hey-api/client-fetch';
import type { HomeError, HomeResponse, IdError, IdResponse, LivezError, LivezResponse, NodeInfoError, NodeInfoResponse, AgentAliveError, AgentAliveResponse, AgentConfigError, AgentConfigResponse, AgentDebugError, AgentDebugResponse, AgentNodeError, AgentNodeResponse, AgentVersionError, AgentVersionResponse, ApiServerDebugError, ApiServerDebugResponse, OrchestratorListJobsData, OrchestratorListJobsError, OrchestratorListJobsResponse, OrchestratorPutJobData, OrchestratorPutJobError, OrchestratorPutJobResponse, OrchestratorGetJobData, OrchestratorGetJobError, OrchestratorGetJobResponse, OrchestratorStopJobData, OrchestratorStopJobError, OrchestratorStopJobResponse, OrchestratorJobExecutionsData, OrchestratorJobExecutionsError, OrchestratorJobExecutionsResponse, OrchestratorListHistoryData, OrchestratorListHistoryError, OrchestratorListHistoryResponse, OrchestratorLogsData, OrchestratorJobResultsData, OrchestratorJobResultsError, OrchestratorJobResultsResponse, OrchestratorListNodesData, OrchestratorListNodesError, OrchestratorListNodesResponse, OrchestratorUpdateNodeData, OrchestratorUpdateNodeError, OrchestratorUpdateNodeResponse, OrchestratorGetNodeData, OrchestratorGetNodeError, OrchestratorGetNodeResponse, ApiServerVersionData, ApiServerVersionError, ApiServerVersionResponse } from './types.gen';

export const client = createClient(createConfig());

export class Utils {
    public static home<ThrowOnError extends boolean = false>(options?: Options<unknown, ThrowOnError>) {
        return (options?.client ?? client).get<HomeResponse, HomeError, ThrowOnError>({
            ...options,
            url: '/'
        });
    }
    
    /**
     * Returns the id of the host node.
     */
    public static id<ThrowOnError extends boolean = false>(options?: Options<unknown, ThrowOnError>) {
        return (options?.client ?? client).get<IdResponse, IdError, ThrowOnError>({
            ...options,
            url: '/api/v1/id'
        });
    }
    
    public static livez<ThrowOnError extends boolean = false>(options?: Options<unknown, ThrowOnError>) {
        return (options?.client ?? client).get<LivezResponse, LivezError, ThrowOnError>({
            ...options,
            url: '/api/v1/livez'
        });
    }
    
    /**
     * Returns the info of the node.
     */
    public static nodeInfo<ThrowOnError extends boolean = false>(options?: Options<unknown, ThrowOnError>) {
        return (options?.client ?? client).get<NodeInfoResponse, NodeInfoError, ThrowOnError>({
            ...options,
            url: '/api/v1/node_info'
        });
    }
    
}

export class Ops {
    public static agentAlive<ThrowOnError extends boolean = false>(options?: Options<unknown, ThrowOnError>) {
        return (options?.client ?? client).get<AgentAliveResponse, AgentAliveError, ThrowOnError>({
            ...options,
            url: '/api/v1/agent/alive'
        });
    }
    
    /**
     * Returns the current configuration of the node.
     */
    public static agentConfig<ThrowOnError extends boolean = false>(options?: Options<unknown, ThrowOnError>) {
        return (options?.client ?? client).get<AgentConfigResponse, AgentConfigError, ThrowOnError>({
            ...options,
            url: '/api/v1/agent/config'
        });
    }
    
    /**
     * Returns debug information on what the current node is doing.
     */
    public static agentDebug<ThrowOnError extends boolean = false>(options?: Options<unknown, ThrowOnError>) {
        return (options?.client ?? client).get<AgentDebugResponse, AgentDebugError, ThrowOnError>({
            ...options,
            url: '/api/v1/agent/debug'
        });
    }
    
    /**
     * Returns the info of the node.
     */
    public static agentNode<ThrowOnError extends boolean = false>(options?: Options<unknown, ThrowOnError>) {
        return (options?.client ?? client).get<AgentNodeResponse, AgentNodeError, ThrowOnError>({
            ...options,
            url: '/api/v1/agent/node'
        });
    }
    
    /**
     * Returns the build version running on the server.
     * See https://github.com/bacalhau-project/bacalhau/releases for a complete list of `gitversion` tags.
     */
    public static agentVersion<ThrowOnError extends boolean = false>(options?: Options<unknown, ThrowOnError>) {
        return (options?.client ?? client).get<AgentVersionResponse, AgentVersionError, ThrowOnError>({
            ...options,
            url: '/api/v1/agent/version'
        });
    }
    
}

export class ComputeNode {
    /**
     * Returns debug information on what the current node is doing.
     */
    public static apiServerDebug<ThrowOnError extends boolean = false>(options?: Options<unknown, ThrowOnError>) {
        return (options?.client ?? client).get<ApiServerDebugResponse, ApiServerDebugError, ThrowOnError>({
            ...options,
            url: '/api/v1/compute/debug'
        });
    }
    
}

export class Orchestrator {
    /**
     * Returns a list of jobs.
     * Returns a list of jobs.
     */
    public static listJobs<ThrowOnError extends boolean = false>(options?: Options<OrchestratorListJobsData, ThrowOnError>) {
        return (options?.client ?? client).get<OrchestratorListJobsResponse, OrchestratorListJobsError, ThrowOnError>({
            ...options,
            url: '/api/v1/orchestrator/jobs'
        });
    }
    
    /**
     * Submits a job to the orchestrator.
     * Submits a job to the orchestrator.
     */
    public static putJob<ThrowOnError extends boolean = false>(options: Options<OrchestratorPutJobData, ThrowOnError>) {
        return (options?.client ?? client).put<OrchestratorPutJobResponse, OrchestratorPutJobError, ThrowOnError>({
            ...options,
            url: '/api/v1/orchestrator/jobs'
        });
    }
    
    /**
     * Returns a job.
     * Returns a job.
     */
    public static getJob<ThrowOnError extends boolean = false>(options: Options<OrchestratorGetJobData, ThrowOnError>) {
        return (options?.client ?? client).get<OrchestratorGetJobResponse, OrchestratorGetJobError, ThrowOnError>({
            ...options,
            url: '/api/v1/orchestrator/jobs/{id}'
        });
    }
    
    /**
     * Stops a job.
     * Stops a job.
     */
    public static stopJob<ThrowOnError extends boolean = false>(options: Options<OrchestratorStopJobData, ThrowOnError>) {
        return (options?.client ?? client).delete<OrchestratorStopJobResponse, OrchestratorStopJobError, ThrowOnError>({
            ...options,
            url: '/api/v1/orchestrator/jobs/{id}'
        });
    }
    
    /**
     * Returns the executions of a job.
     * Returns the executions of a job.
     */
    public static jobExecutions<ThrowOnError extends boolean = false>(options: Options<OrchestratorJobExecutionsData, ThrowOnError>) {
        return (options?.client ?? client).get<OrchestratorJobExecutionsResponse, OrchestratorJobExecutionsError, ThrowOnError>({
            ...options,
            url: '/api/v1/orchestrator/jobs/{id}/executions'
        });
    }
    
    /**
     * Returns the history of a job.
     * Returns the history of a job.
     */
    public static listHistory<ThrowOnError extends boolean = false>(options: Options<OrchestratorListHistoryData, ThrowOnError>) {
        return (options?.client ?? client).get<OrchestratorListHistoryResponse, OrchestratorListHistoryError, ThrowOnError>({
            ...options,
            url: '/api/v1/orchestrator/jobs/{id}/history'
        });
    }
    
    /**
     * Streams the logs for a current job/execution via WebSocket
     * Establishes a WebSocket connection to stream output from the job specified by `id`
     * The stream will continue until either the client disconnects or the execution completes
     */
    public static logs<ThrowOnError extends boolean = false>(options: Options<OrchestratorLogsData, ThrowOnError>) {
        return (options?.client ?? client).get<void, unknown, ThrowOnError>({
            ...options,
            url: '/api/v1/orchestrator/jobs/{id}/logs'
        });
    }
    
    /**
     * Returns the results of a job.
     * Returns the results of a job.
     */
    public static jobResults<ThrowOnError extends boolean = false>(options: Options<OrchestratorJobResultsData, ThrowOnError>) {
        return (options?.client ?? client).get<OrchestratorJobResultsResponse, OrchestratorJobResultsError, ThrowOnError>({
            ...options,
            url: '/api/v1/orchestrator/jobs/{id}/results'
        });
    }
    
    /**
     * Returns a list of orchestrator nodes.
     * Returns a list of orchestrator nodes.
     */
    public static listNodes<ThrowOnError extends boolean = false>(options?: Options<OrchestratorListNodesData, ThrowOnError>) {
        return (options?.client ?? client).get<OrchestratorListNodesResponse, OrchestratorListNodesError, ThrowOnError>({
            ...options,
            url: '/api/v1/orchestrator/nodes'
        });
    }
    
    /**
     * Update an orchestrator node.
     * Update an orchestrator node.
     */
    public static updateNode<ThrowOnError extends boolean = false>(options: Options<OrchestratorUpdateNodeData, ThrowOnError>) {
        return (options?.client ?? client).post<OrchestratorUpdateNodeResponse, OrchestratorUpdateNodeError, ThrowOnError>({
            ...options,
            url: '/api/v1/orchestrator/nodes'
        });
    }
    
    /**
     * Get an orchestrator node
     * Get an orchestrator node
     */
    public static getNode<ThrowOnError extends boolean = false>(options: Options<OrchestratorGetNodeData, ThrowOnError>) {
        return (options?.client ?? client).get<OrchestratorGetNodeResponse, OrchestratorGetNodeError, ThrowOnError>({
            ...options,
            url: '/api/v1/orchestrator/nodes/{id}'
        });
    }
    
}

export class Misc {
    /**
     * Returns the build version running on the server.
     * See https://github.com/bacalhau-project/bacalhau/releases for a complete list of `gitversion` tags.
     */
    public static apiServerVersion<ThrowOnError extends boolean = false>(options: Options<ApiServerVersionData, ThrowOnError>) {
        return (options?.client ?? client).post<ApiServerVersionResponse, ApiServerVersionError, ThrowOnError>({
            ...options,
            url: '/api/v1/version'
        });
    }
    
}