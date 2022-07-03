## [실습] AWS 네트워크 환경 구성

### VPC 생성 및 서브넷 구성
1. VPC 생성
+ my-vpc

2. 서브넷 생성
+ pub-subnet-1
+ priv-subnet-1

### 퍼블릭, 프라이빗 서브넷 구성
1. (퍼브릭 서브넷) 인터넷 게이트웨이 생성 및 VPC 연결

2. (퍼브릭 서브넷) 라우팅 테이블 생성
+ pub-rt

3. (퍼브릭 서브넷) 라우팅 테이블 수정
+ 라우팅: 0.0.0.0/0 --> 인터넷 게이트웨이
+ 서브넷 연결: pub-subnet-1

4. NAT 게이트웨이 생성
+ EIP 생성

5. (프라이빗 서브넷) 라우팅 테이블 생성
+ priv-rt

6. (프라이빗 서브넷) 라우팅 테이블 라우팅 수정
+ 라우팅: 0.0.0.0/0 --> NAT 게이트웨이
+ 서브넷 연결: priv-subnet-1

7. 인스턴스 2대 생성 
+ pub-server (퍼브릭 서브넷)
+ priv-server (프라이빗 서브넷)

8. 연결성 분석기로 연결 상태 확인
+ IGW -> pub-server
+ IGW -> priv-server
+ pub-server -> priv-ser
+ priv-server -> IGW

## [실습정리] 자원삭제
1. 인스턴스 삭제

2. NAT 게이트웨이 삭제

3. EIP 삭제

4. 인스턴스 종료 확인 후 VPC 삭제
