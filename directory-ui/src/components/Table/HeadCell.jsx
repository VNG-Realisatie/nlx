import { node, oneOf } from 'prop-types'
import styled from 'styled-components'

const HeadCell = styled.th`
    display: table-cell;
    padding: .65rem;

    font-size: .85rem;
    line-height: 1.25rem;
    font-weight: 600;
    color: #ADB3C6;
    overflow: hidden;
    text-overflow: ellipsis;
    text-transform: uppercase;
    white-space: nowrap;
    
    text-align: ${ p => p.align};
`

HeadCell.propTypes = {
  children: node,
  align: oneOf(['left', 'center', 'right'])
}

HeadCell.defaultProps = {
  align: 'left'
}

export default HeadCell
