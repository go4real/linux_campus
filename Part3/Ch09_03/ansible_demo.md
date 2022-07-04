## [실습] Ansible로 사용자 관리

0. Ansible cli 도구 설치
python3 -m pip install --user ansible

1. ansible 실습을 위한 디렉토리 생성
```
mkdir ansible_demo
cd ansible_demo
```

2. ansible 사용자용 키페어 생성 
 + ssh-keygen 명령을 사용하여 ansible_key, ansible_key.pub 파일 생성
 ```
 ssh-keygen -C "ansible ssh key" -q -f ./ansible_key -t rsa -b 2048 -N ""
 ```

3. AWS 웹 콘솔에서 신규 인스턴스 생성
 + 인스턴스 이름: ansible_demo_1
 + OS: Ubuntu Linux
 + PublicKeypair 부분은 ansible_key.pub 파일의 전체 내용으로 교체
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
  - name: ansible
    groups: [ sudo, adm ]
    sudo: [ "ALL=(ALL) NOPASSWD:ALL" ]
    shell: /bin/bash
    ssh_authorized_keys: 
    - PublicKeypair
```

4. 호스트 파일 생성 및 타겟 호스트 IP 추가
+ vim demo_hosts
```
10.0.0.15
```

5. ansible 설정 파일 생성
+ vim ansible.cfg
```
[defaults]
inventory = /home/ubuntu/ansible_demo/demo_hosts
remote_user = ansible
private_key_file = /home/ubuntu/ansible_demo/ansible_key
```

6. 타겟 호스트 접속 확인
```
ansible all -m ping
```

7. 플레이북 작성
+ vim user_mgmt.yml
```
---
- name: user module demo
  hosts: all
  become: true
  tasks:
    - name: Create a new user 
      ansible.builtin.user:
        name: fast
        password: "{{ '1234qwer' | password_hash('sha512') }}"
        groups:
          - adm
          - sudo
        state: "present"
        shell: "/bin/bash"
        system: false
        create_home: true
        home: "/home/fast"
        comment: "Ansible demo user"

    - name: Disable SSH password authentication
      ansible.builtin.lineinfile:
        path: /etc/ssh/sshd_config
        regexp: '^PasswordAuthentication'
        line: PasswordAuthentication yes

    - name: Restart sshd service
      ansible.builtin.service:
        name: sshd
        state: restarted
```

8. 플레이북 실행
```
ansible-playbook user_mgmt.yml
```

