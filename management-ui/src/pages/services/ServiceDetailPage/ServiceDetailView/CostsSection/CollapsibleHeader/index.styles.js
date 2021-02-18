// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import styled from 'styled-components'
import { DetailHeading } from '../../../../../../components/DetailView'

export const StyledDetailHeading = styled(DetailHeading)`
  display: flex;
`

export const StyledLabel = styled.small`
  flex: 1;
  padding-right: ${(p) => p.theme.tokens.spacing05};
  text-align: right;
  font-weight: normal;

  &&:first-letter {
    text-transform: capitalize;
  }
`
