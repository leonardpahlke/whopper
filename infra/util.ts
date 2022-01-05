import InfraConfig from "./config";

/**
 * Abstract class used for infrastructure definitions
 */
export default abstract class Infra {
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
    const idString = id === "" ? "" : `${separator}${id}`;
    return `${this.config.projectName}${separator}${this.config.env}${separator}${serviceName}${idString}`;
  }
}

