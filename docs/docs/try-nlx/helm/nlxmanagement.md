---
id: nlx-management
title: Install NLX Management
---

# 4. Install NLX Management

We are now going to install NLX management. NLX management is an API and a web interface with which an NLX installation can be managed.

If you have **not** installed Postgres via the Bitnami chart as described in this guide then you need to make sure that a database called `nlx_management` exists.

## Install internal certificate

Run the following command to install the internal NLX management certificate on the Kubernetes cluster
```
kubectl -n nlx apply -f management-internal-tls.yaml
```

## NLX management chart

We are now going to create a configuration file for the NLX management installation. Open the file `nlx-management-values.yaml` in an editor (e.g. nano, nvim, notepad, etc), edit the values below and then save the file:

- `<hostname nlx-management>` replace this with the hostname on which your NLX management should run eg: management.mijn-organization.nl
   **IMPORTANT: this DNS entry must already exist and resolve to the ingress controller**
- `<postgres-password>` replace this with the Postgres password you saved earlier.
- The values `<file: [name]>` must be replaced by the contents of the named file. You have all these files in your working directory.
   - Copy the contents of the files **excluding** the '-----BEGIN XXXXXXXXX-----' and '-----END XXXXXXXXX-----' lines.
   - Paste the content between the start and end lines and make sure the alignment is the same as the start and end lines
   - Save the modified file

Next, we install NLX Management:

```
helm repo add commonground https://charts.commonground.nl

helm repo update

helm -n nlx upgrade --install management -f nlx-management-values.yaml commonground/nlx-management
```

Now open your browser and go to NLX management (should be located at the hostname you entered) and see if the NLX management start page is displayed. You cannot log in yet.

## Create NLX management administrator

Now we need to create an administrator for NLX management. Open the `job-create-administrator.yaml` file in an editor, edit the values below and save the file:

- `<postgres_password>` with the Postgres password you saved earlier.
- `It is also possible to change admin@example.com (user) and password (password).`

Then run the following command:

```
kubectl -n nlx apply -f job-create-administrator.yaml
```

An administrator account has now been created that you can use to log in to NLX management.

Open NLX management and log in with the following data (or with the data you have adjusted)

email: admin@example.com
password: password

## Accepting the Terms of Service (ToS)

You will have to accept the Terms of Service before you can make use of the NLX network. You can do so by logging in to NLX Management. Please review the terms carefully before accepting them.
