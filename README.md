# Climate Whopper

Project under development

TODO: README...

## Idea:

...

## Getting started

Required installations: `go`, `node`, `pulumi`, `gcloud`

1. Install project dependencies

```bash
# TODO: ...
npm install
go get ./...
```

2. Login to google cloud via `gcloud` CLI

```bash
gcloud auth login
# if you don't have project yet, create one
gcloud config set project <YOUR_GCP_PROJECT_HERE>
gcloud auth application-default login
```

3. Configure project

### Configuration

The project can be configured infrastructure and application wise.

#### Infrastructure configuration

TODO: ...Infrastructure configuration

The infrastructure uses pulumi and therefore the provided pulumi configuration.

gcloud config set project <GCP-PROJECT-ID>

To get kubernetes credentials you can run

```bash
gcloud container clusters get-credentials <cluster> --zone <cluster zone> --project <gcp project>
```

#### Application configuration

TODO: ...Application configuration

## Architecture:

TODO: ...Architecture

Components

Crawler Controller:

Crawler Engine Containers & pod's:

-   Downloader
-   Parser
-   Translator
-   Analyzer

Data Analyzer:

-   Jupyter notebook pod

## Project structure

TODO: ...Project structure
