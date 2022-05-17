### 프로젝트 리포지토리: https://github.com/go4real/Django-Poll-App.git
<br/><br/>
1. ECS 앱 생성  __프로젝트 루트에서 실행__
```
copilot init
# Application name: poll-app

## poll-db 서비스 생성
# Workload type: Backend Service
# Service name: poll-db
# Dockerfile: Use an existing image instead
# Image: postgres
```

2. copilot 디렉토리 내용 확인
+ postgres 포트 설정 
+ copilot/poll-db/manifest.yml

```
...
image:
  location: postgres
  port: 5432    # 공개할 포트정보 추가
...

variables:                    # Pass environment variables as key value pairs.
  POSTGRES_DB: poll
  POSTGRES_USER: fast
  POSTGRES_PASSWORD: 1234qwer
...

```

3. dev 환경 생성
```
copilot env init 
# Environment name: dev
# Credential source: [profile default]
# 다른 항목은 기본설정 사용
```
+ 3분 이상 시간 소요  --> AWS Web Console에서 인프라 생성 내용 확인 (ECS, CloudFormation 등)

4. poll-db 서비스 배포 
```
copilot deploy
```
+ 웹 콘솔에서 서비스 상태 확인 


<br/>
