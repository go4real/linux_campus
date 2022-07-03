## [실습] systemd 서비스 구동 앱의 로그 출력하기

0. pip3 패키지 설치
```
sudo apt update
sudo apt install python3-pip
```

1. 파이썬 패키지 설치
 + uvicorn[standard], gunicorn
```
sudo pip3 install "uvicorn[standard]" gunicorn fastapi
```

2. fastapi 앱 파일 생성
 + mkdir ~/systemd_demo
 + vim /home/ubuntu/systemd_demo/main.py
 + 자동 줄바꿈 해제 :set paste 후 복사 붙여넣기
```
from fastapi import FastAPI
from datetime import datetime
import logging
from logging.handlers import SysLogHandler

log = logging.getLogger('demo')
log.setLevel(logging.DEBUG)
handler = SysLogHandler(address='/dev/log',
        facility=SysLogHandler.LOG_LOCAL7)
handler.setFormatter(logging.Formatter('hello: %(message)s'))
log.addHandler(handler)

app = FastAPI()


@app.get("/")
def read_root():
    now = datetime.now()
    current_time = now.strftime("%H:%M:%S")
    log.debug('[{}] Hello World'.format(current_time))
    return {"Hello": "World"}
```

3. system unit 파일 생성
 + sudo vim /etc/systemd/system/hello.service
```
[Unit]
Description=HelloWorld
After=network.target

[Service]
WorkingDirectory=/home/ubuntu/systemd_demo
ExecStart=gunicorn main:app --workers 4 --worker-class uvicorn.workers.UvicornWorker --bind 0.0.0.0:8000
Restart=on-failure
RestartSec=5s
SyslogIdentifier=gunicorn
SyslogFacility=local7

[Install]
WantedBy=multi-user.target
```

4. hello service 시작 및 동작 상태 확인
```
sudo systemctl daemon-reload
sudo systemctl restart hello
sudo systemctl status hello
```

5. 테스트 
```
curl localhost:8000
sudo journalctl -u hello -n 5
```

6. rsyslog 설정  (/var/log/에 파일로 저장)
 + sudo vim /etc/rsyslog.d/30-hello.conf
 ```
 local7.*             /var/log/hello.log
 ```

7. rsyslog 데몬 재시작
```
sudo systemctl restart rsyslog
```

8. 로그 이벤트 생성 및 로그 파일 저장 내용 확인
```
sudo systemctl restart hello
curl localhost:8000
cat /var/log/hello.log
```

9. 로그 로테이트 설정
+ sudo vim /etc/logrotate.d/hello
```
/var/log/hello.log {
  daily
  rotate 14
  missingok
  notifempty
  copytruncate
  compress
  delaycompress
}
```

## [실습정리] 자원삭제

1. 인스턴스 삭제
