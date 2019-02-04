import { node, oneOf } from 'prop-types'
import styled from 'styled-components'

const TableHeadCell = styled.th`
    display: table-cell;
    border-bottom: 1px solid #EAEAEA;
    padding: .65rem;

    font-size: .85rem;
    line-height: 1.25rem;
    font-weight: 700;
    color: #ADB3C6;
    overflow: hidden;
    text-overflow: ellipsis;
    text-transform: uppercase;
    white-space: nowrap;
    
    text-align: ${ p => p.align};

    &:last-child {
        padding-right: 1.5rem;
    }
`

TableHeadCell.propTypes = {
  children: node,
  align: oneOf(['left', 'center', 'right']),
}

TableHeadCell.defaultProps = {
  align: 'left',
}

export default TableHeadCell
