# 실행환경 구성
1. AWS Copilot cli 설치
+ https://aws.github.io/copilot-cli/docs/getting-started/install/
```
curl -Lo copilot https://github.com/aws/copilot-cli/releases/latest/download/copilot-linux && chmod +x copilot && sudo mv copilot /usr/local/bin/copilot && copilot --help
```
+ copilot 자동 완성 기능 설정
```
source <(copilot completion bash)
copilot completion bash > copilot.sh
sudo mv copilot.sh /etc/bash_completion.d/copilot
```

2. AWS EC2에 적용할 IAM role 생성 및 적용

3. AWS CLI 설치
+ https://docs.aws.amazon.com/ko_kr/cli/latest/userguide/getting-started-install.html
```
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install
```

# 4. AWS CLI 설정  (ap-northeast-2)
```
aws configure
```

5. Docker Desktop
+ https://docs.docker.com/engine/install/ubuntu/
+ https://docs.docker.com/engine/install/linux-postinstall/
```
sudo groupadd docker
sudo usermod -aG docker $USER
newgrp docker   # vs remote server 재시작 또는 sudo reboot
```

6. Git에서 새로 코드를 받으시는 분들
```
git clone https://github.com/go4real/Django-Poll-App.git
git checkout ecs-base
```

