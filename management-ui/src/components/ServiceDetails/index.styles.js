// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import RemoveButton from '../RemoveButton'

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
