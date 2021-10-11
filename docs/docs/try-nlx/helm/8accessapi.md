---
id: 8accessapi
title: 8. Access API through a client
---

# 8. Access API through a client

## Set port forwarding

You have installed the NLX outway and the example service. The outway is currently only available from within the cluster. In order to make a request to our API via the outway, the outway must be available for your local machine. For this, you create a port forward.

Before we can create a port forward we need the name of the outway pod. You can find it by running the following command:

```
kubectl -n nlx get pods
```

You will now see a list containing a line that starts with `outway-nlx-outway-xxxxxxxxxxxx`. Copy that whole name.

Now open a new terminal window and use the following command:
- In the command below, replace `<name of your outway pod>` with the name you just copied and run the command

```
kubectl -n nlx port-forward pod/<name of your outway pod> 8080:8080
```

## Request access to the service

In order to access your service, you need to access that service. After all, you exploring NLX and therefore manage access to the service via NLX.

Requesting and getting access via NLX is very easy:
- Open the directory in NLX Management
- Select the service `basisregister-fictieve-kentekens`.
- In the detail screen that opens, you will see the yellow button `Request access`, click on this.
- Wait a moment and now refresh your browser a few times until you see the text "access requested".
- Now you have to accept the access request by opening the service page and clicking on `basisregister-fictieve-kentekens`
- In the detail screen that opens you will see 'Access requests'.
- Click here to see access requests. If all went well you will see your own. If not, then refresh your browser again.
- You will now see your access request. Accept it by clicking the blue checkmark. If this is successful, your organization should be visible under the heading "Organizations with access".

You now have access to the service via NLX.

## Query API

Due to the port forwarding, your outway is now accessible from your local machine.

Replace in the command below:
- `<your organization>` for the organization name you entered in your external certificate.

```
curl http://localhost:8080/<your organization>/basisregister-fictieve-kentekens/voertuigen
```

Now run the command to query the API via the NLX Outway.

After running the command you should see the following result:

```json
{
  "aantal":6,
  "resultaten":
    [
      {"burgerservicenummer":"663678651","datum_tenaamstelling":"29-01-2018","eerste_kleur":"GRIJS","europese_voertuigcategorie":"M1","handelsbenaming":"MAZDA 3","kenteken":"RT774D","merk":"MAZDA","voertuigsoort":"Personenauto"},
      {"burgerservicenummer":"425749708","datum_tenaamstelling":"16-11-2016","eerste_kleur":"GRIJS","europese_voertuigcategorie":"M1","handelsbenaming":"TOYOTA YARIS HYBRID","kenteken":"KN958B","merk":"TOYOTA","voertuigsoort":"Personenauto"},
      {"burgerservicenummer":"248496086","datum_tenaamstelling":"17-04-2009","eerste_kleur":"GRIJS","europese_voertuigcategorie":"M1","handelsbenaming":"CORSA-C","kenteken":"81HZFB","merk":"OPEL","voertuigsoort":"Personenauto"},
      {"burgerservicenummer":"187788248","datum_tenaamstelling":"15-06-2015","eerste_kleur":"GRIJS","europese_voertuigcategorie":"M1","handelsbenaming":"HEARSE","kenteken":"GJ713R","merk":"CADILLAC","voertuigsoort":"Personenauto"},
      {"burgerservicenummer":"581536630","datum_tenaamstelling":"13-07-2017","eerste_kleur":"BLAUW","europese_voertuigcategorie":"M1","handelsbenaming":"AGILA","kenteken":"50HSZS","merk":"OPEL","voertuigsoort":"Personenauto"},
      {"burgerservicenummer":"750461986","datum_tenaamstelling":"06-11-2000","eerste_kleur":"BLAUW","europese_voertuigcategorie":"M1","handelsbenaming":"SOVEREIGN HE","kenteken":"KS98DN","merk":"JAGUAR","voertuigsoort":"Personenauto"}
    ]
}
```