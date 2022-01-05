import { getStack, all, Config } from "@pulumi/pulumi";
import InfraConfig from "./infra/config";
import { GcpInfra } from "./infra/gcp-infra";
import { K8sInfra } from "./infra/k8s-infra";

// pulumi stacks 
const localStack = "local"
const prdStack = "prd"

// get stack name
const stack = getStack();

if (stack !== localStack && stack !== prdStack) {
  throw new Error("Stack unrecognized configure stack infrastructure first")
}

// create infrastructure config used to get pulumi config
// this class is being injected into gcp-infra & k8s-infra
const infraConfig = new InfraConfig();

// create empty
let clusterKubeconfig = all([]).apply(([]) => "");

// manage kubernetes infrastructure
if (stack === prdStack) {
  // use the GKE cluster
  const gcpInfraOut = new GcpInfra(infraConfig, {}).create();
  clusterKubeconfig = gcpInfraOut.kubeconfig
} else if (stack === localStack) {
  // use local cluster
  const pulumiConfig = new Config()
  clusterKubeconfig = pulumiConfig.requireSecret("kconfig");
}

// Configure Kubernetes
// * namespace
new K8sInfra(infraConfig, {
  kubeconfig: clusterKubeconfig,
}).create();
