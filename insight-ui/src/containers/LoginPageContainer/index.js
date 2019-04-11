import React, { Component } from 'react'
import { bool, shape, string } from 'prop-types'
import {connect} from 'react-redux'

import { fetchIrmaLoginInformationRequest, IRMA_LOGIN_STATUS_DONE } from "../../store/actions";
import ScanQRCodePage from '../../components/ScanQRCodePage'
import ErrorPage from '../../components/ErrorPage'

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

  componentWillReceiveProps(nextProps) {
    const { organization, loginStatus } = nextProps
    const { organization: prevOrganization } = this.props

    if (loginStatus && loginStatus.response === IRMA_LOGIN_STATUS_DONE) {
      this.props.history.push(`/organization/${organization.name}/logs`)
    }

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
    const { loginRequestInfo, loginStatus } = this.props

    return loginStatus && loginStatus.error ?
      <ErrorPage title="Failed to load information">
        <p>{loginStatus.response}</p>
      </ErrorPage> :
      loginRequestInfo && loginRequestInfo.qrCodeValue ?
        <ScanQRCodePage qrCodeValue={loginRequestInfo.qrCodeValue} /> :
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
  loginStatus: shape({
    error: bool.isRequired,
    response: string.isRequired
  })
}

const mapStateToProps = ({ organizations, loginRequestInfo, loginStatus }, ownProps) => {
  const { organizationName } = ownProps.match.params
  return {
    organization: organizations.find(organization => organization.name === organizationName),
    loginRequestInfo,
    loginStatus
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
