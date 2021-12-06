import { InfraConfig } from "./infra/util";
import { GcpInfra } from "./infra/gcp-infra";
import { K8sInfra } from "./infra/k8s-infra";

// create infrastructure config used to get pulumi config
// this class is being injected into gcp-infra & k8s-infra
const infraConfig = new InfraConfig();

// GCP infrastructure
// * GKE cluster
const gcpInfraOut = new GcpInfra(infraConfig, {}).create();

// Kubernetes on GKE cluster
// * namespace
const k8sInfraOut = new K8sInfra(infraConfig, {
    kubeconfig: gcpInfraOut.kubeconfig,
}).create();
