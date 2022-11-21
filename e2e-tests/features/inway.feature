@inway
Feature: inway

    @execution:serial
    Scenario: Unset organisation inway
        Given "Gemeente Stijns" is up and running
        When "Gemeente Stijns" unsets its organization inway
        Then the default inway of "Gemeente Stijns" is no longer the organization inway
