# ping-pong

## Exercise 1.09

[https://github.com/pasiol/ping-pong/tree/1.09]

    kubectl delete -f https://raw.githubusercontent.com/pasiol/todo-project/1.08/manifests/ingress.yaml
    ingress.networking.k8s.io "todo-project-ingress" deleted
    kubectl apply -f https://raw.githubusercontent.com/pasiol/ping-pong/main/manifests/deployment.yaml
    deployment.apps/ping-pong created

    kubectl get pods
    NAME                            READY   STATUS              RESTARTS   AGE
    todo-project-86bd654c5c-drr9p   1/1     Running             0          62m
    log-output-6897c6f44-q9zfw      1/1     Running             0          38m
    ping-pong-5d99cfc6cb-tp5nx      0/1     ContainerCreating   0          12s

    kubectl logs ping-pong-5d99cfc6cb-tp5nx
    2021/11/15 19:00:17 pingopong starting in port 0.0.0.0:8888.

    kubectl apply -f https://raw.githubusercontent.com/pasiol/ping-pong/main/manifests/service.yaml
    service/ping-pong-svc created

    kubectl get svc
    NAME               TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
    kubernetes         ClusterIP   10.43.0.1       <none>        443/TCP    70m
    todo-project-svc   ClusterIP   10.43.151.224   <none>        8000/TCP   53m
    ping-pong-svc      ClusterIP   10.43.211.238   <none>        8888/TCP   8s

    kubectl apply -f https://raw.githubusercontent.com/pasiol/ping-pong/main/manifests/ingress.yaml
    ingress.networking.k8s.io/ping-pong-ingress created

    kubectl get ing
    NAME                CLASS    HOSTS   ADDRESS                            PORTS   AGE
    ping-pong-ingress   <none>   *       172.18.0.2,172.18.0.3,172.18.0.4   80      8s

    curl http://172.18.0.2/pingpong
    Ping / Pongs: 1

    kubectl logs ping-pong-5d99cfc6cb-tp5nx
    2021/11/15 19:00:17 pingopong starting in port 0.0.0.0:8888.
    2021/11/15 19:04:02 Ping / Pongs: 1