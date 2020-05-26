// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import { ReactComponent as NLXManagementLogo } from '../../components/NLXManagementLogo/nlxManagementLogo.svg'
import { ReactComponent as ExternalLink } from './external-link.svg'
import { ReactComponent as OrganizationIcon } from './organization.svg'
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
  background: #313131;
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

export const StyledExternalLink = styled(ExternalLink)`
  height: 1rem;
  margin-left: ${(p) => p.theme.tokens.spacing01};
`

export const StyledOrganization = styled.p`
  border-top: 1px ${(p) => p.theme.tokens.colorPaletteGray700} solid;
  border-bottom: 1px ${(p) => p.theme.tokens.colorPaletteGray700} solid;
  padding: ${(p) => p.theme.tokens.spacing07} 0;
  margin: ${(p) => p.theme.tokens.spacing08} 0
    ${(p) => p.theme.tokens.spacing09};
`

export const StyledOrganizationIcon = styled(OrganizationIcon)`
  fill: ${(p) => p.theme.tokens.colorPaletteGray500};
  margin-right: ${(p) => p.theme.tokens.spacing03};
  vertical-align: text-bottom;
`
