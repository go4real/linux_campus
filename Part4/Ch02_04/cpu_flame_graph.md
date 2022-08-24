## [실습] CPU 성능 데이터 시각화

0. 새로운 EC2 생성
 + Ubuntu20.04, t2.micro

1. perf 도구 설치 
```
sudo apt update
sudo apt install -y linux-tools-common linux-tools-generic linux-tools-`uname -r`
```

2. Node.js v18.x 설치
```
curl -fsSL https://deb.nodesource.com/setup_18.x | sudo -E bash -
sudo apt-get install -y nodejs
node -v
```

3. 테스트 프로그램 작성
+ vim hello.js 
```
const http = require('node:http');
const crypto = require('crypto');

const hostname = '127.0.0.1';
const port = 3000;

const server = http.createServer((req, res) => {
    calc();
    res.statusCode = 200;
    res.setHeader('Content-Type', 'text/plain');
    res.end('Hello, World!\n');
});

function calc() {
    for(i=0; i<1000; i++)
        crypto.createCipher('aes-128-cbc', Math.random().toString(32));
}

server.listen(port, hostname, () => {
    console.log(`Server running at http://${hostname}:${port}/`);
});
```

4. tmux 실행
+ 창 2개 분할 (ctrl+b, ")

5. 백그라운드로 node 프로그램 실행 및 로드 생성
```
node --perf-basic-prof hello.js &
while :; do curl localhost:3000; sleep 0.3; done
```

6. perf 명령 실행
```
sudo su
perf record -F 99 -p `pgrep -n node` -g -- sleep 30
```

7. perf 결과 시각화  
+ 커널 트레이스 정보 포함하기 위해 root 권한으로 실행
```
git clone https://github.com/brendangregg/FlameGraph.git
perf script | ./FlameGraph/stackcollapse-perf.pl | ./FlameGraph/flamegraph.pl > perf.svg
```

8. 파일 내용 확인
+ scp로 파일 복사 후 브라우저로 svg 파일 열기 
