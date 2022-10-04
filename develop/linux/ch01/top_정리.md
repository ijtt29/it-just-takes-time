# top 명령어에 대해서

top의 내용들을 알아보자

uptime, loadavg, task

## uptime (현재 시간과 컴퓨터를 커놓은 시간)

uptime 명령어 확인
`whatis uptime`

`uptime -p`
> 시간만 확인할 수 있음


## load average (시스템 부하상태)
> 시스템이 얼마만큼의 부하가 있나? 를 확인할 수 있다.

1분 5분 15분 시스템 부하평균값을 나타낸다(EMA)
> EMA: 지수이동 평군값

프로세스의 *R* *D* 상태에 따라서 부하정도 측정한다.
* *R* 실행중인 프로세스
* *D* 디스크 I/O 처리완료를 기다리는 프로세스

컴퓨터의 코어 갯수 대비 계산을 하여 확인한다.
> `nproc`
> 현재 VM의 코어수 조회


/proc/loadavg 다음 경로에서도 확인할 수 있다.
```
cat /proc/loadavg

result: 0.00 0.00 0.00 3/120 28814
# 1분 시스템 부하평균값
# 5분 시스템 부하평균값
# 15분 시스템 부하평균값
# (실행중인프로세스) / (전체프로세스)
# 마지막으로 생성된 프로세스 PID

위의 내용은
man proc 로 명령어에 대한 설명을 확인 한다음 proc/loadavg에 나와있다.
```


## task 상태별 정보

total: 전체 프로세스

sleeping
* I와 S 상태









load average : 시스템의 부하정도 확인가능

nproc : 현재 VM의 CPU 코어갯수 확인