## [실습] 확장 가능한 디스크 볼륨을 구성 - 파트3

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
