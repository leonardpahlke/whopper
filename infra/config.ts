import * as pulumi from "@pulumi/pulumi";

// Infrastructure input
export type IInfraVars = {
  zone: string;
  region: string;
  initialNodeCount: number; // 2
  machineType: string; // "n1-standard-1",
};

/**
 * InfraConfig used to translate pulumi configuration and get injected into k8s infra set-ups
 */
export default class InfraConfig {
  projectName: string;

  env: string;

  vars: IInfraVars;

  constructor() {
    const pulumiConfig = new pulumi.Config();
    this.vars = pulumiConfig.requireObject<IInfraVars>("vars");

    this.projectName = pulumiConfig.get("name") ?? "whopper-infra";
    this.env = pulumiConfig.get("env") ?? "prd";
  }
}
