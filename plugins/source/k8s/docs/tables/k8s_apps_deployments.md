
# Table: k8s_apps_deployments
Deployment enables declarative updates for Pods and ReplicaSets.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|context|text|Name of the context from k8s configuration.|
|kind|text|Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds|
|api_version|text|APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources|
|name|text|Name must be unique within a namespace|
|generate_name|text|GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed|
|namespace|text|Namespace defines the space within which each name must be unique|
|self_link|text|SelfLink is a URL representing this object. Populated by the system. Read-only.  DEPRECATED Kubernetes will stop propagating this field in 1.20 release and the field is planned to be removed in 1.21 release.|
|uid|text|UID is the unique in time and space value for this object|
|resource_version|text|An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed|
|generation|bigint|A sequence number representing a specific generation of the desired state. Populated by the system|
|deletion_grace_period_seconds|bigint|Number of seconds allowed for this object to gracefully terminate before it will be removed from the system|
|labels|jsonb|Map of string keys and values that can be used to organize and categorize (scope and select) objects|
|annotations|jsonb|Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata|
|owner_references|jsonb|List of objects depended by this object|
|finalizers|text[]|Must be empty before the object is deleted from the registry|
|zzz_cluster_name|text|Deprecated: ClusterName is a legacy field that was always cleared by the system and never used; it will be removed in the future. The name in the database is changed to help clients detect accidental use.|
|managed_fields|jsonb|ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow|
|replicas|integer|Number of desired pods|
|selector_match_labels|jsonb|matchLabels is a map of {key,value} pairs|
|template|jsonb|Template describes the pods that will be created.|
|strategy_type|text|Type of deployment|
|strategy_rolling_update_max_unavailable_type|bigint||
|strategy_rolling_update_max_unavailable_int_val|integer||
|strategy_rolling_update_max_unavailable_str_val|text||
|strategy_rolling_update_max_surge_type|bigint||
|strategy_rolling_update_max_surge_int_val|integer||
|strategy_rolling_update_max_surge_str_val|text||
|min_ready_seconds|integer|Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)|
|revision_history_limit|integer|The number of old ReplicaSets to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10.|
|paused|boolean|Indicates that the deployment is paused.|
|progress_deadline_seconds|integer|The maximum time in seconds for a deployment to make progress before it is considered to be failed|
|status_observed_generation|bigint|The generation observed by the deployment controller.|
|status_replicas|integer|Total number of non-terminated pods targeted by this deployment (their labels match the selector).|
|status_updated_replicas|integer|Total number of non-terminated pods targeted by this deployment that have the desired template spec.|
|status_ready_replicas|integer|Total number of ready pods targeted by this deployment.|
|status_available_replicas|integer|Total number of available pods (ready for at least minReadySeconds) targeted by this deployment.|
|status_unavailable_replicas|integer|Total number of unavailable pods targeted by this deployment|
|status_collision_count|integer|Count of hash collisions for the Deployment|
