[![codecov](https://codecov.io/gh/epam/edp-component-operator/branch/master/graph/badge.svg?token=88JVA1GA73)](https://codecov.io/gh/epam/edp-component-operator)

# EDP Component Operator

| :heavy_exclamation_mark: Please refer to [EDP documentation](https://epam.github.io/edp-install/) to get the notion of the main concepts and guidelines. |
| --- |

Get acquainted with the EDP Component Operator and the installation process as well as the local development.

## Overview

EDP Component Operator is an extension for K8S/OpenShift that is used for managing the EDP CI/CD components in the Admin Console tool. At the moment, it is delivered as a Custom Resource Definition without any controllers. Some of EDP components are predefined and installed automatically.

_**NOTE:** Operator is platform-independent, that is why there is a unified instruction for deploying._

Additionally, there is the ability to use the custom components and afterwards consume the respective component data from Admin Console.

The basic Custom Resource for EDP Component can be as follows ([edp_components_crd.yaml](deploy-templates/crds/v1.edp.epam.com_edpcomponents.yaml)):

```yaml
apiVersion: v1.edp.epam.com/v1
kind: EDPComponent
metadata:
  name: example-edpcomponent
spec:
  type: example-edpcomponent
  url: https://example-jenkins
  icon: base64encoded_icon
  visible: true   # ensures whether the component is visible on the Admin Console Overview page or not

```

For example: [docker-registry.yaml](documentation/docker-registry.yaml), [kubernetes.yaml](documentation/kubernetes.yaml).

_**NOTE:** Pay attention to the icon field that should be a base64 encoded._

As soon as the CR is created, it will appear on the main page of the Admin Console tool:

![admin-console](readme-resource/admin_console_main.png "admin-console")

## Prerequisites

1. Linux machine or Windows Subsystem for Linux instance with [Helm 3](https://helm.sh/docs/intro/install/) installed;
2. Cluster admin access to the cluster;
3. EDP project/namespace is deployed by following the [Install EDP](https://epam.github.io/edp-install/operator-guide/install-edp/) instruction.

## Installation
In order to install the EDP Component Operator, follow the steps below:

1. To add the Helm EPAMEDP Charts for local client, run "helm repo add":
     ```bash
     helm repo add epamedp https://epam.github.io/edp-helm-charts/stable
     ```
2. Choose available Helm chart version:
     ```bash
     helm search repo epamedp/edp-component-operator -l
     NAME                            CHART VERSION   APP VERSION     DESCRIPTION
     epamedp/edp-component-operator  0.11.0          0.11.0          A Helm chart for EDP Component Operator
     epamedp/edp-component-operator  0.10.0          0.10.0          A Helm chart for EDP Component Operator
     ```

    _**NOTE:** It is highly recommended to use the latest released version._

3. Deploy operator:

Install operator in the <edp-project> namespace with the helm command; find below the installation command example:
```bash
helm install edp-component-operator epamedp/edp-component-operator --version <chart_version>
```

## Local Development

In order to develop the operator, first set up a local environment. For details, please refer to the [Local Development](https://epam.github.io/edp-install/developer-guide/local-development/) page.

Development versions are also available, please refer to the [snapshot helm chart repository](https://epam.github.io/edp-helm-charts/snapshot/) page.

### Related Articles

* [Install EDP](https://epam.github.io/edp-install/operator-guide/install-edp/)
