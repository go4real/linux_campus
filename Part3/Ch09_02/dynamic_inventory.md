# 동적 인벤토리 구성
1. EC2 인스턴스 생성
+ EC2 인스턴스 2대 (Amazon Linux, Ubuntu)
+ 태그 webserver-01, dbserver-01

<br/>

2. boto3 및 ansible aws 플러그인 설치
```
sudo pip3 install boto3
ansible-galaxy collection install amazon.aws
```

<br/>

3. 인벤토리 파일 생성 (dynamic_aws_ec2.yaml)
```
plugin: aws_ec2
filters:
  instance-state-name: running
keyed_groups:
  # tag_Name_Value 형식의 그룹을 생성
  - prefix: tag
    key: tags
```

<br/>

4. 인벤토리 정보 확인
```
ansible-inventory -i dynamic_aws_ec2.yaml --graph
```

<br/>
