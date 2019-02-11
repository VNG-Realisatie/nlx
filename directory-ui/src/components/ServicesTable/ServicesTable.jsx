import React from 'react'
import { Table, TableBody, TableRow, TableHead, TableHeadCell } from '../Table'
import ServicesTableRow from "./ServicesTableRow";

const services = [
  { status: 'online', organization: 'Bitkode', name: 'DemoService', apiType: 'OpenAPI2', apiAddress: 'Icon' },
  { status: 'online', organization: 'Bitkode', name: 'DemoService', apiType: 'OpenAPI2', apiAddress: 'Icon' },
  { status: 'offline', organization: 'Bitkode', name: 'DemoService', apiType: 'OpenAPI2', apiAddress: 'Icon' },
  { status: 'online', organization: 'Bitkode', name: 'DemoService', apiType: 'OpenAPI2', apiAddress: 'Icon' }
]

export const ServicesTable = () => {
  return (
    <Table>
      <TableHead>
        <TableRow>
          <TableHeadCell style={({width: '30px'})} align="center" />
          <TableHeadCell>Organization</TableHeadCell>
          <TableHeadCell>Service</TableHeadCell>
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
