# Helm Chart

회사에서 Staging환경의 EKS 버전을 1.21 -> 1.22로 업그레이드 이후, 과거에 배포되어 관리가 잘되어오지 않던 Helm Chart에서 배포가 되지 않는 문제가 발생했다.

쿠버네티스 버전이 1.21 -> 1.22로 변경되면서 기존에 사용되던 beta버전읜 인그레스 API지원이 중단된것이 원인이되었고,
해결을 위해 Ingress버전은 바꾸어 helm upgrade를 진행했지만 배포가 되지 않았다.

## 원인
이유는 helm 으로 배포하게되면 배포 버전에 해당하는 secret이 생성되는데, 이 secret안에는 배포했을때의 메니페스트 정보가 존재한다.
* https://helm.sh/docs/topics/kubernetes_apis/#updating-api-versions-of-a-release-manifest

## 해결
따라서 secret안에 존재하는 메니페스트의 Ingress 버전또한 배포한 Ingress API버전에 맞게 변경해주어야 새로운 Ingress API 버전으로 배포를 할 수있게 되었다.

## 회고
helm으로 어플리케이션을 배포했을때 revision이 존재하지만, 어떻게 관리되는지에 대한 관심을 크게 가지지 않았던것 같다.
이번기회로 helm으로 배포된 어플리케이션들의 revision들이 어떻게 관리되는지 알게되어 다음에 이러한 문제가 발생했을때 대응이 빨라지지 않을까 기대해 본다.

