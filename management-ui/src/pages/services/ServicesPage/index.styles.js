// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import { IconPlus } from '../../../icons'

export const StyledActionsBar = styled.div`
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
`

export const StyledIconPlus = styled(IconPlus)`
  margin-right: ${(p) => p.theme.tokens.spacing03};
`
