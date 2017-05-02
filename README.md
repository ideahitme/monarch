# Monarch

[Monarch butterfly image]

*s/monarch/leader/g*

Kubernetes component which implements **monarch** election for resources in Kubernetes cluster. 

1. Opens API which given a query and identifier of the pod in the form of `Name of Deployment/Namespace of Deployment/Name of Pod` returns if the pod is a **monarch**
2. Automatically parses deployments in the kubernetes clusters with the annotation `` and automatically maintains **monarch** information
3. Maintains its data in its own separate configmap
4. Returns false for the queries on pods which are not registered (i.e. do not have correct annotation) - this is to prevent race condition

Notes: 

Use [Downward API](https://github.com/kubernetes/kubernetes/blob/release-1.0/docs/user-guide/downward-api.md) to extract Pod name in run-time