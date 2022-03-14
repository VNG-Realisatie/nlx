package httpapi.authz

default allow = false

# Allow users with valid token
allow {
    authValues := split(input.headers["Proxy-Authorization"][0], " ")

    count(authValues) == 2

    authType := authValues[0]
    authToken := authValues[1]

    authType == "Bearer"

    # Loop over object values and look for token match
    # Data object is loaded via users.json
    authToken == data.users[k]
}
