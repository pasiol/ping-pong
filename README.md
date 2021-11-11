# ping-pong

Exercise 2.01

- test locally

    APP_OUTPUT_FILE="./pingpong.txt" APP_LOG_FILE="./pingpong.log" APP_PORT=8000 go run main.go

[https://github.com/pasiol/log-output/tree/2.01]
[https://github.com/pasiol/ping-pong/tree/2.01]

    pasiol@lab:~$ kubectl apply -f https://raw.githubusercontent.com/pasiol/ping-pong/2.01/manifests/deploymentBusyBox.yaml
    pod/busybox1 created
    pasiol@lab:~$ kubectl apply -f https://raw.githubusercontent.com/pasiol/ping-pong/2.01/manifests/deployment.yaml
    deployment.apps/ping-pong created
    pasiol@lab:~$ kubectl apply -f https://raw.githubusercontent.com/pasiol/ping-pong/2.01/manifests/service.yaml
    service/ping-pong-svc created
    pasiol@lab:~$ kubectl apply -f https://raw.githubusercontent.com/pasiol/ping-pong/2.01/manifests/ingress.yaml
    ingress.networking.k8s.io/ping-pong-ingress created
    pasiol@lab:~$ kubectl get ing
    NAME                CLASS    HOSTS   ADDRESS                            PORTS   AGE
    ping-pong-ingress   <none>   *       172.19.0.2,172.19.0.4,172.19.0.5   80      55s

    NAME                         READY   STATUS    RESTARTS   AGE
    busybox1                     1/1     Running   0          25m
    ping-pong-5d99cfc6cb-m9v9l   1/1     Running   0          24m
    pasiol@lab:~$ kubectl exec -it ping-pong-5d99cfc6cb-m9v9l -- curl http://172.19.0.2/pingpong
    Ping / Pongs: 20

    pasiol@lab:~$ kubectl apply -f https://raw.githubusercontent.com/pasiol/log-output/2.01/manifests/persistentVolume.yaml
    persistentvolume/log-output-pv created
    pasiol@lab:~$ kubectl apply -f https://raw.githubusercontent.com/pasiol/log-output/2.01/manifests/persistentVolumeClaim.yaml
    persistentvolumeclaim/log-output-claim created
    pasiol@lab:~$ kubectl apply -f https://raw.githubusercontent.com/pasiol/log-output/2.01/manifests/deployment.yaml
    deployment.apps/log-output-dep created
    pasiol@lab:~$ kubectl apply -f https://raw.githubusercontent.com/pasiol/log-output/2.01/manifests/service.yaml
    service/log-output-svc created
    pasiol@lab:~$ kubectl apply -f https://raw.githubusercontent.com/pasiol/log-output/2.01/manifests/ingress.yaml
    ingress.networking.k8s.io/log-output-ingress created

    pasiol@lab:~$ kubectl get ing
    NAME                 CLASS    HOSTS   ADDRESS                            PORTS   AGE
    ping-pong-ingress    <none>   *       172.19.0.2,172.19.0.4,172.19.0.5   80      23m
    log-output-ingress   <none>   *       172.19.0.2,172.19.0.4,172.19.0.5   80      11m
    pasiol@lab:~$ kubectl exec -it busybox1 -- wget -qO - http://172.19.0.2
    2021-11-11T18:06:36.73809423Z 3e1f8bb6-562b-49d5-ab6b-896c5b109004
    Ping / Pongs: 21

    