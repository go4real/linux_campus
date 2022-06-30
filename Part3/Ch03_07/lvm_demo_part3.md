## [실습] 확장 가능한 디스크 볼륨을 구성 - 파트3

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
sudo file -s /dev/xvdg
# /dev/xvdg: data          <- 파일시스템이 없는 경우
```

<br/>

4. 물리 볼륨을 생성 및 확인
```
sudo pvcreate /dev/xvdg
sudo pvs
```

<br/>

5. 볼륨 그룹 확장
```
sudo vgextend Data /dev/xvdg
sudo vgs
```

6. 논리 볼륨 확장
```
sudo lvextend -L +29G /dev/Data/data1
sudo resize2fs /dev/Data/data1
```
1. 파일시스템 테이블 항목 삭제 및 마운트 해제
```
vim /etc/fstab
sudo umount /data1
```

<br/>

2. 논리 볼륨 삭제
```
sudo lvremove /dev/Data/data1
```

<br/>

3. 볼륨 그룹 삭제
```
sudo vgremove Data
```

<br/>

4. 물리 볼륨  삭제
```
sudo pvremove /dev/xvdf
sudo pvremove /dev/xvdg
```

<br/>

5. AWS 웹 콘솔에서 EBS 볼륨 연결 해제 및 삭제
