# Containerd 커멘드 정리


Containerd 네임스페이스 목록 조회
```
ctr ns ls
```


특정 네임스페이스의 Container 목록 조회
```
ctr -n k8s.io c ls
```


스냅샷 조회
```
ctr -n k8s.io snapshot ls
```
> 스냅샷은 해당결로(var/lib/containerd/) 에 위치한다.
