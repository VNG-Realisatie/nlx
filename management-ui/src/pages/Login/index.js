// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import React from 'react'
import PageTemplate from '../../components/PageTemplate'
import LoginForm from '../../components/LoginForm'
import { func } from 'prop-types'

const Login = ({ login }) => (
    <PageTemplate>
        <LoginForm login={login} />
    </PageTemplate>
)
Login.propTypes = {
    login: func.isRequired,
}

export default Login
