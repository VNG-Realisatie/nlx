# NLX walkthrough

During this walkthrough we'll setup an NLX outway make a request to a service that is available on the NLX demo network.

This guide is written for Mac/Linux users and assumes you have some experience with the terminal/shell. Windows 10 users may be able to follow this tutorial using [ubuntu for windows](https://tutorials.ubuntu.com/tutorial/tutorial-ubuntu-on-windows). However, the windows platform is not officially supported, we advise to use a VM with Ubuntu installed.

## Certificate

To run the outway, we need a certificate for encrypted/TLS communications with other NLX services on the network.

The demo network is public, anyone can create a signed certificate using the "certportal". The certportal will simply sign any certificate with the demo network key. This allows anyone to get access to the demo network quickly without needing to contact a service desk for a valid certificate.

Lets start by creating a directory structure for this walkthrough.

```bash
mkdir -p $HOME/nlx-walkthrough/certs
```

Change directory to the nlx-walkthrough folder.

```bash
cd $HOME/nlx-walkthrough
```

We need to create a key and certificate signing request (csr) for our organization.

```bash
openssl req -utf8 -nodes -sha256 -newkey rsa:4096 -keyout certs/org.key -out certs/org.csr
```

You will need to provide information to the questions asked. The only question that is important right now is `Organization Name (eg, company)` this field is used in the logs of NLX. Please use your name or make up a random name. e.g. "Louis-van-Gaal". Make sure this is something unique so you can easily find your own logs later.

View the contents of `certs/org.csr`:

```bash
cat certs/org.csr
```

Copy the complete CSR, including `-----BEGIN CERTIFICATE REQUEST-----` and `-----END CERTIFICATE REQUEST-----`.

Go to https://certportal.demo.nlx.io and paste the text of the csr you've just coped. Click "Request Certificate". Copy the resulting certificate text.

Create a file, `certs/org.crt`, and paste the certificate. Or click download and move the `certificate.crt` to `certs/org.crt`.

__In a production environment one would not be able to create their own certificates, but instead retrieves them from the central authority, which asserts the Organization Name field is properly set. The certportal is only available in test and demo environments.__

We now have a certificate by which we can identify ourselves to other NLX services, but we still need the rootCA certificate, which we use to verify other NLX services.

```bash
wget https://certportal.demo.nlx.io/root.crt -O certs/root.crt
```

We're now ready to start outways and inways.

## Outway

The easiest way to run an outway is by using docker. Please make sure you have a recent stable version of docker [installed](https://docs.docker.com/install/). 

Start by downloading the latest version of the container image from the docker hub.

```bash
docker pull nlxio/outway:latest
```

Next, we want docker will to start a container based on this image.

```bash
docker run \
    --tty --interactive \
    --volume=$HOME/nlx-walkthrough/certs:/certs:ro \
    --publish=2018:80 \
    nlxio/outway:latest \
    /usr/local/bin/nlx-outway \
    --log-type=development \
    --log-level=debug \
    --directory-address=directory-api.demo.nlx.io:443 \
    --tls-nlx-root-cert=/certs/root.crt \
    --tls-org-cert=/certs/org.crt \
    --tls-org-key=/certs/org.key \
    --disable-logdb
```

We give docker several arguments:

- `--tty` and `--interactive` tells docker to connect the terminal you're using with the container so you may view logs.
- `--volume` tells docker to make the `certs` folder, where we just put the certificates, available inside the container.
- `--publish` connects port 2018 on the host machine to port 80 inside the container. This way, we can send requests to the outway.
- `nlxio/outway:latest` is the name of our docker image (`nlxio/outway`) as stored in the docker registry and the version we want to use (`latest`). The `--` tells docker that all arguments after this one are meant for the outway process, not for docker itself.
- `/usr/local/bin/nlx-outway` is the binary in the container that we want to execute.

All arguments after the image and executable name are passed to the process (nlx-outway) itself..

- `-log-*` setup developmeng logs so we can see what's going on
- `--directory-address` is set to point the outway to the demo environment directory service. This means we can use this outway to connect to services in the demo environment.
- `--tls-*` tells outway where to find the root, org-cert and org-key files.
- `--disable-logdb` tells outway to not write transaction log records, because we have not set up a database for it. Transactionlogs are still written at remote services we will be using.

Make a request through the outway. VNG-Realisatie is hosting a demo api on the network, lets try that.

```bash
curl "localhost:2018/vng-realisatie/demo-api/"
```

You now have a outway running and connected to the demo network of NLX. To make requests to a service, use the URL: `localhost:2018/<organization>/<service>`.

Take a look at the NLX directory at https://directory.demo.nlx.io to find other services and their API documentation.

