# PoC for OpenShift

- Clone this repo

```bash
cd goflow
git checkout enrich_k8s
oc create -f k8s
oc get svc goflow -ojsonpath='{.spec.clusterIP}'
oc edit networks.operator.openshift.io cluster
```

```yaml
apiVersion: operator.openshift.io/v1
kind: Network
metadata:
 name: cluster
spec:
  # ...
  exportNetworkFlows:
    netFlow:
      collectors:
        - <put service IP here>:2056
```

```bash
oc logs svc/goflow --follow
```
