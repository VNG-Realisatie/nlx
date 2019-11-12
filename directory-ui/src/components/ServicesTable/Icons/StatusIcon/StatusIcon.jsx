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

  ${p => p.status === 'unknown' && css`
    border-color:  #CAD0E0;
  `}

  ${p => p.status === 'up' && css`
    border-color:  #63D19E;
  `}

  ${p => p.status === 'degraded' && css`
    border-color:  #FEBF24;
  `}

  ${p => p.status === 'down' && css`
    border-color:  #ff0000;
  `}
`

StatusICon.propTypes = {
  status: oneOf(['unknown', 'up', 'degraded', 'down']).isRequired
}

export default StatusICon
