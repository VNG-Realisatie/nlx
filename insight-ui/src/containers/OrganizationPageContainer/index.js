// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React, { Component } from 'react'
import { Route, Redirect } from 'react-router-dom'
import { shape, string } from 'prop-types'
import { connect } from 'react-redux'

import LoginPageContainer from '../LoginPageContainer'
import LogsPageContainer from '../LogsPageContainer'

export class OrganizationPageContainer extends Component {
  render() {
    const {
      organization,
      match: { url },
    } = this.props
    return organization ? (
      <div>
        <Route
          path={url}
          exact
          render={() => <Redirect to={`${url}/login`} />}
        />
        <Route
          path={`${url}/login`}
          render={(props) => (
            <LoginPageContainer {...props} organization={organization} />
          )}
        />
        <Route
          path={`${url}/logs`}
          render={(props) => (
            <LogsPageContainer {...props} organization={organization} />
          )}
        />
      </div>
    ) : null
  }
}

OrganizationPageContainer.propTypes = {
  organization: shape({
    name: string.isRequired,
    insight_irma_endpoint: string.isRequired, // eslint-disable-line camelcase
    insight_log_endpoint: string.isRequired, // eslint-disable-line camelcase
  }),
  match: shape({ url: string.isRequired }).isRequired,
}

const mapStateToProps = (
  { organizations, loginRequestInfo, loginStatus },
  ownProps,
) => {
  const { organizationName } = ownProps.match.params
  return {
    organization: organizations.find(
      (organization) => organization.name === organizationName,
    ),
  }
}

export default connect(mapStateToProps)(OrganizationPageContainer)
