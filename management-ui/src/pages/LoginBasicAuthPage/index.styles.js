// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import { Alert } from '@commonground/design-system'
import { Form } from 'formik'
import NLXManagementLogo from '../../components/NLXManagementLogo'
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
