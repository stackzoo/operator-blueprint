#  OPERATOR BLUEPRINT

[![Go Report Card](https://goreportcard.com/badge/github.com/stackzoo/operator-blueprint)](https://goreportcard.com/report/github.com/stackzoo/operator-blueprint) 
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

This repo contains an example of a custom *kubernetes* operator, made with <a href="https://github.com/kubernetes-sigs/kubebuilder">kubebuilder</a>.

## References
> NOTE: Before delving into the programming of kubernetes operators, a thorough knowledge of the functioning of *k8s* is recommended.
<br/>

Useful references:
- kubernetes official <a href="https://kubernetes.io/docs/concepts/extend-kubernetes/operator/">docs</a>
- kubebuilder official <a href="https://book.kubebuilder.io/">docs</a>
- O'REILLY <a href="https://www.oreilly.com/library/view/programming-kubernetes/9781492047094/">programming kubernetes</a>

## Prerequisites
- `make`
- `kubectl`
- `docker`
- `go`
- `kubebuilder`

## Getting Started
You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster.
<br/>
**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).

<br/>

The *Makefile* already contains all the functions that you need.
<br/>
to see all make targets run:
```sh
make help
```

### Running on the cluster
1. Install Instances of Custom Resources:

```sh
kubectl apply -f config/samples/
```

2. Build and push your image to the location specified by `IMG`:

```sh
make docker-build docker-push IMG=<some-registry>/operator-blueprint:tag
```

3. Deploy the controller to the cluster with the image specified by `IMG`:

```sh
make deploy IMG=<some-registry>/operator-blueprint:tag
```

### Uninstall CRDs
To delete the CRDs from the cluster:

```sh
make uninstall
```

### Undeploy controller
UnDeploy the controller from the cluster:

```sh
make undeploy
```

## How it works
This project aims to follow the Kubernetes [Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/).
<br/>
It uses [Controllers](https://kubernetes.io/docs/concepts/architecture/controller/),
which provide a reconcile function responsible for synchronizing resources until the desired state is reached on the cluster.

<br/>
This is a very simple and vanilla operator, basically it does nothing more than delete all the pods in the namespace specified in the CRD.
<br/>
Let's take a look at the example manifest inside config/samples:

```yaml
apiVersion: examples.stackzoo.io/v1alpha1
kind: PodBuster
metadata:
  labels:
    app.kubernetes.io/name: podbuster
    app.kubernetes.io/instance: podbuster-sample
    app.kubernetes.io/part-of: operator-blueprint
    app.kubernetes.io/managed-by: kustomize
    app.kubernetes.io/created-by: operator-blueprint
  name: podbuster-sample
spec:
  namespace: test
```

When you apply this manifest to the cluster (after deploying the operator) the custom controller will delete every pods inside the *test* namespace.

### Test It Out

1. Start a local kind cluster:
```sh
make kind-up
```



2. Install the CRDs into the cluster:
```sh
make install
```

3. Prepare resources on the cluster:
```sh
make prepare-resources
```

4. Run your controller (this will run in the foreground, so switch to a new terminal if you want to leave it running):
```sh
make run
```

5. Apply the CRD manifest:
```sh
kubectl apply -f config/samples/examples_v1alpha1_podbuster.yaml
```

6. Check the operator logs in the first terminal:
```sh
2023-03-20T14:04:09+01:00       INFO    Operator Blueprint      {"controller": "podbuster", "controllerGroup": "examples.stackzoo.io", "controllerKind": "PodBuster", "PodBuster": {"name":"podbuster-sample","namespace":"default"}, "namespace": "default", "name": "podbuster-sample", "reconcileID": "3e93bdbf-1eed-47e1-92db-5ad9786f90a2", "Deleting pod": "busybox"}
2023-03-20T14:04:09+01:00       INFO    Operator Blueprint      {"controller": "podbuster", "controllerGroup": "examples.stackzoo.io", "controllerKind": "PodBuster", "PodBuster": {"name":"podbuster-sample","namespace":"default"}, "namespace": "default", "name": "podbuster-sample", "reconcileID": "3e93bdbf-1eed-47e1-92db-5ad9786f90a2", "Deleting pod": "nginx"}
```

7. When you are done, stop the local kind cluster run:
```sh
make kind-down
```

### Modifying the API definitions
If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```sh
make manifests
```

## License

Copyright 2023 stackzoo.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

