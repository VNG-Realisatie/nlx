// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import styled from 'styled-components'
import {
  mobileNavigationHeight,
  mediaQueries,
} from '@commonground/design-system'
import { ReactComponent as LogoVng } from './vng.svg'

export const Wrapper = styled.footer`
  padding: ${(p) => p.theme.tokens.spacing07} 0;
  background: #154967 url('/footer-bg-small.svg') no-repeat center bottom;

  ${mediaQueries.smDown`
    margin-bottom: ${mobileNavigationHeight};
  `}
`

export const FooterContent = styled.div`
  display: flex;
  justify-content: flex-end;
  align-items: center;
  text-align: left;

  > a:not(:last-child) {
    align-items: center;
    display: flex;
    color: ${(p) => p.theme.tokens.colors.colorPaletteGray200};
  }
`

export const List = styled.ul`
  display: flex;
  flex-direction: column;
  padding: 0;
  margin: 0;
  list-style-type: none;

  ${mediaQueries.smUp`
    flex-direction: row;
  `}
`

export const Item = styled.li`
  padding: 0;
  margin: ${(p) => p.theme.tokens.spacing02} 0;

  ${mediaQueries.smUp`
    margin: 0 ${(p) => p.theme.tokens.spacing06} 0 0;
  `}

  a {
    color: ${(p) => p.theme.tokens.colorBackground};
  }
`

export const StyledLogoVng = styled(LogoVng)`
  width: 100px;
`
