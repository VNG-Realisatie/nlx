import React, { Component } from 'react'
import { bool, shape, string } from 'prop-types'
import {connect} from 'react-redux'

import { resetLoginInformation, fetchIrmaLoginInformationRequest, IRMA_LOGIN_STATUS_DONE } from '../../store/actions'
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
  
  componentWillMount() {
    this.props.resetLoginInformation()
  }

  componentDidUpdate(prevProps) {
    const { organization, loginStatus, proof, history } = this.props
    const { organization: prevOrganization } = prevProps 

    if (loginStatus && loginStatus.response === IRMA_LOGIN_STATUS_DONE && proof && proof.loaded) {
      history.push(`/organization/${organization.name}/logs`)
      return
    }

    if (organization === prevOrganization) {
      return
    }

    this.fetchIrmaLoginInformation(organization)
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
  }).isRequired,
  loginRequestInfo: shape({
    qrCodeValue: string
  }),
  loginStatus: shape({
    error: bool.isRequired,
    response: string.isRequired
  }),
  proof: shape({
    loaded: bool,
    error: bool,
    value: string,
    message: string
  })
}

const mapStateToProps = ({ loginRequestInfo, loginStatus, proof }) => {
  return {
    loginRequestInfo,
    loginStatus,
    proof
  }
}

const mapDispatchToProps = dispatch => ({
  fetchIrmaLoginInformation: ({ insight_log_endpoint, insight_irma_endpoint }) =>
    dispatch(fetchIrmaLoginInformationRequest({ insight_log_endpoint, insight_irma_endpoint })),
  resetLoginInformation: () =>
    dispatch(resetLoginInformation())
})

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(LoginPageContainer)
