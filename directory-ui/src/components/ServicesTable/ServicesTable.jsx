import React, { Component } from 'react'
import { Table, TableBody, TableRow, TableHead, TableHeadCell, SortableTableHeadCell } from "../Table";
import ServicesTableRow from './ServicesTableRow';
import { ASCENDING, DESCENDING } from './../Table/SortableTableHeadCell';

const services = [
  { status: 'online', organization: 'Bitkode', name: 'DemoService', apiType: 'OpenAPI2', apiAddress: 'Icon' },
  { status: 'online', organization: 'Bitkode', name: 'DemoService', apiType: 'OpenAPI2', apiAddress: 'Icon' },
  { status: 'offline', organization: 'Bitkode', name: 'DemoService', apiType: 'OpenAPI2', apiAddress: 'Icon' },
  { status: 'online', organization: 'Bitkode', name: 'DemoService', apiType: 'OpenAPI2', apiAddress: 'Icon' }
]

class ServicesTable extends Component {
  constructor(props) {
    super(props)

    this.state = {
      sortBy: null,
      sortOrder: null
    }

    this.toggleSorting = this.toggleSorting.bind(this)
  }

  toggleSorting(columnName) {
    const { sortOrder } = this.state

    const property = columnName
    const direction = sortOrder === null ? ASCENDING :
      sortOrder === ASCENDING ? DESCENDING : ASCENDING

    this.setState({
      sortBy: property,
      sortOrder: direction
    })
  }

  render() {
    const { sortBy, sortOrder } = this.state

    return (
      <Table>
        <TableHead>
          <TableRow>
            <TableHeadCell style={({width: '30px'})} align="center" />
            <SortableTableHeadCell direction={sortBy === 'organization' ? sortOrder : null}
                                   onClick={() => this.toggleSorting('organization')}>
              Organization
            </SortableTableHeadCell>
            <SortableTableHeadCell direction={sortBy === 'name' ? sortOrder : null}
                                   onClick={() => this.toggleSorting('name')}>
              Service
            </SortableTableHeadCell>
            <TableHeadCell align="right">API type</TableHeadCell>
            <TableHeadCell/>
          </TableRow>
        </TableHead>

        <TableBody>
          {
            services
              .map((service, i) =>
                <ServicesTableRow key={i}
                                  status={service.status}
                                  name={service.name}
                                  organization={service.organization}
                                  apiType={service.apiType}
                                  apiAddress={service.apiAddress}
                />
              )
          }
        </TableBody>
      </Table>
    )
  }
}

export default ServicesTable
