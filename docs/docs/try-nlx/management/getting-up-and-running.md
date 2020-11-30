---
id: getting-up-and-running
title: Getting up and running
---


## Start NLX using Docker Compose

Now we have prepared all the requirements to run NLX, we can start all components using Docker Compose.

> Next to the certificates you created in [retrieve a demo certificate](../retrieve-a-demo-certificate.md) you also need certificates from an internal PKI to encrypt traffic between NLX components (such as the Management API and the Inway). The demo already has a working PKI so you don't have to set this up yourself.

First, let's clone the NLX project. It contains the Docker Compose file and its dependencies.

```bash
git clone https://gitlab.com/commonground/nlx/nlx-try-me.git
```

After the repository is cloned, move into it:

```bash
cd nlx-try-me
```

Set the hostname of the Inway (where `my-organization.nl:443` should be replaced with your own hostname).

```bash
echo "INWAY_SELF_ADDRESS=my-organization.nl:443" > .env
```

Then, start all components by running:

```bash
docker-compose up
```

This will start [Dex](https://github.com/dexidp/dex) (Identity Provider), ETCD and the required NLX components.

The NLX components are configured using environment variables which in this guide are set in `docker-compose.yml`

Below you is an overview of the environment variables per NLX component:

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

<Tabs
  defaultValue="inway"
  values={[
    { label: 'Inway', value: 'inway', },
    { label: 'Outway', value: 'outway', },
    { label: 'Management API', value: 'management-api', },
    { label: 'Management UI', value: 'management-ui', },
  ]}
>
<TabItem value="inway">

#### Environment variables

- `DIRECTORY_REGISTRATION_ADDRESS` This address is used by the inway to anounce itself to the directory.
- `INWAY_NAME` Alias the Inway by a name instead of it's unique identifier.
- `SELF_ADDRESS` The address of the inway so it can be reached by the NLX network.
- `MANAGEMENT_API_ADDRESS` The address of the Management API.
- `TLS_NLX_ROOT_CERT` This is the location of the root certificate.
- `TLS_ORG_CERT` This is the location of the organization certificate.
- `TLS_ORG_KEY` This is the location of the organization private key.
- `POSTGRES_DSN` Connection-string to the PostgreSQL database.
- `DISABLE_LOGDB` The value 1 will disable the transaction logs, the value 0 will enable them.
- `LOG_LEVEL` Log level of the application. Options: debug, info, warn.

</TabItem>

<TabItem value="outway">

#### Environment variables

- `DIRECTORY_INSPECTION_ADDRESS` This address is used by the outway to retrieve information about services from the directory.
- `TLS_NLX_ROOT_CERT` This is the location of the root certificate.
- `TLS_ORG_CERT` This is the location of the organization certificate.
- `TLS_ORG_KEY` This is the location of the organization private key.
- `DISABLE_LOGDB` The value 1 will disable the transaction logs, the value 0 will enable them.
- `LOG_LEVEL` Log level of the application. Options: debug, info, warn.

</TabItem>

<TabItem value="management-api">

#### Environment variables

- `DIRECTORY_INSPECTION_ADDRESS` This address is used by the outway to retrieve information about services from the directory.
- `DIRECTORY_REGISTRATION_ADDRESS` This address is used by the inway to anounce itself to the directory.
- `TLS_NLX_ROOT_CERT` This is the location of the root certificate.
- `TLS_ORG_CERT` This is the location of the organization certificate.
- `TLS_ORG_KEY` This is the location of the organization private key.
- `TLS_ROOT_CERT` This is the location of the root certificate from the internal PKI.
- `TLS_CERT` This is the location of the API certificate from the internal PKI.
- `TLS_KEY` This is the location of the private key from the internal PKI.
- `DISABLE_LOGDB` The value 1 will disable the transaction logs, the value 0 will enable them.
- `SECRET_KEY` Secret key that is used for signing sessions
- `OIDC_CLIENT_ID` The OIDC client ID
- `OIDC_CLIENT_SECRET` The OIDC client secret
- `OIDC_DISCOVERY_URL` The OIDC discovery URL
- `OIDC_REDIRECT_URL` The OIDC redirect URL
- `SESSION_COOKIE_SECURE` Use 'secure' cookies
- `LOG_LEVEL` Log level of the application. Options: debug, info, warn.

</TabItem>

<TabItem value="management-ui">

#### Environment variables

- `MANAGEMENT_API_ADDRESS` Address of the Management API.

</TabItem>
</Tabs>

At last, let's verify if all the components are up and running:

```
docker-compose ps
```

It might take a while for all components to become healthy.
If after a while one or more components aren't running you can inspect the logs for any errors.


## Dex <small>(Identity Provider)</small>

The Management UI supports the OpenID Connect protocol for authentication and authorization.
In the demo we provide Dex, which is a configurable Identity Provider.

On Linux based operating systems this works out-of-the-box.
If you're using MacOS or Windows you will need to add the hostname for Dex to the known hosts.

<Tabs
  defaultValue="mac_os"
  values={[
    { label: 'MacOS', value: 'mac_os', },
    { label: 'Windows', value: 'windows', },
  ]}
>

<TabItem value="mac_os">

```bash
sudo sh -c "echo '127.0.0.1 dex.nlx.localhost' >> /etc/hosts"
```

</TabItem>

<TabItem value="windows">

```powershell
Add-Content -Path C:\Windows\System32\drivers\etc\hosts -Value "127.0.0.1`tdex.nlx.localhost" -Force
```

</TabItem>
</Tabs>

Now let's verify that the local hostname for Dex points to the host:

<Tabs
  defaultValue="mac_os"
  values={[
    { label: 'MacOS', value: 'mac_os', },
    { label: 'Linx/Windows', value: 'linux_windows', },
  ]}
>
<TabItem value="linux_windows">

```bash
ping dex.nlx.localhost -4 -c 1
```

</TabItem>

<TabItem value="mac_os">

```bash
ping dex.nlx.localhost -c 1
```

</TabItem>
</Tabs>

The output should be:
```bash
#
# PING dex.nlx.localhost (127.0.0.1) 56(84) bytes of data.
# 64 bytes from localhost (127.0.0.1): icmp_seq=1 ttl=64 time=0.026 ms
# 
# --- dex.nlx.localhost ping statistics ---
# 1 packets transmitted, 1 received, 0% packet loss, time 0ms
# rtt min/avg/max/mdev = 0.026/0.026/0.026/0.000 ms
```

## Access the Management UI

You can access the Management UI by opening `http://localhost:8080` in your browser.
When you do you should see the login screen:

![Login screen](/img/nlx-management-login-screen.png "Login screen")

Clicking on the login button leads you to Dex which acts as an OpenID Connect Identity Provider.
For demo purposes we configured Dex to accept a static username/password but in production you would use your own Identity Provider.

You can login with the demo credentials:

- **Username**: admin@nlx.local
- **Password**: development

After logging in you will be asked to grant access.
Click on "Grant Access" to get access to the Management UI.


## Management UI overview

On the left you will find the main navigation which separates the UI in several pages:

- **Inways**: Lists all available inways.
- **Services**: Shows a list of your services. You can also register new services here.
- **Directory**: Lists all available services in the [demo directory](https://directory.demo.nlx.io/). This is also the place where you can request access to another service.
- **Settings**: Shows all global settings. Currently only the insight and organization inway settings.

![Overview](/img/nlx-management-overview.png "Overview")


## Set the organization inway

In order to receive access requests you have to set a default inway for your organization.
You can do that by going to the settings page, selecting the "Inway-01" and clicking on "Save settings".

![Settings screen](/img/nlx-management-settings-screen.png "Settings screen")


## In sum

So far we have:
- Started all components using docker-compose
- Granted access to the Management UI
- Set a default organization inway

Next up, let's [consume an API](./consume-an-api.md).
