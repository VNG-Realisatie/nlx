// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import { Alert } from '@commonground/design-system'
import { Form } from 'formik'
import NLXManagementLogo from '../../components/NLXManagementLogo'
import { IconExternalLink, IconOrganization } from '../../icons'
import nlxPattern from './nlx-pattern.svg'

export const StyledNLXManagementLogo = styled(NLXManagementLogo)`
  width: 100px;
`

export const Wrapper = styled.div`
  display: flex;
  align-items: center;
  height: 100%;
`

export const StyledAlert = styled(Alert)`
  margin: ${(p) => p.theme.tokens.spacing08} 0;
`

export const StyledSidebar = styled.aside`
  background: ${(p) => p.theme.tokens.colorBackgroundAlt};
  background-image: url(${nlxPattern});
  flex: 0 0 532px;
  text-align: center;
  font-size: 14pt;
  padding-top: 10rem;
  height: 100%;
`

export const Content = styled.main`
  max-width: 70rem;
  padding: 0 10rem;
  margin: auto 0;
`

export const StyledForm = styled(Form)`
  display: flex;
  flex-direction: column;
  width: max-content;

  > button {
    align-self: flex-end;
    justify-self: flex-end;
    width: fit-content;
  }
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
