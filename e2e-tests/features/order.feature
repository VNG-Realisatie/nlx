@management @order
Feature: Order

    Scenario: Create an order
        Given "Gemeente Stijns" is logged into NLX management
            And "Gemeente Stijns" has access to "basis-register-fictieve-kentekens" of "RvRD"
            And "RvRD" has set its default Inway as organization Inway
        When "Gemeente Stijns" creates an order with reference "order-ref-1" for "Vergunningsoftware BV" including the service "basis-register-fictieve-kentekens" of "RvRD"
        Then an order of "Gemeente Stijns" with reference "order-ref-1" for "Vergunningsoftware BV" with service "basis-register-fictieve-kentekens" of "RvRD" is created

    @ignore
    Scenario: Use an order to access service
        Given "Vergunningsoftware BV" has the parkeerrechten admin running
            And "Vergunningsoftware BV" has an active order with reference "order-ref-1" from "Gemeente Stijns" for the following services:
                | Organisation      | Service                           |
                | Gemeente Stijns   | parkeerrechten                    |
                | RvRD              | basis-register-fictieve-kentekens |
                | RvRD              | basis-register-fictieve-personen  |
        When the parkeerrechten admin "Vergunningsoftware BV" adds a parkeerrecht for licence "KN958B" and validity year "2025" with as delegator "Gemeente Stijns" and using order reference "order-ref-1"
        Then "Vergunningsoftware BV" see a entry in the parkeerrechten admin with license "KN958B" and validity year "2025"

    @ignore
    Scenario: Use an order to access service when order is revoked
        Given "Vergunningsoftware BV" has the parkeerrechten admin running
            And "Vergunningsoftware BV" has an revoked order with reference "order-ref-1" from "Gemeente Stijns" for the following services:
                | Organisation      | Service                           |
                | Gemeente Stijns   | parkeerrechten                    |
                | RvRD              | basis-register-fictieve-kentekens |
                | RvRD              | basis-register-fictieve-personen  |
        When the parkeerrechten admin "Vergunningsoftware BV" adds a parkeerrecht for licence "KN958B" and validity year "2025" with as delegator "Gemeente Stijns" and using order reference "order-ref-1"
        Then "Vergunningsoftware BV" gets an error when using the parkeerrechten admin

    @ignore
    Scenario: Use an order to access service when order is expired
        Given "Vergunningsoftware BV" has the parkeerrechten admin running
            And "Vergunningsoftware BV" has an expired order with reference "order-ref-1" from "Gemeente Stijns" for the following services:
                | Organisation      | Service                           |
                | Gemeente Stijns   | parkeerrechten                    |
                | RvRD              | basis-register-fictieve-kentekens |
                | RvRD              | basis-register-fictieve-personen  |
        When the parkeerrechten admin "Vergunningsoftware BV" adds a parkeerrecht for licence "KN958B" and validity year "2025" with as delegator "Gemeente Stijns" and using order reference "order-ref-1"
        Then "Vergunningsoftware BV" gets an error when using the parkeerrechten admin

    @ignore
    Scenario: Use an order to access service when delegator has no access to service
        Given "Vergunningsoftware BV" has the parkeerrechten admin running
            And "Gemeente Stijns" has no access to service "basis-register-fictieve-kentekens" from organization "RvRD"
            And "Vergunningsoftware BV" has an active order with reference "order-ref-1" from "Gemeente Stijns" for the following services:
                | Organisation      | Service                           |
                | Gemeente Stijns   | parkeerrechten                    |
                | RvRD              | basis-register-fictieve-kentekens |
                | RvRD              | basis-register-fictieve-personen  |
        When the parkeerrechten admin "Vergunningsoftware BV" adds a parkeerrecht for licence "KN958B" and validity year "2025" with as delegator "Gemeente Stijns" and using order reference "order-ref-1"
        Then "Vergunningsoftware BV" gets an error when using the parkeerrechten admin

    @ignore
    Scenario: Revoke an order
        Given "Gemeente Stijns" is logged in to NLX management
            And "Gemeente Stijns" has and order for "Vergunningsoftware BV" with as reference "order-ref-1" and the following services:
                | Organisation      | Service                           |
                | Gemeente Stijns   | parkeerrechten                    |
                | RvRD              | basis-register-fictieve-kentekens |
                | RvRD              | basis-register-fictieve-personen  |
            And the parkeerrechten admin of "Vergunningsoftware BV" can retrieve the parkeerrechten using order reference "order-ref-1" and delegator "Gemeente Stijns"
        When "Gemeente Stijns" revokes the order with reference "order-ref-1" for "Vergunningssoftware BV"
        Then the parkeerrechten admin of "Vergunningssoftware BV" can no longer retrieve the parkeerrechtenusing order reference "order-ref-1" and delegator "Gemeente Stijns"
