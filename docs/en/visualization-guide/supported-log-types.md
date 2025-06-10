# Supported Log Types

This greatly helps us to understand the detail of a resource but it's quite hard to understand it. This is an example k8s audit log recorded in a GKE cluster.

K8s audit logs are recorded by kube-apiserver when a client requested some changes on a kubernetes resource. It can contain the full manifest of the resource or sometimes it only contains the patch diff from the previous resource.

KHI diff(KHI shows timeline and the diff of resource from the Kubernetes audit log)

NOTE: KHI's resource timeline reflects changes logged by kube-apiserver. Resources without modifications within the query period will not appear on KHI.
