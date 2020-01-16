// Copyright © VNG Realisatie 2019
// Licensed under the EUPL

import React, { useEffect, useState } from 'react'
import { ThemeProvider } from 'styled-components'

import theme from './theme'
import GlobalStyles from './components/GlobalStyles'

import AuthenticatedApp from './AuthenticatedApp'
import UnauthenticatedApp from './UnauthenticatedApp'

const LOGIN_URL = '/api/auth/login'
const LOGOUT_URL = '/api/auth/logout'
const App = () => {
    const [auth, setAuth] = useState(null)
    const [csrfToken, setCsrfToken] = useState(null)

    useEffect(() => {
        const fetchToken = async () => {
            const result = await fetch(LOGIN_URL)
            if (result.ok) {
                setCsrfToken(
                    result.ok ? result.headers.get('X-CSRF-Token') : null,
                )
                try {
                    const data = await result.json()
                    setAuth(data)
                } catch (e) {
                    setAuth(null)
                }
            }
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

    const logout = () => {
        const submitLogout = async () => {
            const result = await fetch(LOGOUT_URL)
            if (result.ok) {
                setAuth(null)
            }
        }
        submitLogout()
    }

    // TODO handle logout
    return (
        <ThemeProvider theme={theme}>
            <GlobalStyles />
            {auth ? (
                <AuthenticatedApp logout={logout} />
            ) : (
                <UnauthenticatedApp login={login} />
            )}
        </ThemeProvider>
    )
}

export default App
