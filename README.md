# ping-pong

Exercise 1.11

[https://github.com/pasiol/log-output/tree/1.11]
[https://github.com/pasiol/ping-pong/tree/1.11]

    docker exec k3d-k3s-default-agent-0 mkdir -p /tmp/kube
    pasiol@lab:~$ kubectl apply -f https://raw.githubusercontent.com/pasiol/log-output/1.11/manifests/persistentVolume.yaml
    persistentvolume/log-output-pv created
    pasiol@lab:~$ kubectl apply -f https://raw.githubusercontent.com/pasiol/log-output/1.11/manifests/persistentVolumeClaim.yaml
    persistentvolumeclaim/log-output-claim created
    pasiol@lab:~$ kubectl apply -f https://raw.githubusercontent.com/pasiol/log-output/1.11/manifests/deployment.yaml
    deployment.apps/log-output-dep created
    pasiol@lab:~$ kubectl apply -f https://raw.githubusercontent.com/pasiol/log-output/1.11/manifests/service.yaml
    service/log-output-svc created
    pasiol@lab:~$ kubectl apply -f https://raw.githubusercontent.com/pasiol/ping-pong/1.11/manifests/deployment.yaml
    deployment.apps/ping-pong created
    pasiol@lab:~$ kubectl apply -f https://raw.githubusercontent.com/pasiol/ping-pong/1.11/manifests/service.yaml
    service/ping-pong-svc created
    pasiol@lab:~$ kubectl apply -f https://raw.githubusercontent.com/pasiol/log-output/1.11/manifests/ingress.yaml
    ingress.networking.k8s.io/log-output-ingress created
    pasiol@lab:~$ kubectl get ing
    NAME                 CLASS    HOSTS   ADDRESS                            PORTS   AGE
    log-output-ingress   <none>   *       172.19.0.2,172.19.0.4,172.19.0.5   80      49
    pasiol@lab:~$ curl http://172.19.0.2/pingpong
    Ping / Pongs: 1
    pasiol@lab:~$ curl http://172.19.0.2/pingpong
    Ping / Pongs: 2
    pasiol@lab:~$ curl http://172.19.0.2/
    2021-11-08T18:49:15.736664851Z fad9bfc5-6485-4f49-907a-0ed4ef4d53ef
    Ping / Pongs: 2