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

import * as k8s from "@pulumi/kubernetes";
import * as pulumi from "@pulumi/pulumi";
import Infra from "./abstract-infra";
import InfraConfig from "./config";

/**
 * Kubernetes infrastructure input
 */
export type IK8sInput = {
  kubeconfig: pulumi.Output<string>;
};

/**
 * Kubernetes infrastructure output
 */
export type IK8sOutput = {
  namespaceName: pulumi.Output<string>;
};

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
    const clusterProvider = new k8s.Provider(this.GetName("k8s", "provider"), {
      kubeconfig: this.in.kubeconfig,
    });

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
