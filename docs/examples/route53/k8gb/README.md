# Steps to test k8gb on KKP

1. Create 3 user clusters on different providers.(Or you can create 3 clusters on different region on AWS)
2. Add nginx application on those 3 clusters with default settings. You can add them while creating a new cluster as well.
3. Modify each cluster k8gb [helm value file](docs/examples/route53/k8gb/k8gb-cluster-eu-west-1.yaml) for keys `dnsZone`, `edgeDNSZone` and `hostedZoneID`. Only changing part of each file is `clusterGeoTag` and `extGslbClustersGeoTags`.

4. Deploy k8gb on each cluster using above helm value file.

```bash
   helm -n k8gb upgrade -i k8gb k8gb/k8gb --create-namespace -f < helm-value-file-with-path >
```

5. Create a secret containing AWS credentials for external DNS to work.

```bash
    export EXTERNALDNS_NS=k8gb
    kubectl create secret generic external-dns \
    --namespace ${EXTERNALDNS_NS:-"default"} --from-file docs/examples/route53/k8gb/credentials
```

6. Deploy a test app using make target.

```bash
    make deploy-test-apps
```

7. Modify faiover gslb resource by applying below files.

```bash
   k apply -f docs/examples/route53/k8gb/gslb-failover.yaml
```

8. Modify roundrobin gslb resource by applying below files.

```bash
    k apply -f docs/examples/route53/k8gb/gslb-roundrobin.yaml
```

9. Create weightRoundRobin resource by applying below files.

```bash
    k apply -f docs/examples/route53/k8gb/gslb-weighted-roundrobin.yaml
```


10. Here, we have deployed three gslb targets. Check test application availability by running below curl command. Replace failover.test.mihirs.xyz with the domain you specified in Gslb spec.

   - failover.test.mihirs.xyz
   - roundrobin.test.mihirs.xyz
   - weightroundrobin.test.mihirs.xyz

```bash
curl -s ffailover.test.mihirs.xyz | grep message
  "message": "eu-west-1",
```
Notice that traffic was routed to eu-west-1.

11. Emulate the failure in eu-west-1.

```bash
kubectl -n test-gslb scale deploy frontend-podinfo --replicas=0
```

12. Observe Gslb status change.

```bash
kubectl -n test-gslb get gslb test-gslb-failover -o yaml | grep status -A6
status:
  geoTag: us-east-1
  healthyRecords:
    failover.test.mihirs.xyz:
    - 35.168.91.100
  serviceHealth:
    failover.test.mihirs.xyz: Healthy
```

IP in healthyRecords should change to the IP address of NLB in us-east-1

13. Check failover to us-east-1

```bash
curl -s failover.test.k8gb.io| grep message
  "message": "us-east-1",
```

Notice that traffic is properly failed over to us-east-1

14. Test RoundRobin by below command. It saves output in output-1.txt

```bash
end=$((SECONDS+600))  # Run for 600 seconds (10 minutes)
while [ $SECONDS -lt $end ]; do
  echo "Request made at: $(date +"%Y-%m-%d %H:%M:%S")" >> output.txt
  curl -H 'Cache-Control: no-cache' -s -w "\nURL Effective: %{url_effective}\n" http://roundrobin.test.mihirs.xyz >> output-1.txt
  echo -e "\n----------\n" >> output.txt  # Add a separator between requests
  sleep 1  # Add delay between requests
done
```

15. Test weightRoundRobin by below command. It saves output in output-2.txt

```bash
end=$((SECONDS+600))  # Run for 600 seconds (10 minutes)
while [ $SECONDS -lt $end ]; do
  echo "Request made at: $(date +"%Y-%m-%d %H:%M:%S")" >> output.txt
  curl -H 'Cache-Control: no-cache' -s -w "\nURL Effective: %{url_effective}\n" http://weightroundrobin.test.mihirs.xyz >> output-2.txt
  echo -e "\n----------\n" >> output.txt  # Add a separator between requests
  sleep 1  # Add delay between requests
done
```
