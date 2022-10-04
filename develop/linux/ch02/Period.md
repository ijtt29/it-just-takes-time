# 커널 버전 변경하기
```
sudo apt install -y linux-image-5.4.0-65-generic
sudo reboot
```

특정 커널 버전으로 부팅하기
```
ls /boot/
vim /etc/default/grub
```

```
# /etc/default/grub
GRUB_DEFAULT='Advanced options for Ubuntu>Ubuntu, with Linux 5.4.0-65-generic'
```

```
sudo update-grub
sudo reboot
```

# 스케줄러 Period (분배할 CPU 총시간) 란?

*period: (분배할 총시간) 분패할 cpu실행시간단위 (ex: 100ms)*


예시

mysql, nginx, django, spring의 4가지 프로세스가 실행중일때 CPU를 어떻게 나눠사용할것인가?

각 프로세스들을 공평하게 CPU를 나눠 사용한다고 했을때 비율을 어떻게 정할것인가

보통 비율의 계산식은 *분자(할당시간)/분모(총시간) * 100* 인데 CPU의 총시간은 무한대이다.

따라서 이를 위해 period라는 개념이 나오게 되었다.
> ex: 총시간을 100ms로 정하여 나눈다.


확인방법

```
# 기본값 100ms
cat /sys/fs/cgroup/cpu/cpu.cfs_period_us
```

> 모든 프로세스는 기본 cgroup 설정에 적용이 되어있다.
> cgoup (HW 자원 독립 / 제한 기능)
CPU 기본 cgroup 설정에 포함되어있는 프로세스 PID 들 확인

```
cat /sys/fs/cgroup/cpu/tasks
```

유저 프로세스 기준 cgroup 에서 현재 프로세스 PID 확인
```
echo $$ : 현재 프로세스 PID 조회
cat /sys/fs/cgroup/cpu/user.slice/tasks | grep $$
```


### 할당시간 

타임슬라이스(timeslice) vs 런타임(runtime) 
> timeslice: 할당받은 실행'할' 시간
> runtime  : 할당받은 실행'한' 시간

runtime이 timeslice보다 클경우 CPU를 많이 점유하고 있다고 할 수 있다.

CPU사용을 뺏기는 상황은 인터럽트에 의해서도 발생할 수 있음.

보통은 timeslice를 전부 소진해서 뺏기면 정상이지만, 우선순위에 따라서 도중에 뺏길수도 있다.



## 만약 프로세스가 엄청 많다면? (100000개?)

그렇다면 Period의 단위가 줄어는다. (모든 프로세스가 CPU를 나눠사용한다면 timeslice가 작아질것이기 때문)

*sched_min_granularity*
프로세스의 숫자가 많아져서 프로세스당 timeslice가 극도로 작아지는 현상을 막기 위한 값.
``` 
ex: 2,250,000ms == 2.25ms
cat /proc/sys/kernel/sched_min_granularity_ns
```


```
cat /sys/fs/cgroup/cpu/cpu.cfs_period_us
# result: 100000
```
위의 경우 분배할 CPU의 총시간 period 기본값은 100ms 이기 때문에,
총 100ms 중에서 한 프로세스당 최소 2.25ms 는 보장해둔다고 해석할 수 있다.
> 프로세스가 무한히 많아도 보장한다.


## 런타임(runtime) 체크
현재 프로세스 PID 확인
```
echo $$
```

현재 프로세스(Bash) 기준 스케줄링 정보 확인하기
현재 프러세스의 실행시간 (runtime) 확인하기 *se.sum_exec_runtime*
```
cat /proc/$$/sched
```