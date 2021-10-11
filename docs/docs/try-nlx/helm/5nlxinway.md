---
id: 5nlxinway
title: 5. Install NLX Inway
---

# 5. Install NLX Inway

## Install internal certificate

We are going to use the cert-manager we installed earlier to create an internal certificate.

Run the following command to install the certificate on the Kubernetes cluster:

```
kubectl -n nlx apply -f inway-internal-tls.yaml
```

## Install Inway chart

Now we create the values file with the settings for Helm. Open the file `nlx-inway-values.yaml` in an editor, edit the values below and save the file:

- `<inway-name>` the name of your inway. The name must contain only alphanumeric characters and dashes. eg: `my-inway`
- `<self-address>` the address at which your inway can be reached, including the port on which the inway is available (443 by default). This address must match the address (the Common Name) of your organization certificate you created earlier. eg: `inway.mijn-organisation.nl:443`
- The value `<file: ca.crt>` must be replaced by the contents of the ca.crt file. This file is in your working directory.
   - Copy the contents of the files **excluding** the '-----BEGIN CERTIFICATE-----' and '-----END CERTIFICATE-----' lines.
   - Paste the content between the start and end lines and make sure the alignment is the same as the start and end lines
   - Save the modified file

Then we install the Inway by running:

```
helm -n nlx upgrade --install inway -f nlx-inway-values.yaml commonground/nlx-inway
```

Check with `kubectl -n nlx get pods` whether the Inway pod is healthy.

You should now see something like this:

```
inway-nlx-inway-55687b9fc6-r9jqx                 1/1     Running     0          2m30s
management-api-create-user-ntgjj                 0/1     Completed   0          40m
management-nlx-management-api-5c7b9bb84f-xsbh4   1/1     Running     1          17h
management-nlx-management-ui-c5ddb4c69-bsd6p     1/1     Running     2          17h
postgresql-postgresql-0                          1/1     Running     0          18h
```

To verify the status of the Inway, go to NLX management and see if your inway (with the name you chose) is in your list of inways.

## Link domain to your inway

Run the following command:

```
kubectl get -n nlx svc
```

In the result, you will see a service for the Inway called `nlx-inway`. Copy the value stated at `EXTERNAL-IP` and link this IP address to your inway domain.

After linking your domain, run the following command:
```
nslookup <the domain address of your inway>
```

In the result with the address equal to the `EXTERNAL-IP` of the service `nlx-inway`

*Note: it may take some time before the domain is linked. You can continue with this guide, but eventually, the domain must be linked before your service can be accessed via the inway.*

## Set up organization inway

All traffic related to access requests, assignments, etc. runs through the organization inway. You have to set this in the settings page in NLX management. There you select your inway as an organization inway and save the new settings.