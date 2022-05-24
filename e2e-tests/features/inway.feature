@inway
Feature: inway

    @ignore
    Scenario: Delete an inway
        Given "Gemeente Stijns" is logged in to NLX management
            And "Gemeente stijns" has an inway named "gemeente-stijns-nlx-inway" which offer the service "parkeerrechten"
        When "Gemeente Stijns" removes the the inway "gemeente-stijns-nlx-inway"
        Then the inway is no longer visible in the inway overview of the management interface of "Gemeente Stijns"

    @execution:serial
    Scenario: Unset organisation inway
        Given "Gemeente Stijns" is up and running
        When "Gemeente Stijns" unsets its organization inway
        Then the default inway of "Gemeente Stijns" is no longer the organization inway
