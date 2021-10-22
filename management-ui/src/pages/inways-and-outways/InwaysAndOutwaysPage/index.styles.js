// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import styled from 'styled-components'
import { NavLink } from 'react-router-dom'
import { IconInway, IconOutway } from '../../../icons'

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

export const StyledIconInway = styled(IconInway)`
  fill: ${(p) => p.theme.colorPaletteGray50};
  margin-right: ${(p) => p.theme.tokens.spacing03};
`

export const StyledIconOutway = styled(IconOutway)`
  fill: ${(p) => p.theme.colorFocus};
  margin-right: ${(p) => p.theme.tokens.spacing03};
`
