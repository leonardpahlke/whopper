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

    gcpProjectName: string;

    env: string;

    vars: IInfraVars;

    constructor() {
        const pulumiConfig = new pulumi.Config();
        this.vars = pulumiConfig.requireObject<IInfraVars>("vars");

        this.projectName = pulumiConfig.get("name") ?? "whopper";
        this.env = pulumiConfig.require("env");
        this.gcpProjectName = pulumiConfig.require("project");
    }
}
