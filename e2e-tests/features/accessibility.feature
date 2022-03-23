@management @accessibility
Feature: Accessibility

    Scenario: 404 page
        Given "Gemeente Stijns" is up and running
        When "Gemeente Stijns" opens a non-existing page
        Then the page is accessible
