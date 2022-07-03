## [실습] 리눅스 커맨드 파이썬으로 구현

### tail -f 구현하기
1. 파이썬 스크립트 작성
+ vim py_tail
```
#!/usr/bin/python3
import sys
import time

def tail(filename):
    with open(filename) as f:
        f.seek(0,2)
        while True:
            line = f.readline()
            if not line:
                time.sleep(0.1)
                continue
            yield line

filename = sys.argv[1]
for line in tail(filename):
    print(line)
```

2. 실행 권한 부여
```
chmod +x py_tail
```

3. /usr/local/bin 디렉토리로 이동
```
sudo cp py_tail /usr/local/bin
```

4. 테스트
```
py_tail /var/log/auth.log
```

### wc 구현하기
1. 파이썬 스크립트 작성
+ vim py_wc
```
#!/usr/bin/python3
import sys
import time

def wc(filename):
    chars = words = lines = 0
    with open(filename) as f:
        for line in f:
            lines += 1
            words += len(line.split())
            chars += len(line)
            
    print("{} {} {} {}".format(lines, words, chars, filename))

filename = sys.argv[1]
wc(filename)
```

2. 실행 권한 부여
```
chmod +x py_wc
```

3. /usr/local/bin 디렉토리로 이동
```
sudo cp py_wc /usr/local/bin
```

4. 테스트
```
wc py_tail
py_wc py_tail
```
