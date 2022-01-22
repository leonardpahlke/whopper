// Copyright 2022 Leonard Vincent Simon Pahlke
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

