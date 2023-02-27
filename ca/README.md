How to renew the NLX root cert on Demo
---

First download an existing certificate from the Demo environment.
We will use this to verify if the newly generated CA is able to verify the existing certificates.

```shell
cd ./ca
kubectl get secret -n nlx-demo gemeente-riemer-organization-tls -o jsonpath="{.data['tls\.crt']}" | base64 -d > gemeente-riemer-tls.crt
```

Now let's create our new CA. Run the following command:

```shell
cfssl gencert \
    -ca-key root.key \
    -config "config.json" \
    -initca \
    "ca.json" | cfssljson -bare "root"
```

This will output two files (root.csr and root.pem) in the `./ca` directory.

Verify if the CA is still able to verify the existing certificate we've copied from Gemeente Riemer.

```shell
openssl verify -CAfile root.pem gemeente-riemer-tls.crt
```

Copy the content of `root.pem`.

Replace the existing root certificate PEM in the Helm value files.

Path of the property to replace: `global.tls.organization.rootCertificatePEM`.

For the following files:

```
helm/deploy/rvrd/values-demo.yaml
helm/deploy/gemeente-stijns/values-demo.yaml
helm/deploy/gemeente-riemer/values-demo.yaml
helm/deploy/vergunningsoftware-bv/values-demo.yaml
helm/deploy/shared/values-demo.yaml
```

Path of the property to replace: `ca.issuer.certificatePEM`.

For the following files:

```
helm/deploy/shared/values-demo.yaml
```

TODO: describe which secrets to remove
