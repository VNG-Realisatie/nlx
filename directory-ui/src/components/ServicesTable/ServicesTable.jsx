import React from 'react'
import { Table, TableBody, TableRow, TableHead, TableHeadCell, TableBodyCell } from '../Table'
import StatusIcon from "./StatusIcon";

export const ServicesTable = () => {
  return (
    <Table>
      <TableHead>
        <TableRow>
          <TableHeadCell style={({width: '100px'})} align="center">Status</TableHeadCell>
          <TableHeadCell>Organization</TableHeadCell>
          <TableHeadCell>Service</TableHeadCell>
          <TableHeadCell>API type</TableHeadCell>
          <TableHeadCell align="center">API address</TableHeadCell>
        </TableRow>
      </TableHead>

      <TableBody>
        <TableRow>
          <TableBodyCell align="center"><StatusIcon status="online" /></TableBodyCell>
          <TableBodyCell>Bitkode</TableBodyCell>
          <TableBodyCell>DemoService</TableBodyCell>
          <TableBodyCell>OpenAPI2</TableBodyCell>
          <TableBodyCell align="center">Icon</TableBodyCell>
        </TableRow>
        <TableRow>
          <TableBodyCell align="center"><StatusIcon status="offline" /></TableBodyCell>
          <TableBodyCell>Bitkode</TableBodyCell>
          <TableBodyCell>DemoService</TableBodyCell>
          <TableBodyCell>OpenAPI2</TableBodyCell>
          <TableBodyCell align="center">Icon</TableBodyCell>
        </TableRow>
        <TableRow>
          <TableBodyCell align="center"><StatusIcon status="online" /></TableBodyCell>
          <TableBodyCell>Bitkode</TableBodyCell>
          <TableBodyCell>DemoService</TableBodyCell>
          <TableBodyCell>OpenAPI2</TableBodyCell>
          <TableBodyCell align="center">Icon</TableBodyCell>
        </TableRow>
        <TableRow>
          <TableBodyCell align="center"><StatusIcon status="online" /></TableBodyCell>
          <TableBodyCell>Bitkode</TableBodyCell>
          <TableBodyCell>DemoService</TableBodyCell>
          <TableBodyCell>OpenAPI2</TableBodyCell>
          <TableBodyCell align="center">Icon</TableBodyCell>
        </TableRow>
      </TableBody>
    </Table>
  )
}
