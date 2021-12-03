---
id: create-certificate
title: Create Certificates
---

# 2. Create Certificates

## Install Cert Manager

All NLX components within an organization communicate with each other using internal TLS certificates. These certificates can be managed automatically with the help of [cert-manager](https://cert-manager.io/).

Install cert-manager on the cluster with:

```
helm repo add jetstack https://charts.jetstack.io
helm repo update

kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.4.0/cert-manager.crds.yaml

helm install \
  cert-manager jetstack/cert-manager \
  --namespace cert-manager \
  --create-namespace \
  --version v1.4.0
```

## Create CA Issuer
Now we create a CA Issuer for NLX.

### Create private key

Create the private key with:
```
openssl genrsa -out ca.key 2048
```

Check if `ca.key` is created by running:

```
ls
```

### Create Certificate

For Linux:

```
openssl req -x509 -new -nodes -key ca.key -subj "/CN=NLX" -days 3650 -reqexts v3_req -extensions v3_ca -out ca.crt
```

For macOS:

```
openssl req -x509 -new -nodes -key ca.key -subj "/CN=NLX" -days 3650 -reqexts v3_req -extensions v3_ca -out ca.crt -config /usr/local/etc/openssl@1.1/openssl.cnf
```

Check if `ca.crt` is created by running:

```
ls
```

### Create the secret

Let's create the Kubeternetes TLS secret now:

```
kubectl create secret tls internal-ca \
   --cert=ca.crt \
   --key=ca.key \
   --namespace=nlx
```

We now install the internal-issuer on the cluster:

```
kubectl apply -f internal-issuer.yaml
```

Then see if the internal Issuer is done by running:

```
kubectl get issuer --namespace nlx
```

The expected result:

```
NAME       READY   AGE
internal   True    ??
```

## Create the external certificate

Traffic between organizations takes place via an external certificate. For the NLX demo environment, you can easily create a certificate with the certificate portal.

**Note**: If you wish to choose your own organization serial number or keep the serial number when generating a CSR, you need to add the following lines to your openssl config:

```toml
[ req_distinguished_name ]
serialNumber = Serial Number
```

Generate a private key and certificate request for the organization by running:

```
openssl req -utf8 -nodes -sha256 -newkey rsa:4096 -keyout org.key -out org.csr
```

You now have to answer a series of questions. Below is an example of the answers. If the answer is `<skip>` then you can skip this question by pressing enter.

```
- Country Name (2 letter code) []: NL
- State or Province Name (full name) []: Noord holland
- Locality Name (eg, city) []: Haarlem
- Organization Name (eg, company) []: mijn-organisatie (url friendly name with a maximum length of 100 characters)
  (remember this because we will need it later)
- Organizational Unit Name (eg, section) []: <skip>
- Common Name (eg, fully qualified host name) []:
  (host name where your inway can be reached), for example:
  inway.mijn.organisatie.nl
- Email Address []: <skip>
- Serial Number []: (optional), a serial number with a maximum length of 20 characters. Also make sure this value is unique for the network in the [directory overview](https://directory.demo.nlx.io) as we do not check for uniqueness  
- A challenge password []: <skip>
```


A CSR has now been created (Certificate Signing Request). Check the content with the following command:

```
openssl req -in org.csr -text
```

The result should look similar.
```
Subject: C=NL, ST=Noord holland, L=Haarlem, O=mijn-organisatie, CN=inway.mijn.organisatie.nl
```

Print the contents of the CSR to the terminal with:

```
cat org.csr
```

Copy the CSR from the terminal including `-----BEGIN CERTIFICATE REQUEST-----` and `-----END CERTIFICATE REQUEST-----`.


Now open the [NLX Certificate Portal](https://certportal.demo.nlx.io/) and paste the contents of the CSR into the text field. Click on `Request Certificate` in the portal to request your certificate.

If this is successful, you will see `Download certificate` at the bottom of the page,

Click on this and save the certificate as `org.crt` your working directory

Check if `org.crt` is created by running:

```
ls
```

Then create a Kubernetes TLS secret by running:

```
kubectl create secret tls external-tls \
   --cert=org.crt \
   --key=org.key \
   --namespace=nlx
```

Your certificate now exists as secret in Kubernetes. We will use this secret when we install NLX management and the NLX inway.

## Obtaining your Subject Serial Number

The Subject Serial Number of your certificate, added by the Certificate Portal, is the primary identifier of your organization within NLX.

To obtain your serial number, see the Subject part of the certificate by running:

```bash
openssl x509 -in org.crt -text | grep Subject:
```

Example of the output: `Subject: C=nl, ST=zuid-holland, L=gemeente-stijns, O=my-organization, OU=my-organization-unit, CN=an-awesome-organization.nl/serialNumber=01234567890123456789`.

The value after `serialNumber=` in the Subject's CN field is the Subject Serial Number. Save this, because it will later be used to access your own APIs when using the Outway.

For details about this, see the [organization identification](/reference-information/organization-identification) page.
