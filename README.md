# Usage
## Create ecs cluster
```shell
aws --profile $AWS_PROFILE cloudformation deploy \
--no-fail-on-empty-changeset \
--capabilities CAPABILITY_NAMED_IAM \
--template-file cloudformation.yml \
--stack-name ecs-test
```

## Get outputs
```shell
export WEB_IMAGE=$(aws ecr describe-repositories | jq -r '.repositories[] | select(.repositoryName == "ecs-test") | "\(.repositoryUri):latest"')
export TARGET_GROUP_ARN=$(aws --profile $AWS_PROFILE cloudformation describe-stacks --stack-name ecs-test | jq -r '.Stacks[0].Outputs[] | select(.OutputKey == "TargetGroup") | .OutputValue')
export SUBNET_0=$(aws --profile $AWS_PROFILE cloudformation describe-stacks --stack-name ecs-test | jq -r '.Stacks[0].Outputs[] | select(.OutputKey == "Subnet0") | .OutputValue')
export SUBNET_1=$(aws --profile $AWS_PROFILE cloudformation describe-stacks --stack-name ecs-test | jq -r '.Stacks[0].Outputs[] | select(.OutputKey == "Subnet1") | .OutputValue')
export LOG_GROUP=$(aws --profile $AWS_PROFILE cloudformation describe-stacks --stack-name ecs-test | jq -r '.Stacks[0].Outputs[] | select(.OutputKey == "ECSLogGroup") | .OutputValue')
export SECURITY_GROUP=$(aws --profile $AWS_PROFILE cloudformation describe-stacks --stack-name ecs-test | jq -r '.Stacks[0].Outputs[] | select(.OutputKey == "SecurityGroup") | .OutputValue')
```

# Commands
## Build docker image and push
```shell
docker-compose -f docker-compose.yml -f docker-compose.ecs.yml build
docker-compose -f docker-compose.yml -f docker-compose.ecs.yml push
```

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
ecs-cli compose \
    --project-name test-service \
    --file docker-compose.yml \
    --file docker-compose.ecs.yml \
    --cluster-config test-config \
    --ecs-profile test-profile \
    service up \
    --cluster-config test-config \
    --ecs-profile test-profile
```

## Deploy service(Target group)
`test-service` という名前でECSタスク・サービスをデプロイする。
docker-compose.ymlに定義されている `web` を指定したターゲットグループに登録する。

```shell
ecs-cli compose \
    --project-name test-service \
    --file docker-compose.yml \
    --file docker-compose.ecs.yml \
    service up \
    --cluster-config test-config \
    --ecs-profile test-profile \
    --target-groups "targetGroupArn=$TARGET_GROUP_ARN,containerPort=8080,containerName=web"
```

## Stop service
```shell
ecs-cli compose \
    --project-name test-service \
    --file docker-compose.yml \
    --file docker-compose.ecs.yml \
    service down \
    --cluster-config test-config \
    --ecs-profile test-profile
```

## docker-compose.ymlを生成する
```shell
ecs-cli local create --task-def-remote test-service:1
```
