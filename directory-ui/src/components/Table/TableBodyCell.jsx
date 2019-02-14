import { node, oneOf } from 'prop-types'
import styled from 'styled-components'

const TableBodyCell = styled.td`
    display: table-cell;
    padding: .6rem;

    font-size: 1rem;
    line-height: 1rem;
    font-weight: 400;
    
    text-align: ${p => p.align};
`

TableBodyCell.propTypes = {
  children: node,
  align: oneOf(['left', 'center', 'right']).isRequired,
}

TableBodyCell.defaultProps = {
  align: 'left',
}

export default TableBodyCell
