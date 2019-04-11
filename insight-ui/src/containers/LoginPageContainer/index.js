import React, { Component } from 'react'
import { shape, string } from 'prop-types'
import {connect} from 'react-redux'

import { fetchIrmaLoginInformationRequest } from '../../store/actions'
import LoginPage from '../../components/LoginPage'

export class LoginPageContainer extends Component {
  fetchIrmaLoginInformation(organization) {
    if (!organization) {
      return
    }

    this.props.fetchIrmaLoginInformation({
      insight_irma_endpoint: organization.insight_irma_endpoint,
      insight_log_endpoint: organization.insight_log_endpoint,
    });
  }

  componentWillUpdate(nextProps) {
    const { organization } = nextProps
    const { organization: prevOrganization } = this.props

    if (organization === prevOrganization) {
      return
    }

    this.fetchIrmaLoginInformation(nextProps.organization)
  }

  componentDidMount() {
    const { organization } = this.props

    if (!organization) {
      return
    }

    this.fetchIrmaLoginInformation(organization)
  }

  render() {
    const { loginRequestInfo } = this.props

    return loginRequestInfo && loginRequestInfo.qrCodeValue ?
      <LoginPage qrCodeValue={loginRequestInfo.qrCodeValue} /> :
      null
  }
}

LoginPageContainer.propTypes = {
  organization: shape({
    name: string.isRequired,
    insight_irma_endpoint: string.isRequired,
    insight_log_endpoint: string.isRequired
  }),
  loginRequestInfo: shape({
    qrCodeValue: string
  }),
}

const mapStateToProps = ({ organizations, loginRequestInfo }, ownProps) => {
  const { organizationName } = ownProps.match.params
  return {
    organization: organizations.find(organization => organization.name === organizationName),
    loginRequestInfo: loginRequestInfo,
  }
}

const mapDispatchToProps = dispatch => ({
  fetchIrmaLoginInformation: ({ insight_log_endpoint, insight_irma_endpoint }) =>
    dispatch(fetchIrmaLoginInformationRequest({ insight_log_endpoint, insight_irma_endpoint }))
})

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(LoginPageContainer)
