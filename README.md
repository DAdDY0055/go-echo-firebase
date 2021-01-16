# go-echo-firebase
GoでechoとFirebaseAuthを利用したAPIサンプル

# USAGE
## Dockerイメージの作成

```
% docker build -t echo-firebase .
```

```
% docker-compose up -d
```

Dockerイメージの起動

```
% docker run -p 8080:8080 echo-firebase
```

`localhost:8080`でアクセス

DBアクセス

```
% mysql -u root -h 127.0.0.1 -P 33060 -proot echoAPI
```

## Modelファイルの生成

[sqlboiler](https://github.com/volatiletech/sqlboiler#configuration)でModel自動生成用のコマンド取得のため以下を実行

```
% go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql
```

DBの定義を変更した後、以下コマンドを実行

```
% sqlboiler mysql --output models --pkgname models --wipe
```

→ `models`ディレクトリ配下にSQLBoilerでDB接続用のモデルファイルが自動生成される
