# Crossplane AzureDevops Provider

<p float="left">
  <img src="images/crossplane.png" width="300" />
  <img src="images/azdevops.png" width="300" />
</p>


<br/>
<br/>

![CI](https://github.com/r3drun3/provider-azuredevops/actions/workflows/ci.yml/badge.svg)

<br/>

`provider-azuredevops` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/upbound/upjet) code
generation tools and exposes XRM-conformant managed resources for the
AzureDevops API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/r3drun3/provider-azuredevops):
```
up ctp provider install r3drun3/provider-azuredevops:v0.1.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-azuredevops
spec:
  package: r3drun3/provider-azuredevops:v0.1.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/r3drun3/provider-azuredevops).

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Local development and demo
clone this repo:  
```console
git clone https://github.com/R3DRUN3/provider-azuredevops.git && cd provider-azure
```

<br/>

Start a local `kubernetes` cluster (we will use `kind`):  
```console
kind create cluster --config ./k8s-local/kind-config.yml
```
<br/>

Apply Custom Resource Definitions:  
```console
kubectl apply -f package/crds
```

<br/>

Run the provider:  
```console
make run
```
<br/>

Under `examples/providerconfig` create a kubernetes secrets with your azuredevops org url and access token:  
```console
cat <<EOF > examples/providerconfig/secret.yaml
apiVersion: v1
kind: Secret
metadata:
  name: example-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "personal_access_token": "your azure devops personal access token",
      "org_service_url": "your azure devops org url"
    }
EOF
```

<br/>

Now open a new terminal and apply all the required resources:  
```console
kubectl create namespace crossplane-system --dry-run=client -o yaml | kubectl apply -f - \
&& kubectl apply -f examples/providerconfig/ \
&& kubectl apply -f examples/project/project.yaml
```

<br/>
This will create a new azure devops project in your organization:  
<div style="width: 60%; height: 60%">

  ![](images/az-devops-project.png)
</div>

<br/>

If you want to delete the created resource on azure, simply run:  
```console
kubectl delete -f examples/project/project.yaml
```

<br/>


## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/r3drun3/provider-azuredevops/issues).
