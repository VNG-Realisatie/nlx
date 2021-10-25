// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//

import styled from 'styled-components'
import Table from '../../../../../components/Table'
import { IconOutway } from '../../../../../icons'

export const StyledIconTd = styled(Table.Td)`
  width: 3rem;
`

export const StyledOutwayIcon = styled(IconOutway)`
  fill: ${(p) => p.theme.colorFocus};
`
