## [실습] 인스턴스 ID 기반 SSH 연결 스크립트 작성

1. 테스트에 사용할 인스턴스 생성
+ 타겟 호스트 1대
+ Ubuntu 20.04 이미지 사용
+ Bastion 호스트와 동일한 VPC 사용

2. Bastion 호스트에 접속
+ 인스턴스에 적절한 권한을 가진 ROLE 설정 확인 
+ AWS CLI 및 jq 패키지 설치 확인

3. ssh_by_id 파일 작성
+ vim ssh_by_id
```
#!/bin/bash
#
# Usage
# $ ./ssh_by_id i-12345

if [ $# -ne 1 ]; then
  echo "Usage"
  echo "$ ./ssh_by_id i-12345"
  exit 1
fi

instance_id=$1

# get EC2 data
availability_zone=$(aws ec2 describe-instances --instance-ids $instance_id | jq -r .Reservations[0].Instances[0].Placement.AvailabilityZone)
ip_address=$(aws ec2 describe-instances --instance-ids $instance_id | jq -r .Reservations[0].Instances[0].PrivateIpAddress)

# generate RSA key pair
tmpfile="/tmp/ssh.`echo $RANDOM | base64 | head -c 6`"
ssh-keygen -C "temp ssh key" -q -f $tmpfile -t rsa -b 2048 -N ""
public_key=${tmpfile}.pub
private_key=$tmpfile
chmod 400 $private_key

# register public key
aws ec2-instance-connect send-ssh-public-key \
  --instance-id  $instance_id \
  --instance-os-user ubuntu \
  --ssh-public-key file://$public_key \
  --availability-zone $availability_zone > /dev/null

# ssh into ec2 instance with private key
ssh -i $private_key ubuntu@$ip_address
```

4. ssh_by_id 파일에 실행 권한 부여 및 복사
```
chmod +x ssh_by_id
sudo cp ssh_by_id /usr/local/bin
```

5. ssh_by_id 테스트
```
ssh_by_id i-12345678
```
