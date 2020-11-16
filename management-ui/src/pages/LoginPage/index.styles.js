// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import NLXManagementLogo from '../../components/NLXManagementLogo'
import { IconExternalLink, IconOrganization } from '../../icons'
import nlxPattern from './nlx-pattern.svg'

export const StyledNLXManagementLogo = styled(NLXManagementLogo)`
  width: 100px;
`

export const StyledContainer = styled.main`
  display: flex;
  align-items: center;
  height: 100%;
`

export const StyledSidebar = styled.div`
  background: ${(p) => p.theme.tokens.colorBackgroundAlt};
  background-image: url(${nlxPattern});
  flex: 0 0 532px;
  text-align: center;
  font-size: 14pt;
  padding-top: 10rem;
  height: 100%;
`

export const StyledContent = styled.div`
  flex: 1;
  padding: 0 10rem 10rem 10rem;
  max-width: 70rem;
`

export const StyledIconExternalLink = styled(IconExternalLink)`
  margin-left: ${(p) => p.theme.tokens.spacing02};
`

export const StyledOrganization = styled.p`
  border-top: 1px ${(p) => p.theme.tokens.colorPaletteGray700} solid;
  border-bottom: 1px ${(p) => p.theme.tokens.colorPaletteGray700} solid;
  padding: ${(p) => p.theme.tokens.spacing07} 0;
  margin: ${(p) => p.theme.tokens.spacing08} 0
    ${(p) => p.theme.tokens.spacing09};
`

export const StyledIconOrganization = styled(IconOrganization)`
  fill: ${(p) => p.theme.tokens.colorPaletteGray500};
`
