@inway
Feature: inway

    @ignore
    Scenario: Delete an inway
        Given "Gemeente Stijns" is logged in to NLX management
            And "Gemeente stijns" has an inway named "gemeente-stijns-nlx-inway" which offer the service "parkeerrechten"
        When "Gemeente Stijns" removes the the inway "gemeente-stijns-nlx-inway"
        Then the inway is no longer visible in the inway overview of the management interface of "Gemeente Stijns"

    @ignore
    Scenario: Unset organisation inway
        Given "Gemeente Stijns" is logged in to NLX management
            And "Gemeente Stijns" has an inway named "gemeente-stijns-nlx-inway" which is configured as organization inway
            And in the directory the organization inway of "Gemeente Stijns" is set to the inway "gemeente-stijns-nlx-inway"
        When "Gemeente Stijns" unsets its organization inway
        Then the organization inway is no longer set in the management interface of "Gemeente Stijns"
            And the inway "gemeente-stijns-nlx-inway" is no longer the organization inway of "Gemeente Stijns" in the directory
            