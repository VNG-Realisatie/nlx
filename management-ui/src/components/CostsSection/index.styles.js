// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//

import styled from 'styled-components'
import Table from '../../../../../components/Table'

export const TdPrice = styled(Table.Td)`
  text-align: right;
`

export const StyledLabel = styled.small`
  flex: 1;
  padding-right: ${(p) => p.theme.tokens.spacing05};
  text-align: right;
  font-weight: normal;
`
