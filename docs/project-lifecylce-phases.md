## Project lifecycle
The project is divided into three major lifecycle phases. 
1. During the bootstrapping phase, the project infrastructure and services are set up. The phase is completed when the system can be used normally for the first time.
2. In the second phase, the operational phase, the system is used after the initial setup. During this phase, the infrastructure and the system are used and managed (including updates).
3. The last phase, the clean-up phase, is initiated when the entire project system is deleted. The phase is complete when all resources initially deployed and after updates across the infrastructure systems in use have been deleted.

A general overview is provided below.

...
TODO: add project lifecycle overview

The phases are described in the next chapters with additional diagrams.

Translated with www.DeepL.com/Translator (free version)

## Bootstrapping Phase
```mermaid
%%{init: {
    'theme': 'base', 
    'themeVariables': {
        'primaryColor': '#053238',
        'secondaryColor': '#950EBA',
        'primaryTextColor': '#ffffff',
        'primaryBorderColor': '#ffffff',
        'secondaryBorderColor': '#ffffff',
        'fontFamily': 'arial black',
        'titleColor': '#053238',
        'background': '#ffffff',
        'clusterBkg': '#ffffff',
        'clusterBorder': '#6238FA'
        }
    }
}%%
stateDiagram-v2
    state fork_state <<fork>>
    [*] --> CloneRepository
    CloneRepository --> fork_state
    fork_state --> Go_Install
    fork_state --> Npm_Install

    state join_state <<join>>
    Npm_Install --> join_state
    Go_Install --> join_state

    join_state --> create_operator_k8s_cluster
    create_operator_k8s_cluster --> pull_kubeconfig
    pull_kubeconfig --> set_pulumi_config
    set_pulumi_config --> deploy_pulumi_kubernetes_operator
    deploy_pulumi_kubernetes_operator --> [*]
```

TODO: walk over the steps one by one

## Operation Phase

```mermaid
%%{init: {
    'theme': 'base', 
    'themeVariables': {
        'primaryColor': '#053238',
        'secondaryColor': '#950EBA',
        'primaryTextColor': '#ffffff',
        'primaryBorderColor': '#ffffff',
        'secondaryBorderColor': '#ffffff',
        'fontFamily': 'arial black',
        'titleColor': '#053238',
        'background': '#ffffff',
        'clusterBkg': '#ffffff',
        'clusterBorder': '#6238FA'
        }
    }
}%%
flowchart TD
    start[START] -->|`git checkout main`| branch
    subgraph LOCAL MACHINE
        branch[Branch off main 'FeatureX Branch'] -->|`git checkout -b featurex`| update[Make changes];
        update -->|`git commit -a -S -m 'update msg'`| push[Push changes];
    end

    %% deploy -.- cluster[Kubernetes Cluster]
    %% more_tests -.- cluster

    subgraph GITHUB ACTIONS
        push -->|`git push --set-upstream featurex`| test[Verify Code with\n Unit Tests and Linters];
        test -->|`task verify test-unit`| test_ok{Ok?}
        test_ok -->|Yes| deploy[Deploy pulumi 'FeatureX' stack]
        test_ok -->|No| update
        deploy -->|`pulumi up -s featurex`| more_tests[Run integration and smoke tests]
        more_tests -->|`task test-integration test-smoke`| more_tests_ok{Ok?}
        more_tests_ok -->|No| update
        more_tests_ok -->|Yes| merge_ok{Merge\n conflict or\n commit?}
        merge_ok -->|Yes| update
    end

    subgraph GITHUB
        merge_ok -->|No| approve{Approve\n changes?}
        approve -->|Yes| merge[Merge]
        approve -->|No| update
        merge -->|`git merge featurex`| delete_pulumi[Delete pulumi\n resources and stack]
        delete_pulumi -->|`pulumi destroy -s featurex`| cleanup[Delete feature branch\n and close PR]
    end
    cleanup --> finish[END]
```

TODO: walk over the steps one by one

## Cleanup

```mermaid
%%{init: {
    'theme': 'base', 
    'themeVariables': {
        'primaryColor': '#053238',
        'secondaryColor': '#950EBA',
        'primaryTextColor': '#ffffff',
        'primaryBorderColor': '#ffffff',
        'secondaryBorderColor': '#ffffff',
        'fontFamily': 'arial black',
        'titleColor': '#053238',
        'background': '#ffffff',
        'clusterBkg': '#ffffff',
        'clusterBorder': '#6238FA'
        }
    }
}%%
stateDiagram-v2
    state fork_state <<fork>>
    [*] --> fork_state
    fork_state --> DestroyPulumiStack
    fork_state --> DisableGithubAutomations 
    DestroyPulumiStack --> DestroyK8sCluster
    
    state join_state <<join>>
    DestroyK8sCluster --> join_state
    DisableGithubAutomations --> join_state
    join_state --> [*]
```

TODO: walk over the steps one by one