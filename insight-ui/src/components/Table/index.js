import { node } from 'prop-types'
import styled from 'styled-components'

import Body from './Body'
import HeadCell from './HeadCell'
import BodyCell from './BodyCell'
import Head from './Head'
import Row from './Row'

const Table = styled.table`
  border-collapse: separate;
  border-spacing: 0;
  display: table;
  margin: 0 auto;
`

Table.propTypes = {
  children: node,
}

Table.Body = Body
Table.HeadCell = HeadCell
Table.BodyCell = BodyCell
Table.Head = Head
Table.Row = Row

export default Table
