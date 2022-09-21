# 리눅스 교육 정리

## OS 가 하는일

1. 어플리케이션 관리
2. 하드웨어 자원 관리
    * CPU
    * 메모리
    * 디스크
    * 네트워크


## OS의 구성

Core 부분
* PM (Process Management)
* MM (Memory Management)
* irq / exception 처리, locking

인터럽트(irq): 하드웨어와 CPU와의 관계에서 발생할 수 있는것 (전기적 신호)
exception: 페이지 폴트/ 시스템 콜

> 문제 해결의 시작은 언제 어떻게 돌아가는지?

I/O 처리
* 네트워크 (L4:TCP, L3:IP, L2:DD)
* 스토리지 (VFS/FS/Block) VFS제일중요
* 디바이스 드라이버

기타 (핵심기능은 아님)
* security
* tools
* sounds

## CPU가 실행하는 함수들
### 1. 유저 함수
### 2. 라이브러리 함수 
### 3. 커널 함수
Entry(갑자기 호출되는 시작 지점)
1. 예외처리: 시스템콜 (open, read, write), 페이지폴트(물리메모리 가상메모리 매핑안됬을때 실행) 등
2. 인터럽트: 
    * 네트워크 패킷도착
    * USB 연결
    * ssd 디스크 I/O 끝난상황

3. 커널태스크 (커널스레드)
> 후반부로(미룬) 작업
    * ksoftirqd
    * kworker