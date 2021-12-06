import * as pulumi from "@pulumi/pulumi";

// Infrastructure input
export interface IInfraVars {
    zone: string;
    region: string;

    initialNodeCount: number; // 2
    machineType: string; // "n1-standard-1",
}

/**
 * InfraConfig used to translate pulumi configuration and get injected into k8s infra set-ups
 */
export class InfraConfig {
    projectName: string;
    env: string;
    vars: IInfraVars;

    constructor() {
        let pulumiConfig = new pulumi.Config();
        this.vars = pulumiConfig.requireObject<IInfraVars>("vars");

        this.projectName = pulumiConfig.get("name") ?? "whopper-infra";
        this.env = pulumiConfig.get("env") ?? "prd";
    }
}

/**
 * Abstract class used for infrastructure definitions
 */
export abstract class Infra {
    protected config: InfraConfig;
    constructor(config: InfraConfig) {
        this.config = config;
    }

    // create, used to define the infrastructure with pulumi libraries
    abstract create(): void;

    /**
     * GetName is a method which is getting used across the project to create uniform construct names
     * @param serviceName Name of the service or something similar to quickly identify the resource
     * @param id Optional identifier if multiple resources with the same serviceName are being created (default '""')
     * @param separator Optional separator that is used to split the name (default '-')
     * @returns A string which is getting used as resource name / identifier
     */
    GetName(serviceName: string, id = "", separator = "-"): string {
        const idString = id == "" ? "" : `${separator}${id}`;
        return `${this.config.projectName}${separator}${this.config.env}${separator}${serviceName}${idString}`;
    }
}
