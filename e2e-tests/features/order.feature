@management @order
Feature: Order

    Scenario: Create an order
        Given "Gemeente Stijns" is up and running
            And "RvRD" is up and running
            And "Vergunningsoftware BV" is up and running
            And the Outway "gemeente-stijns-nlx-outway" of "Gemeente Stijns" has access to "basisregister-fictieve-kentekens" of "RvRD"
        When "Gemeente Stijns" creates an order with reference "order-ref-1" for "Vergunningsoftware BV" including the service "basisregister-fictieve-kentekens" of "RvRD" via Outway "gemeente-stijns-nlx-outway"
        Then an order of "Gemeente Stijns" with reference "order-ref-1" for "Vergunningsoftware BV" with service "basisregister-fictieve-kentekens" of "RvRD" is created

    Scenario: Use an order to access service
        Given "Vergunningsoftware BV" is up and running
            And "RvRD" is up and running
            And "Gemeente Stijns" is up and running
            And the Outway "gemeente-stijns-nlx-outway" of "Gemeente Stijns" has access to "basisregister-fictieve-kentekens" of "RvRD"
            And "Vergunningsoftware BV" has an active order for Outway "vergunningsoftware-bv-nlx-outway" with reference "order-ref-1" from "Gemeente Stijns" for service "basisregister-fictieve-kentekens" of "RvRD" via Outway "gemeente-stijns-nlx-outway"
        When the Outway "vergunningsoftware-bv-nlx-outway" of "Vergunningsoftware BV" calls the service "basisregister-fictieve-kentekens" of "RvRD" with valid authorization details via the order of "Gemeente Stijns" with reference "order-ref-1"
        Then "Vergunningsoftware BV" receives a successful response

    Scenario: Use an order to access service when order is revoked
        Given "Vergunningsoftware BV" is up and running
            And "RvRD" is up and running
            And "Gemeente Stijns" is up and running
            And the Outway "gemeente-stijns-nlx-outway" of "Gemeente Stijns" has access to "basisregister-fictieve-kentekens" of "RvRD"
            And "Vergunningsoftware BV" has a revoked order for Outway "vergunningsoftware-bv-nlx-outway" with reference "order-ref-1" from "Gemeente Stijns" for service "basisregister-fictieve-kentekens" of "RvRD" via Outway "gemeente-stijns-nlx-outway"
        When the Outway "vergunningsoftware-bv-nlx-outway" of "Vergunningsoftware BV" calls the service "basisregister-fictieve-kentekens" of "RvRD" with valid authorization details via the order of "Gemeente Stijns" with reference "order-ref-1"
        Then "Vergunningsoftware BV" receives an order revoked response

    Scenario: Use an order to access a service when access of delegator is revoked
        Given "Vergunningsoftware BV" is up and running
            And "RvRD" is up and running
            And "Gemeente Stijns" is up and running
            And the Outway "gemeente-stijns-nlx-outway" of "Gemeente Stijns" has access to "basisregister-fictieve-kentekens" of "RvRD"
            And "Vergunningsoftware BV" has an active order for Outway "vergunningsoftware-bv-nlx-outway" with reference "order-ref-1" from "Gemeente Stijns" for service "basisregister-fictieve-kentekens" of "RvRD" via Outway "gemeente-stijns-nlx-outway"
            And "RvRD" revokes access of Outway "gemeente-stijns-nlx-outway" from "Gemeente Stijns" to "basisregister-fictieve-kentekens"
        When the Outway "vergunningsoftware-bv-nlx-outway" of "Vergunningsoftware BV" calls the service "basisregister-fictieve-kentekens" of "RvRD" with valid authorization details via the order of "Gemeente Stijns" with reference "order-ref-1"
        Then "Vergunningsoftware BV" receives a delegator does not have access response

    Scenario: Revoke an order
        Given "Gemeente Stijns" is up and running
            And "RvRD" is up and running
            And "Vergunningsoftware BV" is up and running
            And the Outway "gemeente-stijns-nlx-outway" of "Gemeente Stijns" has access to "basisregister-fictieve-kentekens" of "RvRD"
            And "Vergunningsoftware BV" has an active order for Outway "vergunningsoftware-bv-nlx-outway" with reference "order-ref-1" from "Gemeente Stijns" for service "basisregister-fictieve-kentekens" of "RvRD" via Outway "gemeente-stijns-nlx-outway"
        When "Gemeente Stijns" revokes the order "order-ref-1" for "Vergunningsoftware BV"
        Then "Vergunningsoftware BV" has a revoked incoming order from "Gemeente Stijns" with reference "order-ref-1"

    Scenario: Use an order to access service when order is expired
        Given "Vergunningsoftware BV" is up and running
            And "RvRD" is up and running
            And "Gemeente Stijns" is up and running
            And the Outway "gemeente-stijns-nlx-outway" of "Gemeente Stijns" has access to "basisregister-fictieve-kentekens" of "RvRD"
            And "Vergunningsoftware BV" has an expired order for Outway "vergunningsoftware-bv-nlx-outway" with reference "order-ref-1" from "Gemeente Stijns" for service "basisregister-fictieve-kentekens" of "RvRD" via Outway "gemeente-stijns-nlx-outway"
        When the Outway "vergunningsoftware-bv-nlx-outway" of "Vergunningsoftware BV" calls the service "basisregister-fictieve-kentekens" of "RvRD" with valid authorization details via the order of "Gemeente Stijns" with reference "order-ref-1"
        Then "Vergunningsoftware BV" receives an order expired response

    @ignore
    Scenario: Use an order to access service when delegator has no access to service
        Given "Vergunningsoftware BV" has the Outway "vergunningsoftware-bv-nlx-outway" running
            And "Vergunningsoftware BV" has an active order with reference "order-ref-1" from "Gemeente Stijns" for service "basisregister-fictieve-kentekens" of "RvRD"
            And "Gemeente Stijns" has no access to service "basisregister-fictieve-kentekens" of "RvRD"
        When the Outway "vergunningsoftware-bv-nlx-outway" of "Vergunningsoftware BV" calls the service "basisregister-fictieve-kentekens" via the order of "Gemeente Stijns" with reference "order-ref-1"
        Then "Vergunningsoftware BV" receives a error response
