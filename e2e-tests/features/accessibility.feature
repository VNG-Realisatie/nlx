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

    Scenario: Directory page
        Given "Gemeente Stijns" is up and running
        When "Gemeente Stijns" opens the directory page
        Then the page is accessible

    Scenario: Inways page
        Given "Gemeente Stijns" is up and running
        When "Gemeente Stijns" opens the inways page
        Then the page is accessible

    Scenario: Inways detail page
        Given "Gemeente Stijns" is up and running
        When "Gemeente Stijns" opens the inway detail page of the default inway
        Then the page is accessible with the tabindex disabled

    Scenario: Services page
        Given "Gemeente Stijns" is up and running
        When "Gemeente Stijns" opens the services page
        Then the page is accessible
