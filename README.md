# go-echo-firebase
GoでechoとFirebaseAuthを利用したAPIサンプル

# USAGE
Dockerイメージの作成

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
