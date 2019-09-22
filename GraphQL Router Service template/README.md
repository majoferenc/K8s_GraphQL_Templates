# GraphQL Node.js backend

Experimental microservice with Node.js.

### Features

- GraphQL API with GraphiQL playground
- Skaffold CI/CD
- Draft CI/CD
- Docker image
- Docker-compose ready
- K8s Helm chart

### Deployment

You can run service automatic Docker image build and K8s Helm chart push after every source change with:

```sh
skaffold dev
```

If you want also publish Docker image to your container registry you can use:

```sh
draft up
```
