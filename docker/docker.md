# docker環境の作り方

## Dockerfile

- FROMでOSを指定

```dockerfile
FROM ubuntu:18.04
```

- RUNで実行するコマンドを指定<br>
- ARGでファイル内で使用する変数を宣言

```dockerfile
RUN mkdir /var/run/sshd

ARG ROOT_PASSWORD
RUN echo root:${ROOT_PASSWORD} | chpasswd
```

## コンテナ立ち上げ
Dockerfileがある階層で実行すること

```dockerfile
docker build --build-arg ROOT_PASSWORD="y20010410" -t ssh_container .
```

## コンテナ実行
```bash
docker run --network host -itd --rm --name ssh_container ssh_container 
```

## コンテナ接続
```bash
ssh root@localhost
```

## userの作成
```bash
useradd python_test
```

## pythonの実行
```python
python3 sample.py < sample.txt > result.txt 2>&1
```

## コンテナのコマンドを実行
```shell
docker exec
```

## コンテナにファイルを渡す
```shell
docker cp
```

## shellを実行
```shell
. docker.sh 渡すファイル名 実行する言語 user_id 引数ファイル
```
