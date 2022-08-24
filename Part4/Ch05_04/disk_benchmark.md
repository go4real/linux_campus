## [실습] 디스크 벤치마크 도구로 성능 측정

0. 새로운 EC2 생성
 + Ubuntu20.04, t3.small
 + 추가 EBS 볼륨 생성 (gp3, 100GB, 3000IOPS, 125Throughput)

1. fio 도구 설치 
```
sudo apt update
sudo apt install -y fio
```

2. 연속 쓰기 
 + 10개의 job을 생성해서, 1M 크기로 연속 쓰기 진행
 ```
 sudo fio --name=write_test \
 --filename=/dev/xvdb --filesize=100G \
 --time_based --ramp_time=2s --runtime=1m \
 --ioengine=libaio --direct=1 --verify=0 --randrepeat=0 \
 --bs=1M --iodepth=64 --rw=write --numjobs=10 --offset_increment=10G
 ```

3. 결과 확인
 + EBS Throughput 설정: 125

4. 랜덤 읽기 
 + 4K 크기로 램덤 읽기 수행
```
sudo fio --name=read_test \
--filename=/dev/xvdb --filesize=100G \
--time_based --ramp_time=2s --runtime=1m \
--ioengine=libaio --direct=1 --verify=0 --randrepeat=0 \
--bs=4K --iodepth=256 --rw=randread
```

5. 결과 확인
 + EBS IOPS 설정: 3000


