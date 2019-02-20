import React from 'react'
import { func, string, arrayOf, shape } from 'prop-types'
import Table from '../Table'
import ServicesTableRow from './ServicesTableRow'

const ServicesTable = ({ services, sortBy, sortOrder, onToggleSorting, ...props }) =>
  <Table {...props}>
    <Table.Head>
      <Table.Row>
        <Table.HeadCell style={({ width: '56px' })} align="center" />
        <Table.SortableHeadCell
                               direction={sortBy === 'organization' ? sortOrder : null}
                               onClick={() => onToggleSorting('organization')}>
          Organization
        </Table.SortableHeadCell>
        <Table.SortableHeadCell direction={sortBy === 'name' ? sortOrder : null}
                               onClick={() => onToggleSorting('name')}>
          Service
        </Table.SortableHeadCell>
        <Table.HeadCell align="right">API type</Table.HeadCell>
        <Table.HeadCell style={({ width: '48px' })}/>
        <Table.HeadCell style={({ width: '48px' })}/>
      </Table.Row>
    </Table.Head>

    <Table.Body>
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
    </Table.Body>
  </Table>

ServicesTable.propTypes = {
  services: arrayOf(shape({
    status: string,
    organization: string,
    name: string,
    apiType: string,
    apiAddress: string
  })),
  sortBy: string,
  sortOrder: string,
  onToggleSorting: func
}

ServicesTable.defaultProps = {
  services: [],
  onToggleSorting: () => {}
}

export default ServicesTable
