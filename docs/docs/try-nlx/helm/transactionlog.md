---
id: transaction-log
title: Setup the Transaction Log (optional)
---

# 9. Setup the Transaction Log (optional)

The Inway and the Outway are able to log metadata of the requests they process, these log records can be viewed using NLX Management. In this step you will learn how to setup the Transaction Log.

## Preparation

Start by downloading the required files

```
curl --location \
    --remote-name https://gitlab.com/commonground/nlx/nlx/-/raw/master/technical-docs/nlx-helm-installation-guide/txlog-api-internal-tls.yaml \
    --remote-name https://gitlab.com/commonground/nlx/nlx/-/raw/master/technical-docs/nlx-helm-installation-guide/txlog-api-job-mmigrations.yaml \
    --remote-name https://gitlab.com/commonground/nlx/nlx/-/raw/master/technical-docs/nlx-helm-installation-guide/txlog-api-values.yaml
```

## Certificate

Run the following command to install a certificate for the transaction-log API on the Kubernetes cluster:

```
kubectl -n nlx apply -f txlog-api-internal-tls.yaml
```

## Setup the database

The transaction logs are stored in a Postgres database. We will use the [Postgres instance](./postgresql.md) we created earlier. Open the `txlog-api-job-migrations.yaml` file in an editor, replace the value `<postgres_password>` with the Postgres password you saved earlier and save the file.

The following command will start a job in Kubernetes to create the database `nlx_transaction_log` and the required database schema:
```
kubectl -n nlx apply -f txlog-api-job-migrations.yaml
```

Check with `kubectl -n nlx get jobs` if the job has succeeded.

A similar line should now be displayed:

```
NAME                COMPLETIONS   DURATION   AGE
transaction-log     1/1           3s         30m
```

## Install the Transaction Log API

The transaction logs can be viewed in NLX Management. NLX Management communicates with the Transaction Log API to retrieve the logs from the database. Now let's install the Transaction Log API on the Kubernetes cluster.

First open the `txlog-api-values.yaml`, edit the values below and save the file:

- `<postgres-password>` replace this with the Postgres password you saved earlier.
- The values `<file: ca.crt>` must be replaced by the contents of the file `ca.crt`. You have this file in your working directory.
   - Copy the contents of the files **excluding** the '-----BEGIN XXXXXXXXX-----' and '-----END XXXXXXXXX-----' lines.
   - Paste the content between the start and end lines and make sure the alignment is the same as the start and end lines
   - Save the modified file

Run the following commands to install the Transaction Log API on the cluster:

```
helm repo add commonground https://charts.commonground.nl

helm repo update

helm -n nlx upgrade --install txlog-api -f txlog-api-values.yaml commonground/nlx-txlog-api
```

Check if the Transaction Log API is running:

```
kubectl -n nlx get pods
```

A similar line should now show up:

```
txlog-api-nlx-txlog-api-7ff48948f8-pkdfm   1/1     Running     1          1m30s
```

## Update the NLX deployment

Now that the Transaction Log API is running, we need to update our existing NLX components so they start using the Transaction Log.

Open the `nlx-management-values.yaml` file in an editor, uncomment the line `#txlogAPIAddress: txlog-api-nlx-txlog-api:8443` and save it. The first view lines of the file should now look like this:

```yaml
config:
  directoryHostname: directory-api.demo.nlx.io
  enableBasicAuth: true
  txlogAPIAddress: txlog-api-nlx-txlog-api:8443
```

Update the NLX Management deployment:

```
helm -n nlx upgrade --install management -f nlx-management-values.yaml commonground/nlx-management
```

Open a browser and go to NLX Management. In the menu, on the left side of the screen, select `Transactie logs`, if everything is configured correctly you should now see an empty overview.

The next step will be to update the Outway and Inway deployments so that these components will write to the transaction log after receiving a request.

Let's start with updating the Inway, open `nlx-inway-values.yaml` in a file editor and find this section:

```
transactionLog:
  enabled: false
```

Change it to the following and make sure to replace `<postgres-password>` with the Postgres password your copied earlier:

```
transactionLog:
  enabled: true
  hostname: postgresql
  username: postgres
  password: <postgres-password>
  ## sslMode disabled is not recommended for a production environment
  sslMode: disable
  database: nlx_transaction_log
```

Update the Outway by opening `nlx-outway-values.yaml` in a file editor and repeat the same process.

We have now updated the value files of both the Inway and Outway. Next up is redeploying the Inway and Outway so that the new values are used. Run the following command:

```
helm -n nlx upgrade --install inway -f nlx-inway-values.yaml commonground/nlx-inway \
helm -n nlx upgrade --install outway -f nlx-outway-values.yaml commonground/nlx-outway
```

Check with whether the Inway and Outway pods are healthy:

```
kubectl -n nlx get pods
```

You should now see something like this:

```
inway-nlx-inway-55687b9fc6-r9jqx                  1/1     Running     0          2m30s
outway-nlx-outway-5c69944c9-4jzrr                 1/1     Running     0          2m31s
```

## Test the transaction log

Now we will make a request through our Outway to fill the transaction log we just setup. To make a request please repeat the step ['Query API' of the step 'Access API through a client'](./accessapi.md#query-api).

After making the request open your browser and navigate to NLX Management, login and open the `Transaction logs`. You should now see two entries, one of the Outway sending the request and one of the Inway receiving the request.
