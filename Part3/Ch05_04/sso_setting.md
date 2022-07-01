## [실습] Bastion Host에 MFA 설정하기 

0. 데모용 EC2 인스턴스 생성
 + Ubuntu 20.04 

1. google authenticator 패키지 설치
```
sudo apt update
sudo apt install -y libpam-google-authenticator
```

2. MFA 앱 설치
 + https://chrome.google.com/webstore/category/extensions


3. google authenticator 초기화 
```
google-authenticator
```

4. PAM(Pluggable Authentication Module) 정보 백업
```
sudo cp /etc/pam.d/sshd /etc/pam.d/sshd.bak
```

5. google authenticator 관련 인증 정보 추가
 + sudo vim /etc/pam.d/sshd
```
# Standard Un*x authentication.
#@include common-auth   # 주석처리

# Standard Un*x password updating.
@include common-password
auth required pam_google_authenticator.so 
auth required pam_permit.so
```

6. sshd 설정 정보 백업
```
sudo cp /etc/ssh/sshd_config /etc/ssh/sshd_config.bak
```

7. sshd 설정 정보 변경
 + sudo vim /etc/ssh/sshd_config
```
# Change to yes to enable challenge-response passwords (beware issues with
# some PAM modules and threads)
ChallengeResponseAuthentication yes

AuthenticationMethods publickey,keyboard-interactive
```

8. ssh 서비스 재시작
```
sudo systemctl restart sshd.service
```

9. ssh 연결 후 MFA 정보 확인

10. 데모용 EC2 인스턴스 삭제
