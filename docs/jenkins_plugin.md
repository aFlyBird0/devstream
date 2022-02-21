## 1 jenkins Plugin

This plugin installs [jenkins](https://jenkins.io) in an existing Kubernetes cluster using the Helm chart.

## 2 Usage:

```yaml
tools:
# name of the instance with jenkins
- name: jenkins-dev
  plugin:
    # kind of the plugin
    kind: jenkins
    # version of the plugin
    version: 0.0.1
  # options for the plugin
  options:
    # need to create the namespace or not, default: false
    create_namespace: false
    # Helm repo information
    repo:
      # name of the Helm repo
      name: jenkins
      # url of the Helm repo
      url: https://charts.jenkins.io
    # Helm chart information
    chart:
      # name of the chart
      chart_name: jenkins/jenkins
      # release name of the chart
      release_name: dev
      # k8s namespace where jenkins will be installed
      namespace: jenkins
      # whether to wait for the release to be deployed or not
      wait: true
      # the time to wait for any individual Kubernetes operation (like Jobs for hooks). This defaults to 5m0s
      timeout: 5m
      # whether to perform a CRD upgrade during installation
      upgradeCRDs: true
```

Currently, all the parameters in the example above are mandatory.