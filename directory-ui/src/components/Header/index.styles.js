// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import styled from 'styled-components'
import { NLXLogo, Icon, mediaQueries } from '@commonground/design-system'

export const StyledIcon = styled(Icon)`
  fill: ${(p) => p.theme.tokens.colorPaletteGray600};
`

export const LogoWrapper = styled.header`
  ${(p) => `background: ${p.theme.tokens.colorBackground};`}

  ${(p) =>
    p.homepage &&
    mediaQueries.smDown`background: linear-gradient(90deg, #d6eef9 0%, #b3d0e1 100%);`}
`

export const StyledNLXLogo = styled(NLXLogo)`
  height: 27px;
  margin: ${(p) => p.theme.tokens.spacing07} 0;
`

export const NavigationWrapper = styled.div`
  position: relative;
  z-index: 10;
  background-color: ${(p) => p.theme.tokens.colorBrand1};
`
