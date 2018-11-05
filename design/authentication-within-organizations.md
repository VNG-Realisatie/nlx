# Authentication within organizations

## Introduction
NLX provides a generic way of handling authentication of API-requests between organizations. This is done by using mutual TLS and leveraging specific fields within the certificate. In this way a service is able to identity the organization that performs a request. Just like NLX prescribes a pattern for the authentication of traffic between organizations, we also would like to prescribe a pattern for handling authentication of services offered within organizations. Authentication within organizations differs from authentication between organizations. While traffic between organizations is authenticated on an organizational level, authentication within organizations is done on a user level. This imposes different requirements on the required solution.

One of the largest benefits of having authentication logic for NLX on user-level as well in NLX, is that not every application has to have its own authentication code. Also performing security audits on the infrastructure of the whole organization becomes easier.

## The roles of applications, services and NLX
NLX arises from the [Common Ground vision](https://github.com/VNG-Realisatie/common-ground). The Common Ground vision promotes an API-first approach where services initially offer a programmable interface (API). Applications are then build on top of this API. These applications offer a more friendly interface to users and implement functionality for a specific application. It should also be possible for power-users to implement their own application on top of the API's by talking directly to the NLX outway.

This pattern results in a separation of concerns between applications and services. Services are primarily responsible for holding the data (state) and providing authorization mechanisms to make sure a requester (organization or user) is permitted to request a specific resource. Applications' primary concern is to offer a friendly interface to users to view and mutate data. This makes an application stateless.

## A pluggable authentication service for outways
Organizations have different mechanisms for user authentication. Therefore NLX offers a pluggable authentication interface that is easily extendable to fit the specific requirements of an organization. When the outway is configured to use an external authentication service, the headers and method of every request is routed to the authentication service. The outway only allows the request to a service when the authentication service returns a HTTP 200 status code. When a request is not allowed, for example because a user is not authenticted, the service should return a HTTP 403 status code. When the NLX outway is not able to reach the authentication service, the outway returns a HTTP 503 status code to the original requester and the API request is not served.

The external authentication service of an organization is also responsible for adding user-specific metadata (claims) to the request that can be used by services on the network to perform further authorization. This metadata can exist of certain claims, for example a unique username or organization-specific roles. The unique user attribute can be set using the `X-NLX-Requester-User` header. Other claims can be made in a JSON serialized header `X-NLX-Requester-Claims`.

Note that user-specific headers are only available on traffic that happens within organizations. When a request is made between organizations, the request is first logged on the outway of the requesting organization (including the user-specific metadata). Then the user-specific data is stripped from the request and finally the request travels further through the network.

## Generic authorization features on outways and inways
Services are mainly responsible for handling authorization on the NLX network. Depending on the specific service requirements they implement organization- or user-based authorization rules.

NLX also provides some generic authorization features that can be used to apply defense in depth:

- The NLX outway only allows requests to be made when they are authenticated. When a user is not authenticated the request will be denied.
- The NLX inway allows a per-service configuration. The administrator can choose to configure a service to be accessible only from within the organization (default) or make a service accessible to the NLX network and publish it in the central directory.
