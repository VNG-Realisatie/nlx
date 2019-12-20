// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import React, { useEffect, useState } from 'react'
import { ThemeProvider } from 'styled-components'

import theme from './theme'
import GlobalStyles from './components/GlobalStyles'

import AuthenticatedApp from './AuthenticatedApp'
import UnauthenticatedApp from './UnauthenticatedApp'

const LOGIN_URL = '/api/auth/login'
const App = () => {
    const [auth, setAuth] = useState(null)
    const [csrfToken, setCsrfToken] = useState(null)

    useEffect(() => {
        const fetchToken = async () => {
            const result = await fetch(LOGIN_URL)
            setCsrfToken(result.ok ? result.headers.get('X-CSRF-Token') : null)
        }
        if (!auth) {
            fetchToken()
        }
    }, [auth])

    const login = ({ username, password }) => {
        const submitLogin = async () => {
            const result = await fetch(LOGIN_URL, {
                method: 'POST',
                credentials: 'same-origin',
                headers: {
                    'X-CSRF-Token': csrfToken,
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ username, password }),
            })
            if (result.ok) {
                const body = await result.json()
                setAuth(body)
            }
        }
        submitLogin()
    }

    // TODO handle logout
    return (
        <ThemeProvider theme={theme}>
            <GlobalStyles />
            {auth ? <AuthenticatedApp /> : <UnauthenticatedApp login={login} />}
        </ThemeProvider>
    )
}

export default App
