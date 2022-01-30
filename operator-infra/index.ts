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
import * as gcp from "@pulumi/gcp";

const pulumiConfig = new pulumi.Config();

// get pulumi config variables
const initialNodeCount = pulumiConfig.getNumber("initialNodeCount") ?? 1;
const machineType = pulumiConfig.get("machineType") ?? "n1-standard-1";
const k8sClusterName = pulumiConfig.get("clusterName") ?? "operator-cluster";
const gcpProjectName = pulumiConfig.require("project");
const gcpDefaultZone = pulumiConfig.require("zone");

// get latest gke master version
const engineVersion = gcp.container
    .getEngineVersions({ project: gcpProjectName, location: gcpDefaultZone })
    .then((v) => v.latestMasterVersion);

// Create a GKE cluster
const cluster = new gcp.container.Cluster(k8sClusterName, {
    initialNodeCount,
    minMasterVersion: engineVersion,
    nodeVersion: engineVersion,
    location: gcpDefaultZone,
    nodeConfig: {
        machineType,
        oauthScopes: [
            "https://www.googleapis.com/auth/compute",
            "https://www.googleapis.com/auth/devstorage.read_only",
            "https://www.googleapis.com/auth/logging.write",
            "https://www.googleapis.com/auth/monitoring",
        ],
    },
});

// Export some variables which can be used to connect to the cluster

// eslint-disable-next-line import/prefer-default-export
export const clusterName: pulumi.Output<string> = cluster.name;

// eslint-disable-next-line import/prefer-default-export
export const clusterZone: pulumi.Output<string> = cluster.location;
