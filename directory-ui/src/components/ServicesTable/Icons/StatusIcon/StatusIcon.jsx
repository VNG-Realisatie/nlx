// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import { oneOf } from 'prop-types'
import styled, {css}from 'styled-components'

const StatusICon = styled.div`
  display: inline-block;
  width: 10px;
  height: 10px;
  border-radius: 50%;
  border: 2px solid transparent;

  ${p => p.status === 'online' && css`
    border-color:  #63D19E;
  `}

  ${p => p.status === 'offline' && css`
    border-color:  #CAD0E0;
  `}
`

StatusICon.propTypes = {
  status: oneOf(['online', 'offline']).isRequired
}

export default StatusICon
