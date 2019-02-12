import styled from 'styled-components'
import TableRow from '../Table/TableRow'

const offlineBackgroundColor = '#F7F9FC'
const offlineTextColor = '#A3AABF'
const onlineBackgroundColor = '#FFFFFF'
const onlineTextColor = '#2D3240'

export const StyledServiceTableRow = styled(TableRow)`
  background-color: ${
    p => p.status === 'offline' ? 
      offlineBackgroundColor : onlineBackgroundColor
  }
  color: ${
    p => p.status === 'offline' ? 
      offlineTextColor : onlineTextColor 
  }
`

export const StyledApiTypeLabel = styled.span`
  padding: .3rem .5rem;
  border-radius: 3px;
  border: 1px solid #CAD0E0;
  opacity: ${
  p => p.status === 'offline' ?
    .4 : 1
  }
`
