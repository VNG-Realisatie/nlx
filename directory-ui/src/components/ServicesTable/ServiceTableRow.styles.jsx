import styled from 'styled-components'
import Table from '../Table'

const offlineBackgroundColor = '#F7F9FC'
const offlineTextColor = '#A3AABF'
const onlineBackgroundColor = '#FFFFFF'
const onlineTextColor = '#2D3240'

export const StyledServiceTableRow = styled(Table.Row)`
  td {
    background-color: ${
      p => p.status === 'offline' ?
        offlineBackgroundColor : onlineBackgroundColor
      }
  }
  color: ${
    p => p.status === 'offline' ?
      offlineTextColor : onlineTextColor
  }
`

export const StyledApiTypeLabel = styled.span`
  display: inline-flex;
  font-size: 12px;
  line-height: 20px;
  padding: 3px 8px 5px;
  border-radius: 3px;
  border: 1px solid #CAD0E0;
  opacity: ${
  p => p.status === 'offline' ?
    .4 : 1
  }
`
