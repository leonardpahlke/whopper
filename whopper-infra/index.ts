import * as pulumi from "@pulumi/pulumi";
import * as kx from "@pulumi/kubernetesx";
import * as k8s from "@pulumi/kubernetes";
import InfraConfig from "./config";
import { GcpInfra } from "./gcp-infra";
import { K8sInfra } from "./k8s-infra";
import * as operator from "./operator";

const prdStack = "prd"
const stack = pulumi.getStack();

//
// PULUMI KUBERNETES OPERATOR
//  The pulumi operator gets deployed to a seperate kubernetes cluster and checks periodically for 
//  new resources in a referenced git repository. Resources can be anything defined in pulumi including
//  other kubernetes cluster and cross cloud provider services. 
//
if (stack === prdStack) {
    // access pulumi configuration
    const config = new pulumi.Config();

    const pulumiAccessToken = config.requireSecret("pulumiAccessToken");
    const stackName = config.require("stackName");
    const stackProjectRepo = config.require("stackProjectRepo");
    const stackCommit = config.require("stackCommit");
    const operatorKubeconfig = config.requireSecret("operatorKubeconfig");
    const operatorPodName = config.require("operatorPodName")

    const provider = new k8s.Provider("k8s", { kubeconfig: operatorKubeconfig });

    // Create the Pulumi Kubernetes Operator
    const pulumiOperator = new operator.PulumiKubernetesOperator(operatorPodName, {
        namespace: "default",
        provider,
    });

    // Create the API token as a Kubernetes Secret.
    const apiAccessToken = new kx.Secret("accesstoken", {
        stringData: { accessToken: pulumiAccessToken },
    });

    // Create a Blue/Green app deployment in-cluster.
    const appStack = new k8s.apiextensions.CustomResource("app-stack", {
        apiVersion: 'pulumi.com/v1',
        kind: 'Stack',
        spec: {
            envRefs: {
                PULUMI_ACCESS_TOKEN: {
                    type: "Secret",
                    secret: {
                        name: apiAccessToken.metadata.name,
                        key: "accessToken",
                    },
                },
            },
            stack: stackName,
            projectRepo: stackProjectRepo,
            commit: stackCommit,
            destroyOnFinalize: true,
        }
    }, { dependsOn: pulumiOperator.deployment });
    appStack.id.get()
}

//
// PROJECT INFRASTRUCTURE RESOURCES
//  Infrastructure setup which is getting deployed by the pulumi k8s operator.
//  The infrastructure describes the system which is used to process articles.
//
// infraConfig accesses pulumi configuration
const infraConfig = new InfraConfig();

// manage kubernetes infrastructure
const gcpInfraOut = new GcpInfra(infraConfig, {}).create();
const clusterKubeconfig = gcpInfraOut.kubeconfig

// Configure Kubernetes cluster
new K8sInfra(infraConfig, {
    kubeconfig: clusterKubeconfig,
}).create();
