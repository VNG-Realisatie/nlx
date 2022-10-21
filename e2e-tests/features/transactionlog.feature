@transactionlog
Feature: Transaction Log

    Scenario: Transaction Log outgoing entry
        Given "Gemeente Stijns" is up and running
            And "RvRD" is up and running
            And the Outway "gemeente-stijns-nlx-outway" of "Gemeente Stijns" has access to "basisregister-fictieve-kentekens" of "RvRD"
        When the Outway "gemeente-stijns-nlx-outway" of "Gemeente Stijns" calls the service "basisregister-fictieve-kentekens" from "RvRD" with valid authorization
        Then "Gemeente Stijns" sees an outgoing transaction log entry for the service "basisregister-fictieve-kentekens" of "RvRD"

    Scenario: Transaction Log incoming entry
        Given "Gemeente Stijns" is up and running
            And "RvRD" is up and running
            And the Outway "gemeente-stijns-nlx-outway" of "Gemeente Stijns" has access to "basisregister-fictieve-kentekens" of "RvRD"
        When the Outway "gemeente-stijns-nlx-outway" of "Gemeente Stijns" calls the service "basisregister-fictieve-kentekens" from "RvRD" with valid authorization
        Then "RvRD" sees an incoming transaction log entry from "Gemeente Stijns" for the service "basisregister-fictieve-kentekens"

    Scenario: Transaction Log outgoing delegation entry
        Given "Vergunningsoftware BV" is up and running
            And "RvRD" is up and running
            And "Gemeente Stijns" is up and running
            And the Outway "gemeente-stijns-nlx-outway" of "Gemeente Stijns" has access to "basisregister-fictieve-kentekens" of "RvRD"
            And "Vergunningsoftware BV" has an active order for Outway "vergunningsoftware-bv-nlx-outway" with reference "order-ref-1" from "Gemeente Stijns" for service "basisregister-fictieve-kentekens" of "RvRD" via Outway "gemeente-stijns-nlx-outway"
        When the Outway "vergunningsoftware-bv-nlx-outway" of "Vergunningsoftware BV" calls the service "basisregister-fictieve-kentekens" of "RvRD" with valid authorization details via the order of "Gemeente Stijns" with reference "order-ref-1"
        Then "Vergunningsoftware BV" sees an outgoing delegation transaction log entry for the service "basisregister-fictieve-kentekens" of "RvRD" on behalf of "Gemeente Stijns"

    Scenario: Transaction Log incoming delegation entry
        Given "Vergunningsoftware BV" is up and running
            And "RvRD" is up and running
            And "Gemeente Stijns" is up and running
            And the Outway "gemeente-stijns-nlx-outway" of "Gemeente Stijns" has access to "basisregister-fictieve-kentekens" of "RvRD"
            And "Vergunningsoftware BV" has an active order for Outway "vergunningsoftware-bv-nlx-outway" with reference "order-ref-1" from "Gemeente Stijns" for service "basisregister-fictieve-kentekens" of "RvRD" via Outway "gemeente-stijns-nlx-outway"
        When the Outway "vergunningsoftware-bv-nlx-outway" of "Vergunningsoftware BV" calls the service "basisregister-fictieve-kentekens" of "RvRD" with valid authorization details via the order of "Gemeente Stijns" with reference "order-ref-1"
        Then "RvRD" sees an incoming delegation transaction log entry from "Vergunningsoftware BV" for the service "basisregister-fictieve-kentekens" on behalf of "Gemeente Stijns"
