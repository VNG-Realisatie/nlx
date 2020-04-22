// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Route, Redirect } from 'react-router-dom'
import PropTypes from 'prop-types'

import UserContext from '../user-context'

const LoginRoutePath = '/login'

class AuthenticatedRoute extends React.Component {
  static propTypes = {
    unauthenticatedPath: PropTypes.string.isRequired,
    children: PropTypes.node,
  }

  static defaultProps = {
    unauthenticatedPath: LoginRoutePath,
  }

  static contextType = UserContext

  render() {
    const { unauthenticatedPath, children, ...rest } = this.props
    const { isReady, user } = this.context

    if (!isReady) {
      return null
    }

    if (user) {
      return <Route {...rest}>{children}</Route>
    }

    return <Redirect to={unauthenticatedPath} />
  }
}

export default AuthenticatedRoute
export { LoginRoutePath }
