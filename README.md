```sh
kubectl run -i --tty ubuntu --image=ubuntu --restart=Never
# kubectl exec -ti $(kubectl get pods | grep ubuntu | cut -d" " -f1) bash
apt update && apt install curl mysql-client
curl -LO https://github.com/josephspurrier/gowebapp/raw/master/config/mysql.sql
mysql -u root -p -h tomcat-mysql < mysql.sql

kubectl get svc tomcat -o template --template='{{range .spec.ports}}{{.nodePort}}{{end}}' && echo
```
