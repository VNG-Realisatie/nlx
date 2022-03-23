@management @accessibility
Feature: Accessibility

    Scenario: 404 page
        Given "Gemeente Stijns" is up and running
        When "Gemeente Stijns" opens a non-existing page
        Then the page is accessible

    Scenario: Login page
        Given "Gemeente Stijns" is up and running
        When "Gemeente Stijns" opens the login page
        Then the page is accessible
