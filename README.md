## Quick walkthrough

* `kubectl create -f local-volumes.yaml`
* `kubectl create secret generic mysql-pass --from-file=password.txt`
* Set proper hostname in nodeSelector inside `mysql-deployment.yaml`
* `kubectl create -f mysql-deployment.yaml`
* Use `kubectl get events` to resolve infinite `Pending` issues
* `kubectl create -f tomcat-deployment.yaml`
* `kubectl create -f nginx-deployment.yaml`
* Run `sudo mkdir -p /data1/etc/nginx` on worker k8s node which runs nginx pod
* Copy nginx template: `sudo cp mysite.template /data1/etc/nginx`
* `kubectl create -f ingress_rc.yaml`
* `kubectl create -f exposed_80_port.yaml`
* Visit [http://k8s-node-1/](http://k8s-node-1/) URL

Add MySQL scheme into MySQL:

```sh
kubectl run -i --tty ubuntu --image=ubuntu --restart=Never
# kubectl exec -ti $(kubectl get pods | grep ubuntu | cut -d" " -f1) bash
apt update && apt install -y curl mysql-client
curl -LO https://github.com/josephspurrier/gowebapp/raw/master/config/mysql.sql
mysql -u root -p -h tomcat-mysql gowebapp < mysql.sql
```

kubectl get svc tomcat -o template --template='{{range .spec.ports}}{{.nodePort}}{{end}}' && echo

kubectl label --overwrite node %NODENAME% %LABEL%=%VALUE%

kubectl exec -ti kube-dns-v11-32r8f -c etcd --namespace=kube-system sh

etcdctl get /skydns/local/cluster/svc/default/frontend_mysql_1/4ae33895
{"host":"10.9.68.2","priority":10,"weight":10,"ttl":30,"targetstrip":0}
```

```
it is neseccary to add this string into ingress nginx.conf:

proxy_set_header Host $host;

$ kubectl exec -ti nginx-ingress-guml5 sh
$ vim /etc/nginx/nginx.conf

As a result config should look like:

location / {
  proxy_set_header Host $host;
  proxy_pass http://%something is here%;
}
```
