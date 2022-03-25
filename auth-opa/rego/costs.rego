package httpapi.inway

default allow = false

# Allow organizations with enough budget left
allow {
    serialNumber := input.headers["X-Nlx-Request-Organization"][0]

    data.organizations[serialNumber].budget_left >= input.service.request_costs
}
