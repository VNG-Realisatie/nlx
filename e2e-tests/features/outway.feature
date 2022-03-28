@outway
Feature: Outway

    Scenario: Access a service using an Outway
        Given "Gemeente Stijns" is up and running
            And "RvRD" is up and running
            And the Outway "gemeente-stijns-nlx-outway" of "Gemeente Stijns" has access to "basisregister-fictieve-kentekens" of "RvRD"
        When the Outway "gemeente-stijns-nlx-outway" of "Gemeente Stijns" calls the service "basisregister-fictieve-kentekens" from "RvRD"
        Then "Gemeente Stijns" receives a successful response

    Scenario: Access a service using an Outway without access
        Given "Gemeente Stijns" is up and running
            And "RvRD" is up and running
            And "RvRD" has a service named "basisregister-fictieve-kentekens"
        When the Outway "gemeente-stijns-nlx-outway" of "Gemeente Stijns" calls the service "basisregister-fictieve-kentekens" from "RvRD"
        Then "Gemeente Stijns" receives an unauthorized response

