# Authentication within organizations

## Introduction
NLX provides a generic way of handling authentication of API-requests between organizations. This is done by using mutual TLS and leveraging specific fields within the certificate. In this way a service is able to identity the organization that performs a request. Just like NLX prescribes a pattern for the authentication of traffic between organizations, we also would like to prescribe a pattern for handling authentication of services offered within organizations.

Authentication within organizations differs from authentication between organizations. While traffic between organizations is authenticated on an organizational level, authentication within organizations is done on a user level. This imposes different requirements on the required solution.

## The roles of applications, services and NLX
NLX arises from the [Common Ground vision](https://github.com/VNG-Realisatie/common-ground). The Common Ground vision promotes an API-first approach where services initially offer a programmable interface (API). Applications are then build on top of this API. These applications offer a more friendly interface to users and implement functionality for a specific application. It should also be possible for power-users to implement their own application on top of the API's.

This pattern results in a separation of concerns between applications and services. Services are primarily responsible for holding the data (state) and providing authorization mechanisms to make sure a requestor (organization or user) is permitted to request a specific resource. Applications' primary concern is to offer a friendly interface to users to view and mutate data. This makes an application stateless.

## A pluggable authentication service for outways
Organizations have different mechanisms for user authentication. Therefore NLX offers a pluggable authentication interface that is easily extendable to fit the specific requirements of an organization. When the outway is configured to use an external authentication service, the authentication header of every request is routed to the authentication service