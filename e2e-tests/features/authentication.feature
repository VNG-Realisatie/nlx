@management @authentication @execution:serial @unauthenticated
Feature: Authentication

    Scenario: Using OpenID Connect with incorrect credentials
        Given "Gemeente Stijns" is logged out of NLX Management
        When "Gemeente Stijns" authenticates using the wrong credentials
        Then the authentication page for "Gemeente Stijns" should display an error

    Scenario: Using OpenID Connect with correct credentials
        Given "Gemeente Stijns" is logged out of NLX Management
        When "Gemeente Stijns" authenticates using the right credentials
        Then the Inways page of "Gemeente Stijns" should be visible

    Scenario: Using Basic Auth with incorrect credentials
        Given "Vergunningsoftware BV" is logged out of NLX Management
        When "Vergunningsoftware BV" authenticates using the wrong credentials
        Then the authentication page for "Vergunningsoftware BV" should display an error

    Scenario: Using Basic Auth with correct credentials
        Given "Vergunningsoftware BV" is logged out of NLX Management
        When "Vergunningsoftware BV" authenticates using the right credentials
        Then the Inways page of "Vergunningsoftware BV" should be visible
