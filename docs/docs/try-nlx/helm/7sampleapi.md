---
id: 7sampleapi
title: 7. Install sample API
---

# 7. Install sample API

We are now ready to install a service via the new NLX Inway. We have an example API available for this that we are going to download and install: `basisregister-fictieve-kentekens`.

Download and install the Helm chart from the basisregister-fictieve-kentekens by running:


```
git clone https://gitlab.com/commonground/nlx/nlx.git helm/deploy/rvrd/charts/basisregister-fictieve-kentekens

helm -n nlx upgrade --install brfk helm/deploy/rvrd/charts/basisregister-fictieve-kentekens
```

Check with `kubectl -n nlx get pods` if the Service pod is healthy.

A similar line should now be displayed:

```
brfk-basisregister-fictieve-kentekens-6854dd6f86-rlqgk   1/1     Running     0          49s
```

## Make example API available through NLX

Now we open NLX Management and go to the Services page. Here we add a service by clicking the yellow `Service toevoegen` button. Enter the following values and save the service:

- Servicenaam: basisregister-fictieve-kentekens
- API endpoint URL: http://brfk-basisregister-fictieve-kentekens
    - This URL is not publicly available but can be reached by your inway.
- Inways: Check your Inway here

Now wait a minute and check if you can find this service in the directory in your NLX management.