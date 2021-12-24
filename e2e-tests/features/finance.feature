@management @finance
Feature: finance

    @ignore
    Scenario: Export finance to csv
        Given "Gemeente Stijns" is logged in to NLX management
            And "Gemeente Stijns" has an inway named "gemeente-stijns-nlx-inway" which offer the service "parkeerrechten" with transactionlog and costs configured
            And 10 requests have been made to the service "parkeerrechten" service of "Gemeente Stijns"
        When "Gemeente Stijns" exports the finances
        Then a csv export is downloaded and it contains 10 requests
