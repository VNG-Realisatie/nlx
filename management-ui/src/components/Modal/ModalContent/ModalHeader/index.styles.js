// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import { IconClose } from '../../../../icons'

export const Header = styled.header`
  display: flex;
  justify-content: space-between;
  width: 100%;
  padding: ${(p) => {
    const { spacing03, spacing05, spacing07 } = p.theme.tokens
    const noTitleTop = p.hasCloseButton ? spacing03 : spacing05
    return p.hasTitle
      ? `${spacing05} ${spacing05} ${spacing05} ${spacing07}`
      : `${noTitleTop} ${spacing05} ${spacing05} ${spacing07}`
  }};
  background-color: ${(p) => p.theme.tokens.colorBackgroundAlt};
`

export const Title = styled.h1`
  padding: ${(p) => `${p.theme.tokens.spacing03} 0`};
  margin: 0 ${(p) => p.theme.tokens.spacing05} 0 0;
`

// TODO: When moving modal to design-system, make a reusable close button which
// combines the common styling between the drawer's close button and this one
export const CloseButton = styled.button`
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: ${(p) => {
    const { spacing03 } = p.theme.tokens
    return `${spacing03}`
  }};
  /* Minimum sizing for icon-only buttons */
  min-width: 2.25rem;
  min-height: 2.25rem;
  border: none;
  margin-left: auto;
  vertical-align: middle;
  font-size: ${(p) => p.theme.tokens.fontSizeMedium};
  font-weight: ${(p) => p.theme.tokens.fontWeightSemiBold};
  line-height: ${(p) => p.theme.tokens.lineHeightText};
  text-align: center;
  text-decoration: none;
  cursor: pointer;
  user-select: none;
  background-color: transparent;

  &:focus {
    outline: 2px solid ${(p) => p.theme.tokens.colorFocus};
  }
`

export const StyledIconClose = styled(IconClose)`
  width: 32px;
  height: 32px;
  fill: ${(p) => p.theme.tokens.colorPaletteGray400};
`
