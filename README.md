# Commands
## Create ecs-cli profile
`test-profile` という名前でecs-cliのprofileを作成する。

```shell
ecs-cli configure profile --profile-name test-profile --access-key ${ACCESS_KEY} --secret-key ${SECRET_KEY}
```

## Create cluster config
`test-config` という名前でconfigを作成する。
デフォルトの起動タイプはFargate。

```shell
ecs-cli configure --cluster ecs-test --default-launch-type FARGATE --config-name test-config --region us-east-1
```

## Deploy service
`test-service` という名前でECSタスク・サービスをデプロイする。

```shell
ecs-cli compose --project-name test-service service up \                                                                                                                     44s 159ms
    --cluster-config test-config \
    --ecs-profile private_gitlab \
```

## Deploy service(Target group)
`test-service` という名前でECSタスク・サービスをデプロイする。
docker-compose.ymlに定義されている `web` を指定したターゲットグループに登録する。

```shell
ecs-cli compose --project-name test-service service up \                                                                                                                     44s 159ms
    --cluster-config test-config \
    --ecs-profile private_gitlab \
    --target-group-arn ${TARGET_GROUP_ARN}
    --container-name web \
    --container-port 80
```

## docker-compose.ymlを生成する
```shell
ecs-cli local create --task-def-remote test-service:1
```
