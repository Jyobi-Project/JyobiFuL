#!/bin/sh

# 実行用ディレクトリを作成
docker exec -it -w /root/$2_test/ ssh_container mkdir $3_test

# 実行ファイルをdockerに保存
docker cp ./$1 ssh_container:/root/$2_test/$3_test/

# 標準入力用のファイルをdockerに保存
docker cp ./$4 ssh_container:/root/$2_test/$3_test/

# プログラムを実行し、result.txtに保尊
docker exec -i -w /root/$2_test/$3_test/ ssh_container $2 $1 < $4 > $3_result.txt 2>&1

# 作業ディレクトリを削除
docker exec ssh_container rm -rf /root/$2_test/$3_test
