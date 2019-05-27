// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React, { Component } from 'react'
import { arrayOf, shape, string, func } from 'prop-types'
import {connect} from 'react-redux'

import {fetchOrganizationsRequest } from '../../store/actions'
import Sidebar from '../../components/Sidebar'

export class SidebarContainer extends Component {
  constructor(props) {
    super(props)

    this.state = {
      query: ''
    }

    this.onSearchQueryChanged = this.onSearchQueryChanged.bind(this)
  }

  componentWillMount() {
    this.props.fetchOrganizationsRequest()
  }

  onSearchQueryChanged(query) {
    this.setState({ query })
  }

  getFilteredOrganizationsByQuery(organizations, query = '') {
    return organizations
      .map(organization => organization.name)
      .filter(organization => organization.includes(query.toLowerCase()))
  }

  render() {
    const { organizations } = this.props
    const { query } = this.state
    return <Sidebar onSearchQueryChanged={this.onSearchQueryChanged}
                    organizations={this.getFilteredOrganizationsByQuery(organizations, query)} />
  }
}

SidebarContainer.propTypes = {
  organizations: arrayOf(shape({
    name: string.isRequired
  })),
  fetchOrganizationsRequest: func
}

SidebarContainer.defaultProps = {
  organizations: [],
  fetchOrganizationsRequest: () => {}
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
