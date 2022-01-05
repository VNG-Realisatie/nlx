// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useContext } from 'react'
import { Navigate } from 'react-router-dom'
import { string, node } from 'prop-types'
import UserContext from '../user-context'

const LoginRoutePath = '/login'

const AuthenticatedRoute = ({ unauthenticatedPath, children }) => {
  const { isReady, user } = useContext(UserContext)

  if (!isReady) {
    return null
  }

  if (user) {
    return children
  }

  return <Navigate to={unauthenticatedPath} />
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
