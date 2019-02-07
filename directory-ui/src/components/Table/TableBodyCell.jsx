import { node, oneOf } from 'prop-types'
import styled from 'styled-components'

const TableBodyCell = styled.td`
    display: table-cell;
    padding: .65rem;
    border: 2px solid transparent;

    font-size: 1rem;
    line-height: 1.9rem;
    font-weight: 400;
    color: #2D3240;
    
    text-align: ${ p => p.align};

    &:last-child {
        padding-right: 1.5rem;
    }
`

TableBodyCell.propTypes = {
  children: node,
  align: oneOf(['left', 'center', 'right']).isRequired,
}

TableBodyCell.defaultProps = {
  align: 'left',
}

export default TableBodyCell
