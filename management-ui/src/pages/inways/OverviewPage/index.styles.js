// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import styled from 'styled-components'
import { NavLink } from 'react-router-dom'

export const ActionsBar = styled.div`
  display: flex;
  margin-bottom: 2rem;
`

export const ActionsBarButton = styled(NavLink)`
  &.active {
    background-color: ${(p) => p.theme.colorBackgroundButtonSecondarySelected};
    color: ${(p) => p.theme.colorTextButtonSecondarySelected};
  }
`
