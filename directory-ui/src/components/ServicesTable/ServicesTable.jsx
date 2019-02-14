import React from 'react'
import { func, string, arrayOf, shape } from 'prop-types'
import { Table, TableBody, TableRow, TableHead, TableHeadCell, SortableTableHeadCell } from '../Table'
import ServicesTableRow from './ServicesTableRow'

const ServicesTable = ({ services, sortBy, sortOrder, onToggleSorting, ...props }) =>
  <Table {...props}>
    <TableHead>
      <TableRow>
        <TableHeadCell style={({ width: '30px' })} align="center" />
        <SortableTableHeadCell style={({ width: '210px' })}
                               direction={sortBy === 'organization' ? sortOrder : null}
                               onClick={() => onToggleSorting('organization')}>
          Organization
        </SortableTableHeadCell>
        <SortableTableHeadCell direction={sortBy === 'name' ? sortOrder : null}
                               onClick={() => onToggleSorting('name')}>
          Service
        </SortableTableHeadCell>
        <TableHeadCell style={({ width: '120px' })} align="right">API type</TableHeadCell>
        <TableHeadCell style={({ width: '24px' })}/>
        <TableHeadCell style={({ width: '24px' })}/>
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
