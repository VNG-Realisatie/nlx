// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import styled from 'styled-components'
import { DetailHeading } from '../../../../../../../components/DetailView'

export const StyledContainer = styled(DetailHeading)`
  display: flex;
`

export const StyledLabel = styled.p`
  flex: 1;
  text-align: right;
  font-weight: normal;
  margin: 0;
`

export const StateDetail = styled.div`
  flex: 1 1 auto;
  display: flex;
  flex-direction: column;
  justify-content: center;
  min-height: 50px;
  font-weight: normal;
`
