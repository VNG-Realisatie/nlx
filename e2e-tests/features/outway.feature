@outway
Feature: Outway

    Scenario: Access a service using an Outway
        Given "Gemeente Stijns" is logged into NLX management
            And "Gemeente Stijns" has the default Outway running
            And "RvRD" has set its default Inway as organization Inway
            And "Gemeente Stijns" has access to "basis-register-fictieve-kentekens" of "RvRD"
        When the default Outway of "Gemeente Stijns" calls the service "basis-register-fictieve-kentekens" from "RvRD"
        Then "Gemeente Stijns" receives a successful response

    @ignore
    Scenario: Access a service using an outway when having no access
        Given "Gemeente Stijns" has an outway
            And "Gemeente Stijns" has no access to the service "basis-register-fictieve-kentekens" of "RvRD"
        When the outway of "Gemeente Stijns" calls the service "basis-register-fictieve-kentekens" from "RvRD"
        Then "Gemeente Stijns" does not receive a successful response
