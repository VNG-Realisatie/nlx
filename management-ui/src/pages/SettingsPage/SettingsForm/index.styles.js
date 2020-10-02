// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled, { css } from 'styled-components'
import EmptyContentMessage from '../../../components/EmptyContentMessage'
import LoadingMessage from '../../../components/LoadingMessage'

export const Form = styled.form`
  margin-bottom: ${(p) => p.theme.tokens.spacing10};
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
