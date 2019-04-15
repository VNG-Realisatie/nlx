import React, { Component } from 'react'
import { arrayOf, shape, string, func } from 'prop-types'
import {connect} from 'react-redux'

import {fetchOrganizationsRequest } from '../../store/actions'
import Sidebar from '../../components/Sidebar'

export class SidebarContainer extends Component {
  componentWillMount() {
    this.props.fetchOrganizationsRequest()

    this.setState({
      query: ''
    })

    this.onSearchQueryChanged = this.onSearchQueryChanged.bind(this)
  }

  onSearchQueryChanged(query) {
    this.setState({ query })
  }

  getOrganizationsForSidebar() {
    const { organizations } = this.props
    const { query } = this.state

    return organizations
      .map(organization => organization.name)
      .filter(organization => organization.includes(query.toLowerCase()))
  }

  render() {
    return <Sidebar onSearchQueryChanged={this.onSearchQueryChanged}
                    organizations={this.getOrganizationsForSidebar()} />
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
