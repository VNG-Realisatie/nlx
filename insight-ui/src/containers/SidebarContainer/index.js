import React, { Component } from 'react'
import { arrayOf, shape, string } from 'prop-types'
import {connect} from 'react-redux'

import {fetchOrganizationsRequest } from '../../store/actions'
import Sidebar from '../../components/Sidebar'

export class SidebarContainer extends Component {
  componentWillMount() {
    this.props.fetchOrganizationsRequest()
  }

  render() {
    const { organizations } = this.props
    return <Sidebar organizations={organizations.map(organization => organization.name)} />
  }
}

SidebarContainer.propTypes = {
  organizations: arrayOf(shape({
    name: string.isRequired
  }))
}

SidebarContainer.defaultProps = {
  organizations: []
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
