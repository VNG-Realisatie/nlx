@outway
Feature: Outway

    @ignore
    Scenario: Access a service using an outway
        Given "Gemeente Stijns" has an outway
            And "Gemeente Stijns" has an active access request to the service "basis-register-fictieve-kentekens" of "RvRD" 
        When the outway of "Gemeente Stijns" calls the service "basis-register-fictieve-kentekens" from "RvRD"
        Then "Gemeente Stijns" receives a successful response

    @ignore
    Scenario: Access a service using an outway when having no access
        Given "Gemeente Stijns" has an outway
            And "Gemeente Stijns" has no access to the service "basis-register-fictieve-kentekens" of "RvRD" 
        When the outway of "Gemeente Stijns" calls the service "basis-register-fictieve-kentekens" from "RvRD"
        Then "Gemeente Stijns" does not receive a successful response
        