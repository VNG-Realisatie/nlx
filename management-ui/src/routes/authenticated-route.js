// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useContext } from 'react'
import { Redirect, Route } from 'react-router-dom'
import { string, node } from 'prop-types'
import UserContext from '../user-context'

const LoginRoutePath = '/login'

const AuthenticatedRoute = ({ unauthenticatedPath, children, ...props }) => {
  const { isReady, user } = useContext(UserContext)

  if (!isReady) {
    return null
  }

  if (user) {
    return <Route {...props}>{children}</Route>
  }

  return <Redirect to={unauthenticatedPath} />
}

AuthenticatedRoute.propTypes = {
  unauthenticatedPath: string.isRequired,
  children: node,
}

AuthenticatedRoute.defaultProps = {
  unauthenticatedPath: LoginRoutePath,
}

export default AuthenticatedRoute
export { LoginRoutePath }
