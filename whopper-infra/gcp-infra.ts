import * as gcp from "@pulumi/gcp";
import * as pulumi from "@pulumi/pulumi";
import Infra from "./abstract-infra";
import InfraConfig from "./config";

/**
 * GCP infrastructure input
 */
export type IGcpInput = {};

/**
 * GCP infrastructure output
 */
export type IGcpOutput = {
  clusterName: pulumi.Output<string>;
  kubeconfig: pulumi.Output<string>;
};

/**
 * GCP infrastructure class that implements the abstract class infra
 * * gcp services:
 *  - GKE
 */
export class GcpInfra extends Infra {
  private in: IGcpInput;

  constructor(config: InfraConfig, input: IGcpInput) {
    super(config);
    this.in = input;
  }

  create(): IGcpOutput {
    const engineVersion = gcp.container
      .getEngineVersions()
      .then((v) => v.latestMasterVersion);
    // Create a GKE cluster
    const cluster = new gcp.container.Cluster(this.GetName("cluster"), {
      initialNodeCount: this.config.vars.initialNodeCount,
      minMasterVersion: engineVersion,
      nodeVersion: engineVersion,
      nodeConfig: {
        machineType: this.config.vars.machineType,
        oauthScopes: [
          "https://www.googleapis.com/auth/compute",
          "https://www.googleapis.com/auth/devstorage.read_only",
          "https://www.googleapis.com/auth/logging.write",
          "https://www.googleapis.com/auth/monitoring",
        ],
      },
    });

    // Export the Cluster name
    // const clusterName: pulumi.Output<string> = cluster.name;

    // Kubeconfig
    const kubeconfig = pulumi
      .all([cluster.name, cluster.endpoint, cluster.masterAuth])
      .apply(([name, endpoint, masterAuth]) => {
        const context = `${gcp.config.project}_${gcp.config.zone}_${name}`;
        return `apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: ${masterAuth.clusterCaCertificate}
    server: https://${endpoint}
  name: ${context}
contexts:
- context:
    cluster: ${context}
    user: ${context}
  name: ${context}
current-context: ${context}
kind: Config
preferences: {}
users:
- name: ${context}
  user:
    auth-provider:
      config:
        cmd-args: config config-helper --format=json
        cmd-path: gcloud
        expiry-key: '{.credential.token_expiry}'
        token-key: '{.credential.access_token}'
      name: gcp
`;
      });

    // set gcp-infra output
    return {
      clusterName: cluster.name,
      kubeconfig,
    };
  }
}
