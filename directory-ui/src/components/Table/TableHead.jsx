import { node } from 'prop-types'
import styled from 'styled-components'

const TableHead = styled.thead`
    display: table-header-group;
`

TableHead.propTypes = {
  children: node,
}

export default TableHead
