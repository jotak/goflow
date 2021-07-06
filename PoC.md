## OpenShift

- Pre-requisite: have a running OpenShift cluster with ovn-k, version 4.8 at least.

- Clone this repo

```bash
cd goflow
git checkout enrich_k8s
oc apply -f k8s/goflow.yaml
COLLECTOR_IP=`oc get svc goflow -ojsonpath='{.spec.clusterIP}'` && echo $COLLECTOR_IP
oc patch networks.operator.openshift.io cluster --type='json' -p "$(sed -e "s/COLLECTOR_IP/$COLLECTOR_IP/" k8s/net-cluster-patch.json)"
```

- Make sure patch was applied:

```bash
oc get networks.operator.openshift.io cluster -ojsonpath='{.spec.exportNetworkFlows}' && echo
```

- Check the logs:

```bash
oc logs svc/goflow --follow
```

## Upstream ovn-kubernetes

E.g. using KIND: https://github.com/ovn-org/ovn-kubernetes/blob/master/docs/kind.md

```bash
cd goflow
git checkout enrich_k8s
kubectl apply -f k8s/goflow.yaml
COLLECTOR_IP=`kubectl get svc goflow -ojsonpath='{.spec.clusterIP}'` && echo $COLLECTOR_IP
kubectl set env daemonset/ovnkube-node -c ovnkube-node -n ovn-kubernetes OVN_NETFLOW_TARGETS="$COLLECTOR_IP:2056"
```


## Next steps

- Deploy ES, feed ES (or Loki)
- Use informers to watch kube resources for enrichment
