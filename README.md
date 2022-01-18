# Whopper

!!! Project under active development !!!

Overview: ...

---

## Installation and project setup

Before creating the infrastructure and start processing articles make sure to walk over the following sections of this chapter.
- [Whopper](#whopper)
  - [Installation and project setup](#installation-and-project-setup)
    - [Required installations](#required-installations)
    - [Used environment variables](#used-environment-variables)
    - [Install local dependencies](#install-local-dependencies)
  - [Project lifecycle](#project-lifecycle)
    - [Bootstrapping](#bootstrapping)
    - [Operation Phase](#operation-phase)
    - [Cleanup](#cleanup)

After that go a head to the next chapter [Project lifecycle](#project-lifecycle).

### Required installations

-  `go` for application logic is used: [installation guide]()
-  `pulumi` as infrastructure as code (IaC) tool is used: [installation guide]()
-  `node` and `npm` is used to describe Pulumi IaC: [installation guide]()
-  `gcloud` CLI is used to access resources created in gcp: [installation guide]()
-  `kubectl` is used to interact with created kubernetes clusters: [installation guide]()
-  `task` is used instead of a makefile to install, verify, build and in general maintain the project: [installation guide]()
-  `protoc` is used to compile proto3 api file: [installation guide](https://grpc.io/docs/protoc-installation/)

You can check if installations are met by running `sh ./scripts/check-installations.sh` or `task check-installations`.

```
$ task -l
task: Available tasks for this project:
* bootstrap: 	Initialize project and create it resources
* cleanup: 	delete created it-resources
* install: 	Install dependencies and setup project
* k-op: 	kubectl wrapper to interact with operator cluster
* update: 	update project dependencies
* verify: 	Run all verify and test tasks
* verify-go: 	Run golang verify and test tasks
```

### Used environment variables

A couple of environment variables are **required**.

- Export you pulumi access token as environment variable `export PULUMI_ACCESS_TOKEN=XXXXXXX`.
- Make sure `GOPATH` is set; check via `echo $GOPATH`


### Install local dependencies

To install depenencies locally use the task file run `task install`
This will do the following this for you:
- Install go libraries
- Install npm libraries locally
- Clone project [googleapi]() which is used to build parts of the whopper api (see [api proto3 file](./api/whopper.proto))
- Compile proto3 golang code

---

## Project lifecycle

The lifecycle of this project is desirbed in more detail in one of the docs documents [see docs/project-lifecycle](./docs/project-lifecylce-phases.md).

![Project lifecycle overview](./assets/consider-cloud-native-ops.png)

A brief summary of each phase is given in the sections:
- [Bootstrapping](#bootstrapping)
- [Operation Phase](#operation-phase)
- [Cleanup](#cleanup)

### Bootstrapping
Setup infrastructure for the fist time.
There are two different types of infrastructure deployments. 

1. **Production deployment**: The production deployment is on going and uses the pulumi kubernetes operator. The operator watches the `main` branch for updates. Production infrastructure is build to stay with rolling blue-green updates.
2. **Feature test deployment**: The feature test deployment is a short term infrastructure test. It is used whenever a commit is pushed to any other branch then main. The infrastructure is deployed, tested and then destroyed. If these phases succeed without any error the feature is considered "bug free" and can be merged to `main`.

There are two tasks available adressing both deployment types:

```bash
task init-pulumi-project 
```

TODO: ...

### Operation Phase
operation

```bash
TODO: ...
```

To access kubernetes cluster 

### Cleanup
destroy created it resources

```bash
TODO: ...
```
