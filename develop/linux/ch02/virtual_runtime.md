# virtual runtime

virtual runtime을 쓰는 이유

스케쥴러가 실행하는 프로세스를 어떻게 선택할거냐


우선순위 vs 실행한 시간



## virtual runtime

> 실행한시간(runtime)이 가장 적은건 부터 선택한다.


vruntime += 실행한시간(runtime) * 평균가중치 / 나의가중치


