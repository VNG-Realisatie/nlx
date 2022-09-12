@management @service
Feature: Service

    Scenario: Create a service
        Given "Gemeente Stijns" is up and running
        When "Gemeente Stijns" create a service named "MyService" and exposed via the default Inway
        Then the service "MyService" of "Gemeente Stijns" is created

    Scenario: Delete a service
        Given "Gemeente Stijns" is up and running
            And "Gemeente Stijns" has a service named "MyServiceToDelete"
        When "Gemeente Stijns" removes the service "MyServiceToDelete"
        Then the service "MyServiceToDelete" of "Gemeente Stijns" is no longer available

    Scenario: Revoke access to the service
        Given "RvRD" is up and running
            And "Gemeente Stijns" is up and running
            And the Outway "gemeente-stijns-nlx-outway" of "Gemeente Stijns" has access to "basisregister-fictieve-kentekens" of "RvRD"
        When "RvRD" revokes access of "Gemeenste Stijns" to "basisregister-fictieve-kentekens"
        Then the Outway "gemeente-stijns-nlx-outway" of "Gemeente Stijns" no longer has access to the service "basisregister-fictieve-kentekens" from "RvRD"
