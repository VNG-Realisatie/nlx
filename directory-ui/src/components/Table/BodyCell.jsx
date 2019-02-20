import { node, oneOf } from 'prop-types'
import styled from 'styled-components'

const BodyCell = styled.td`
    display: table-cell;
    padding: .6rem;

    font-size: 14px;
    line-height: 22px;
    font-weight: 400;

    text-align: ${p => p.align};
`

BodyCell.propTypes = {
  children: node,
  align: oneOf(['left', 'center', 'right'])
}

BodyCell.defaultProps = {
  align: 'left',
}

export default BodyCell
