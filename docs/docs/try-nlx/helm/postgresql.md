---
id: postgresql
title: Install PostgreSQL
---

# 3. Install PostgreSQL

NLX uses PostgreSQL to store various data, including the configuration of the various components and transaction logging. In this section, we install PostgreSQL using the [Bitnami Helm chart](https://bitnami.com/stack/postgresql/helm).

Install PostgreSQL by running:

```
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update

helm install \
  postgresql bitnami/postgresql \
  --namespace nlx \
  --version 10.6.0 \
  --set postgresqlDatabase=nlx_management
```

Then check if PostgreSQL is running properly by running:

```
kubectl get pods -n nlx
```

The output should look like this (make sure the status is `Running`).

```
NAME                      READY   STATUS    RESTARTS   AGE
postgresql-postgresql-0   1/1     Running   0          ?s
```

Make sure you get the password for PostgreSQL and write it down somewhere, we'll need this password later in this guide. Get the password with the following command:

```
kubectl get secret --namespace nlx postgresql -o jsonpath="{.data.postgresql-password}" | base64 -d
```

*Save this password somewhere without the closing `%` sign*
