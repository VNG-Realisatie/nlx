import React, { Component } from 'react'
import {connect} from 'react-redux'

import {fetchOrganizationsRequest } from '../../store/actions'
import Sidebar from '../../components/Sidebar'

class SidebarContainer extends Component {
  componentWillMount() {
    this.props.fetchOrganizationsRequest()
  }
  render() {
    const { organizations } = this.props
    return <Sidebar organizations={organizations.map(organization => organization.name)} />
  }
}

const mapStateToProps = ({ organizations }) =>
  ({ organizations })

const mapDispatchToProps = dispatch => ({
  fetchOrganizationsRequest: () => dispatch(fetchOrganizationsRequest())
})

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(SidebarContainer)
