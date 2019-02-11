import styled from 'styled-components'
import TableRow from '../Table/TableRow'

const offlineBackgroundColor = '#F7F9FC'
const offlineTextColor = '#A3AABF'
const onlineBackgroundColor = '#FFFFFF'
const onlineTextColor = '#2D3240'

export default styled(TableRow)`
  background-color: ${
    p => p.status === 'offline' ? 
      offlineBackgroundColor : onlineBackgroundColor
  }
  color: ${
    p => p.status === 'offline' ? 
      offlineTextColor : onlineTextColor 
  }
`
