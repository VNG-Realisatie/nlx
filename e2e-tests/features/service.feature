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

    @ignore
    Scenario: Revoke access to the service
        Given "RvRD" is logged into NLX management
            And "RvRD" has given access to "Gemeenste Stijns" for the service "voorbeeld-websockets"
            And the websocket chat of "Gemeente Stijns" can establish a connection
        When "RvRD" revokes access of "Gemeenste Stijns" to "voorbeeld-websockets"
        Then the websocket chat of "Gemeente Stijns" cannot establish a connection

