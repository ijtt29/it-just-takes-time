#  k8s API Server CA DNS 추가

개인 로컬 쿠버네티스는 컴퓨터의 VM에 설치되어 구동되어있다.

개인 공부 및 배포 목적이고 집 이외의 환경에서도 접속하여 작업하기 위하여 iptime의 DNS로 접속할 수 있도록 설정하고 싶어 api server의 설정을 수정하는 내용을 정리한다.


## 수정한 yaml
```
kubectl -n kube-system get configmap kubeadm-config -o jsonpath='{.data.ClusterConfiguration}' > kubeadm.yaml
```

## yaml 수정
```
apiServer:
  certSANs:           # 추가
  - glmimo.iptime.org # 추가
  extraArgs:
    authorization-mode: Node,RBAC
  timeoutForControlPlane: 4m0s
apiVersion: kubeadm.k8s.io/v1beta3
certificatesDir: /etc/kubernetes/pki
clusterName: kubernetes
controllerManager: {}
dns: {}
etcd:
  local:
    dataDir: /var/lib/etcd
imageRepository: k8s.gcr.io
kind: ClusterConfiguration
kubernetesVersion: v1.24.6
networking:
  dnsDomain: cluster.local
  podSubnet: 192.168.0.0/16
  serviceSubnet: 10.96.0.0/12
scheduler: {}

```


## master node에서 실행
apiserver.crt 백업
```
sudo mv /etc/kubernetes/pki/apiserver.{crt,key} ~
```

kubeadm을 이용해 새로운 인증서를 생성
```
sudo kubeadm init phase certs apiserver --config kubeadm.yaml
```

클러스터 설정 업데이트 
```
sudo kubeadm init phase upload-config kubeadm --config kubeadm.yaml
```