import * as k8s from "@pulumi/kubernetes";
import * as pulumi from "@pulumi/pulumi";
import { Infra, InfraConfig } from "./util";

/**
 * Kubernetes infrastructure input
 */
export interface IK8sInput {
    kubeconfig: pulumi.Output<string>;
}

/**
 * Kubernetes infrastructure output
 */
export interface IK8sOutput {
    namespaceName: pulumi.Output<string>;
}

/**
 * Kubernetes infrastructure class that implements the abstract class infra
 * * kubernetes setup:
 *  - namespace
 */
export class K8sInfra extends Infra {
    private in: IK8sInput;
    constructor(config: InfraConfig, input: IK8sInput) {
        super(config);
        this.in = input;
    }

    create(): IK8sOutput {
        // Create a Kubernetes provider instance that uses our cluster from above.
        const clusterProvider = new k8s.Provider(
            this.GetName("k8s", "provider"),
            {
                kubeconfig: this.in.kubeconfig,
            }
        );

        // Create a Kubernetes Namespace
        const ns = new k8s.core.v1.Namespace(
            this.GetName("ns"),
            {},
            { provider: clusterProvider }
        );

        // Export the Namespace name
        const namespaceName = ns.metadata.apply((m) => m.name);

        // set k8s-infra output
        return { namespaceName };
    }
}
