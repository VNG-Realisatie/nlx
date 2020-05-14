// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled, { css } from 'styled-components'
import { Alert } from '@commonground/design-system'
import EmptyContentMessage from '../EmptyContentMessage'
import LoadingMessage from '../LoadingMessage'

export const Form = styled.form`
  margin-bottom: ${(p) => p.theme.tokens.spacing10};
`

export const CheckboxGroup = styled.div`
  label {
    display: block;
  }
  label:not(:first-child) {
    display: block;
    margin-top: ${(p) => p.theme.tokens.spacing03};
    line-height: ${(p) => p.theme.tokens.lineHeightText};
  }
`

const inwaysMessage = css`
  height: ${(p) => p.theme.tokens.spacing08};
  line-height: ${(p) => p.theme.tokens.spacing08};
`

export const InwaysEmptyMessage = styled(EmptyContentMessage)`
  ${inwaysMessage}
  text-align: left;
`

export const InwaysLoadingMessage = styled(LoadingMessage)`
  ${inwaysMessage}
  justify-content: flex-start;
`

export const VisibilityAlert = styled(Alert)`
  margin-top: ${(p) => p.theme.tokens.spacing06};
  margin-bottom: ${(p) => p.theme.tokens.spacing07};
  width: 30rem;
`

export const ServiceNameWrapper = styled.div`
  margin-top: 0;
  margin-bottom: ${(p) => p.theme.tokens.spacing10};
`
