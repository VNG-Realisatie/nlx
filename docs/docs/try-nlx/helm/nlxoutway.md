---
id: nlx-outway
title: Install NLX Outway
---

# 6. Install NLX Outway

## Install internal certificate

We are going to use the cert-manager we installed earlier to create an internal certificate.

Run the following command to install the certificate on the Kubernetes cluster:

```
kubectl -n nlx apply -f outway-internal-tls.yaml
```

## Install Outway Chart

Now we create the values file with the settings for Helm. Open the file `nlx-outway-values.yaml` in an editor, edit the values below and save the file:

- The value `<file: ca.crt>` must be replaced by the contents of the ca.crt file. This file is in your working directory.
   - Copy the contents of the files **excluding** the '-----BEGIN CERTIFICATE-----' and '-----END CERTIFICATE-----' lines.
   - Paste the content between the start and end lines and make sure the alignment is the same as the start and end lines
   - Save the modified file

Then we install the Outway by running:

```
helm -n nlx upgrade --install outway -f nlx-outway-values.yaml commonground/nlx-outway
```

Check with `kubectl -n nlx get pods` if the service pod is healthy.

A similar line should now show up:

```
brfk-basisregister-fictieve-kentekens-6854dd6f86-rlqgk   1/1     Running     1          4d5h
inway-nlx-inway-55687b9fc6-r9jqx                         1/1     Running     3          4d5h
management-api-create-user-ntgjj                         0/1     Completed   0          4d5h
management-nlx-management-api-5c7b9bb84f-xsbh4           1/1     Running     2          4d23h
management-nlx-management-ui-c5ddb4c69-bsd6p             1/1     Running     3          4d23h
outway-nlx-outway-5c69944c9-4jzrr                        1/1     Running     0          3m8s
postgresql-postgresql-0                                  1/1     Running     1          4d23h
```
