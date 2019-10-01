// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React, { PureComponent } from 'react'
import { bool, string, array, func } from 'prop-types'
import ServicesTable from '../../components/ServicesTable/ServicesTable';
import { ASCENDING, DESCENDING } from '../../components/Table/SortableHeadCell';

class ServicesTableContainer extends PureComponent {
  constructor(props) {
    super(props)

    this.state = {
      sortBy: props.sortBy,
      sortOrder: props.sortOrder,
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
    const { services, filterQuery, filterByOnlineServices, onServiceClickedHandler } = this.props
    const { sortBy, sortOrder } = this.state
    const filteredServices = this.filterServices(services, filterQuery, filterByOnlineServices)
    const sortedFilteredServices = this.sortServices(filteredServices, sortBy, sortOrder)

    return (
      <ServicesTable services={sortedFilteredServices}
                     sortBy={sortBy}
                     sortOrder={sortOrder}
                     onToggleSorting={property => this.onToggleSorting(property)}
                     onServiceClickedHandler={onServiceClickedHandler}
      />
    )
  }
}

ServicesTableContainer.propTypes = {
  filterQuery: string,
  filterByOnlineServices: bool,
  services: array,
  sortBy: string,
  sortOrder: string,
  onServiceClickedHandler: func,
}

ServicesTableContainer.defaultProps = {
  filterQuery: '',
  filterByOnlineServices: false,
  services: [],
  sortBy: null,
  sortOrder: null,
  onServiceClickedHandler: () => {},
}

export default ServicesTableContainer
