import * as pulumi from "@pulumi/pulumi";
import * as gcp from "@pulumi/gcp";

const pulumiConfig = new pulumi.Config();

// get pulumi config variables
const initialNodeCount = pulumiConfig.getNumber("initialNodeCount") ?? 1;
const machineType = pulumiConfig.get("machineType") ?? "n1-standard-1";
const clusterName = pulumiConfig.get("clusterName") ?? "operator-cluster";

// get latest gke master version
const engineVersion = gcp.container
    .getEngineVersions()
    .then((v) => v.latestMasterVersion);

    // Create a GKE cluster
const cluster = new gcp.container.Cluster(clusterName, {
    initialNodeCount,
    minMasterVersion: engineVersion,
    nodeVersion: engineVersion,
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

// Export the name of the created cluster
export default cluster.name;
