# ping-pong

## Exercise 1.09

    pasiol@lab:~$ kubectl apply -f https://raw.githubusercontent.com/pasiol/ping-pong/main/manifests/deployment.yaml
    deployment.apps/ping-pong created
    pasiol@lab:~$ kubectl get pods
    NAME                          READY   STATUS    RESTARTS   AGE
    log-output-5ff9857984-n7tbg   1/1     Running   0          170m
    ping-pong-5958c444d8-zqjcm    1/1     Running   0          15s
    pasiol@lab:~$ kubectl logs ping-pong-5958c444d8-zqjcm
    2021/11/06 16:19:29 go-pingopong starting in port 0.0.0.0:8888.
    pasiol@lab:~$ kubectl apply -f https://raw.githubusercontent.com/pasiol/ping-pong/main/manifests/service.yaml
    service/ping-pong-svc created
    pasiol@lab:~$ kubectl get svc
    NAME             TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
    kubernetes       ClusterIP   10.43.0.1       <none>        443/TCP          34h
    log-output-svc   NodePort    10.43.253.246   <none>        1234:30080/TCP   171m
    ping-pong-svc    ClusterIP   10.43.237.184   <none>        8888/TCP         18s
    pasiol@lab:~$ kubectl apply -f https://raw.githubusercontent.com/pasiol/ping-pong/main/manifests/ingress.yaml
    ingress.networking.k8s.io/ping-pong-ingress created
    pasiol@lab:~$ kubectl get ing
    NAME                CLASS    HOSTS   ADDRESS                            PORTS   AGE
    ping-pong-ingress   <none>   *       172.19.0.2,172.19.0.3,172.19.0.4   80      6s
    pasiol@lab:~$ curl http://172.19.0.2/pingpong
    Ping / Pongs: 2
    pasiol@lab:~$ kubectl logs ping-pong-5958c444d8-zqjcm
    2021/11/06 16:19:29 go-pingopong starting in port 0.0.0.0:8888.
    2021/11/06 16:23:21 Ping / Pongs: 1
    2021/11/06 16:24:52 Ping / Pongs: 2