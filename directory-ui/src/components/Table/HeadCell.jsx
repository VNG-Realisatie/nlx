// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import { node, oneOf } from 'prop-types'
import styled from 'styled-components'

const HeadCell = styled.th`
    display: table-cell;
    padding: 12px 16px 12px 0;

    font-size: 12px;
    line-height: 20px;
    font-weight: 600;
    color: #A3AABF;
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
