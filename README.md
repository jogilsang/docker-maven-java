# repo-app-source
repo-app-source

- Github (Enable Other service EX : Gitlab)
- CodeBuild
  - Dockerfile
  - buildspec.yml
    - $AWS_ACCOUNT_ID
    - $AWS_DEFAULT_REGION
    - $ECR_REPO_NAME
- CodeDeploy
  - BuildArtifact
    - appspec.yaml
      - $CONTAINER_NAME
      - $CONTAINER_PORT
    - taskdef.json
      - $HOST_PORT
      - $CONTAINER_PORT
      - $CONTAINER_NAME
      - $TASK_FAILY
    - imageDetail.json
