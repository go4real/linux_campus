## [실습] 앱 구동용 서비스 유저 만들기

1. Dockerfile 생성
 + vim Dockerfile
```
FROM ubuntu:20.04
CMD ["sleep", "infinity"]
```

<br/>

2. 컨테이너 이미지 생성
```
docker build --tag user_test .
```

<br/>

3. 컨테이너 실행 및 사용자 정보 확인
```
docker run --rm -d --name user_test user_test
docker exec -it user_test id
```

<br/>

4. 사용자 정보 추가
 + vim Dockerfile
```
FROM ubuntu:20.04
RUN groupadd -r app &&\
    useradd -r -g app -d /home/app -c "Docker image user" app
USER app
CMD ["sleep", "infinity"]
```

5. 컨테이너 이미지 생성
```
docker build --tag user_test .
```

6. 컨테이너 실행 및 사용자 정보 확인
```
docker stop user_test
docker run --rm -d --name user_test user_test
docker exec -it user_test id
```
