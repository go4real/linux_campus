## [실습] 새로운 ssh 접속 키 등록하기

0. 데모용 인스턴스 생성
 + EC2 생성 : ssh_demo

1. 새로운 키페어 생성 및 다운로드 
 + EC2 > 네트워크 및 보안 > 키 페어
 + 예, new-ec2-key
 ```
 cp ~/Downloads/new-ec2-key.pem ~/linux_campus && cd ~/linux_campus
 ```

2. 인스턴스 중지

3. 등록할 퍼블릭 키 생성 
```
chmod 400 ~/linux_campus/new-ec2-key.pem
ssh-keygen -y -f ~/linux_campus/new-ec2-key.pem
```

4. 인스턴스 중지 및 사용자 데이터 편집
 + 작업 > 인스턴스 설정 > 사용자 데이터 편집
 + PublicKeypair 부분은 3번의 ssh-keygen 실행 결과 전체를 입력
```
Content-Type: multipart/mixed; boundary="//"
MIME-Version: 1.0

--//
Content-Type: text/cloud-config; charset="us-ascii"
MIME-Version: 1.0
Content-Transfer-Encoding: 7bit
Content-Disposition: attachment; filename="cloud-config.txt"

#cloud-config
cloud_final_modules:
- [users-groups, once]
users:
  - name: ubuntu
    ssh_authorized_keys: 
    - PublicKeypair
```

5. 인스턴스 시작

6. 새로운 개인키로 서버 접근
```
ssh -i ~/linux_campus/new-ec2-key.pem ubuntu@IP_ADDRESS
```
7. 인스턴스 중지

8. 사용자 데이터 삭제

9. 인스턴스 시작


