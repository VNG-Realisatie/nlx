---
id: access-api
title: Access API through a client
---

# 8. Access API through a client

## Set port forwarding

You have installed the NLX Outway and the example service.
The Outway is currently only available from within the cluster.
In order to make a request to our API via the outway,
the outway must be available for your local machine. For this, you create a port forward.

Before we can create a port forward we need the name of the Outway pod.
You can find it by running the following command:

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

In order to access your service, you need to request access to that service.

Requesting and getting access via NLX:

1. Navigate to the 'Directory' in NLX Management.
2. Select the service `basisregister-fictieve-kentekens` from the list
3. Expand the 'Outways zonder toegang' section by clicking on its title
4. Click on 'Toegang aanvragen' for the Outway 'outway-nlx-outway'
5. Now navigate to back the 'Services' page and again select the service `SwaggerPetStore`.
6. You should see one access request under the section 'Toegangsverzoeken'.
7. Expand the section and click on 'Accepteren' to accept the access request.
8. You now have access to the service.

You now have access to the service via NLX.

## Query API

Due to the port forwarding, your Outway is now accessible from your local machine.

Replace in the command below:
- `<your subject serial number>` for the certificate's serial number you entered in your external certificate. How to find this value is described in [Create Certificates](./create-certificate)

```
curl http://localhost:8080/<your subject serial number>/basisregister-fictieve-kentekens/voertuigen
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
