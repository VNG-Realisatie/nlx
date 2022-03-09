@outway
Feature: Outway

    Scenario: Access a service using an Outway
        Given "Gemeente Stijns" is up and running
            And "RvRD" is up and running
            And the Outway "gemeente-stijns-nlx-outway" of "Gemeente Stijns" has access to "basis-register-fictieve-kentekens" of "RvRD"
        When the default Outway of "Gemeente Stijns" calls the service "basis-register-fictieve-kentekens" from "RvRD"
        Then "Gemeente Stijns" receives a successful response

    Scenario: Access a service using an Outway without access
        Given "Gemeente Stijns" is up and running
        And "RvRD" is up and running
        And the Outway "gemeente-stijns-nlx-outway" of "Gemeente Stijns" has access to "basis-register-fictieve-kentekens" of "RvRD"
        When the Outway "gemeente-stijns-outway-2" of "Gemeente Stijns" calls the service "basis-register-fictieve-kentekens" from "RvRD"
        Then "Gemeente Stijns" receives a unauthorized response

    @ignore
    Scenario: Access a service using an outway when having no access
        Given "Gemeente Stijns" has an outway
            And "Gemeente Stijns" has no access to the service "basis-register-fictieve-kentekens" of "RvRD"
        When the Outway "gemeente-stijns-nlx-outway" of "Gemeente Stijns" calls the service "basis-register-fictieve-kentekens" from "RvRD"
        Then "Gemeente Stijns" does not receive a successful response