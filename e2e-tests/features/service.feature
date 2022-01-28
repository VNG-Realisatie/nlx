@management @service
Feature: Service

    Scenario: Create a service
        Given "Gemeente Stijns" is logged into NLX management
            And "Gemeente Stijns" has accepted the Terms of Service
            And "Gemeente Stijns" has the default Inway running
            And "Gemeente Stijns" has set its default Inway as organization Inway
        When "Gemeente Stijns" create a service named "MyService" and exposed via the default Inway
        Then the service "MyService" of "Gemeente Stijns" is created

    @ignore
    Scenario: Revoke access to the service
        Given "RvRD" is logged into NLX management
            And "RvRD" has given access to "Gemeenste Stijns" for the service "voorbeeld-websockets"
            And the websocket chat of "Gemeente Stijns" can establish a connection
        When "RvRD" revokes access of "Gemeenste Stijns" to "voorbeeld-websockets"
        Then the websocket chat of "Gemeente Stijns" cannot establish a connection

    @ignore
    Scenario: Delete a service
        Given "Gemeente Stijns" is logged into NLX management
            And "Gemeente Stijns" offers the service "parkeerrechten" using the inway "gemeente-stijns-nlx-inway"
        When "Gemeente Stijns" removes the service "parkeerrechten"
        Then the service "parkeerrechten" is no longer available under My Services of the management interface of "Gemeente Stijns"

