// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import RemoveButton from '../../../components/RemoveButton'
import { ServiceVisibilityAlert } from '../../../components/ServiceVisibilityAlert'

export const StyledInwayName = styled.span`
  flex-grow: 1;
`

export const StyledActionsBar = styled.div`
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  padding-bottom: ${(p) => p.theme.tokens.spacing05};
`

export const StyledRemoveButton = styled(RemoveButton)`
  margin-left: auto;
`

export const StyledServiceVisibilityAlert = styled(ServiceVisibilityAlert)`
  margin-top: ${(p) => p.theme.tokens.spacing06};
  margin-bottom: ${(p) => p.theme.tokens.spacing07};
  width: 100%;
`
