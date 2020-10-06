// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import styled from 'styled-components'
import HeadCell from './HeadCell'

export const StyledSortableTableHeadCell = styled(HeadCell)`
  cursor: pointer;

  color: ${(p) => p.isSorting && '#2D3240'};
`

export const StyledArrow = styled.svg`
  width: 8px;
  height: 10px;
  position: relative;
  left: 5px;
  top: 1px;
`
