# systemd 실습 
1. 실습용 스크립트 파일 작성 
+ vim /usr/local/bin/my-uptime.sh
```
#!/bin/sh
echo "Uptime: `uptime`" >> /tmp/uptime.log
```

<br/>

2. systemd 서비스 생성 -> service 타입의 Unit파일 생성
+ vim /etc/systemd/system/my-uptime.service 
```
[Unit]
Description=systemd practice.

[Service]
ExecStart=/bin/sh /usr/local/bin/my-uptime.sh

[Install]
WantedBy=multi-user.target
```

<br/>

3. systemd 데몬 리로드 (systemd 목록 수정시 항상 실행)
```
systemctl daemon-reload
```
<br/>

4. 시작 서비스에 my_uptime.service 등록 
```
sudo systemctl enable my_uptime.service
```
<br/>

5. 시스템 재시작
```
sudo reboot
```
<br/>

6. 로그 정보 확인
```
cat /tmp/uptime.log
```
