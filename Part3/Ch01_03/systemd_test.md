# systemd 실습 
1. 실행중인 서비스 목록 확인
```
sudo systemctl list-units -t service
```

<br/>

2. cron 서비스 상태 확인 
```
sudo systemctl status cron.service
vim /lib/systemd/system/cron.service 
```

<br/>

3. cron 프로세스 강제 종료 및 재시작 확인
```
ps aux | grep cron
sudo kill -9 CRON_PROCESS_ID
ps aux | grep cron
```

<br/>

4. cron 서비스 중단 
```
sudo systemctl stop cron.service
ps aux | grep cron
```

5. cron 서비스 재시작
```
sudo systemctl restart cron.service
ps aux | grep cron
```
