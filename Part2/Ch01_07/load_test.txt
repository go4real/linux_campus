## Prod용 환경 구성
1. container insights 활성화 포함
```
copilot env init --name prod --profile default --container-insights
```

2. 티어별 서비스 배포
```
# copilot deploy -e prod
copilot svc deploy --app poll-app --env prod --name poll-db
copilot svc deploy --app poll-app --env prod --name poll-backend
copilot svc deploy --app poll-app --env prod --name poll-frontend
```

3. ab툴 사용을 위한 패키지 설치
```
sudo apt install apache2-utils
```

4. 로드테스트 요청이 데이터베이스를 접근하도록 백엔드 코드 수정 
- main 브랜치를 사용하는분의 경우 다음 코드를 주석 처리
- polls/views.py
```
#@login_required()
def polls_list(request):
```


5. ab 테스트 실행
ab -n 5000 -c 25 http://{ELB_URL}/polls/list

6. AWS CloudWatch에서 모니터링된 값 확인 
