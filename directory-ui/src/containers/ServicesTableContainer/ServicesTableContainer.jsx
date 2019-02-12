import React, { Component } from 'react'
import { bool, string, array } from 'prop-types'
import ServicesTable from '../../components/ServicesTable/ServicesTable';
import { ASCENDING, DESCENDING } from '../../components/Table/SortableTableHeadCell';

class ServicesTableContainer extends Component {
  constructor(props) {
    super(props)

    this.state = {
      sortBy: null,
      sortOrder: null,
    }
  }

  onToggleSorting(property) {
    const { sortOrder } = this.state

    const direction = sortOrder === null ? ASCENDING :
      sortOrder === ASCENDING ? DESCENDING : ASCENDING

    this.setState({
      sortBy: property,
      sortOrder: direction
    })
  }

  filterServicesByOnlineStatus(services) {
    return services.filter(service => service.status === 'online')
  }

  filterServicesByQuery(services, query) {
    return services
      .filter(service =>
        service.organization.toLowerCase().includes(query) ||
        service.name.toLowerCase().includes(query)
      )
  }

  filterServices(services, query, filterByOnlineServices) {
    let result

    result = filterByOnlineServices ?
      this.filterServicesByOnlineStatus(services) :
      services

    result = query ?
      this.filterServicesByQuery(result, query) :
      result

    return result
  }

  sortServices(services, sortBy, sortOrder) {
    if (!sortBy) {
      return services
    }

    let result

    result = services
      .sort((a, b) => {
        const aValue = a[sortBy].toLowerCase()
        const bValue = b[sortBy].toLowerCase()
        return aValue > bValue ? 1 : aValue < bValue ? -1 : 0
      })

    if (sortOrder === DESCENDING) {
      result.reverse()
    }

    return result
  }

  render() {
    const { services, filterQuery, filterByOnlineServices } = this.props
    const { sortBy, sortOrder } = this.state
    const filteredServices = this.filterServices(services, filterQuery, filterByOnlineServices)
    const sortedFilteredServices = this.sortServices(filteredServices, sortBy, sortOrder)

    return (
      <ServicesTable services={sortedFilteredServices}
                     sortBy={sortBy}
                     sortOrder={sortOrder}
                     onToggleSorting={property => this.onToggleSorting(property)}
      />
    )
  }
}

ServicesTableContainer.propTypes = {
  filterQuery: string,
  filterByOnlineServices: bool,
  services: array
}

ServicesTableContainer.defaultProps = {
  filterQuery: '',
  filterByOnlineServices: false,
  services: []
}

export default ServicesTableContainer
