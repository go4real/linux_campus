## [실습] 새로운 볼륨 추가하기

1. 스토리지 볼륨(EBS) 생성
+ AWS 웹 콘솔에서 새로운 볼륨 추가
+ https://ap-northeast-2.console.aws.amazon.com/ec2/v2/home?region=ap-northeast-2#Volumes:

<br/>

2. 가상 머신에 생성한 볼륨 연결
+ AWS 웹 콘솔에서 ebs 볼륨 선택후 "볼륨 연결" 진행

<br/>

3. EC2 터미널에 접속해서 디바이스 상태 및 파일 시스템 존재 여부 확인
```
lsblk
sudo file -s /dev/xvdf
# /dev/xvdf: data          <- 파일시스템이 없는 경우
# sudo file -s /dev/xvda1  <- 파일시스템이 있는 경우 
```

<br/>

4. 파일 시스템 생성 및 마운트 
```
sudo mkfs -t ext4 /dev/xvdf
sudo file -s /dev/xvdf

sudo mkdir /data
sudo mount /dev/xvdf /data
```

<br/>

5. 파일시스템 테이블 수정
```
sudo cp /etc/fstab /etc/fstab.orig
sudo blkid
sudo vim /etc/fstab
# 예, UUID=374b00fb-03ed-4b4b-9f67-1b140daa7de1  /data  ext4  defaults,discard,nofail  0  2
```

6. 테스트
```
sudo umount /data
sudo mount -a
```
