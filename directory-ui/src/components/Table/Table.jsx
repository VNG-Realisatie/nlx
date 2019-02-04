import { node } from 'prop-types'
import styled from 'styled-components'

const Table = styled.table`
    border-collapse: collapse;
    border-spacing: 0;
    display: table;
    width: 100%;
`

Table.propTypes = {
  children: node,
}

export default Table
