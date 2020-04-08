// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import NLXManagementLogo from '../../components/NLXManagementLogo'
import { ReactComponent as ExternalLink } from './external-link.svg'
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
`

export const StyledExternalLink = styled(ExternalLink)`
  height: 1rem;
  margin-left: ${(p) => p.theme.tokens.spacing01};
`
